# we-life
> gin框架we-life

参考文档 https://gin-gonic.com/docs/
## 开发及发布
### 工具
* golang(1.12+)

#### 如何开始
1. 下载并安装[Golang](https://golang.org/dl/)

#### 项目位置

$GOPATH/srcjaden/we-life

#### 启动模板前的准备

拉取 common-lib 代码。

#### 运行项目
```
本地环境(默认读取local.json)
go run main.go
```

#### 发布项目
```
make deploy tag=v1.0.1 env={dev,stage,prod}
```

#### 目录结构
```
├─business                    -- 外部服务
│  └─business-example
├─common                      -- 公共
├─conf                        -- 配置
├─controller                  -- 控制器
├─deploy                      -- 部署文件
│  └─k8s 
│      ├─dev
│      ├─prod
│      └─stage
├─domain                      -- 内部服务
│  └─example 
├─errcode                     -- 错误码
├─initialize                  -- 初始化
├─logs                        -- 日志文件
├─middleware                  -- 中间件
├─router                      -- 路由
├─test                        -- 测试
└─utils                       -- 工具包
```



