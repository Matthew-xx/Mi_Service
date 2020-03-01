package main

import (
	"./rabbitMq"
	"fmt"
)

func main()  {
	rabbitmq := rabbitMq.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PubilshSimple("hello imooc!")
	fmt.Println("发送成功")
}


