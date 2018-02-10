package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"melon_micro/proto/pusher"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// TransferData 運輸訊息
type TransferData struct {
	ChatRoom string `json:"chat_room"`
	Data     string `json:"data"`
}

// Pusher 即時推送
type Pusher struct {
	// RegisteringList 註冊列表
	registerList map[uint64]*pusher.RegisterRequest
	// mu 讀寫鎖
	mu *sync.RWMutex
}

// Ping 測試連線
func (p *Pusher) Ping(ctx context.Context, req *pusher.PingRequest, res *pusher.PongResponse) (err error) {
	res.ServiceName = "pusher"
	res.Environment = PusherConfig.App.Env
	return
}

// Register 註冊server
func (p *Pusher) Register(ctx context.Context, req *pusher.RegisterRequest, res *pusher.RegisterResponse) (err error) {
	if req.Host == "" || req.Port == "" || req.ChatRoom == "" {
		res.ErrorCode = 1
		res.ErrorText = "Host or Port or ChatRoom Required!"
		return
	}

	defer func() {
		if catchErr := recover(); catchErr != nil {
			err = catchErr.(error)
		}
		res.ErrorCode = 0
	}()

	res.ErrorCode = -1
	timer := time.NewTimer(time.Second * 3)
	isTimeout := false
	isNotOK := true
	var ID uint64
	go func() {
		<-timer.C
		isTimeout = true
	}()
	p.mu.Lock()
	for isNotOK {
		if isTimeout {
			break
		}
		ID = uint64(time.Now().Nanosecond())
		// 檢查ID有沒有重複
		_, isNotOK = p.registerList[ID]
		if isNotOK {
			continue
		}
		p.registerList[ID] = req
	}
	p.mu.Unlock()

	if isTimeout {
		res.ErrorCode = 2
		res.ErrorText = "回應超時"
		return
	}

	res.Id = ID
	return
}

// Push 推送訊息
func (p *Pusher) Push(ctx context.Context, req *pusher.PushRequest, res *pusher.PushResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			err = catchErr.(error)
		}
		res.ErrorCode = 0
	}()

	res.ErrorCode = -1
	redisConn, connErr := redis.Dial("tcp", PusherConfig.Cache.Host+PusherConfig.Cache.Port)
	if connErr != nil {
		res.ErrorCode = 1
		res.ErrorText = fmt.Sprintf("(Dial: %s) Redis連線錯誤：%v", PusherConfig.Cache.Host+PusherConfig.Cache.Port, connErr)
		return
	}
	defer redisConn.Close()
	transferData := TransferData{
		ChatRoom: req.ChatRoom,
		Data:     req.Message,
	}
	transferJSON, marshalErr := json.Marshal(transferData)
	if marshalErr != nil {
		res.ErrorCode = 2
		res.ErrorText = fmt.Sprintf("JSON_ENCODE ERROR : %v", marshalErr)
	}

	reply, doErr := redisConn.Do("Publish", PusherConfig.Cache.Channel, transferJSON)
	if doErr != nil {
		res.ErrorCode = 3
		res.ErrorText = reply.(string)
		return
	}
	return
}

// RegisterList 註冊列表
func (p *Pusher) RegisterList(ctx context.Context, req *pusher.RegisterListRequest, res *pusher.RegisterListResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			err = catchErr.(error)
		}
	}()

	if req.ChatRoom == "" {
		p.mu.RLock()
		res.RegisterList = p.registerList
		p.mu.RUnlock()
		return
	}

	tempRegisterList := map[uint64]*pusher.RegisterRequest{}
	p.mu.RLock()
	for ID, registerInfo := range p.registerList {
		if registerInfo.ChatRoom == req.ChatRoom {
			tempRegisterList[ID] = registerInfo
		}
	}
	p.mu.RUnlock()
	res.RegisterList = tempRegisterList
	return
}

// ChatRoomList 聊天室列表
func (p *Pusher) ChatRoomList(ctx context.Context, req *pusher.ChatRoomListRequest, res *pusher.ChatRoomListResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			err = catchErr.(error)
		}
	}()

	tempChatRoomList := []string{}

	p.mu.RLock()
	for ID, registerInfo := range p.registerList {
		if req.Id == 0 || req.Id == ID {
			tempChatRoomList = append(tempChatRoomList, registerInfo.ChatRoom)
		}
	}
	p.mu.RUnlock()
	res.ChatRoomList = UniqueStringSlice(tempChatRoomList)
	return
}

