package oceangkevng

/*
Copyright 2021 Upbound Inc.
*/

import (
	"encoding/base64"
	"fmt"
	"reflect"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("spotinst_ocean_gke_launch_spec", func(r *ujconfig.Resource) {

		// we need to override the default group that upjet generated for
		// this resource, which would be "spotinst"
		r.ShortGroup = "oceangkevng"
		r.Kind = "OceanGkeVng"

		r.TerraformConversions = append(r.TerraformConversions, &metadataConverter{})
	})
}

var _ ujconfig.TerraformConversion = &metadataConverter{}

type metadataConverter struct {
}

func (m *metadataConverter) Convert(params map[string]any, r *ujconfig.Resource, mode ujconfig.Mode) (map[string]any, error) {
	if mode == ujconfig.FromTerraform {
		metadata := params["metadata"].([]map[string]any)
		fmt.Println("~~~~~", reflect.TypeOf(metadata).Name())

		for _, mi := range metadata {
			mi["value"] = base64.StdEncoding.EncodeToString([]byte(mi["value"].(string)))
		}

		params["metadata"] = metadata

	} else if mode == ujconfig.ToTerraform {
		metadata := params["metadata"].([]map[string]any)

		for _, mi := range metadata {
			decodeString, err := base64.StdEncoding.DecodeString(mi["value"].(string))
			if err != nil {
				fmt.Println("Error decoding base64 encoded string")
				continue
			}

			mi["value"] = string(decodeString)
		}

		params["metadata"] = metadata
	}

	return params, nil
}
