package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/GetSession/handler"
	"Mi_house/GetSession/subscriber"

	GetSession "Mi_house/GetSession/proto/GetSession"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetSession.RegisterGetSessionHandler(service.Server(), new(handler.GetSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSession", service.Server(), new(subscriber.GetSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
