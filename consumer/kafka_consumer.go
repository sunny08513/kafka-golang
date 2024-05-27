package consumer

import (
	"context"
	"fmt"
	c "kafka-golang/constant"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.Broker1Address, c.Broker2Address},
		Topic:   c.Topic,
		GroupID: "my-group",
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
