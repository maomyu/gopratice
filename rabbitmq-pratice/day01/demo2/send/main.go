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
	q, err := ch.QueueDeclare(
		"hello", //队列的名字
		true,    //是否持久化
		false,   //对列没有用到的时候是否删除
		false,   //是否设置排他，true为是。如果设置为排他，则队列仅对首次声明他的连接可见，并在连接断开时自动删除
		false,   //是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息
		nil,     //arguments
	)
	failOnError(err, "Failed to declare q queue")
	fmt.Println("队列的名字：", q.Name)
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set Qos")
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     //发送到交换机的名字
		q.Name, // 路由，即队列的名字
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
