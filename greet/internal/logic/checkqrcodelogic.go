package logic

import (
	"context"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckQrCodeLogic {
	return &CheckQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckQrCodeLogic) CheckQrCode(req *types.CheckQrCodeReq) (resp *types.CheckQrCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
