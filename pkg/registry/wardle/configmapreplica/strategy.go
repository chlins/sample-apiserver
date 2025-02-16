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

package configmapreplica

import (
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/sample-apiserver/pkg/apis/wardle"
)

// NewStrategy creates and returns a ConfigMapReplicaStrategy instance
func NewStrategy(typer runtime.ObjectTyper) ConfigMapReplicaStrategy {
	return ConfigMapReplicaStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a ConfigMapReplica
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*wardle.ConfigMapReplica)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a ConfigMapReplica")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchConfigMapReplica is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchConfigMapReplica(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *wardle.ConfigMapReplica) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type ConfigMapReplicaStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (ConfigMapReplicaStrategy) NamespaceScoped() bool {
	return true
}

func (ConfigMapReplicaStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (ConfigMapReplicaStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (ConfigMapReplicaStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (ConfigMapReplicaStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (ConfigMapReplicaStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (ConfigMapReplicaStrategy) Canonicalize(obj runtime.Object) {
}

func (ConfigMapReplicaStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
