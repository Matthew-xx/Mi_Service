package main

import (
	"Mi_house/GetImageCd/handler"
	"Mi_house/GetImageCd/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	GetImageCd "Mi_house/GetImageCd/proto/GetImageCd"
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
		micro.Name("go.micro.srv.GetImageCd"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetImageCd.RegisterGetImageCdHandler(service.Server(), new(handler.GetImageCd))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetImageCd", service.Server(), new(subscriber.GetImageCd))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetImageCd", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
