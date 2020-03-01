package main

import "../rabbitMq"

func main()  {
	imoocTwo := rabbitMq.NewRabbitMQRouting("exImooc","imooc_two")
	imoocTwo.ReceiveRouting()
}
