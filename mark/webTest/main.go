package main

import (
        "github.com/micro/go-micro/util/log"
        "github.com/micro/go-plugins/registry/etcdv3"
        "net/http"

        "github.com/micro/go-micro/registry"
        "github.com/micro/go-micro/web"
        "mark/webTest/handler"
)

func main() {
        reg := etcdv3.NewRegistry(func(op *registry.Options){
                op.Addrs = []string{
                        "127.0.0.1:2379",
                }
        })
	// create new web service
        service := web.NewService(
                web.Registry(reg),
                //服务名称
                web.Name("go.micro.web.webTest"),
                web.Version("latest"),
                //设置服务的端口号
                web.Address(":8089"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler 注册web页面
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler 注册请求
	service.HandleFunc("/webTest/call", handler.WebTestCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
