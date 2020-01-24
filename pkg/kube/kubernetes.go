package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
	RestConfig *rest.Config
	Clientset  *kubernetes.Clientset
}

func NewByInClusterConfig() (*Kubernetes, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return new(config)
}

func NewByKubeConfig(configPath string) (*Kubernetes, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	return new(config)
}

func new(config *rest.Config) (*Kubernetes, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Kubernetes{
		RestConfig: config,
		Clientset:  clientset,
	}, nil
}
