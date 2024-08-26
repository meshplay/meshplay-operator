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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	val1     int32 = 1
	val60    int64 = 60
	val11000 int32 = 11000

	valtrue bool = true

	MeshplayLabel = map[string]string{
		"app": "meshplay",
	}

	MeshSyncLabel = map[string]string{
		"app":       MeshplayLabel["app"],
		"component": "meshsync",
	}

	MeshplayAnnotation = map[string]string{
		"meshplay/component-type": "management-plane",
	}

	Deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "meshplay-meshsync",
			Labels:      MeshSyncLabel,
			Annotations: MeshplayAnnotation,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &val1,
			Selector: &metav1.LabelSelector{
				MatchLabels: MeshSyncLabel,
			},
			Template: PodTemplate,
		},
	}

	PodTemplate = corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "meshplay-meshsync",
			Labels:      MeshSyncLabel,
			Annotations: MeshplayAnnotation,
		},
		Spec: corev1.PodSpec{
			ServiceAccountName:            "meshplay-operator",
			ShareProcessNamespace:         &valtrue,
			TerminationGracePeriodSeconds: &val60,
			Containers: []corev1.Container{
				{
					Name:            "meshsync",
					Image:           "khulnasoft/meshsync:stable-latest",
					ImagePullPolicy: corev1.PullAlways,
					Ports: []corev1.ContainerPort{
						{
							Name:          "client",
							HostPort:      val11000,
							ContainerPort: val11000,
						},
					},
					Command: []string{
						"./meshplay-meshsync", "--broker-url", "$(BROKER_URL)",
					},
					Env: []corev1.EnvVar{
						{
							Name:  "BROKER_URL",
							Value: "http://localhost:4222",
						},
					},
				},
			},
		},
	}
)
