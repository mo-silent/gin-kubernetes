package main

import (
	"flag"
	"net/http"
	"path/filepath"
	"time"

	"gitee.com/MoGD/gin-kubernetes/core"
	_ "gitee.com/MoGD/gin-kubernetes/docs"
	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/initialize"
	"k8s.io/client-go/util/homedir"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title K8S Client API
// @version 0.0.1
// @description This is a k8s cluster management service
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// 传入的参数
	if home := homedir.HomeDir(); home != "" {
		global.KUBECONFIG = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		global.KUBECONFIG = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// create the k8sClient
	// global.K8SCLIENT = initialize.InitK8sClient(global.KUBECONFIG)
	// global.DynamicK8SCLIENT = initialize.InitDynamicK8sClient(global.KUBECONFIG)

	// Generate variables by importing a configuration file from viper
	global.VP = core.Viper()
	// 初始化路由
	router := initialize.InitRouters()
	// Custom HTTP configuration
	s := &http.Server{
		Addr:           global.CONFIG.System.Addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// server start to listen
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
