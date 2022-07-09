package system

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type User struct{}

type UserGetter interface {
	User() RouterSystemInterface
}

func newUsers() *User {
	return &User{}
}

func (u *User) InitRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("user")
	userApi := v1.ApiV1Enter().ApiV1System().User()
	{
		// userGroup.POST("login", userApi.Login)     // 用户登录
		// userGroup.POST("captcha", userApi.Captcha) // 用户验证码
		userGroup.POST("create", userApi.Create)   // 创建 user
		userGroup.DELETE("delete", userApi.Delete) // 删除 user
		userGroup.PUT("update", userApi.Update)    // 更新 user
		userGroup.GET("get", userApi.Get)          // 获取 user 信息
		userGroup.GET("list", userApi.List)        // 获取所有 user 信息
	}
}
