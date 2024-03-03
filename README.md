

# go env
```cmd
C:\Users\admin>go version
go version go1.17.11 windows/amd64

set GOENV=D:\projects\gopath\goenv.conf
go env -w GO111MODULE=on
go env -w GOCACHE=D:\projects\gopath\cache
go env -w GOPATH=D:\projects\gopath
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOROOT="D:\Program Files\Go"
go env -w GOMODCACHE=D:\projects\gopath\pkg\mod


GOENV='/root/.config/go/env'
GOCACHE='/root/.cache/go-build'
GOMODCACHE='/root/go/pkg/mod'
GOPATH='/root/go'
GOROOT='/opt/go'

```

```cmd
.
├── application.yaml
├── application.yaml.bak
├── config
│   └── config.go
├── controller
│   └── getMeAllFields.go
├── databases
│   └── mysql.go
├── go.mod
├── go.sum
├── log
│   └── log.go
├── main.go
├── model
│   └── MarketEmotionSQL.go
├── README.md
└── router
    └── router.go
    
    
databases下为数据库初始化链接
model: 数据库操作
controller: 业务逻辑处理

```

# 使用依赖注入（dependency injection）
