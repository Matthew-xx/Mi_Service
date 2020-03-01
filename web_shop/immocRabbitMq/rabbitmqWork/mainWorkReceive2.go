package main

import "../rabbitMq"

func main()  {
	rabbitmq := rabbitMq.NewRabbitMQSimple(""+"imoocSimple")
	rabbitmq.ConsumeSimple()
}
