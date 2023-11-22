package b

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
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
// @Description: implements define.IBService , define.IBRPCService
type Service struct {
	rpc.Module
}

func (s *Service) BuyItem(ctx context.Context, args *bmodels.BuyParam, reply *bmodels.ItemData) (err error) {
	reply.PropId = args.PropId
	reply.Count = args.Count * 2
	reply.Name = "买一送一特价商品"
	log.DebugTag("example", "args:%v,reply:%v", args, reply)
	return
}

func (s *Service) GetName() string {
	return "b"
}

func (s *Service) GetVersion() string {
	return "v1.0"
}

func (s *Service) Startup() (bool, error) {
	return true, nil
}

func Startup() {
	c := rpc.NewRPCServer(). //创建一个rpc服务
					WithService(&Service{}). //启动一个service的服务
					WithCache(tgf.CacheModuleClose).
					WithRandomServicePort(8010, 8020).
					Run()
	select {
	case <-c:
		log.InfoTag("service", "service is down ")
	}
}
