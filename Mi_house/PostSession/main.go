package main

import (
	"Mi_house/PostSession/handler"
	"Mi_house/PostSession/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTSESSION "Mi_house/PostSession/proto/PostSession"
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
		micro.Name("go.micro.srv.PostSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTSESSION.RegisterPostSessionHandler(service.Server(), new(handler.PostSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostSession", service.Server(), new(subscriber.PostSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
