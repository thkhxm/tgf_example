package a

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	bservice "github.com/thkhxm/tgf_example/common/api/b"
	"github.com/thkhxm/tgf_example/common/pb"
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

// Service
// @Description: implements define.IAService , define.IARPCService
type Service struct {
	rpc.Module
}

func (s *Service) Buy(ctx context.Context, args *rpc.Args[*pb.DefaultRequest], reply *rpc.Reply[*pb.DefaultRequest]) (err error) {
	res, e := rpc.SendRPCMessage(ctx, bservice.BuyItem.New(&bmodels.BuyParam{
		PropId: "1001",
		Count:  3,
	}, &bmodels.ItemData{}))
	if e != nil {
		log.Error("rpc error %v", e)
		return e
	}
	log.InfoTag("example", "Buy Item res: %v", res)
	return
}

func (s *Service) GetName() string {
	return "a"
}

func (s *Service) GetVersion() string {
	return "v1.0"
}

func (s *Service) Startup() (bool, error) {
	return true, nil
}

func Startup() {
	c := rpc.NewRPCServer(). //创建一个rpc服务
					WithGatewayWS("8032", "/example"). //启动一个网关
					WithService(&Service{}).           //启动一个service的服务
					WithWhiteService("Buy").           //添加该rpc到白名单,无需登录即可访问
					WithCache(tgf.CacheModuleClose).
					WithRandomServicePort(8000, 8010).
					Run()
	select {
	case <-c:
		log.InfoTag("service", "service is down ")

	}
}
