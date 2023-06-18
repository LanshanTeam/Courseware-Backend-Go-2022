package logic

import (
	"context"

	"lanshan/Courseware-Backend-Go-2022/class15/demo/rpc/internal/svc"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayWorldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayWorldLogic {
	return &SayWorldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayWorldLogic) SayWorld(in *pb.SayWorldReq) (*pb.SayWorldResp, error) {
	// todo: add your logic here and delete this line
	if in.Word == "hello" {
		return &pb.SayWorldResp{Word: "hello word"}, nil
	}
	return &pb.SayWorldResp{Word: "it's nothing"}, nil
}
