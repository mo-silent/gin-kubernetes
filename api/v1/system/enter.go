package system

import "github.com/gin-gonic/gin"

// ApiV1SystemInterface system common interface, define unified method
type ApiV1SystemInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// ApiV1SystemEnterInterface system interface, all system-relate method calls
type ApiV1SystemEnterInterface interface {
	UserGetter
	BaseGetter
}

// ApiV1SystemEnter system api unified enter
type ApiV1SystemEnter struct{}

// User return user api struct
func (s *ApiV1SystemEnter) User() ApiV1SystemInterface {
	return newUsers()
}

// Base return base api struct
func (s *ApiV1SystemEnter) Base() BaseInterface {
	return newBases()
}
