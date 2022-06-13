package request

type UpdateMessage struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	ReplicasNumber int32  `json:"replicasNumber"`
	NewImage       string `json:"newImage"`
}
