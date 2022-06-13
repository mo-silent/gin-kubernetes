package v1

import (
	"gitee.com/MoGD/gin-kubernetes/api/v1/k8s"
)

type UnifiedApiGroup struct {
	K8SApiGroup k8s.K8SApiGroup
}

var UnifiedApiGroupEnter = new(UnifiedApiGroup)
