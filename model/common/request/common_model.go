package request

// Containers
// kubernetes deployment containers variable
type Containers struct {
	Name   string
	Images string
	Ports  ContainerPorts
}

// ContainerPorts
// kubernetes container ports variable
type ContainerPorts struct {
	Ports []map[string]interface{}
}

type Metadata struct {
	Metadata map[string]interface{}
}
