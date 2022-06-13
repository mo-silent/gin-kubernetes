package router

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type DeploymentRouter struct{}

// InitDeploymentRouter 初始化 deployment 路由
func (deploy *DeploymentRouter) InitDeploymentRouter(Router *gin.RouterGroup) {
	deploymentGroup := Router.Group("deployment")
	deploymentApi := v1.ApiGroupEnter.DeploymentApi
	{
		deploymentGroup.GET("getDeployment/:namespace", deploymentApi.GetDeployment) // 获取 deployment 信息
		deploymentGroup.GET("listDeployment", deploymentApi.ListDeployment)          // 获取命名空间下的所有 deployment 信息
		deploymentGroup.POST("createDeployment", deploymentApi.CreateDeployment)     // 创建 deployment
		deploymentGroup.DELETE("deleteDeployment", deploymentApi.DeleteDeployment)   // 删除 deployment
	}
}
