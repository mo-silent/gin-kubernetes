package router

type K8SRouterGroup struct {
	PodRouter
	DeploymentRouter
}

var KubeRouterGroupEnter = new(K8SRouterGroup)
