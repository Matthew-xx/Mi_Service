package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pd "pro_package"
)

type server struct {}

//为了注册服务时的&server{}能类型匹配，建两个与hello文件里面的函数（因为HelloServerClient实现了下面两个方法
func (this *server) SayHello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloRsp, err error){

	return &pd.HelloRsp{Msg:"hello "+in.Name},nil
}

func (this *server) SayName(ctx context.Context, in *pd.NameReq) (out *pd.NameRsp, err error){

	return &pd.NameRsp{Msg:in.Name+" moring"},nil
}

func main()  {
	ln,err := net.Listen("tcp",":10086")
	if err != nil {
		fmt.Println("网络错误",err)
	}

	//创建grpc服务
	srv := grpc.NewServer()
	//注册服务
	pd.RegisterHelloServerServer(srv,&server{})

	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误",err)
	}
}
