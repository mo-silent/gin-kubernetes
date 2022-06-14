package k8s

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/request"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/retry"
)

// DeploymentInterface deployment interface
type DeploymentInterface interface {
	K8SCommonInterface
}

// DeploymentGetter getter deployment
type DeploymentGetter interface {
	Deployment() DeploymentInterface
}

// DeploymentApi deploy api enter
type DeploymentApi struct{}

// newDeployments return DeploymentApi
func newDeployments() *DeploymentApi {
	return &DeploymentApi{}
}

// Create
// @Tags Deployment
// @Summary 创建 Deployment
// @Produce application/json
// @Param data body request.DeploymentRequest true "Deployment simple configuration"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/create [post]
func (deploy *DeploymentApi) Create(c *gin.Context) {
	// 获取 deployment 信息
	var deployReq request.DeploymentRequest
	_ = c.ShouldBindJSON(&deployReq)
	fmt.Println(deployReq)
	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "demo-deployment",
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name":          "http",
										"protocol":      "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	var DeploymentRes = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	result, err := global.DynamicK8SCLIENT.Resource(DeploymentRes).Namespace(deployReq.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: "Created deployment fail!",
		})
		// panic(err)
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("Created deployment %q.\n", result.GetName()),
	})
}

// Delete
// @Tags Deployment
// @Summary 删除单个 Deployment
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "deployment 名称"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/delete [delete]
func (deploy *DeploymentApi) Delete(c *gin.Context) {
	// 获取命名空间和 deployment 名称
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	// err := global.DynamicK8SCLIENT.Resource(DeploymentRes).Namespace(namespace).Delete(context.TODO(), name, deleteOptions)
	if err := global.K8SCLIENT.AppsV1().Deployments(namespace).Delete(context.TODO(), name, deleteOptions); err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: fmt.Sprintf("delete namespace %v deployment %v fail!", namespace, name),
		})
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("delete deployment %v success", name),
	})

}

// Update
// @Tags Deployment
// @Summary 更新 Deployment 的镜像版本和副本集
// @Produce application/json
// @Param data body request.DeployUpdateMessage true "Deployment configuration information that needs to be changed"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/update [put]
func (deploy *DeploymentApi) Update(c *gin.Context) {
	// 获取更新信息
	var updateMessage request.DeployUpdateMessage
	_ = c.ShouldBindJSON(&updateMessage)
	deploymentsClient := global.K8SCLIENT.AppsV1().Deployments(updateMessage.Namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(context.TODO(), updateMessage.Name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = &updateMessage.ReplicasNumber                   // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = updateMessage.NewImage // change nginx version
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: fmt.Sprintf("update namespace %v deployment %v fail!", updateMessage.Namespace, updateMessage.Name),
		})
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("Updated deployment %v...", updateMessage.Name),
	})

}

// Get
// @Tags Deployment
// @Summary 获取单个 Deployment
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "deployment 名称"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/get [get]
func (deploy *DeploymentApi) Get(c *gin.Context) {
	// 获取命名空间和 deployment 名称
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")

	result, err := global.K8SCLIENT.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: fmt.Sprintf("get namespace %v deployment %v fail!", namespace, name),
		})
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("get deployment %v, value %v", name, result),
	})

}

// List
// @Tags Deployment
// @Summary 获取命名空间下的所有 Deployment
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Success 200 {object} response.CommonResponse
// @Router /deployment/list [get]
func (deploy *DeploymentApi) List(c *gin.Context) {
	// 获取命名空间
	namespace := c.DefaultQuery("namespace", "default")

	result, err := global.K8SCLIENT.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: fmt.Sprintf("list namespace %v deployment fail!", namespace),
		})
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Message: result,
	})

}
