package router

import "gitee.com/MoGD/gin-kubernetes/router/k8s"

// RouterInterface router unified entrance
type RouterInterface interface {
	K8SRouter() k8s.K8SRouterGroupInterface
}

type RouterGroup struct {
	K8SRouterGroup *k8s.K8SRouter
}

// K8SRouter return k8s router enter instance
func (r *RouterGroup) K8SRouter() k8s.K8SRouterGroupInterface {
	return r.K8SRouterGroup
}

// RouterGroupEnter return unified router instance
func RouterGroupEnter() *RouterGroup {
	return new(RouterGroup)
}
