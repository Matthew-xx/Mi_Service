package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd "pro_package"
)
func main()  {
	//客户端连接服务器
	conn,err :=grpc.Dial("localhost:10086",grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常",err)
	}
	defer conn.Close()

	//获得grpc句柄
	c :=pd.NewHelloServerClient(conn)
	//通过句柄调用函数
	req,err := c.SayHello(context.Background(),&pd.HelloReq{Name:"mark"})
	if err != nil {
		fmt.Println("sayhello服务调用失败",err)
	}
	fmt.Println("调用sayhello返回：",req.Msg)

	res,err := c.SayName(context.Background(),&pd.NameReq{Name:"马克思"})
	if err != nil {
		fmt.Println("sayname服务调用失败",err)
	}
	fmt.Println("调用sayname返回：",res.Msg)
}
