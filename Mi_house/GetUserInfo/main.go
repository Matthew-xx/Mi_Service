package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/GetUserInfo/handler"
	"Mi_house/GetUserInfo/subscriber"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	GETUSERINFO "Mi_house/GetUserInfo/proto/GetUserInfo"
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
		micro.Name("go.micro.srv.GetUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETUSERINFO.RegisterGetUserInfoHandler(service.Server(), new(handler.GetUserInfo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserInfo", service.Server(), new(subscriber.GetUserInfo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
