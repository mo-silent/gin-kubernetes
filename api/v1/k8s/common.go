package k8s

import "github.com/gin-gonic/gin"

// k8s generic function definitions
type K8SCommonInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}
