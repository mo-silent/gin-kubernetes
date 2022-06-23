package initialize

import (
	"gitee.com/MoGD/gin-kubernetes/global"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// InitK8sClient 初始化 k8s client
// Return *kubernetes.Clientset
func InitK8sClient() *kubernetes.Clientset {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", global.CONFIG.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}

// InitDynamicK8sClient 初始化动态 k8s client
// Return dynamic.Interface
func InitDynamicK8sClient() dynamic.Interface {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", global.CONFIG.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the client
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}
