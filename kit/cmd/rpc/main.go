package main

import (
	"github.com/thkhxm/tgf/util"
	"github.com/thkhxm/tgf_example/common/service"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/8/31
//***************************************************

func main() {
	util.SetAutoGenerateAPICodePath("../common/api")
	util.SetGenerateFileNameSuffix("rpc")
	//util.SetAutoGenerateAPICSCode("E:\\unity\\project\\t2\\Assets\\HotFix\\Code", "HotFix.Code")
	util.GeneratorRPC[service.IBRPCService]("b", "v1.0", "bservice", "b")
	//util.GeneratorRPC[service.IUserRPCService](user.ModuleName, user.Version, "userservice", "user")
	//util.GeneratorRPC[service.IFightRPCService](fight.ModuleName, fight.Version, "fightservice", "fight")
	//util.GeneratorRPC[service.IMatchRPCService](match.ModuleName, match.Version, "matchservice", "match")
	//util.GeneratorRPC[service.IPropRPCService](prop.ModuleName, prop.Version, "propservice", "prop")
	//util.GeneratorRPC[service.IMapLevelRPCService](maplevel.ModuleName, maplevel.Version, "maplevelservice", "maplevel")
	//util.GeneratorRPC[service.IDungeonRPCService](dungeon.ModuleName, dungeon.Version, "dungeonservice", "dungeon")
}
