package main

import (
	"../rabbitMq"
	"fmt"
	"strconv"
)

func main()  {
	rabbitmq := rabbitMq.NewRabbitmqPubSub(""+"newProduct")
	for i:=0; i<100; i++ {
		rabbitmq.PublishPub("订阅模式生产第"+strconv.Itoa(i)+"条"+"数据")
	}
	fmt.Println("发送成功！")
}

