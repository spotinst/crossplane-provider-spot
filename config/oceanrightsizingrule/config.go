package oceanrightsizingrule

/*
Copyright 2021 Upbound Inc.
*/

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("spotinst_ocean_right_sizing_rule", func(r *ujconfig.Resource) {

		// we need to override the default group that upjet generated for
		// this resource, which would be "spotinst"
		r.ShortGroup = "oceanrightsizingrule"
		r.Kind = "OceanRightSizingRule"
	})
}
