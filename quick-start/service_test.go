package quick_start_test

import (
	"github.com/thkhxm/tgf/robot"
	"github.com/thkhxm/tgf_example/common/pb"
	quick_start "github.com/thkhxm/tgf_example/quick-start"
	"google.golang.org/protobuf/proto"
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

func TestStartup(t *testing.T) {
	//启动服务
	quick_start.Startup()
}

func TestHallService_SayHello(t *testing.T) {
	//创建一个websocket的机器人
	client := robot.NewRobotWs("/example")
	//连接到指定路径
	client.Connect("127.0.0.1:8032")
	//注册消息响应处理
	client.RegisterCallbackMessage("hall.SayHello", func(iRobot robot.IRobot, bytes []byte) {
		result := &pb.DefaultRequest{}
		err := proto.Unmarshal(bytes, result)
		if err != nil {
			t.Logf("proto unmarshal error %v", err)
			return
		}
		t.Logf("say hello call back , he say : %s", result.Val)
	})
	//发送消息
	client.SendMessage("hall", "SayHello", &pb.DefaultRequest{Val: "tim"})
	//挂起等待
	select {}
}
