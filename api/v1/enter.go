package v1

import (
	"gitee.com/MoGD/gin-kubernetes/api/v1/k8s"
	"gitee.com/MoGD/gin-kubernetes/api/v1/system"
)

// ApiV1Interface api v1 enter interface
type ApiV1Interface interface {
	ApiV1K8s() k8s.ApiV1K8sEnter
	ApiV1System() system.ApiV1SystemEnterInterface
}

// ApiV1Unified api v1 enter
type ApiV1Unified struct {
	ApiV1K8SEnter    *k8s.ApiV1K8SEnter
	ApiV1SystemEnter *system.ApiV1SystemEnter
}

// ApiV1K8S return k8s.ApiV1K8SEnter
func (u *ApiV1Unified) ApiV1K8S() k8s.ApiV1K8sEnter {
	return u.ApiV1K8SEnter
}

// ApiV1System return system.ApiV1SystemEnter
func (u *ApiV1Unified) ApiV1System() system.ApiV1SystemEnterInterface {
	return u.ApiV1SystemEnter
}

// ApiV1Enter init ApiV1Unified
func ApiV1Enter() *ApiV1Unified {
	return new(ApiV1Unified)
}
