package main

import "../rabbitMq"

func main()  {
	imoocOne := rabbitMq.NewRabbitMQRouting("exImooc","imooc_one")
	imoocOne.ReceiveRouting()
}
