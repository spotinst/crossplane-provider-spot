//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachWorkloadsObservation) DeepCopyInto(out *AttachWorkloadsObservation) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]NamespacesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachWorkloadsObservation.
func (in *AttachWorkloadsObservation) DeepCopy() *AttachWorkloadsObservation {
	if in == nil {
		return nil
	}
	out := new(AttachWorkloadsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachWorkloadsParameters) DeepCopyInto(out *AttachWorkloadsParameters) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]NamespacesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachWorkloadsParameters.
func (in *AttachWorkloadsParameters) DeepCopy() *AttachWorkloadsParameters {
	if in == nil {
		return nil
	}
	out := new(AttachWorkloadsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetachWorkloadsNamespacesObservation) DeepCopyInto(out *DetachWorkloadsNamespacesObservation) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]NamespacesLabelsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]NamespacesWorkloadsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetachWorkloadsNamespacesObservation.
func (in *DetachWorkloadsNamespacesObservation) DeepCopy() *DetachWorkloadsNamespacesObservation {
	if in == nil {
		return nil
	}
	out := new(DetachWorkloadsNamespacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetachWorkloadsNamespacesParameters) DeepCopyInto(out *DetachWorkloadsNamespacesParameters) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]NamespacesLabelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]NamespacesWorkloadsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetachWorkloadsNamespacesParameters.
func (in *DetachWorkloadsNamespacesParameters) DeepCopy() *DetachWorkloadsNamespacesParameters {
	if in == nil {
		return nil
	}
	out := new(DetachWorkloadsNamespacesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetachWorkloadsObservation) DeepCopyInto(out *DetachWorkloadsObservation) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]DetachWorkloadsNamespacesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetachWorkloadsObservation.
func (in *DetachWorkloadsObservation) DeepCopy() *DetachWorkloadsObservation {
	if in == nil {
		return nil
	}
	out := new(DetachWorkloadsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetachWorkloadsParameters) DeepCopyInto(out *DetachWorkloadsParameters) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]DetachWorkloadsNamespacesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetachWorkloadsParameters.
