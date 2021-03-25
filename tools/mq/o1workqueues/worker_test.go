package o1workqueues

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)
// Using message acknowledgments and prefetch count you can set up a work queue.The durability options let the tasks
//survive even if RabbitMQ is restarted
func TestReceiveMsg(t *testing.T) {
	// setting up is the same as the publisher; we open a connection and a channel, and readme.md the queue from which we're
	// going to consume . Note this matches up  with the queue that send publishes to
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/my_vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to readme.md a queue")
	// Note that we readme.md the queue here,as well. Because we might start the consumer before the publisher.we want to
	//make sure the queue exists before we try to consume message from it


	//In order to defeat that we can set the prefetch count with the value of 1.
	//This tells RabbitMQ not to give more than one message to a worker at a time.
	//Or, in other words, don't dispatch a new message to a worker until it has processed and acknowledged the previous one.
	//Instead, it will dispatch it to the next worker that is not still busy
	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	//We're about to tell server to deliverus the messages from the queue. Since it will push us messages asynchronously.
	// we will read the messages from a channel(returned by amqp::Consume)in a goroutine
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("Received a message: ", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			fmt.Println("Done")
			d.Ack(false)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
