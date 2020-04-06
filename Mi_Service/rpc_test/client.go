package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	//建立网络连接
	cli,err :=rpc.DialHTTP("tcp","127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络错误")
	}

	var pd int

	err = cli.Call("Pandaner.Getinfo",10086,&pd)
	if err != nil {
		fmt.Println("call失败")
	}
	fmt.Println("最后得到值：",pd)
}

