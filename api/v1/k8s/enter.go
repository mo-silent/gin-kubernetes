package k8s

import "github.com/gin-gonic/gin"

// ApiV1K8sInterface k8s generic function definitions
type ApiV1K8sInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// ApiV1K8sEnter k8s api unified entrance interface
type ApiV1K8sEnter interface {
	PodGetter
	DeploymentGetter
	NamespaceGetter
}

// ApiV1K8SEnter k8s api unified entrance
type ApiV1K8SEnter struct {
}

// Deployment Deployment return a DeploymentInterface
func (g *ApiV1K8SEnter) Deployment() ApiV1K8sInterface {
	return newDeployments()
}

// Pod return a PodInterface
func (g *ApiV1K8SEnter) Pod() ApiV1K8sInterface {
	return newPods()
}

func (g *ApiV1K8SEnter) Namespace() ApiV1K8sInterface {
	return newNamespace()
}
