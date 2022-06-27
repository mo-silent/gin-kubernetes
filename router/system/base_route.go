package system

import (
	v1 "gitee.com/MoGD/gin-kubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type Base struct{}

type BaseGetter interface {
	Base() SystemRouterInterface
}

func newBases() *Base {
	return &Base{}
}

func (b *Base) InitRouter(Router *gin.RouterGroup) {
	baseGroup := Router.Group("base")
	baseApi := v1.ApiV1Enter().ApiV1System().Base()
	{
		baseGroup.POST("login", baseApi.Login)     // 用户登录
		baseGroup.POST("captcha", baseApi.Captcha) // 用户验证码
		baseGroup.POST("initdata", baseApi.InitData)
	}
}
