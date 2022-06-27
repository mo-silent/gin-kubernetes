package system

import (
	"errors"
	"fmt"
	"net/http"

	"gitee.com/MoGD/gin-kubernetes/conf"
	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"gitee.com/MoGD/gin-kubernetes/model/system"
	"gitee.com/MoGD/gin-kubernetes/model/system/request"
	"gitee.com/MoGD/gin-kubernetes/service"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// BaseInterface base api enter, public api
type BaseInterface interface {
	Login(c *gin.Context)
	Captcha(c *gin.Context)
	InitData(c *gin.Context)
}

// ApiV1Base api v1 base struct
type ApiV1Base struct{}

// BaseGetter base api enter
type BaseGetter interface {
	Base() BaseInterface
}

// NewBases return new api v1 base struct
func newBases() *ApiV1Base {
	return &ApiV1Base{}
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// Login user login api
// @Tags Base
// @Summary 用户登录
// @Produce application/json
// @Param data body request.Login true "用户名，密码"
// @Success 200 {object} response.CommonResponse
// @Router /base/login [post]
func (b *ApiV1Base) Login(c *gin.Context) {
	var user request.Login
	_ = c.ShouldBindJSON(&user)

	if !store.Verify(user.CaptchaId, user.CaptchaValue, true) {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "验证码错误",
		})
		return
	}

	var userRes system.User

	if errors.Is(global.DB.Where("username = ?", user.Username).First(&userRes).Error, gorm.ErrRecordNotFound) ||
		bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(user.Password)) != nil {
		c.JSON(http.StatusNotFound, response.CommonResponse{
			Msg: "用户不存在或密码错误",
		})
		return
	}

	claims := service.NewClaims()
	token, err := claims.CreateToken(user.Username)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "token 获取失败",
		})
		return
	}
	data := struct {
		Token string
	}{
		token,
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Data: data,
		Msg:  "登录成功",
	})

}

// Captcha generate a verification code
// @Tags Base
// @Summary 生成验证码,返回包括随机数id,base64,验证码长度
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.CommonResponse
// @Router /base/captcha [post]
func (b *ApiV1Base) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	var driver = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "验证码获取失败",
		})
		return
	}
	data := struct {
		CaptchaId     string
		PicPath       string
		CaptchaLength int
	}{
		id,
		b64s,
		6,
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Data: data,
		Msg:  "验证码获取成功",
	})

}

// Captcha generate a verification code
// @Tags Base
// @Summary 生成验证码,返回包括随机数id,base64,验证码长度
// @accept application/json
// @Produce application/json
// @Param data body conf.Mysql true "初始化数据库参数"
// @Success 200 {object} response.CommonResponse
// @Router /base/initdata [post]
func (b *ApiV1Base) InitData(c *gin.Context) {
	// if global.DB != nil {
	// 	c.JSON(http.StatusOK, response.CommonResponse{
	// 		Msg: "已经存在数据库配置",
	// 	})
	// 	return
	// }
	var dbInfo conf.Mysql
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "参数校验不通过",
		})
		return
	}
	if err := service.InitData(global.DB); err != nil {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "初始化数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, response.CommonResponse{
		Msg: "数据初始化成功",
	})
}
