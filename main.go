package main

import (
	"context"
	"github.com/skullkon/go-kafka/consumer"
	"github.com/skullkon/go-kafka/producer"
)

func main() {
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go producer.Produce(ctx)
	consumer.Consume(ctx)
}
