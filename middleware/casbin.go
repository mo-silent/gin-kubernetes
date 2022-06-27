package middleware

import (
	"fmt"
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"gitee.com/MoGD/gin-kubernetes/service"
	"github.com/gin-gonic/gin"
)

// CasbinHandler casbin middleware to determine permissions
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("x-token")   // 获取 token
		claims := service.NewClaims()              // 获取 claims
		userLogin, err := claims.ParseToken(token) //解析 token

		if err != nil {
			fmt.Println("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		}
		casbin := service.NewCasbin()
		// waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户名
		sub := userLogin.Username
		e := casbin.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.CONFIG.System.Env == "develop" ||
			success || userLogin.Username == "admin" {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, response.CommonResponse{
				Msg: "权限不足",
			})
			c.Abort()
			return
		}
	}
}
