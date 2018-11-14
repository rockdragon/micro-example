package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/rockdragon/micro_example/handler"
	"github.com/rockdragon/micro_example/subscriber"
	"github.com/rockdragon/micro_example/utils"

	business "github.com/rockdragon/micro_example/proto/business"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(utils.SrvName),
		micro.Version(utils.Version),
	)

	// Initialise service
	service.Init()

	// Register Handler
	business.RegisterAPIHandler(service.Server(), new(handler.API))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(utils.SrvName, service.Server(), new(subscriber.API))

	// Register Function as Subscriber
	micro.RegisterSubscriber(utils.SrvName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
