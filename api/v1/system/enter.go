package system

import "github.com/gin-gonic/gin"

// SystemCommonInterface system common interface, define unified method
type SystemCommonInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// SystemInterface system interface, all system-relatea method calls
type SystemInterface interface {
	UserGetter
	BaseGetter
}

// ApiV1SystemEnter system api unified enter
type ApiV1SystemEnter struct{}

// User return user api struct
func (s *ApiV1SystemEnter) User() SystemCommonInterface {
	return newUsers()
}

// Base return base api struct
func (s *ApiV1SystemEnter) Base() BaseInterface {
	return newBases()
}
