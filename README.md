> Author mogd 2022-06-10
> \
> Update mogd 2022-07-09

# gin-kubernetes

Go 语言的 gin 框架开发的一个 K8S CMDB 管理平台

> 个人项目，并非生产环境，没有 kubesphere 和 Rancher 等开源平台好
> /
> 文档后续会详细补充，当前还是开发状态

## 项目结构
```
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
├── middleware
├── model
│   ├── request
│   └── response
├── router
├── service
└── utils
    ├── timer
    └── upload
```

## issue

### initdata api

1. 数据库初始化会重写，已经初始化的会报错，没加判断
2. 数据初始化接口传入的参数并不会重写到 config.yaml 中，还没完善

> 计划一：重写初始化方式，通过 context 来实现
> \
> 计划二：添加 viper 修改数据库配置信息函数

### deployment create api

1. deployment 的结构体还没完善，接口传参其实是无效的，api 内写死了 deployment 的定义，创建的是 nginx deployment

> 计划一：修正 pod 的定义
> \
> 计划二：添加 volume 定义
> \
> 计划三：修正 deployment 定义