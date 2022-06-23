package k8s

import "github.com/gin-gonic/gin"

// K8SRouterInterface k8s generic router function
type K8SRouterInterface interface {
	InitRouter(Router *gin.RouterGroup)
}

// K8SRouterGroupInterface k8s router unified entrance
type K8SRouterGroupInterface interface {
	PodGetter
	DeploymentGetter
	NamespaceGetter
}

type K8SRouter struct {
}

// Depolyment return deployment router instance
func (r *K8SRouter) Depolyment() K8SRouterInterface {
	return newDeployments()
}

// Pod return pod router instance
func (r *K8SRouter) Pod() K8SRouterInterface {
	return newPods()
}

// Namespace return namespace router instance
func (r *K8SRouter) Namespace() K8SRouterInterface {
	return newNamespace()
}
