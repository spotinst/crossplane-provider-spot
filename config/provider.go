/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/spotinst/crossplane-provider-spot/config/oceanaks"
	"github.com/spotinst/crossplane-provider-spot/config/oceanaksvng"
	"github.com/spotinst/crossplane-provider-spot/config/oceanaws"
	"github.com/spotinst/crossplane-provider-spot/config/oceanawslaunchspec"

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

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spot.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		oceanaws.Configure,
		oceanawslaunchspec.Configure,
		oceanaks.Configure,
		oceanaksvng.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
