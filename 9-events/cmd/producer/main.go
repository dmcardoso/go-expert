package main

import (
	"github.com/dmcardoso/go-expert/9-events/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World 2!", "amq.direct")
}
