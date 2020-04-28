package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Mi_house/PostAvatar/handler"
	"Mi_house/PostAvatar/subscriber"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTAVATAR "Mi_house/PostAvatar/proto/PostAvatar"
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
		micro.Name("go.micro.srv.PostAvatar"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTAVATAR.RegisterPostAvatarHandler(service.Server(), new(handler.PostAvatar))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostAvatar", service.Server(), new(subscriber.PostAvatar))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostAvatar", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
