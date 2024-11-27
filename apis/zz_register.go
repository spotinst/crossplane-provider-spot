// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/spotinst/crossplane-provider-spot/apis/oceanaks/v1alpha1"
	v1alpha1oceanaksvng "github.com/spotinst/crossplane-provider-spot/apis/oceanaksvng/v1alpha1"
	v1alpha1oceanaws "github.com/spotinst/crossplane-provider-spot/apis/oceanaws/v1alpha1"
	v1alpha1oceanawslaunchspec "github.com/spotinst/crossplane-provider-spot/apis/oceanawslaunchspec/v1alpha1"
	v1alpha1oceangcp "github.com/spotinst/crossplane-provider-spot/apis/oceangcp/v1alpha1"
	v1alpha1oceanrightsizingrule "github.com/spotinst/crossplane-provider-spot/apis/oceanrightsizingrule/v1alpha1"
	v1alpha1apis "github.com/spotinst/crossplane-provider-spot/apis/v1alpha1"
	v1beta1 "github.com/spotinst/crossplane-provider-spot/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1oceanaksvng.SchemeBuilder.AddToScheme,
		v1alpha1oceanaws.SchemeBuilder.AddToScheme,
		v1alpha1oceanawslaunchspec.SchemeBuilder.AddToScheme,
		v1alpha1oceangcp.SchemeBuilder.AddToScheme,
		v1alpha1oceanrightsizingrule.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
