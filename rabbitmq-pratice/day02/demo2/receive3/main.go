package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.10.252:5672")
	failOnError(err, "Failed to connect to server")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")
	defer ch.Close()
	// 申请一个交换器
	err = ch.ExchangeDeclare(
		"logs_direct", //交换器的名字
		"direct",      //交换器的类型、这里为广播类型
		true,          //是否持久
		false,         //无用的时候是否自动删除
		false,         //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
		false,         //no-wait
		nil,           //arguments
	)
	// 申请一个队列
	q, err := ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when usused
		true,  // exclusive
		false, //no-wait
		nil,   // arguments
	)
	// 队列绑定
	err = ch.QueueBind(
		q.Name,        //队列的名字
		"err",         //routing key
		"logs_direct", //所绑定的交换器
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// 消费者进行消费
	msgs, err := ch.Consume(
		q.Name, //队列的名称
		"",     //消费者的名称
		false,  // 是否自动应答
		false,  //是否排他
		false,  // 如果设置为true，则生产者与消费者不能是同一连接
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", d.Body)
			str := fmt.Sprintf("%s", d.Body)
			t := time.Duration(len(str) - 70)
			time.Sleep(t * time.Second)
			// 开始手动应答
			d.Acknowledger.Ack(d.DeliveryTag, true)
		}
	}()

	log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	<-forever
}
