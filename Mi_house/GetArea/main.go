package main

import (
	"Mi_house/GetArea/handler"
	"Mi_house/GetArea/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	GETAREA "Mi_house/GetArea/proto/GetArea"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:2379",
		}
	})
	// New Service
	//在web中的handler使用grpc调用了，这里也改grpc
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETAREA.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), new(subscriber.GetArea))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
