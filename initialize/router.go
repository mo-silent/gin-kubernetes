package initialize

import (
	"time"

	"gitee.com/MoGD/gin-kubernetes/middleware"
	"gitee.com/MoGD/gin-kubernetes/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouters 初始化 gin 路由
// Return *gin.Engine
func InitRouters() *gin.Engine {
	Router := gin.Default()
	k8sRouter := router.RouterGroupEnter().K8SRouters()
	systemRouter := router.RouterGroupEnter().SystemRouters()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cross-domain support
	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	Router.Use(mwCORS)

	PublicGroup := Router.Group("")
	{
		systemRouter.Base().InitRouter(PublicGroup) // 注册 base 路由，不需要鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		k8sRouter.Depolyment().InitRouter(PrivateGroup) // 注册 pod 路由
		k8sRouter.Pod().InitRouter(PrivateGroup)        // 注册 deployment 路由
		k8sRouter.Namespace().InitRouter(PrivateGroup)  // 注册 namespace 路由
		systemRouter.User().InitRouter(PrivateGroup)    // 注册 user 路由
	}

	return Router
}