// Deregister 取消註冊
func (p *Pusher) Deregister(ctx context.Context, req *pusher.DeregisterRequest, res *pusher.DeregisterResponse) (err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			err = catchErr.(error)
			return
		}

		res.ErrorCode = 0
		res.ErrorText = fmt.Sprintf("Deregister OK, ID: %d", req.Id)
	}()
	res.ErrorCode = -1
	p.mu.Lock()
	delete(p.registerList, req.Id)
	p.mu.Unlock()
	return
}

// 監聽通道
func (p *Pusher) listenChannel(ip, port, channel string) (err error) {
	defer func() {
		log.Printf("🎃  Listen Channel Stop  🎃")
		if catchErr := recover(); catchErr != nil {
			log.Printf("🔎  because unexpected error (%v) ", catchErr)
		}
	}()

	log.Printf("✨  Channel start connecting (%s) ✨", ip+port)
	redisConn, connErr := redis.Dial("tcp", ip+port)
	if connErr != nil {
		log.Printf("🎃  Channel connected (%s) error (%v) 🎃", ip+port, connErr)
		err = connErr
		return
	}
	defer redisConn.Close()
	log.Printf("✨  Channel connected OK (%s) ✨", ip+port)

	log.Printf("✨  Channel start subscribing (%s) ✨", channel)
	psc := redis.PubSubConn{Conn: redisConn}
	err = psc.PSubscribe(channel)
	if err != nil {
		log.Printf("🎃  Channel subscribe (%s) error (%v) 🎃", channel, err)
		return
	}
	log.Printf("✨  Channel subscribe (%s) OK ✨", channel)

	reply, pingErr := redisConn.Do("PING")
	if pingErr != nil {
		log.Printf("🎃  PING error (%v) 🎃", pingErr)
		err = pingErr
		return
	}

	log.Printf("✨  PING REPLY %v ✨", string(reply.([]interface{})[0].([]byte)))

	log.Printf("✨  Listen Channel Start (%s:%s) ✨", ip+port, channel)

	for {
		switch v := psc.Receive().(type) {
		case redis.PMessage:
			log.Printf("✨  channel (%s) receive transfer data (%v) to decode ✨", channel, string(v.Data))
			var transferData TransferData
			err := json.Unmarshal(v.Data, &transferData)
			if err != nil {
				log.Printf("🎃  channel (%s) decode transfer data (%s) failed for chat room (%s) 🎃", channel, transferData.Data, transferData.ChatRoom)
				continue
			}
			log.Printf("✨  channel (%s) receive message (%s) for chat room (%s) after decode transfer data ✨", channel, transferData.Data, transferData.ChatRoom)
			p.pushToClient(transferData.ChatRoom, transferData.Data)

		default:
			continue
		}
	}
}

// 推送訊息到客戶端
func (p *Pusher) pushToClient(targetChatRoom, message string) {
	wg := new(sync.WaitGroup)
	for ID, severInfo := range p.registerList {
		if severInfo.ChatRoom != targetChatRoom {
			continue
		}
		wg.Add(1)
		go func(ID uint64, severInfo *pusher.RegisterRequest, wg *sync.WaitGroup) {
			url := fmt.Sprintf("http://%s/pusher", severInfo.Ip+severInfo.Port)

			defer func() {
				if catchErr := recover(); catchErr != nil {
					log.Printf(`🎃  #%d chat room (%s) push message (%s) to url (%s) failed 🎃`, ID, targetChatRoom, message, url)
					log.Printf(`🔎  #%d because unexpected error (%v)`, ID, catchErr)
				}
				wg.Done()
			}()

			log.Printf("✨  #%d chat room (%s) ready to push message (%s) to url (%s) ✨", ID, targetChatRoom, message, url)

			payload := strings.NewReader(message)
			req, err := http.NewRequest("PUT", url, payload)
			if err != nil {
				log.Printf(`🎃  #%d Create reuqest failed (%v) 🎃`, ID, err)
				return
			}
			req.Header.Add("host", severInfo.Host)
			req.Header.Add("api-key", severInfo.ApiKey)

			res, curlErr := http.DefaultClient.Do(req)

			if curlErr != nil {
				log.Printf(`🎃  #%d Send reuqest failed (%v) 🎃`, ID, curlErr)
				return
			}
			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)
			if string(body) != "ok" {
				log.Printf(`🎃  #%d Server (%s) not response ok (%s) 🎃`, ID, url, string(body))
				return
			}

			log.Printf("✨  #%d chat room (%s) push message (%s) to url (%s) OK ✨", ID, targetChatRoom, message, url)

		}(ID, severInfo, wg)
	}
	wg.Wait()

}
