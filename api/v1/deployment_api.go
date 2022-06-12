package v1

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
)

type DeploymentApi struct{}

var DeploymentRes = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

// CreateDeployment
// @Tags Deployment
// @Summary 创建 Deployment
// @Produce application/json
// @Param data body request.DeploymentRequest true "Deployment simple configuration"
// @Success 200 {object} response.CommonResponse
// @Router /pod/createDeployment [post]
func (deploy *DeploymentApi) CreateDeployment(c *gin.Context) {
	// 获取 depolyment 信息
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

// DeleteDeployment
// @Tags Deployment
// @Summary 删除单个 Deployment
// @Produce application/json
// @Param   namespace  query  string  false "命名空间" default(default)
// @Param   name    query  string  true "deployment 名称"
// @Success 200 {object} response.CommonResponse
// @Router /pod/deleteDeployment [post]
func (deploy *DeploymentApi) DeleteDeployment(c *gin.Context) {
	// 获取命名空间和 pod 名称
	namespace := c.DefaultQuery("namespace", "default")
	name := c.Query("name")

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	if err := global.DynamicK8SCLIENT.Resource(DeploymentRes).Namespace(namespace).Delete(context.TODO(), name, deleteOptions); err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Message: fmt.Sprintf("delete namespace %v deployment %v fail!", namespace, name),
		})
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Message: fmt.Sprintf("delete deployment %v success", name),
	})

}
