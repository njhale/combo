/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	// Body is the parameterized template string.
	Body string `json:"body"`

	// +optional
	// Parameters is the set of strings within Body to treat as parameters.
	Parameters []string `json:"parameters,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:categories=combo,scope=Cluster

// Template is the Schema for the provisionerclasses API
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TemplateSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Template
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// CombinationSpec defines the desired state of Combination
type CombinationSpec struct {
	// Template is the name of the template to evaluate.
	Template string `json:"template"`

	// Arguments contains the list of values to use for each parameter.
	// +optional
	Arguments map[string][]string `json:"arguments,omitempty"`
}

// CombinationStatus defines the observed state of Combination
type CombinationStatus struct {
	// Conditions represents the current condition of the Combination.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:categories=combo,scope=Cluster
// +kubebuilder:subresource:status

// Combination is the Schema for the bundles API
type Combination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CombinationSpec   `json:"spec"`
	Status CombinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CombinationList contains a list of Combination
type CombinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Combination `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{}, &Combination{}, &CombinationList{})
}
