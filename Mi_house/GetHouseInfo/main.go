package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"Mi_house/GetHouseInfo/handler"
	"Mi_house/GetHouseInfo/subscriber"
	"github.com/micro/go-plugins/registry/etcdv3"

	GETHOUSEINFO "Mi_house/GetHouseInfo/proto/GetHouseInfo"
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
		micro.Name("go.micro.srv.GetHouseInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETHOUSEINFO.RegisterGetHouseInfoHandler(service.Server(), new(handler.GetHouseInfo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), new(subscriber.GetHouseInfo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
