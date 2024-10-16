package oceanaws

/*
Copyright 2021 Upbound Inc.
*/

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("spotinst_ocean_aws", func(r *ujconfig.Resource) {

		// we need to override the default group that upjet generated for
		// this resource, which would be "spotinst"
		r.ShortGroup = "oceanaws"
		r.Kind = "OceanAws"
	})
}
