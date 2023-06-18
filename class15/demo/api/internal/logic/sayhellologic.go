package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/svc"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/types"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/rpc/pb"
)

type SayHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SayHelloLogic) SayHello(req *types.SayHelloReq) (resp *types.SayHelloResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.DemoRpcClient.SayWorld(l.ctx, &pb.SayWorldReq{Word: req.Word})
	if err != nil {
		return &types.SayHelloResp{Word: "not word"}, err
	}

	return &types.SayHelloResp{Word: res.Word}, err
}
