package main

import (
	"context"
	c "kafka-golang/consumer"
	p "kafka-golang/producer"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go p.Produce(ctx)
	c.Consume(ctx)
}
