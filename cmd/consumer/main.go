package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"realTimeCryptoPrice/pkg/rabbitmq"
)

func main(queue, consumer string) {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, queue, consumer)
}
