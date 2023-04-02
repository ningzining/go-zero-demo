package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/greet/internal/global"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
)

type ScanQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScanQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanQrCodeLogic {
	return &ScanQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScanQrCodeLogic) ScanQrCode(req *types.ScanQrCodeReq) (resp *types.ScanQrCodeResp, err error) {
	// todo: add your logic here and delete this line
	global.CodeMap[req.QrCode] = req.UserId
	token := req.QrCode + ":" + req.UserId

	codeResp := types.QrCodeResp{
		State: types.QRCODE_STATE_SCAN,
	}
	conn, ok := global.WsMap[req.QrCode]
	if !ok {
		fmt.Printf("connect not exist")
		return nil, errors.New("connect not exist")
	}

	bytes, err := json.Marshal(codeResp)
	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		fmt.Printf("websocket write message error\n")
	}
	resp = &types.ScanQrCodeResp{
		Token: token,
	}
	return
}
