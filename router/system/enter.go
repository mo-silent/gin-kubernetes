package system

import "github.com/gin-gonic/gin"

// SystemRouterInterface k8s generic router function
type SystemRouterInterface interface {
	InitRouter(Router *gin.RouterGroup)
}

// SystemRouterGroupInterface system router unified interface
type SystemRouterGroupInterface interface {
	UserGetter
	BaseGetter
}

// SystemRouter system router unified enter
type SystemRouter struct{}

// User return user router enter
func (s *SystemRouter) User() SystemRouterInterface {
	return newUsers()
}

// Base return base router enter
func (s *SystemRouter) Base() SystemRouterInterface {
	return newBases()
}
