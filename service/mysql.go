package service

import (
	"fmt"

	"gitee.com/MoGD/gin-kubernetes/model/system"
	adapter "github.com/casbin/gorm-adapter/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// InitData init system admin user and init casbin rule
func InitData(db *gorm.DB) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("gkadmin"), bcrypt.DefaultCost)
	test, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user := []system.User{
		{
			Username: "admin",
			Password: string(password),
			NickName: "超级管理员",
			Phone:    "070-523045",
			Email:    "moguande2018@gmail.com",
		},
		{
			Username: "test",
			Password: string(test),
			NickName: "测试用户",
			Phone:    "070-523045",
			Email:    "moguande2018@gmail.com",
		},
	}
	if err := db.Table("users").Create(&user).Error; err != nil {
		fmt.Printf("user 表初始化超级管理员失败\n error: %v\n", err)
		return err
	}

	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "test", V1: "/user/get", V2: "GET"},
		{Ptype: "p", V0: "test", V1: "/pod/get", V2: "GET"},
		{Ptype: "p", V0: "test", V1: "/namespace/get", V2: "GET"},
		{Ptype: "p", V0: "test", V1: "/deployment/get", V2: "GET"},
	}
	if err := db.Create(&entities).Error; err != nil {
		fmt.Printf("Casbin 表初始化失败\n error: %v\n", err)
		return err
	}
	return nil
}
