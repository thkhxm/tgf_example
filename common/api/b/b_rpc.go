//Auto generated by tgf util
//created at 2023-11-22 12:40:57.3561753 +0800 CST m=+0.033412601

package bservice

import (
	"github.com/thkhxm/tgf/rpc"

	"github.com/thkhxm/tgf_example/common/rpc/b"
)

var bService = &rpc.Module{Name: "b", Version: "v1.0"}

var (
	BuyItem = rpc.ServiceAPI[*bmodels.BuyParam, *bmodels.ItemData]{
		ModuleName:  bService.Name,
		Name:        "BuyItem",
		MessageType: bService.Name + "." + "BuyItem",
	}
)