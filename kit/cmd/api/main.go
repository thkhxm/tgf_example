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
//2023/5/29
//***************************************************

func main() {
	util.SetAutoGenerateAPICodePath("../common/api")
	//util.SetAutoGenerateAPICSCode("E:\\unity\\project\\t2\\Assets\\HotFix\\Code", "HotFix.Code")
	util.GeneratorAPI[service.IAService]("a", "v1.0", "api")
	//util.GeneratorAPI[service.IMatchService](match.ModuleName, match.Version, "api")
	//util.GeneratorAPI[service.IMapLevelService](maplevel.ModuleName, maplevel.Version, "api")
	//util.GeneratorAPI[service.IPropService](prop.ModuleName, prop.Version, "api",
	//	"UpdatePropItemPush")
	//util.GeneratorAPI[service.IFightService](fight.ModuleName, fight.Version, "api",
	//	"BattleMainInfoPush", "BattleAttackResultPush")
	//util.GeneratorAPI[service.IDungeonService](dungeon.ModuleName, dungeon.Version, "api",
	//	"DungeonMainPush")

	util.GenerateCSApiService()
}
