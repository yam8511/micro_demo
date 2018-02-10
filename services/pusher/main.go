package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"melon_micro/proto/pusher"
	"os"
	"sync"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/naoina/toml"
)

// Config 設定檔
type Config struct {
	App struct {
		Env string `toml:"env" json:"env,omitempty"`
	} `toml:"app" json:"app,omitempty"`
	Cache struct {
		Host     string `toml:"host" json:"host,omitempty"`
		Password string `toml:"password" json:"password,omitempty"`
		Port     string `toml:"port" json:"port,omitempty"`
		Channel  string `toml:"channel" json:"channel,omitempty"`
	} `toml:"cache" json:"cache,omitempty"`
}

// PusherConfig 推送設定
var PusherConfig Config

func main() {
	Usage()

	pusherService := &Pusher{
		registerList: map[uint64]*pusher.RegisterRequest{},
		mu:           new(sync.RWMutex),
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	service := micro.NewService(
		micro.Name("pusher"),
		micro.Version("1.0"),
		micro.BeforeStart(func() (err error) {
			err = loadPusherConfig()
			CheckDangerError(err)
			go func() {
				pusherService.listenChannel(PusherConfig.Cache.Host, PusherConfig.Cache.Port, PusherConfig.Cache.Channel)
				wg.Done()
			}()
			// CheckDangerError(err)
			log.Println("🐳  Pusher Service Start 🐳")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("🔥  Pusher Service Stop 🔥")
			return nil
		}),
		micro.WrapHandler(logWrapper),
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	pusher.RegisterPusherHandler(service.Server(), pusherService)

	go func() {
		err := service.Run()
		if err != nil {
			log.Printf("🎃  Pusher Service return an error (%v) 🎃", err)
		}
		wg.Done()
	}()

	wg.Wait()
}

// 讀取設定檔
func loadPusherConfig() (err error) {
	if GetAppEnv() == "" {
		err = errors.New("🎃 請宣告 PROJECT_ENV 【local、development、qatest、gcp-qatest】 🎃")
		return
	}
	configFile := GetAppRoot() + "/config/" + GetAppEnv() + ".toml"
	tomlData, readFileErr := ioutil.ReadFile(configFile)
	if readFileErr != nil {
		err = readFileErr
		return
	}

	err = toml.Unmarshal(tomlData, &PusherConfig)
	return
}

// logWrapper 記錄請求
func logWrapper(fn server.HandlerFunc) (sh server.HandlerFunc) {
	sh = func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[%v] server request: %s", time.Now(), req.Method())
		return fn(ctx, req, rsp)
	}
	return
}

// Setup and the client
func runClient(service micro.Service) {
	// Create new client
	microClient := pusher.NewPusherClient("pusher", service.Client())

	// Ping the service
	rsp, err := microClient.Ping(context.TODO(), &pusher.PingRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	// Print response
	log.Println(rsp.ServiceName)
}
