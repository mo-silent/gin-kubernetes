package system

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/common/response"
	"gitee.com/MoGD/gin-kubernetes/model/system"
	"gitee.com/MoGD/gin-kubernetes/model/system/request"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ApiV1User user api v1 struct
type ApiV1User struct{}

// UserGetter user api enter
type UserGetter interface {
	User() ApiV1SystemInterface
}

// newUsers return user api struct
func newUsers() *ApiV1User {
	return &ApiV1User{}
}

// Create user create and register api
// @Tags User
// @Summary 注册用户
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.User true "用户名，密码，昵称，手机和邮箱"
// @Success 200 {object} response.CommonResponse
// @Router /user/create [post]
func (u *ApiV1User) Create(c *gin.Context) {
	var user system.User
	_ = c.ShouldBindJSON(&user)

	if !errors.Is(global.DB.Unscoped().Where("username = ?", user.Username).First(&system.User{}).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		c.JSON(http.StatusConflict, response.CommonResponse{
			Msg: "用户已存在或已软删除",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	if t, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z"); user.CreatedAt == t {
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
	}

	if err := global.DB.Table("users").Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "用户创建失败",
		})
		fmt.Printf("creat user faild, error: %v\n", err)
		return
	}
	c.JSON(http.StatusCreated, response.CommonResponse{
		Msg: "用户创建成功",
	})

}

// Delete user delete api
// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.GetName true "用户名，密码，昵称，手机和邮箱"
// @Success 200 {object} response.CommonResponse
// @Router /user/delete [delete]
func (u *ApiV1User) Delete(c *gin.Context) {
	var getName request.GetName
	_ = c.ShouldBindJSON(&getName)

	if errors.Is(global.DB.Unscoped().Where("username = ?", getName.Username).First(&system.User{}).Error, gorm.ErrRecordNotFound) { // 判断用户是否存在
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "用户不存在",
		})
		return
	}
	if getName.Forced {
		if global.DB.Unscoped().Where("username = ?", getName.Username).Delete(&system.User{}).Error != nil {
			c.JSON(http.StatusBadGateway, response.CommonResponse{
				Msg: "删除数据库记录失败, 请检查接口",
			})
			return
		}
		c.JSON(http.StatusCreated, response.CommonResponse{
			Msg: "用户删除成功",
		})
		return
	}
	// determine whether the user is soft-deleted
	if errors.Is(global.DB.Where("username = ?", getName.Username).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "用户已软删除, 如果需要删除数据库记录, 请加上 forced 参数",
		})
		return
	}

	global.DB.Where("username = ?", getName.Username).Delete(&system.User{})

	c.JSON(http.StatusCreated, response.CommonResponse{
		Msg: "用户删除成功",
	})
}

// Update user update password api
// @Tags User
// @Summary 用户更新密码
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.NewPassword true "用户名，原密码，新密码"
// @Success 200 {object} response.CommonResponse
// @Router /user/update [put]
func (u *ApiV1User) Update(c *gin.Context) {

	var newUser request.NewPassword
	_ = c.ShouldBindJSON(&newUser)

	var user system.User
	if errors.Is(global.DB.Where("username = ?", newUser.Username).First(&user).Error, gorm.ErrRecordNotFound) ||
		bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password)) != nil { // 判断用户是否存在和原密码
		c.JSON(http.StatusConflict, response.CommonResponse{
			Msg: "用户不存在或原密码错误",
		})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(newUser.NewPassword), bcrypt.DefaultCost)
	newUser.Password = string(password)
	// update user password
	if global.DB.Model(&system.User{}).Where("username = ?", newUser.Username).Update("password", newUser.NewPassword).Error != nil {
		c.JSON(http.StatusBadRequest, response.CommonResponse{
			Msg: "用户密码更新失败",
		})
		return
	}

	c.JSON(http.StatusCreated, response.CommonResponse{
		Msg: "用户密码更新成功",
	})
}

// Get user get api
// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param   username  query  string  true "用户名"
// @Success 200 {object} response.CommonResponse
// @Router /user/get [get]
func (u *ApiV1User) Get(c *gin.Context) {
	username := c.Query("username")
	user := map[string]interface{}{}
	if global.DB.Table("users").Where("username = ?", username).Find(&user).Error != nil {
		c.JSON(http.StatusNotFound, response.CommonResponse{
			Msg: "用户查询失败",
		})
		return
	}
	c.JSON(http.StatusCreated, response.CommonResponse{
		Data: user,
		Msg:  "用户信息获取成功",
	})
}

// List user list api
// @Tags User
// @Summary 列出所有用户
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.CommonResponse
// @Router /user/list [get]
func (u *ApiV1User) List(c *gin.Context) {
	var results []map[string]interface{}

	if global.DB.Table("users").Find(&results).Error != nil {
		c.JSON(http.StatusNotFound, response.CommonResponse{
			Msg: "用户查询失败",
		})
	}
	c.JSON(http.StatusCreated, response.CommonResponse{
		Data: results,
		Msg:  "所有用户获取成功",
	})
}
