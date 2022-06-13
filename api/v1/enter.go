package v1

type ApiGroup struct {
	PodApi
	DeploymentApi
}

var ApiGroupEnter = new(ApiGroup)
