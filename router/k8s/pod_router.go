package k8s

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type PodRouter struct{}

// InitPodRouter 初始化 pod 路由
func (p *PodRouter) InitPodRouter(Router *gin.RouterGroup) {
	podGroup := Router.Group("pod")
	podApi := v1.ApiV1Enter().ApiV1K8S().Pod()
	{
		podGroup.POST("create", podApi.Create)   // 创建 pod
		podGroup.DELETE("delete", podApi.Delete) // 删除 pod
		podGroup.PUT("update", podApi.Update)    // 更新 pod 信息
		podGroup.GET("get", podApi.Get)          // 获取单个 pod 信息
		podGroup.GET("list", podApi.List)        // 获取命名空间的所有 pod 信息

	}
}
