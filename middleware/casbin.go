package middleware

import (
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/service"
	"github.com/gin-gonic/gin"
)

var casbinService = new(service.CasbinService)

// CasbinHandler casbin middleware to determine permissions
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := "test"
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			// response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.JSON(http.StatusConflict, "权限不足")
			c.Abort()
			return
		}
	}
}
