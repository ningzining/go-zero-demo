package logic

import (
	"context"
	"fmt"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	var builder strings.Builder
	builder.WriteString("hello ")
	builder.WriteString(req.Name)
	msg := builder.String()
	fmt.Printf("%s\n", msg)
	resp = &types.Response{
		Message: msg,
	}
	return
}
