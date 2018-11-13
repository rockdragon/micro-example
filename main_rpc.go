package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/rockdragon/micro_example/handler"
	"github.com/rockdragon/micro_example/subscriber"
	"github.com/rockdragon/micro_example/utils"

	example "github.com/rockdragon/micro_example/proto/example"
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
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(utils.SrvName, service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber(utils.SrvName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
