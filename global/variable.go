package global

import (
	"gitee.com/MoGD/gin-kubernetes/conf"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var (
	K8SCLIENT        *kubernetes.Clientset
	DynamicK8SCLIENT dynamic.Interface
	CONFIG           conf.Server
	VP               *viper.Viper
	DB               *gorm.DB
)
