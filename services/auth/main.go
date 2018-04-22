package main

import (
	"context"
	"log"
	pb "proto/auth/micro"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nsq"
)

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
var service micro.Service

func main() {
	service = micro.NewService(
		micro.Name("auth"),
		micro.Version("1.0.0"),
		micro.RegisterInterval(time.Second*5),
		micro.RegisterTTL(time.Second*10),
		micro.BeforeStart(func() error {
			err := SetUpTable()
			if err != nil {
				return err
			}
			log.Println("🐳  Auth Service Start 🐳")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("🔥  Auth Service Stop 🔥")
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
	err := micro.RegisterSubscriber("auth_login", service.Server(), new(LoginSubscriber), server.SubscriberQueue("login"))
	if err != nil {
		log.Printf("🎃  Auth Login Subscribe return an error : %v 🎃", err)
		return
	}

	// 註冊訂閱
	err = micro.RegisterSubscriber("auth_register", service.Server(), new(JoinSubscriber), server.SubscriberQueue("register"))
	if err != nil {
		log.Printf("🎃  Auth Register Subscribe return an error : %v 🎃", err)
		return
	}

	// 註冊服務
	pb.RegisterAuthHandler(service.Server(), new(Auth))

	if err := service.Run(); err != nil {
		log.Printf("🎃  Auth Service return an error : %v 🎃", err)
	}
}

// Setup and the client
func runClient() {
	// service := micro.NewService()
	// service.Init()

	// err := micro.NewPublisher("meet", service.Client()).Publish(context.Background(), &pb.HelloRequest{Name: "Zuolar"})
	// if err != nil {
	// 	log.Fatalln("Publish", err)
	// 	return
	// }
	// time.Sleep(time.Millisecond)

	// // Create new client
	microClient := pb.NewAuthClient("auth", service.Client())

	// Ping the service
	rsp, err := microClient.Ping(context.TODO(), &pb.PingRequest{})
	if err != nil {
		log.Fatalln("Ping", err)
	}

	// Print response
	log.Println("Service", rsp.ServiceName)
}
