package main

import (
	"Mi_house/GetSmsCd/handler"
	"Mi_house/GetSmsCd/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	GETSMSCD "Mi_house/GetSmsCd/proto/GetSmsCd"
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
		micro.Name("go.micro.srv.GetSmsCd"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETSMSCD.RegisterGetSmsCdHandler(service.Server(), new(handler.GetSmsCd))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSmsCd", service.Server(), new(subscriber.GetSmsCd))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSmsCd", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
