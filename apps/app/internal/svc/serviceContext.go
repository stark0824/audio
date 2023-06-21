package svc

import (
	"audio/apps/app/internal/config"
	"audio/apps/app/internal/middleware"
	"audio/apps/rpc/drama/dramaclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	DramaRpcClient dramaclient.Drama
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		DramaRpcClient: dramaclient.NewDrama(zrpc.MustNewClient(c.DramaRpcConf)),
	}
}
