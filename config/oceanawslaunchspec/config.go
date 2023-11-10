package oceanawslaunchspec

/*
Copyright 2021 Upbound Inc.
*/

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("spotinst_ocean_aws_launch_spec", func(r *ujconfig.Resource) {

		// we need to override the default group that upjet generated for
		// this resource, which would be "spotinst"
		r.ShortGroup = "oceanawslaunchspec"
		r.Kind = "OceanAwsLaunchSpec"
	})
}
