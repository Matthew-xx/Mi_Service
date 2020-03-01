package main

import "../rabbitMq"
func main()  {
	rabbitmq := rabbitMq.NewRabbitmqPubSub(""+ "newProduct")
	rabbitmq.ReceiveSub()
}
