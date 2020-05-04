package main

import (
	"Mi_house/GetUserHouses/handler"
	"Mi_house/GetUserHouses/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	GETUSERHOUSES "Mi_house/GetUserHouses/proto/GetUserHouses"
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
		micro.Name("go.micro.srv.GetUserHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETUSERHOUSES.RegisterGetUserHousesHandler(service.Server(), new(handler.GetUserHouses))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), new(subscriber.GetUserHouses))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
