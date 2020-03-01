package rabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//  amqp://账号：密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@192.168.99.100:5672/imooc"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}
//创建rabbitmq结构体实例
func NewRabbitMQ(queueName,exchange,key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName:queueName,Exchange:exchange,Key:key, Mqurl:MQURL}
	var err error
	//创建rabbit连接
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err,"创建连接错误")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err,"获取channel失败")

	return rabbitmq
}

//断开channel和connection
func (r *RabbitMQ) Destory()  {
	r.channel.Close()
	r.conn.Close()
}

//错误处理
func (r *RabbitMQ) failOnError(err error,message string)  {
	if err != nil{
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

//简单模式step1：简单模式下rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ  {

	return NewRabbitMQ(queueName,"","")
}

//简单模式step2：简单模式下生产代码
func (r *RabbitMQ) PubilshSimple(message string) {
	//1,申请队列,如果队列不存在会自动创建，如果存在则跳过创建,保证队列存在消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false,  //是否自动删除
		false,  //是否具有排他性
		false,  //是否阻塞
		nil,  //额外属性
		)
	if err != nil {
		fmt.Println(err)
	}

	//发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //如果为true，会根据exchange类型和routkey规则，
		// 如果无法找到符合条件的队列，那么会把消息返回给发送者
		false,  //如果为true，当exchange发送消息到队列后发现队列上
		// 没有绑定消费者，则会把消息返回给发送者
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		},)

}

//消费代码
func (r *RabbitMQ) ConsumeSimple()  {
	//申请队列
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false,  //是否自动删除
		false,  //是否具有排他性
		false,  //是否阻塞
		nil,  //额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msgs,err := r.channel.Consume(
		r.QueueName,
		"",  //用来区分多个消费者,空为不区分
		true,  //是否自动应答，说明已接收完消息
		false, //是否具有排他性
		false,  //若为true，表示不能将同一个connection中发送的消息
		// 传递给这个connection中的消费者
		false,  //队列是否阻塞，false为阻塞
		nil,
		)
	if err != nil {
		fmt.Println(err)
	}

	//消费消息
	forever := make(chan bool)
	go func() {
		for d:= range msgs{
			//实现要处理消息的逻辑函数
			log.Printf("received a message: %s",d.Body)

		}
	}()

	log.Printf("[*] waiting for messages,to exit press ctrl+c")
	<- forever
}

//订阅模式创建rabbitmq实例
func NewRabbitmqPubSub(exchangeName string) *RabbitMQ {
	//创建rabbitmq实例
	rabbitmq := NewRabbitMQ("",exchangeName,"")
	var err error
	//获取connection
	rabbitmq.conn ,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err,"failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err,"failed to open a channel")

	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange, //
		"fanout",  //订阅模式下，广播类型
		true,
		false,
		false, //true表示这个exchange不可以被client用来推送消息，
		// 仅用来进行exchange和exchange之间的绑定
		false,
		nil,
		)

	r.failOnError(err,"failed to declare an excha"+"nge")

	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		},
		)
}

//订阅模式消费端代码
func (r *RabbitMQ) ReceiveSub()  {
	//试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", //交换机类型
		true,
		false,
		false,
		false,
		nil,
		)
	r.failOnError(err,"failed to declare an exch"+"ange")

	q,err := r.channel.QueueDeclare(
		"",  //随机生产队列名称
		false,
		false,
		true, //排他性
		false,
		nil,
		)
	r.failOnError(err,"failed to declare a queue")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"",  //在pub,sub模式下，key为空
		r.Exchange,
		false,
		nil,
		)

	//消费消息
	messages,err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)

	forever := make(chan bool)

	go func() {
		for d:= range messages{
			log.Printf("received a message:%s",d.Body)
		}
	}()
	fmt.Println("exit press ctrl+c\n")
	<- forever
}

//路由模式
func NewRabbitMQRouting(exchangeName,routingKey string) *RabbitMQ {
	//创建rabbitMQ实例
	rabbitmq := NewRabbitMQ("",exchangeName,routingKey)
	var err error
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err,"failed to connect rabbitmq")

	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err,"failed to open a channel")

	return rabbitmq
}

func (r *RabbitMQ) PublishRouting(message string)  {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",  //与其他mq不一样
		true,
		false,
		false,
		false,
		nil,
		)
	r.failOnError(err,"failed to declare an excha"+"nge")

	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,   //其他方法不用设置
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		},
		)
	if err != nil {
		panic(err)
	}
}

//路由模式接收消息
func (r *RabbitMQ) ReceiveRouting()  {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",  //交换机类型
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err,"failed to declare an excha"+"nge")

	q,err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	r.failOnError(err,"failed to declare a queue")

	//绑定到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,  //需要绑定key
		r.Exchange,
		false,
		nil,
		)

	//消费消息
	messages,err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	forever := make(chan bool)

	go func() {
		for d:= range messages{
			log.Printf("received a message:%s",d.Body)
		}
	}()

	fmt.Println("退出")
	<- forever
}


