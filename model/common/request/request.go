package request

type PodRequest struct {
	Namespace     string            `json:"namespace"`
	PodName       string            `json:"podName"`
	ContainerName string            `json:"containerName"`
	Image         string            `json:"image"`
	Labels        map[string]string `json:"labels"`
	ContainerPort int32             `json:"containerPort"`
}

type DeploymentRequest struct {
	Namespace string                 `json:"namespace"`
	Object    map[string]interface{} `json:"object"`
}
