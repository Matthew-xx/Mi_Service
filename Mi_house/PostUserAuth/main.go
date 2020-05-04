package main

import (
	"Mi_house/PostUserAuth/handler"
	"Mi_house/PostUserAuth/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"

	POSTUSERAUTH "Mi_house/PostUserAuth/proto/PostUserAuth"
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
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTUSERAUTH.RegisterPostUserAuthHandler(service.Server(), new(handler.PostUserAuth))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), new(subscriber.PostUserAuth))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
