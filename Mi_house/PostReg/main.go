package main

import (
	"Mi_house/PostReg/handler"
	"Mi_house/PostReg/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTREG "Mi_house/PostReg/proto/PostReg"
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
		micro.Name("go.micro.srv.PostReg"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTREG.RegisterPostRegHandler(service.Server(), new(handler.PostReg))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), new(subscriber.PostReg))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
