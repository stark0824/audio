# Audio - 有声微服务

### 版本说明

```
go version : 1.20.1
goctl version: 1.4.4
go 框架: go-zero
```



项目目录：

```
├── README.md
├── apps
│   ├── app
│   │   ├── admin.api
│   │   ├── admin.go
│   │   ├── etc
│   │   └── internal
│   ├── model
│   │   ├── drama
│   └── rpc
│       └── drama
├── common
│   └── errorx
│       └── baseerror.go
├── go.mod
└── go.sum
```

#### Api文件

1.编写api.api,使用goctl工具生成

```
goctl api go  -api api.api -dir ./ --style=goZero
```



2.生成rpc服务

```
goctl rpc protoc cp_drama_admin.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero 
```



3.rpc服务生成dockerfile

```
goctl docker -go cpDramaAdmin.go
```



4.测试rpc

```
go get github.com/fullstorydev/grpcui/cmd/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui
```

启动测试：

```
grpcui -plaintext 127.0.0.1:8899
```

cp_user_internal.proto



## 调用外部Rpc

```
protoc --go_out=cp --go-grpc_out=cp message/*.proto 
protoc --go_out=cp --go-grpc_out=cp cp_user_internal.proto
```



## 生成DockerFile

```
cd apps/app && goctl docker -go client.go
//回到项目根目录执行命令，生成镜像
docker build -t audio-api:v1 -f apps/app/Dockerfile  .

cd apps/rpc/drama && goctl docker -go cpDramaAdmin.go
//回到项目根目录执行命令，生成镜像
docker build -t drama-rpc:v1 -f apps/rpc/drama/Dockerfile .
```



## 项目目录

```
.
├── README.md
├── apps
│   ├── app
│   ├── model
│   └── rpc
├── common
│   ├── globalkey
│   ├── result
│   └── xerr
├── docker-compose.yaml
├── go.mod
├── go.sum
└── proto
    ├── pb
    └── user_auth
```



