package global

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var (
	KUBECONFIG       *string
	K8SCLIENT        *kubernetes.Clientset
	DynamicK8SCLIENT dynamic.Interface
)
