/*
Copyright 2023 KhulnaSoft, Inc.

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

package meshsync

import (
	"testing"

	meshplayv1alpha1 "github.com/khulnasoft/meshplay-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetObjects(t *testing.T) {
	m := &meshplayv1alpha1.MeshSync{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "test",
		},
		Spec: meshplayv1alpha1.MeshSyncSpec{
			Size: 1,
		},
	}
	obj := GetObjects(m)
	if obj == nil {
		t.Error("GetObjects returned nil")
	}
	if obj[ServerObject] == nil {
		t.Error("GetObjects returned nil for server object")
	}
}
