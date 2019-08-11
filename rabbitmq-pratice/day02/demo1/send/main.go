package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// 错误检查代码
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
func bodyFrom(str []string) (newstr string) {
	for _, v := range str {
		newstr += v
	}
	return newstr
}
func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.10.252:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 创建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 申请一个交换器
	err = ch.ExchangeDeclare(
		"logs",   //交换器的名字
		"fanout", //交换器的类型、这里为广播类型
		true,     //是否持久
		false,    //无用的时候是否自动删除
		false,    //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
		false,    //no-wait
		nil,      //arguments
	)

	failOnError(err, "Failed to declare q queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set Qos")
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs", //发送到交换机的名字
		"",     // 路由，即队列的名字
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //消息的持久化
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)

	failOnError(err, "Failed to publish a message")
}
