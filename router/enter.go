package router

type RouterGroup struct {
	PodRouter
	DeploymentRouter
}

var RouterGroupEnter = new(RouterGroup)
