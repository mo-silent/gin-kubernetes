package global

import (
	"gitee.com/MoGD/gin-kubernetes/conf"
	"github.com/spf13/viper"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var (
	KUBECONFIG       *string
	K8SCLIENT        *kubernetes.Clientset
	DynamicK8SCLIENT dynamic.Interface
	CONFIG           conf.Server
	VP               *viper.Viper
)
