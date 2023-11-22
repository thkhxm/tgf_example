package service

import (
	bmodels "github.com/thkhxm/tgf_example/common/rpc/b"
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

type IBService interface {
}

type IBRPCService interface {
	BuyItem(ctx context.Context, args *bmodels.BuyParam, reply *bmodels.ItemData) (err error)
}
