package router

import "gitee.com/MoGD/gin-kubernetes/router/k8s"

type UnifiedRouterGroup struct {
	KubeRouterGroup k8s.KubeRouterGroup
}

var UnifiedRouterGroupEnter = new(UnifiedRouterGroup)
