package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"melon_micro/proto/config"
	"os"
	"path/filepath"

	"github.com/micro/go-micro"
	"github.com/naoina/toml"
)

// Database 資料庫設定
type Database struct {
	DB         string `toml:"db" json:"db,omitempty"`
	HostMaster string `toml:"host_master" json:"host_master,omitempty"`
	HostSlave  string `toml:"host_slave" json:"host_slave,omitempty"`
	Port       string `toml:"port" json:"port,omitempty"`
	Username   string `toml:"username" json:"username,omitempty"`
	Password   string `toml:"password" json:"password,omitempty"`
}

// Cache 資訊
type Cache struct {
	Host     string `toml:"host" json:"host,omitempty"`
	Password string `toml:"password" json:"password,omitempty"`
	Port     string `toml:"port" json:"port,omitempty"`
}

// API 資訊
type API struct {
	IP     string `toml:"ip" json:"ip,omitempty"`
	Host   string `toml:"host" json:"host,omitempty"`
	Port   string `toml:"port" json:"port,omitempty"`
	APIKey string `toml:"api_key" json:"api_key,omitempty"`
}

// Config 設定檔
type Config struct {
	App struct {
		Env string `toml:"env" json:"env,omitempty"`
	} `toml:"app" json:"app,omitempty"`
	Database struct {
		GameMaster   Database `toml:"game_master" json:"game_master,omitempty"`
		GameSlave    Database `toml:"game_slave" json:"game_slave,omitempty"`
		TradeMaster  Database `toml:"trade_master" json:"trade_master,omitempty"`
		TradeSlave   Database `toml:"trade_slave" json:"trade_slave,omitempty"`
		MemberMaster Database `toml:"member_master" json:"member_master,omitempty"`
		MemberSlave  Database `toml:"member_slave" json:"member_slave,omitempty"`
		Test         Database `toml:"test" json:"test,omitempty"`
	} `toml:"database" json:"database,omitempty"`
	Cache struct {
		BothMaster   Cache `toml:"both_master" json:"both_master,omitempty"`
		BothSlave    Cache `toml:"both_slave" json:"both_slave,omitempty"`
		GameMaster   Cache `toml:"game_master" json:"game_master,omitempty"`
		GameSlave    Cache `toml:"game_slave" json:"game_slave,omitempty"`
		TradeMaster  Cache `toml:"trade_master" json:"trade_master,omitempty"`
		TradeSlave   Cache `toml:"trade_slave" json:"trade_slave,omitempty"`
		MemberMaster Cache `toml:"member_master" json:"member_master,omitempty"`
		MemberSlave  Cache `toml:"member_slave" json:"member_slave,omitempty"`
		Test         Cache `toml:"test" json:"test,omitempty"`
	} `toml:"cache" json:"cache,omitempty"`
	API struct {
		Durian API `toml:"durian" json:"durian,omitempty"`
		Green  API `toml:"green" json:"green,omitempty"`
	} `toml:"api" json:"api,omitempty"`
}

// Get 取設定值
func (c *Config) Get(ctx context.Context, req *config.ConfigRequest, rsp *config.ConfigResponse) (err error) {
	configFile := GetAppRoot() + "/environment/" + GetAppEnv() + ".toml"
	tomlData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("err", err)
		return
	}

	err = toml.Unmarshal(tomlData, &c)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("Config", c)

	var conf2Json []byte
	conf2Json, err = json.Marshal(c)
	if err != nil {
		log.Println("err", err)
		return
	}

	log.Println("json", string(conf2Json))

	var json2Conf *config.Conf
	err = json.Unmarshal(conf2Json, &json2Conf)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("conf", json2Conf)
	rsp.Config = json2Conf
	// rsp.Config = "OK"
	return nil
}

// Ping 測試連線
func (c *Config) Ping(ctx context.Context, req *config.PingRequest, res *config.PongResponse) (err error) {
	res.ServiceName = "config"
	res.CurrentEnvironment = GetAppEnv()
	return
}

// GetAppRoot 取執行檔的根目錄
func GetAppRoot() string {
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("WARNING", "GetAppRoot：取根目錄失敗，自動抓取 APP_ROOT 的環境變數")
		return os.Getenv("APP_ROOT")
	}
	return root
}

// GetAppEnv 取環境設定
func GetAppEnv() string {
	return os.Getenv("PROJECT_ENV")
}

func main() {
	cfg := new(Config)

	service := micro.NewService(
		micro.Name("config"),
		micro.Version("latest"),
		micro.BeforeStart(func() error {
			log.Println("🐳  Config Service Start 🐳")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("🔥  Config Service Stop 🔥")
			return nil
		}),
	)

	service.Init()
	config.RegisterConfigHandler(service.Server(), cfg)

	if err := service.Run(); err != nil {
		log.Printf("🎃  Config Service return an error : %v 🎃", err)
	}
}
