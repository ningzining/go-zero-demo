package logic

import (
	"context"
	"github.com/google/uuid"
	"strings"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQrCodeLogic {
	return &GetQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQrCodeLogic) GetQrCode(req *types.GetQrCodeReq) (resp *types.GetQrCodeResp, err error) {
	// todo: add your logic here and delete this line
	u := uuid.New().String()
	u = strings.ReplaceAll(u, "-", "")
	u = req.DeviceId + ":" + u
	resp = &types.GetQrCodeResp{
		QrCode: u,
	}
	return
}
