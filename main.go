package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"gitee.com/MoGD/gin-kubernetes/core"
	_ "gitee.com/MoGD/gin-kubernetes/docs"
	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/initialize"
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
	// Generate variables by importing a configuration file from viper
	global.VP = core.Viper()

	global.DB = initialize.InitDB()
	// register tables
	if global.DB != nil {
		initialize.RegisterTable(global.DB)
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}

	// create the k8sClient
	// global.K8SCLIENT = initialize.InitK8sClient(global.KUBECONFIG)
	// global.DynamicK8SCLIENT = initialize.InitDynamicK8sClient(global.KUBECONFIG)

	initRun()

}

// init route run init
func initRun() {
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
	// custom defined print
	fmt.Printf(`
	欢迎使用 gin-kubernetes
	当前版本:1.0.0beta
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, global.CONFIG.System.Addr)
	// server start to listen
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