func (in *DetachWorkloadsParameters) DeepCopy() *DetachWorkloadsParameters {
	if in == nil {
		return nil
	}
	out := new(DetachWorkloadsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsObservation) DeepCopyInto(out *LabelsObservation) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsObservation.
func (in *LabelsObservation) DeepCopy() *LabelsObservation {
	if in == nil {
		return nil
	}
	out := new(LabelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsParameters) DeepCopyInto(out *LabelsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsParameters.
func (in *LabelsParameters) DeepCopy() *LabelsParameters {
	if in == nil {
		return nil
	}
	out := new(LabelsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonthlyRepetitionBasisObservation) DeepCopyInto(out *MonthlyRepetitionBasisObservation) {
	*out = *in
	if in.IntervalMonths != nil {
		in, out := &in.IntervalMonths, &out.IntervalMonths
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.WeekOfTheMonth != nil {
		in, out := &in.WeekOfTheMonth, &out.WeekOfTheMonth
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WeeklyRepetitionBasis != nil {
		in, out := &in.WeeklyRepetitionBasis, &out.WeeklyRepetitionBasis
		*out = make([]WeeklyRepetitionBasisObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonthlyRepetitionBasisObservation.
func (in *MonthlyRepetitionBasisObservation) DeepCopy() *MonthlyRepetitionBasisObservation {
	if in == nil {
		return nil
	}
	out := new(MonthlyRepetitionBasisObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonthlyRepetitionBasisParameters) DeepCopyInto(out *MonthlyRepetitionBasisParameters) {
	*out = *in
	if in.IntervalMonths != nil {
		in, out := &in.IntervalMonths, &out.IntervalMonths
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.WeekOfTheMonth != nil {
		in, out := &in.WeekOfTheMonth, &out.WeekOfTheMonth
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WeeklyRepetitionBasis != nil {
		in, out := &in.WeeklyRepetitionBasis, &out.WeeklyRepetitionBasis
		*out = make([]WeeklyRepetitionBasisParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonthlyRepetitionBasisParameters.
func (in *MonthlyRepetitionBasisParameters) DeepCopy() *MonthlyRepetitionBasisParameters {
	if in == nil {
		return nil
	}
	out := new(MonthlyRepetitionBasisParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesLabelsObservation) DeepCopyInto(out *NamespacesLabelsObservation) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesLabelsObservation.
func (in *NamespacesLabelsObservation) DeepCopy() *NamespacesLabelsObservation {
	if in == nil {
		return nil
	}
	out := new(NamespacesLabelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesLabelsParameters) DeepCopyInto(out *NamespacesLabelsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesLabelsParameters.
func (in *NamespacesLabelsParameters) DeepCopy() *NamespacesLabelsParameters {
	if in == nil {
		return nil
	}
	out := new(NamespacesLabelsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesObservation) DeepCopyInto(out *NamespacesObservation) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LabelsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]WorkloadsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesObservation.
func (in *NamespacesObservation) DeepCopy() *NamespacesObservation {
	if in == nil {
		return nil
	}
	out := new(NamespacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesParameters) DeepCopyInto(out *NamespacesParameters) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LabelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]WorkloadsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesParameters.
func (in *NamespacesParameters) DeepCopy() *NamespacesParameters {
	if in == nil {
		return nil
	}
	out := new(NamespacesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesWorkloadsObservation) DeepCopyInto(out *NamespacesWorkloadsObservation) {
	*out = *in
	if in.RegexName != nil {
		in, out := &in.RegexName, &out.RegexName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadName != nil {
		in, out := &in.WorkloadName, &out.WorkloadName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadType != nil {
		in, out := &in.WorkloadType, &out.WorkloadType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesWorkloadsObservation.
func (in *NamespacesWorkloadsObservation) DeepCopy() *NamespacesWorkloadsObservation {
	if in == nil {
		return nil
	}
	out := new(NamespacesWorkloadsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesWorkloadsParameters) DeepCopyInto(out *NamespacesWorkloadsParameters) {
	*out = *in
	if in.RegexName != nil {
		in, out := &in.RegexName, &out.RegexName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadName != nil {
		in, out := &in.WorkloadName, &out.WorkloadName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadType != nil {
		in, out := &in.WorkloadType, &out.WorkloadType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesWorkloadsParameters.
func (in *NamespacesWorkloadsParameters) DeepCopy() *NamespacesWorkloadsParameters {
	if in == nil {
		return nil
	}
	out := new(NamespacesWorkloadsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationBoundariesObservation) DeepCopyInto(out *RecommendationApplicationBoundariesObservation) {
	*out = *in
	if in.CPUMax != nil {
		in, out := &in.CPUMax, &out.CPUMax
		*out = new(float64)
		**out = **in
	}
	if in.CPUMin != nil {
		in, out := &in.CPUMin, &out.CPUMin
		*out = new(float64)
		**out = **in
	}
	if in.MemoryMax != nil {
		in, out := &in.MemoryMax, &out.MemoryMax
		*out = new(float64)
		**out = **in
	}
	if in.MemoryMin != nil {
		in, out := &in.MemoryMin, &out.MemoryMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationBoundariesObservation.
func (in *RecommendationApplicationBoundariesObservation) DeepCopy() *RecommendationApplicationBoundariesObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationBoundariesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationBoundariesParameters) DeepCopyInto(out *RecommendationApplicationBoundariesParameters) {
	*out = *in
	if in.CPUMax != nil {
		in, out := &in.CPUMax, &out.CPUMax
		*out = new(float64)
		**out = **in
	}
	if in.CPUMin != nil {
		in, out := &in.CPUMin, &out.CPUMin
		*out = new(float64)
		**out = **in
	}
	if in.MemoryMax != nil {
		in, out := &in.MemoryMax, &out.MemoryMax
		*out = new(float64)
		**out = **in
	}
	if in.MemoryMin != nil {
		in, out := &in.MemoryMin, &out.MemoryMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationBoundariesParameters.
func (in *RecommendationApplicationBoundariesParameters) DeepCopy() *RecommendationApplicationBoundariesParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationBoundariesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationIntervalsObservation) DeepCopyInto(out *RecommendationApplicationIntervalsObservation) {
	*out = *in
	if in.MonthlyRepetitionBasis != nil {
		in, out := &in.MonthlyRepetitionBasis, &out.MonthlyRepetitionBasis
		*out = make([]MonthlyRepetitionBasisObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepetitionBasis != nil {
		in, out := &in.RepetitionBasis, &out.RepetitionBasis
		*out = new(string)
		**out = **in
	}
	if in.WeeklyRepetitionBasis != nil {
		in, out := &in.WeeklyRepetitionBasis, &out.WeeklyRepetitionBasis
		*out = make([]RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationIntervalsObservation.
func (in *RecommendationApplicationIntervalsObservation) DeepCopy() *RecommendationApplicationIntervalsObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationIntervalsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationIntervalsParameters) DeepCopyInto(out *RecommendationApplicationIntervalsParameters) {
	*out = *in
	if in.MonthlyRepetitionBasis != nil {
		in, out := &in.MonthlyRepetitionBasis, &out.MonthlyRepetitionBasis
		*out = make([]MonthlyRepetitionBasisParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepetitionBasis != nil {
		in, out := &in.RepetitionBasis, &out.RepetitionBasis
		*out = new(string)
		**out = **in
	}
	if in.WeeklyRepetitionBasis != nil {
		in, out := &in.WeeklyRepetitionBasis, &out.WeeklyRepetitionBasis
		*out = make([]RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationIntervalsParameters.
func (in *RecommendationApplicationIntervalsParameters) DeepCopy() *RecommendationApplicationIntervalsParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationIntervalsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation) DeepCopyInto(out *RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation) {
	*out = *in
	if in.IntervalDays != nil {
		in, out := &in.IntervalDays, &out.IntervalDays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IntervalHoursEndTime != nil {
		in, out := &in.IntervalHoursEndTime, &out.IntervalHoursEndTime
		*out = new(string)
		**out = **in
	}
	if in.IntervalHoursStartTime != nil {
		in, out := &in.IntervalHoursStartTime, &out.IntervalHoursStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation.
func (in *RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation) DeepCopy() *RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationIntervalsWeeklyRepetitionBasisObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters) DeepCopyInto(out *RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters) {
	*out = *in
	if in.IntervalDays != nil {
		in, out := &in.IntervalDays, &out.IntervalDays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IntervalHoursEndTime != nil {
		in, out := &in.IntervalHoursEndTime, &out.IntervalHoursEndTime
		*out = new(string)
		**out = **in
	}
	if in.IntervalHoursStartTime != nil {
		in, out := &in.IntervalHoursStartTime, &out.IntervalHoursStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters.
func (in *RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters) DeepCopy() *RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationIntervalsWeeklyRepetitionBasisParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationMinThresholdObservation) DeepCopyInto(out *RecommendationApplicationMinThresholdObservation) {
	*out = *in
	if in.CPUPercentage != nil {
		in, out := &in.CPUPercentage, &out.CPUPercentage
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPercentage != nil {
		in, out := &in.MemoryPercentage, &out.MemoryPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationMinThresholdObservation.
func (in *RecommendationApplicationMinThresholdObservation) DeepCopy() *RecommendationApplicationMinThresholdObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationMinThresholdObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationMinThresholdParameters) DeepCopyInto(out *RecommendationApplicationMinThresholdParameters) {
	*out = *in
	if in.CPUPercentage != nil {
		in, out := &in.CPUPercentage, &out.CPUPercentage
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPercentage != nil {
		in, out := &in.MemoryPercentage, &out.MemoryPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationMinThresholdParameters.
func (in *RecommendationApplicationMinThresholdParameters) DeepCopy() *RecommendationApplicationMinThresholdParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationMinThresholdParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationOverheadValuesObservation) DeepCopyInto(out *RecommendationApplicationOverheadValuesObservation) {
	*out = *in
	if in.CPUPercentage != nil {
		in, out := &in.CPUPercentage, &out.CPUPercentage
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPercentage != nil {
		in, out := &in.MemoryPercentage, &out.MemoryPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationOverheadValuesObservation.
func (in *RecommendationApplicationOverheadValuesObservation) DeepCopy() *RecommendationApplicationOverheadValuesObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationOverheadValuesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationApplicationOverheadValuesParameters) DeepCopyInto(out *RecommendationApplicationOverheadValuesParameters) {
	*out = *in
	if in.CPUPercentage != nil {
		in, out := &in.CPUPercentage, &out.CPUPercentage
		*out = new(float64)
		**out = **in
	}
	if in.MemoryPercentage != nil {
		in, out := &in.MemoryPercentage, &out.MemoryPercentage
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationApplicationOverheadValuesParameters.
func (in *RecommendationApplicationOverheadValuesParameters) DeepCopy() *RecommendationApplicationOverheadValuesParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationApplicationOverheadValuesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRule) DeepCopyInto(out *RightSizingRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRule.
func (in *RightSizingRule) DeepCopy() *RightSizingRule {
	if in == nil {
		return nil
	}
	out := new(RightSizingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RightSizingRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRuleList) DeepCopyInto(out *RightSizingRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RightSizingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRuleList.
func (in *RightSizingRuleList) DeepCopy() *RightSizingRuleList {
	if in == nil {
		return nil
	}
	out := new(RightSizingRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RightSizingRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRuleObservation) DeepCopyInto(out *RightSizingRuleObservation) {
	*out = *in
	if in.AttachWorkloads != nil {
		in, out := &in.AttachWorkloads, &out.AttachWorkloads
		*out = make([]AttachWorkloadsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DetachWorkloads != nil {
		in, out := &in.DetachWorkloads, &out.DetachWorkloads
		*out = make([]DetachWorkloadsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OceanID != nil {
		in, out := &in.OceanID, &out.OceanID
		*out = new(string)
		**out = **in
	}
	if in.RecommendationApplicationBoundaries != nil {
		in, out := &in.RecommendationApplicationBoundaries, &out.RecommendationApplicationBoundaries
		*out = make([]RecommendationApplicationBoundariesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationIntervals != nil {
		in, out := &in.RecommendationApplicationIntervals, &out.RecommendationApplicationIntervals
		*out = make([]RecommendationApplicationIntervalsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationMinThreshold != nil {
		in, out := &in.RecommendationApplicationMinThreshold, &out.RecommendationApplicationMinThreshold
		*out = make([]RecommendationApplicationMinThresholdObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationOverheadValues != nil {
		in, out := &in.RecommendationApplicationOverheadValues, &out.RecommendationApplicationOverheadValues
		*out = make([]RecommendationApplicationOverheadValuesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestartPods != nil {
		in, out := &in.RestartPods, &out.RestartPods
		*out = new(bool)
		**out = **in
	}
	if in.RuleName != nil {
		in, out := &in.RuleName, &out.RuleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRuleObservation.
func (in *RightSizingRuleObservation) DeepCopy() *RightSizingRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RightSizingRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRuleParameters) DeepCopyInto(out *RightSizingRuleParameters) {
	*out = *in
	if in.AttachWorkloads != nil {
		in, out := &in.AttachWorkloads, &out.AttachWorkloads
		*out = make([]AttachWorkloadsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DetachWorkloads != nil {
		in, out := &in.DetachWorkloads, &out.DetachWorkloads
		*out = make([]DetachWorkloadsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OceanID != nil {
		in, out := &in.OceanID, &out.OceanID
		*out = new(string)
		**out = **in
	}
	if in.RecommendationApplicationBoundaries != nil {
		in, out := &in.RecommendationApplicationBoundaries, &out.RecommendationApplicationBoundaries
		*out = make([]RecommendationApplicationBoundariesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationIntervals != nil {
		in, out := &in.RecommendationApplicationIntervals, &out.RecommendationApplicationIntervals
		*out = make([]RecommendationApplicationIntervalsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationMinThreshold != nil {
		in, out := &in.RecommendationApplicationMinThreshold, &out.RecommendationApplicationMinThreshold
		*out = make([]RecommendationApplicationMinThresholdParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecommendationApplicationOverheadValues != nil {
		in, out := &in.RecommendationApplicationOverheadValues, &out.RecommendationApplicationOverheadValues
		*out = make([]RecommendationApplicationOverheadValuesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestartPods != nil {
		in, out := &in.RestartPods, &out.RestartPods
		*out = new(bool)
		**out = **in
	}
	if in.RuleName != nil {
		in, out := &in.RuleName, &out.RuleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRuleParameters.
func (in *RightSizingRuleParameters) DeepCopy() *RightSizingRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RightSizingRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRuleSpec) DeepCopyInto(out *RightSizingRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRuleSpec.
func (in *RightSizingRuleSpec) DeepCopy() *RightSizingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RightSizingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RightSizingRuleStatus) DeepCopyInto(out *RightSizingRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RightSizingRuleStatus.
func (in *RightSizingRuleStatus) DeepCopy() *RightSizingRuleStatus {
	if in == nil {
		return nil
	}
	out := new(RightSizingRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyRepetitionBasisObservation) DeepCopyInto(out *WeeklyRepetitionBasisObservation) {
	*out = *in
	if in.IntervalDays != nil {
		in, out := &in.IntervalDays, &out.IntervalDays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IntervalHoursEndTime != nil {
		in, out := &in.IntervalHoursEndTime, &out.IntervalHoursEndTime
		*out = new(string)
		**out = **in
	}
	if in.IntervalHoursStartTime != nil {
		in, out := &in.IntervalHoursStartTime, &out.IntervalHoursStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyRepetitionBasisObservation.
func (in *WeeklyRepetitionBasisObservation) DeepCopy() *WeeklyRepetitionBasisObservation {
	if in == nil {
		return nil
	}
	out := new(WeeklyRepetitionBasisObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyRepetitionBasisParameters) DeepCopyInto(out *WeeklyRepetitionBasisParameters) {
	*out = *in
	if in.IntervalDays != nil {
		in, out := &in.IntervalDays, &out.IntervalDays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IntervalHoursEndTime != nil {
		in, out := &in.IntervalHoursEndTime, &out.IntervalHoursEndTime
		*out = new(string)
		**out = **in
	}
	if in.IntervalHoursStartTime != nil {
		in, out := &in.IntervalHoursStartTime, &out.IntervalHoursStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyRepetitionBasisParameters.
func (in *WeeklyRepetitionBasisParameters) DeepCopy() *WeeklyRepetitionBasisParameters {
	if in == nil {
		return nil
	}
	out := new(WeeklyRepetitionBasisParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadsObservation) DeepCopyInto(out *WorkloadsObservation) {
	*out = *in
	if in.RegexName != nil {
		in, out := &in.RegexName, &out.RegexName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadName != nil {
		in, out := &in.WorkloadName, &out.WorkloadName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadType != nil {
		in, out := &in.WorkloadType, &out.WorkloadType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadsObservation.
func (in *WorkloadsObservation) DeepCopy() *WorkloadsObservation {
	if in == nil {
		return nil
	}
	out := new(WorkloadsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadsParameters) DeepCopyInto(out *WorkloadsParameters) {
	*out = *in
	if in.RegexName != nil {
		in, out := &in.RegexName, &out.RegexName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadName != nil {
		in, out := &in.WorkloadName, &out.WorkloadName
		*out = new(string)
		**out = **in
	}
	if in.WorkloadType != nil {
		in, out := &in.WorkloadType, &out.WorkloadType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadsParameters.
func (in *WorkloadsParameters) DeepCopy() *WorkloadsParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadsParameters)
	in.DeepCopyInto(out)
	return out
}