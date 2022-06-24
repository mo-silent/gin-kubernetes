package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type User struct {
	global.GVA_MODEL
	Username string `json:"username" gorm:"uniqueIndex:idx_username,length:255,comment:用户名"`
	Password string `json:"password" gorm:"comment:登录密码"`
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Phone    string `json:"phone"  gorm:"comment:用户手机号"`
	Email    string `json:"email"  gorm:"comment:用户邮箱"`
}

func (User) TableName() string {
	return "users"
}
