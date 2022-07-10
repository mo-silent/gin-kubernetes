package global

import (
	"time"

	"gitee.com/MoGD/gin-kubernetes/conf"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

var (
	K8sClint *kubernetes.Clientset
	CONFIG   conf.Server
	VP       *viper.Viper
	DB       *gorm.DB
)

type MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
