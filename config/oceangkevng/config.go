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

type metadataItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (m *metadataConverter) Convert(params map[string]any, r *ujconfig.Resource, mode ujconfig.Mode) (map[string]any, error) {
	if mode == ujconfig.FromTerraform {
		metadata := params["metadata"].([]metadataItem)
		fmt.Println("~~~~~~", reflect.TypeOf(metadata).Name())
		newMetadata := make([]metadataItem, 0, len(metadata))

		for _, mi := range metadata {
			newMetadata = append(newMetadata, metadataItem{
				Key:   mi.Key,
				Value: base64.StdEncoding.EncodeToString([]byte(mi.Value)),
			})
		}

		params["metadata"] = newMetadata

	} else if mode == ujconfig.ToTerraform {
		metadata := params["metadata"].([]metadataItem)
		newMetadata := make([]metadataItem, 0, len(metadata))

		for _, mi := range metadata {
			decodeString, err := base64.StdEncoding.DecodeString(mi.Value)
			if err != nil {
				fmt.Println("Error decoding base64 encoded string")
				continue
			}

			newMetadata = append(newMetadata, metadataItem{
				Key:   mi.Key,
				Value: string(decodeString),
			})
		}

		params["metadata"] = newMetadata
	}

	return params, nil
}
