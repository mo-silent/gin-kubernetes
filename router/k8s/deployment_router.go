package k8s

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type DeploymentRouter struct{}

// DeploymentGetter deployment router enter
type DeploymentGetter interface {
	Deployment() RouterK8SInterface
}

// newDeployments return deployment router instance
func newDeployments() *DeploymentRouter {
	return &DeploymentRouter{}
}

// InitRouter InitDeploymentRouter init deployment router
func (deploy *DeploymentRouter) InitRouter(Router *gin.RouterGroup) {
	deploymentGroup := Router.Group("deployment")
	deploymentApi := v1.ApiV1Enter().ApiV1K8S().Deployment()
	{
		deploymentGroup.POST("create", deploymentApi.Create)   // 创建 deployment
		deploymentGroup.DELETE("delete", deploymentApi.Delete) // 删除 deployment
		deploymentGroup.PUT("update", deploymentApi.Update)    // 更新 deployment
		deploymentGroup.GET("get", deploymentApi.Get)          // 获取 deployment 信息
		deploymentGroup.GET("list", deploymentApi.List)        // 获取命名空间下的所有 deployment 信息
	}
}
