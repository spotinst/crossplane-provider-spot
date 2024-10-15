// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	oceanaks "github.com/spotinst/crossplane-provider-spot/internal/controller/oceanaks/oceanaks"
	oceanaksvng "github.com/spotinst/crossplane-provider-spot/internal/controller/oceanaksvng/oceanaksvng"
	oceanaws "github.com/spotinst/crossplane-provider-spot/internal/controller/oceanaws/oceanaws"
	oceanawslaunchspec "github.com/spotinst/crossplane-provider-spot/internal/controller/oceanawslaunchspec/oceanawslaunchspec"
	providerconfig "github.com/spotinst/crossplane-provider-spot/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		oceanaks.Setup,
		oceanaksvng.Setup,
		oceanaws.Setup,
		oceanawslaunchspec.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
