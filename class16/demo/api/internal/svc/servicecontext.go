package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/config"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/rpc/demo"
)

type ServiceContext struct {
	Config        config.Config
	DemoRpcClient demo.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		DemoRpcClient: demo.NewDemo(zrpc.MustNewClient(c.DemoRpc)),
	}
}
