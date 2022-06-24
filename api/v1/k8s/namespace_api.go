package k8s

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/util/retry"
)

// NamespaceGetter getter namespace
type NamespaceGetter interface {
	Namespace() K8SCommonInterface
}

// NamespaceApi namespace api enter
type NamespaceApi struct{}

// newNamespace return NamespaceApi
func newNamespace() *NamespaceApi {
	return &NamespaceApi{}
}

// Create
// @Tags Namespace
// @Summary 创建 Namespace
// @Security ApiKeyAuth
// @Produce application/json
// @Param name query string  true "Namespace name"
// @Success 200 {object} response.CommonResponse
// @Router /namespace/create [post]
func (deploy *NamespaceApi) Create(c *gin.Context) {
	name := c.Query("name")
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err := global.K8SCLIENT.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: "create namespace fail!",
		})
		return
		// panic(err.Error())
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: fmt.Sprintf("Created namespace %v.\n", name),
	})
}

// Delete
// @Tags Namespace
// @Summary 删除单个 Namespace
// @Security ApiKeyAuth
// @Produce application/json
// @Param   name    query  string  true "namespace 名称"
// @Success 200 {object} response.CommonResponse
// @Router /namespace/delete [delete]
func (deploy *NamespaceApi) Delete(c *gin.Context) {
	name := c.Query("name")

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	// err := global.DynamicK8SCLIENT.Resource(NamespaceRes).Namespace(namespace).Delete(context.TODO(), name, deleteOptions)
	if err := global.K8SCLIENT.CoreV1().Namespaces().Delete(context.TODO(), name, deleteOptions); err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: fmt.Sprintf("delete namespace %v fail!", name),
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: fmt.Sprintf("delete namespace %v success", name),
	})

}

// Update nothing to do
func (deploy *NamespaceApi) Update(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.CommonResponse{
		Msg: "Forbidden API",
	})
	// 获取更新信息
	// oldName := c.Query("oldname")
	// newName := c.Query("newname")

	// namespacesClient := global.K8SCLIENT.CoreV1().Namespaces()
	// retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	// 	// Retrieve the latest version of Namespace before attempting update
	// 	// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
	// 	result, getErr := namespacesClient.Get(context.TODO(), oldName, metav1.GetOptions{})
	// 	if getErr != nil {
	// 		panic(fmt.Errorf("failed to get latest version of Namespace: %v", getErr))
	// 	}

	// 	result.Name = newName
	// 	_, updateErr := namespacesClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	// 	return updateErr
	// })
	// if retryErr != nil {
	// 	c.JSON(http.StatusForbidden, response.CommonResponse{
	// 		Msg: fmt.Sprintf("update namespace %v to %v fail!", oldName, retryErr),
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, response.CommonResponse{
	// 	Msg: fmt.Sprintf("Updated namespace %v to %v...", oldName, newName),
	// })

}

// Get
// @Tags Namespace
// @Summary 获取单个 Namespace
// @Security ApiKeyAuth
// @Produce application/json
// @Param   name    query  string  true "namespace 名称"
// @Success 200 {object} response.CommonResponse
// @Router /namespace/get [get]
func (deploy *NamespaceApi) Get(c *gin.Context) {
	// 获取 namespace 名称
	name := c.Query("name")

	result, err := global.K8SCLIENT.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: fmt.Sprintf("get namespace %v fail!", name),
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: fmt.Sprintf("get namespace %v, value %v", name, result),
	})

}

// List
// @Tags Namespace
// @Summary 获取所有的 Namespace
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.CommonResponse
// @Router /namespace/list [get]
func (deploy *NamespaceApi) List(c *gin.Context) {
	result, err := global.K8SCLIENT.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusForbidden, response.CommonResponse{
			Msg: "list all namespace fail!",
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonResponse{
		Msg:  "list deploy succeed",
		Data: result,
	})

}
