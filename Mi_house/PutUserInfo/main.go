package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/PutUserInfo/handler"
	"github.com/micro/go-micro/registry"
	"Mi_house/PutUserInfo/subscriber"
	"github.com/micro/go-plugins/registry/etcdv3"

	PUTUSERINFO "Mi_house/PutUserInfo/proto/PutUserInfo"
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
		micro.Name("go.micro.srv.PutUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	PUTUSERINFO.RegisterPutUserInfoHandler(service.Server(), new(handler.PutUserInfo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PutUserInfo", service.Server(), new(subscriber.PutUserInfo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PutUserInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
