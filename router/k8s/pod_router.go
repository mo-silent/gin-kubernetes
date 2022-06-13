package k8s

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type PodRouter struct{}

// InitPodRouter 初始化 pod 路由
func (p *PodRouter) InitPodRouter(Router *gin.RouterGroup) {
	podGroup := Router.Group("pod")
	podApi := v1.UnifiedApiGroupEnter.K8SApiGroup
	{
		podGroup.GET("listNamespacePod/:namespace", podApi.ListNamespacePod) // 获取命名空间的所有 pod 信息
		podGroup.GET("listAllPod", podApi.ListAllPod)                        // 获取所有 pod 信息
		podGroup.POST("getPod", podApi.GetPod)                               // 获取单个 pod 信息
		podGroup.POST("createPod", podApi.CreatePod)                         // 创建 pod
		podGroup.DELETE("deletePod", podApi.DeletePod)                       // 删除 pod
	}
}
