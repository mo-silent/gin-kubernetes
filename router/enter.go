package router

type KubeRouterGroup struct {
	PodRouter
	DeploymentRouter
}

var KubeRouterGroupEnter = new(KubeRouterGroup)
