package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"gitee.com/MoGD/gin-kubernetes/model/common/request"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodApi struct{}

// @Tags Pod
// @Summary 创建 Pod
// @Produce application/json
// @Param data body request.PodReques true "Pod simple configuration"
// @Success 200 {object} response.CommonResponse
// @Router /pod/createPod [post]
func (p *PodApi) CreatePod(c *gin.Context) {
	var podReq request.PodReques
	_ = c.ShouldBindJSON(&podReq)
	// fmt.Println(podReq)
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   podReq.PodName,
			Labels: podReq.Labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  podReq.ContainerName,
					Image: podReq.Image,
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: podReq.ContainerPort,
						},
					},
				},
			},
			RestartPolicy: "Always",
		},
	}
	// 获取 pod 接口
	podClient := global.K8SCLIENT.CoreV1().Pods(podReq.Namespace)
	// 创建 Pod
	_, err := podClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "create pod fail!",
		})
		// panic(err.Error())
	}
	// 循环获取 pod 状态，检查为 Running 状态后，返回 pod 信息
	// for {
	// 	podStatus, _ := podClient.Get(context.TODO(), "demo-pod", metav1.GetOptions{})
	// 	if podStatus.Status.Phase == "Running" {
	// 		c.JSON(http.StatusOK, response.CommonResponse{
	// 			Message: podStatus,
	// 		})
	// 		break
	// 	}
	// }
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: "create pod interface call success! Please watch pod status until the status is running.",
	})

}

// @Tags Pod
// @Summary 删除单个 Pod
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "pod名称"
// @Success 200 {object} response.CommonResponse
// @Router /pod/deletePod [post]
func (p *PodApi) DeletePod(c *gin.Context) {
	// 获取命名空间和 pod 名称
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")
	// 获取 pod 接口
	podClient := global.K8SCLIENT.CoreV1().Pods(namespace)
	// 删除 Pod
	err := podClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "delete pod fail!",
		})
		// panic(err.Error())
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("delete pod %v success", name),
	})

}

// @Tags Pod
// @Summary 获取单个 Pod 信息
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "pod名称"
// @Success 200 {object} response.CommonResponse
// @Router /pod/getPod [post]
func (p *PodApi) GetPod(c *gin.Context) {
	// get namespace
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")
	fmt.Println(namespace, name)
	if namespace == "" || strings.TrimSpace(namespace) == "" {
		namespace = "default"
	}
	// list one pod
	pods, err := global.K8SCLIENT.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "get pod fail!",
		})
		// panic(err.Error())
	}
	//fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: pods,
	})
}

// @Tags Pod
// @Summary 获取单个命名空间中所有的  Pod 信息
// @Produce application/json
// @Param   namespace    path  string  false "命名空间" default(default)
// @Success 200 {object} response.CommonResponse
// @Router /pod/listNamespacePod/{namespace} [get]
func (p *PodApi) ListNamespacePod(c *gin.Context) {
	// get namespace
	namespace := c.Param("namespace")
	if namespace == "{namespace}" || namespace == "" ||
		strings.TrimSpace(namespace) == "" {
		namespace = "default"
	}
	// list pod
	pods, err := global.K8SCLIENT.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "list namespace: " + namespace + " pod fail!",
		})
		// panic(err.Error())
	}
	//fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: pods.Items,
	})
}

// @Tags Pod
// @Summary 获取所有 Pod 信息
// @Produce application/json
// @Success 200 {object} response.CommonResponse
// @Router /pod/listAllPod [get]
func (p *PodApi) ListAllPod(c *gin.Context) {
	// list pod
	pods, err := global.K8SCLIENT.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "list all pod fail!",
		})
		// panic(err.Error())
	}
	//fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: pods.Items,
	})
}
