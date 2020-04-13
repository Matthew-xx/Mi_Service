package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"mark/myTest/handler"
	myTest "mark/myTest/proto/myTest"
	"mark/myTest/subscriber"
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
		micro.Name("go.micro.srv.myTest"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	myTest.RegisterMyTestHandler(service.Server(), new(handler.MyTest))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.myTest", service.Server(), new(subscriber.MyTest))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.myTest", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
