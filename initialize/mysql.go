package initialize

import (
	"fmt"

	"gitee.com/MoGD/gin-kubernetes/global"
	"gitee.com/MoGD/gin-kubernetes/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB init database connection
func InitDB() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.DBname == "" {
		fmt.Println("db name is null, please input db name")
		return nil
	}
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect db %v error\n", m.Dsn())
		return nil
	}
	return db
}

// RegisterTable 注册数据表
func RegisterTable(db *gorm.DB) {
	err := db.AutoMigrate(&system.User{})
	if err != nil {
		fmt.Printf("error: %v, register table failed\n", err)
	}
}
