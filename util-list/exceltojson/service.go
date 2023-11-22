package exceltojson

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/component"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf_example/common/conf"
	"github.com/thkhxm/tgf_example/common/define/constant"
	errorcodes "github.com/thkhxm/tgf_example/common/define/errorcode"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/11/22
//***************************************************

type ExcelConfigService struct {
	rpc.Module
}

func (e *ExcelConfigService) GetName() string {
	return "config"
}

func (e *ExcelConfigService) GetVersion() string {
	return "v1.0"
}

func (e *ExcelConfigService) Startup() (bool, error) {
	// 获取单个配置
	component.GetGameConf[*conf.PropConf]("100001")
	// 获取多个id相同的配置
	component.GetGameConfBySlice[*conf.PropConf]("100002")
	// 修改配置文件之后,调用该函数可以重新刷新配置,做到配置在线热更
	component.InitGameConfToMem()
	//通过键值对配置,生成自定义的常量结构,可以用于错误码,常量Key等配置
	log.Info("print error code %d", errorcodes.SystemErr)
	cs, _ := component.GetGameConf[*conf.ConstantConf](constant.MapLevelInitId)
	log.Info("print constant key %s, value %s", cs.Key, cs.Value)
	//

	return true, nil
}

func Startup() {
	c := rpc.NewRPCServer().
		WithService(&ExcelConfigService{}).
		WithGameConfig("../common/conf/js").
		WithRandomServicePort(8040, 8050).
		WithCache(tgf.CacheModuleClose).
		Run()
	<-c
}
