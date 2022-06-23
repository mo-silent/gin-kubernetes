package k8s

import "github.com/gin-gonic/gin"

// K8SCommonInterface k8s generic function definitions
type K8SCommonInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// K8SInterface k8s api unified entrance interface
type K8SInterface interface {
	PodGetter
	DeploymentGetter
	NamespaceGetter
}

// ApiV1K8SEnter k8s api unified entrance
type ApiV1K8SEnter struct {
}

// Depolyment return a DeploymentInterface
func (g *ApiV1K8SEnter) Deployment() K8SCommonInterface {
	return newDeployments()
}

// Pod return a PodInterface
func (g *ApiV1K8SEnter) Pod() K8SCommonInterface {
	return newPods()
}

func (g *ApiV1K8SEnter) Namespace() K8SCommonInterface {
	return newNamespace()
}
