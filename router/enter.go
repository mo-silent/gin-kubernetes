package router

import (
	"gitee.com/MoGD/gin-kubernetes/router/k8s"
	"gitee.com/MoGD/gin-kubernetes/router/system"
)

// UnifiedRouterInterface router unified entrance
type UnifiedRouterInterface interface {
	K8SRouters() k8s.RouterK8SGroupInterface
	SystemRouters() system.RouterSystemGroupInterface
}

type EnterRouterGroup struct {
	K8SRouter    *k8s.RouterK8S
	SystemRouter *system.RouterSystem
}

// K8SRouters K8SRouter return k8s router enter instance
func (r *EnterRouterGroup) K8SRouters() k8s.RouterK8SGroupInterface {
	return r.K8SRouter
}

// SystemRouters SystemRouter return system router enter instance
func (r *EnterRouterGroup) SystemRouters() system.RouterSystemGroupInterface {
	return r.SystemRouter
}

// EnterRouter return unified router instance
func EnterRouter() *EnterRouterGroup {
	return new(EnterRouterGroup)
}
