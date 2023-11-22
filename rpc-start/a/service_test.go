package a_test

import (
	"github.com/thkhxm/tgf/robot"
	"github.com/thkhxm/tgf_example/common/api"
	"github.com/thkhxm/tgf_example/common/pb"
	"github.com/thkhxm/tgf_example/rpc-start/a"
	"testing"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/11/22
//***************************************************

func TestService_Buy(t *testing.T) {
	//创建一个websocket的机器人
	client := robot.NewRobotWs("/example")
	//连接到指定路径
	client.Connect("127.0.0.1:8032")
	client.Send(api.Buy.MessageType, &pb.DefaultRequest{})
	select {}
}

func TestStartup(t *testing.T) {
	a.Startup()
}
