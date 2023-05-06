// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package demo

import (
	"context"

	"lanshan/Courseware-Backend-Go-2022/class15/demo/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SayWorldReq  = pb.SayWorldReq
	SayWorldResp = pb.SayWorldResp

	Demo interface {
		SayWorld(ctx context.Context, in *SayWorldReq, opts ...grpc.CallOption) (*SayWorldResp, error)
	}

	defaultDemo struct {
		cli zrpc.Client
	}
)

func NewDemo(cli zrpc.Client) Demo {
	return &defaultDemo{
		cli: cli,
	}
}

func (m *defaultDemo) SayWorld(ctx context.Context, in *SayWorldReq, opts ...grpc.CallOption) (*SayWorldResp, error) {
	client := pb.NewDemoClient(m.cli.Conn())
	return client.SayWorld(ctx, in, opts...)
}
