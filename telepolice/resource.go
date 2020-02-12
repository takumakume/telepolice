package telepolice

import corev1 "k8s.io/api/core/v1"

type resouce struct {
	pod    corev1.Pod
	status bool
}
