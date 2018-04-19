package main

import (
	"context"
	"log"
	"os"
	greeter "proto/greeter/micro"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nsq"
)

// Greeter 客人
type Greeter struct{}

// Ping 測試連線
func (c *Greeter) Ping(ctx context.Context, req *greeter.PingRequest, res *greeter.PongResponse) (err error) {
	res.ServiceName = "greeter"
	res.Environment = os.Getenv("PROJECT_ENV")
	return
}

// Hello 打招呼
func (c *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// Meet 遇見
type Meet struct{}

// Nice 打招呼
func (c *Meet) Nice(ctx context.Context, req *greeter.HelloRequest) error {
	log.Println("Hello " + req.Name)
	return nil
}

func logWrapper(fn server.HandlerFunc) (sh server.HandlerFunc) {
	sh = func(ctx context.Context, req server.Request, rsp interface{}) error {
		beginTime := time.Now()
		err := fn(ctx, req, rsp)
		log.Printf("[%v] server request: %s, excusion: %v", time.Now(), req.Method(), time.Since(beginTime))
		return err
	}
	return
}

var isClient bool

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.RegisterInterval(time.Second*5),
		micro.RegisterTTL(time.Second*10),
		micro.BeforeStart(func() error {
			log.Println("🐳  Greeter Service Start 🐳")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("🔥  Greeter Service Stop 🔥")
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
			isClient = c.Bool("run_client")
		}),
	)

	if isClient {
		runClient()
		return
	}

	// 註冊訂閱
	err := micro.RegisterSubscriber("meet", service.Server(), new(Meet), server.SubscriberQueue("guest"))
	if err != nil {
		log.Printf("🎃  Meet Subscribe return an error : %v 🎃", err)
		return
	}

	// 註冊服務
	greeter.RegisterSayHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Printf("🎃  Greeter Service return an error : %v 🎃", err)
	}
}

// Setup and the client
func runClient() {
	service := micro.NewService()
	service.Init()

	err := micro.NewPublisher("meet", service.Client()).Publish(context.Background(), &greeter.HelloRequest{Name: "Zuolar"})
	if err != nil {
		log.Fatalln("Publish", err)
		return
	}
	time.Sleep(time.Millisecond)

	// Create new client
	microClient := greeter.NewSayClient("greeter", service.Client())

	// Ping the service
	rsp, err := microClient.Ping(context.TODO(), &greeter.PingRequest{})
	if err != nil {
		log.Fatalln("Ping", err)
	}

	// Print response
	log.Println("Service", rsp.ServiceName)
}
