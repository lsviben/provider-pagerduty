/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RulesetInitParameters struct {

	// Name of the ruleset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to the team that owns the ruleset. If none is specified, only admins have access.
	Team []TeamInitParameters `json:"team,omitempty" tf:"team,omitempty"`
}

type RulesetObservation struct {

	// The ID of the ruleset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the ruleset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Routing keys routed to this ruleset.
	RoutingKeys []*string `json:"routingKeys,omitempty" tf:"routing_keys,omitempty"`

	// Reference to the team that owns the ruleset. If none is specified, only admins have access.
	Team []TeamObservation `json:"team,omitempty" tf:"team,omitempty"`

	// Type of ruleset. Currently, only sets to global.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RulesetParameters struct {

	// Name of the ruleset.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to the team that owns the ruleset. If none is specified, only admins have access.
	// +kubebuilder:validation:Optional
	Team []TeamParameters `json:"team,omitempty" tf:"team,omitempty"`
}

type TeamInitParameters struct {
}

type TeamObservation struct {

	// The ID of the ruleset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TeamParameters struct {

	// The ID of the ruleset.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Team in team to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

// RulesetSpec defines the desired state of Ruleset
type RulesetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RulesetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RulesetInitParameters `json:"initProvider,omitempty"`
}

// RulesetStatus defines the observed state of Ruleset.
type RulesetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RulesetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ruleset is the Schema for the Rulesets API. Creates and manages an ruleset in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Ruleset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   RulesetSpec   `json:"spec"`
	Status RulesetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RulesetList contains a list of Rulesets
type RulesetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ruleset `json:"items"`
}

// Repository type metadata.
var (
	Ruleset_Kind             = "Ruleset"
	Ruleset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ruleset_Kind}.String()
	Ruleset_KindAPIVersion   = Ruleset_Kind + "." + CRDGroupVersion.String()
	Ruleset_GroupVersionKind = CRDGroupVersion.WithKind(Ruleset_Kind)
)

func init() {
	SchemeBuilder.Register(&Ruleset{}, &RulesetList{})
}
