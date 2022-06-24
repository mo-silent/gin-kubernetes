package router

import (
	"gitee.com/MoGD/gin-kubernetes/router/k8s"
	"gitee.com/MoGD/gin-kubernetes/router/system"
)

// RouterInterface router unified entrance
type RouterInterface interface {
	K8SRouters() k8s.K8SRouterGroupInterface
	SystemRouters() system.SystemRouterGroupInterface
}

type RouterGroup struct {
	K8SRouter    *k8s.K8SRouter
	SystemRouter *system.SystemRouter
}

// K8SRouter return k8s router enter instance
func (r *RouterGroup) K8SRouters() k8s.K8SRouterGroupInterface {
	return r.K8SRouter
}

// SystemRouter return system router enter instance
func (r *RouterGroup) SystemRouters() system.SystemRouterGroupInterface {
	return r.SystemRouter
}

// RouterGroupEnter return unified router instance
func RouterGroupEnter() *RouterGroup {
	return new(RouterGroup)
}
