package main

import (
	"Mi_house/PostHouses/handler"
	"Mi_house/PostHouses/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTHOUSES "Mi_house/PostHouses/proto/PostHouses"
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
		micro.Name("go.micro.srv.PostHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTHOUSES.RegisterPostHousesHandler(service.Server(), new(handler.PostHouses))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHouses", service.Server(), new(subscriber.PostHouses))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHouses", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
