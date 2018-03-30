package main

import (
	"context"
	"log"
	"os"
	greeter "proto/greeter/micro"
	"time"

	"github.com/micro/cli"
	mrpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
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

func logWrapper(fn server.HandlerFunc) (sh server.HandlerFunc) {
	sh = func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[%v] server request: %s", time.Now(), req.Method())
		return fn(ctx, req, rsp)
	}
	return
}

var isClient bool

func main() {
	service := mrpc.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
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

	greeter.RegisterSayHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Printf("🎃  Greeter Service return an error : %v 🎃", err)
	}
}

// Setup and the client
func runClient() {
	service := mrpc.NewService()
	// Create new client
	microClient := greeter.NewSayClient("greeter", service.Client())

	// Ping the service
	rsp, err := microClient.Ping(context.TODO(), &greeter.PingRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	// Print response
	log.Println(rsp.ServiceName)
}
