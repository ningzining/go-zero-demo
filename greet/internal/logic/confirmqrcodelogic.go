package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"go-zero-demo/greet/internal/global"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmQrCodeLogic {
	return &ConfirmQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmQrCodeLogic) ConfirmQrCode(req *types.ConfirmQrCodeReq) (resp *types.ConfirmQrCodeResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := global.CodeMap[req.QrCode]
	if !ok {
		fmt.Printf("token error")
		return nil, errors.New("token error")
	}
	conn, ok := global.WsMap[req.QrCode]
	if !ok {
		fmt.Printf("connect not exist")
		return nil, errors.New("connect not exist")
	}
	codeResp := &types.QrCodeResp{
		State: types.QRCODE_STATE_LOGIN,
		User: types.UserInfo{
			UserId:   userId,
			Username: "default",
		},
	}
	bytes, err := json.Marshal(codeResp)
	_ = conn.WriteMessage(websocket.TextMessage, bytes)
	resp = &types.ConfirmQrCodeResp{}
	return
}
