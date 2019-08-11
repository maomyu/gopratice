package main

import (
	"fmt"
	"log"

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

	q, err := ch.QueueDeclare(
		"hello", //name
		true,    //durable
		false,   //delete when usused
		false,   // exclusive
		false,   //no-wait
		nil,     // arguments
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
			// 开始手动应答
			d.Acknowledger.Ack(d.DeliveryTag, true)
		}
	}()

	log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	<-forever
}
