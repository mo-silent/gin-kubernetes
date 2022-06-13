package initialize

import (
	"gitee.com/MoGD/gin-kubernetes/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouters 初始化 gin 路由
// Return *gin.Engine
func InitRouters() *gin.Engine {
	Router := gin.Default()
	k8sRouter := router.UnifiedRouterGroupEnter.KubeRouterGroup
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	PrivateGroup := Router.Group("")
	{
		k8sRouter.InitPodRouter(PrivateGroup)        // 注册 pod 路由
		k8sRouter.InitDeploymentRouter(PrivateGroup) // 注册 deployment 路由
	}

	return Router
}
