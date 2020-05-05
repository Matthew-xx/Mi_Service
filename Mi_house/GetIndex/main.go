package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/GetIndex/handler"
	"Mi_house/GetIndex/subscriber"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	GetIndex "Mi_house/GetIndex/proto/GetIndex"
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
