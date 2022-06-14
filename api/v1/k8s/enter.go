package k8s

type K8SInterface interface {
	PodGetter
	DeploymentGetter
}

type ApiV1K8SEnter struct {
}

// Depolyment return a DeploymentInterface
func (g *ApiV1K8SEnter) Deployment() DeploymentInterface {
	return newDeployments()
}

// Pod return a PodInterface
func (g *ApiV1K8SEnter) Pod() PodInterface {
	return newPods()
}
