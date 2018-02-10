package main

import (
	"context"
	"melon_micro/proto/pusher"
	"testing"

	"github.com/micro/go-micro/client"
)

var serviceName = "pusher"

// Create new client
var microClient = pusher.NewPusherClient(serviceName, client.NewClient())
var wserverList = []*pusher.RegisterRequest{
	&pusher.RegisterRequest{
		Ip:       "127.0.0.1",
		Host:     "localhost",
		Port:     ":6001",
		ApiKey:   "",
		ChatRoom: "game",
	},
	&pusher.RegisterRequest{
		Ip:       "127.0.0.1",
		Host:     "localhost",
		Port:     ":6002",
		ApiKey:   "",
		ChatRoom: "trade",
	},
	&pusher.RegisterRequest{
		Ip:       "127.0.0.1",
		Host:     "localhost",
		Port:     ":2015",
		ApiKey:   "",
		ChatRoom: "member",
	},
}
var IDList = []uint64{}

// TestPing 測試連線
func TestPing(t *testing.T) {
	// Ping the service
	res, err := microClient.Ping(context.TODO(), &pusher.PingRequest{})
	if err != nil {
		t.Error(err)
		return
	}

	if res.ServiceName != serviceName {
		t.Errorf("Service is not %s, response service is %s", serviceName, res.ServiceName)
		return
	}

	t.Log("OK! Service is running :", res.GetServiceName())
}

func TestRegister(t *testing.T) {
	emptyRequest := new(pusher.RegisterRequest)
	res, err := microClient.Register(context.TODO(), emptyRequest)
	if err != nil {
		t.Error(err)
		return
	}
	if res.ErrorCode == 0 {
		t.Error("空註冊不能成功！")
		return
	}

	for _, request := range wserverList {
		res, err := microClient.Register(context.TODO(), request)
		if err != nil {
			t.Error(err)
			return
		}

		if res.ErrorCode != 0 {
			t.Errorf("Service give an error : %d (%s)", res.ErrorCode, res.ErrorText)
			return
		}

		if res.Id == 0 {
			t.Error("Service doesn't give an ID!")
			return
		}

		IDList = append(IDList, res.Id)

		t.Log("OK! Service give a register ID :", res.Id)
	}
}

func TestRegisterList(t *testing.T) {
	res, err := microClient.RegisterList(context.TODO(), &pusher.RegisterListRequest{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(res.RegisterList) == 0 {
		t.Error("Service give an empty list")
		return
	}

	if len(res.RegisterList) < len(wserverList) {
		t.Error("Service's length of register list should greater than 3 :", res.RegisterList)
		return
	}

	t.Log("OK! Service give a register list :", res.RegisterList)

	res, err = microClient.RegisterList(context.TODO(), &pusher.RegisterListRequest{
		ChatRoom: "game",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if len(res.RegisterList) != 1 {
		t.Error("Service's register list with chat room (game) should be one :", res.RegisterList)
		return
	}

	t.Log("OK! Service give a register list with chat room (game) :", res.RegisterList)
}

func TestChatRoomList(t *testing.T) {
	res, err := microClient.ChatRoomList(context.TODO(), &pusher.ChatRoomListRequest{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(res.ChatRoomList) == 0 {
		t.Error("Service give an empty list")
		return
	}

	t.Log("OK! Service give a channel list :", res.ChatRoomList)

	TestID := IDList[0]
	res, err = microClient.ChatRoomList(context.TODO(), &pusher.ChatRoomListRequest{
		Id: TestID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if len(res.ChatRoomList) != 1 {
		t.Errorf("Service's chat room list (ID:%d) should be one: %v", TestID, res.ChatRoomList)
		return
	}

	t.Logf("OK! Service give a chat room list (ID:%d) : %v", TestID, res.ChatRoomList)
}

func TestPush(t *testing.T) {
	res, err := microClient.Push(context.TODO(), &pusher.PushRequest{
		ChatRoom: "game",
		Message:  "unit test",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if res.ErrorCode != 0 {
		t.Errorf("Service give an error : %d (%s)", res.ErrorCode, res.ErrorText)
		return
	}

	if res.ErrorText == "" {
		t.Log("OK! Service push ok")
	} else {
		t.Log("OK! Service push ok with text", res.ErrorText)
	}
}

func TestDeregister(t *testing.T) {
	for _, id := range IDList {
		res, err := microClient.Deregister(context.TODO(), &pusher.DeregisterRequest{Id: id})
		if err != nil {
			t.Error(err)
			return
		}

		if res.ErrorCode != 0 {
			t.Errorf("Service give an error : %d (%s)", res.ErrorCode, res.ErrorText)
			return
		}

		t.Log("OK! Service deregister ID:", id)
	}
}
