/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"fmt"
	"strings"

	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/crossplane-provider-spot/config/oceanaks"
	"github.com/spotinst/crossplane-provider-spot/config/oceanaksvng"
	"github.com/spotinst/crossplane-provider-spot/config/oceanaws"
	"github.com/spotinst/crossplane-provider-spot/config/oceanawslaunchspec"
	"github.com/spotinst/crossplane-provider-spot/config/oceangke"
	"github.com/spotinst/crossplane-provider-spot/config/oceangkevng"
	"github.com/spotinst/crossplane-provider-spot/config/oceanrightsizingrule"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "spot"
	modulePath     = "github.com/spotinst/crossplane-provider-spot"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF Google v4.77.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, fmt.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(sdkProvider *schema.Provider, generationProvider bool) *ujconfig.Provider {
	var includeList []string
	var sdkPluginIncludeList []string

	// Use the spotinst terraform sdk and not the forked terraform cli version to reconcile
	// gke resources because gke resources contain some fields which don't translate well when
	// trying to write them to the file system but do work when using the sdk
	for _, externalName := range ExternalNameConfigured() {
		if strings.Contains(externalName, "ocean_gke") {
			sdkPluginIncludeList = append(sdkPluginIncludeList, externalName)
		} else {
			includeList = append(includeList, externalName)
		}
	}

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			panic(fmt.Errorf("cannot read the Terraform SDK provider from the JSON schema for code generation: %w", err))
		}

		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spot.upbound.io"),
		ujconfig.WithIncludeList(includeList),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformPluginSDKIncludeList(sdkPluginIncludeList),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		oceanaws.Configure,
		oceanawslaunchspec.Configure,
		oceanaks.Configure,
		oceanaksvng.Configure,
		oceanrightsizingrule.Configure,
		oceangke.Configure,
		oceangkevng.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
