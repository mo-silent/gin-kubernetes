package k8s

import "github.com/gin-gonic/gin"

// RouterK8SInterface k8s generic router function
type RouterK8SInterface interface {
	InitRouter(Router *gin.RouterGroup)
}

// RouterK8SGroupInterface k8s router unified entrance
type RouterK8SGroupInterface interface {
	PodGetter
	DeploymentGetter
	NamespaceGetter
}

type RouterK8S struct {
}

// Deployment Deployment return deployment router instance
func (r *RouterK8S) Deployment() RouterK8SInterface {
	return newDeployments()
}

// Pod return pod router instance
func (r *RouterK8S) Pod() RouterK8SInterface {
	return newPods()
}

// Namespace return namespace router instance
func (r *RouterK8S) Namespace() RouterK8SInterface {
	return newNamespace()
}
