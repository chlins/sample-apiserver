/*
Copyright 2017 The Kubernetes Authors.

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapReplicaList is a list of Flunder objects.
type ConfigMapReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ConfigMapReplica `json:"items"`
}

// ConfigMapReplicaSpec is ConfigMapReplica Spec
type ConfigMapReplicaSpec struct {
	Namespace           string   `json:"namespace"`
	SourceConfigMapName string   `json:"sourceConfigMapName"`
	TargetNamespaces    []string `json:"targetNamespaces"`
}

// ConfigMapReplicaStatus is ConfigMapReplica Status
type ConfigMapReplicaStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapReplica is ConfigMap replica
type ConfigMapReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ConfigMapReplicaSpec   `json:"spec"`
	Status ConfigMapReplicaStatus `json:"status"`
}
