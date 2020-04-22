package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/GetIndex/handler"
	"Mi_house/GetIndex/subscriber"

	GetIndex "Mi_house/GetIndex/proto/GetIndex"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetIndex"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetIndex.RegisterGetIndexHandler(service.Server(), new(handler.GetIndex))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetIndex", service.Server(), new(subscriber.GetIndex))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetIndex", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
