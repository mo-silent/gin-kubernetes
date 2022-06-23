package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gitee.com/MoGD/gin-kubernetes/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [Mogd](https://gitee.com/MoGD)
func Viper(path ...string) *viper.Viper {
	var config, kubeconfig string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(configEnv); configEnv == "" { // 判断 configEnv 常量存储的环境变量是否为空
				config = configDefaultFile
				fmt.Printf("config的路径为%s\n", configDefaultFile)
			} else { // configEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("使用%s环境变量,config的路径为%s\n", configEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	if kubeconfig != "" {
		global.CONFIG.Kubeconfig = kubeconfig
		v.Set("kubeconfig", kubeconfig)
	}
	_ = v.WriteConfig()
	return v
}
