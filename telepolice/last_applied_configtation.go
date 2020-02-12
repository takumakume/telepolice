package telepolice

type KubernetesLastAppliedConfiguration struct {
	Metadata KubernetesMetadata `json:"metadata"`
	Spec     KubernetesSpec     `json:"spec"`
}

type KubernetesMetadata struct {
	SelfLink string `json:"selfLink"`
}

type KubernetesSpec struct {
	Replicas int32 `json:"replicas"`
}
