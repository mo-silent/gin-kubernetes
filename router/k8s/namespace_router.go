package k8s

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type NamespaceRouter struct{}

// NamespaceGetter namespace router enter
type NamespaceGetter interface {
	Namespace() RouterK8SInterface
}

// newNamespace return new namespace instance
func newNamespace() *NamespaceRouter {
	return &NamespaceRouter{}
}

// InitRouter InitNamespaceRouter init namespace router
func (p *NamespaceRouter) InitRouter(Router *gin.RouterGroup) {
	namespaceGroup := Router.Group("namespace")
	namespaceApi := v1.ApiV1Enter().ApiV1K8S().Namespace()
	{
		namespaceGroup.POST("create", namespaceApi.Create)   // 创建 namespace
		namespaceGroup.DELETE("delete", namespaceApi.Delete) // 删除 namespace
		// namespaceGroup.PUT("update", namespaceApi.Update)    // 更新 namespace 信息
		namespaceGroup.GET("get", namespaceApi.Get)   // 获取单个 namespace 信息
		namespaceGroup.GET("list", namespaceApi.List) // 获取命名空间的所有 namespace 信息
	}
}
