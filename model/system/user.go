package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type User struct {
	global.GVA_MODEL
	Username string `json:"userName" gorm:"comment:用户登录名"`
	Password string `json:"password" gorm:"comment:登录密码"`
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Phone    string `json:"phone"  gorm:"comment:用户手机号"`
	Email    string `json:"email"  gorm:"comment:用户邮箱"`
}

func (User) TableName() string {
	return "users"
}
