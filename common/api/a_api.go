//Auto generated by tgf util
//created at 2023-11-22 11:25:16.3481113 +0800 CST m=+0.034760001

package api

import (
	"github.com/thkhxm/tgf/rpc"

	"github.com/thkhxm/tgf_example/common/pb"
)

var aService = &rpc.Module{Name: "a", Version: "v1.0"}

var (
	Buy = rpc.ServiceAPI[*pb.DefaultRequest, *pb.DefaultRequest]{
		ModuleName:  aService.Name,
		Name:        "Buy",
		MessageType: aService.Name + "." + "Buy",
	}
)

const ()
