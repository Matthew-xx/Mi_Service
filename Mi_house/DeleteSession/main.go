package main

import (
	"Mi_house/DeleteSession/handler"
	"Mi_house/DeleteSession/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	DELETESESSION "Mi_house/DeleteSession/proto/DeleteSession"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:2379",
		}
	})
	// New Service
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.DeleteSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	DELETESESSION.RegisterDeleteSessionHandler(service.Server(), new(handler.DeleteSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), new(subscriber.DeleteSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
