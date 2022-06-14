package request

type PodUpdateMessage struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	NewImage  string `json:"newImage"`
}
