package request

type DeployUpdateMessage struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	ReplicasNumber int32  `json:"replicasNumber"`
	NewImage       string `json:"newImage"`
}
