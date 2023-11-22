package service

import (
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf_example/common/pb"
	"golang.org/x/net/context"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/11/22
//***************************************************

type IAService interface {
	Buy(ctx context.Context, args *rpc.Args[*pb.DefaultRequest], reply *rpc.Reply[*pb.DefaultRequest]) (err error)
}

type IARPCService interface {
}
