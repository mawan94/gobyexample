package o1workqueues

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"strings"
	"testing"
)
func bodyFrom(args []string) string  {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	}else {
		s = strings.Join(args[1:], " ")
	}
	return s
}


// we need a helper function to check the return value for each amqp call
func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}

func TestSendMsg(t *testing.T) {


	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/my_vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create a channel ,which is where most of this API for getting things done resides
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// To send, we must readme.md a queue for us to send to; then we can publish msg to the queue
	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to readme.md a queue")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}
