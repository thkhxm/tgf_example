package quick_start

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/log"
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
//2023/11/21
//***************************************************

func Startup() {
	c := rpc.NewRPCServer(). //创建一个rpc服务
					WithGatewayWS("8032", "/example", nil). //启动一个网关
					WithService(&HallService{}).            //启动一个service的服务
					WithWhiteService("SayHello").           //添加该rpc到白名单,无需登录即可访问
					WithCache(tgf.CacheModuleClose).
					Run()
	select {
	case <-c:
		log.InfoTag("service", "service is down ")

	}
}

var hallservice_check rpc.IService = &HallService{}

// HallService
// @Description: implements rpc.IService
type HallService struct {
	rpc.Module
}

func (h *HallService) SayHello(ctx context.Context, args *rpc.Args[*pb.DefaultRequest], reply *rpc.Reply[*pb.DefaultRequest]) (err error) {
	log.InfoTag("debug", "%s say hello ", args.GetData().Val)
	reply.SetData(&pb.DefaultRequest{Val: "good luck"})
	return
}

func (h *HallService) GetName() string {
	return "hall"
}

func (h *HallService) GetVersion() string {
	return "v1.0"
}

func (h *HallService) Startup() (bool, error) {
	return true, nil
}
