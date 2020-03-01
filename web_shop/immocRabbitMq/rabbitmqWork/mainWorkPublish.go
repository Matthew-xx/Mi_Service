package main

import (
	"../rabbitMq"
	"fmt"
	"strconv"
)

func main()  {
	rabiitmq := rabbitMq.NewRabbitMQSimple("imoocSimple")
	for i:=0; i<=100; i++ {
		rabiitmq.PubilshSimple("hello iooc!" + strconv.Itoa(i))
		//time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
