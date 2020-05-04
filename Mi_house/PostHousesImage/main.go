package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/PostHousesImage/handler"
	"Mi_house/PostHousesImage/subscriber"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTHOUSESIMAGE "Mi_house/PostHousesImage/proto/PostHousesImage"
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
		micro.Name("go.micro.srv.PostHousesImage"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTHOUSESIMAGE.RegisterPostHousesImageHandler(service.Server(), new(handler.PostHousesImage))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHousesImage", service.Server(), new(subscriber.PostHousesImage))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHousesImage", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
