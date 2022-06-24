package k8s

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
	"k8s.io/client-go/util/retry"
)

// PodGetter getter interface
type PodGetter interface {
	Pod() K8SCommonInterface
}

// PodApi pod api enter
type PodApi struct{}

// newPods return a PodApi
func newPods() *PodApi {
	return &PodApi{}
}

// Create create pod
// @Tags Pod
// @Summary 创建 Pod
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.PodRequest true "Pod simple configuration"
// @Success 200 {object} response.CommonResponse
// @Router /pod/create [post]
func (p *PodApi) Create(c *gin.Context) {
	var podReq request.PodRequest
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
			Msg: "create pod fail!",
		})
		return
		// panic(err.Error())
	}
	// 循环获取 pod 状态，检查为 Running 状态后，返回 pod 信息
	// for {
	// 	podStatus, _ := podClient.Get(context.TODO(), "demo-pod", metav1.GetOptions{})
	// 	if podStatus.Status.Phase == "Running" {
	// 		c.JSON(http.StatusOK, response.CommonResponse{
	// 			Msg: podStatus,
	// 		})
	// 		break
	// 	}
	// }
	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: "create pod interface call success! Please watch pod status until the status is running.",
	})

}

// Delete
// @Tags Pod
// @Summary 删除单个 Pod
// @Security ApiKeyAuth
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "pod名称"
// @Success 200 {object} response.CommonResponse
// @Router /pod/delete [delete]
func (p *PodApi) Delete(c *gin.Context) {
	// 获取命名空间和 pod 名称
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")
	// 获取 pod 接口
	podClient := global.K8SCLIENT.CoreV1().Pods(namespace)
	// 删除 Pod
	err := podClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: "delete pod fail!",
		})
		return
		// panic(err.Error())
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: fmt.Sprintf("delete pod %v success", name),
	})

}

// Update
// @Tags Pod
// @Summary 更新 Pod
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.PodUpdateMessage true "Pod configuration information that needs to be changed"
// @Success 200 {object} response.CommonResponse
// @Router /pod/update [put]
func (p *PodApi) Update(c *gin.Context) {
	// 获取更新信息
	var updateMessage request.PodUpdateMessage
	_ = c.ShouldBindJSON(&updateMessage)

	podClient := global.K8SCLIENT.CoreV1().Pods(updateMessage.Namespace)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		pods, err := podClient.Get(context.TODO(), updateMessage.Name, metav1.GetOptions{})
		if err != nil {
			c.JSON(http.StatusForbidden, response.CommonResponse{
				Msg: "Update pod return when get pod fail!",
			})
			// panic(err.Error())
		}
		pods.Spec.Containers[0].Image = updateMessage.NewImage

		_, updateErr := podClient.Update(context.TODO(), pods, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: fmt.Sprintf("update namespace %v pod %v fail!", updateMessage.Namespace, updateMessage.Name),
		})
		return
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: fmt.Sprintf("Updated pod %v...", updateMessage.Name),
	})

}

// Get
// @Tags Pod
// @Summary 获取单个 Pod 信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "pod名称"
// @Success 200 {object} response.CommonResponse
// @Router /pod/get [get]
func (p *PodApi) Get(c *gin.Context) {
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
			Msg: "get pod fail!",
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg:  "get pod succeed",
		Data: pods,
	})
}

// List
// @Tags Pod
// @Summary 获取命名空间中所有的  Pod 信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Success 200 {object} response.CommonResponse
// @Router /pod/list [get]
func (p *PodApi) List(c *gin.Context) {
	// get namespace
	namespace := c.DefaultQuery("namespace", "default")
	if namespace == "{namespace}" || namespace == "" ||
		strings.TrimSpace(namespace) == "" {
		namespace = "default"
	}
	// list pod
	pods, err := global.K8SCLIENT.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: "list namespace: " + namespace + " pod fail!",
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg:  "list pod succeed",
		Data: pods.Items,
	})
}

// // ListAllPod
// // @Tags Pod
// // @Summary 获取所有 Pod 信息
// // @Produce application/json
// // @Success 200 {object} response.CommonResponse
// // @Router /pod/listAllPod [get]
// func (p *PodApi) ListAllPod(c *gin.Context) {
// 	// list pod
// 	pods, err := global.K8SCLIENT.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		c.JSON(http.StatusForbidden, response.CommonResponse{
// 			Msg: "list all pod fail!",
// 		})
// 		// panic(err.Error())
// 	}
// 	//fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
// 	c.JSON(http.StatusOK, response.CommonResponse{
// 		Msg: pods.Items,
// 	})
// }
