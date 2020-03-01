package main

import (
	"../rabbitMq"
	"fmt"
	"strconv"
)

func main()  {
	imoocOne := rabbitMq.NewRabbitMQRouting("exImooc","imooc_one")
	imoocTwo := rabbitMq.NewRabbitMQRouting("exImooc","imooc_two")

	for i:=0; i<=10; i++ {
		imoocOne.PublishRouting("hello imooc one"+strconv.Itoa(i))
		imoocTwo.PublishRouting("hello imooc two"+strconv.Itoa(i))
	}
	fmt.Println("发送成功")
}

