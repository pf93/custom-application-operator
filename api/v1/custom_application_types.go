/*
Copyright 2024 pf93.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CustomApplicationSpec defines the desired state of CustomApplication
type CustomApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CustomApplication. Edit customapplication_types.go to remove/update
	Deployment DeploymentTemplate	`json:"deployment,omitempty"`
	Service    ServiceTemplate      `json:"servife,omitempty"`
}

type DeploymentTemplate struct {
	appsv1.DeploymentSpec	`json:",inline"`
}

type ServiceTemplate struct {
	corev1.ServiceSpec	`json:",inline"`
}

// CustomApplicationStatus defines the observed state of CustomApplication
type CustomApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Workflow	appsv1.DeploymentStatus	  `json:"workflow"`
	Network     corev1.ServiceStatus	  `json:"network"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomApplication is the Schema for the customapplications API
type CustomApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomApplicationSpec   `json:"spec,omitempty"`
	Status CustomApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomApplicationList contains a list of CustomApplication
type CustomApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomApplication{}, &CustomApplicationList{})
}
