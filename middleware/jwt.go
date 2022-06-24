package middleware

import (
	"errors"
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"gitee.com/MoGD/gin-kubernetes/service"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.CommonResponse{
				Msg: "未登录或非法访问",
			})
			c.Abort()
			return
		}

		j := service.NewClaims()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		TokenExpired := "Token is expired"
		if err != nil {
			if err == errors.New(TokenExpired) {
				// response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.JSON(http.StatusBadRequest, response.CommonResponse{
					Msg: "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, response.CommonResponse{
				Data: err.Error(),
				Msg:  "请求错误",
			})
			// response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
