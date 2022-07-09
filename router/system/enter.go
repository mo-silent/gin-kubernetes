package system

import "github.com/gin-gonic/gin"

// RouterSystemInterface k8s generic router function
type RouterSystemInterface interface {
	InitRouter(Router *gin.RouterGroup)
}

// RouterSystemGroupInterface system router unified interface
type RouterSystemGroupInterface interface {
	UserGetter
	BaseGetter
}

// RouterSystem system router unified enter
type RouterSystem struct{}

// User return user router enter
func (s *RouterSystem) User() RouterSystemInterface {
	return newUsers()
}

// Base return base router enter
func (s *RouterSystem) Base() RouterSystemInterface {
	return newBases()
}
