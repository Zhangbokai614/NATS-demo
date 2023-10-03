package main

import (
	"nats-test/message"
	"time"
)

const max = 6
const limit = 3

var count = 0

func main() {
	nc := message.Nc
	defer nc.Close()

	for ; count < limit; count++ {
		if err := message.PublishMessage(nc, "hello", "Transistor"); err != nil {
			panic(err)
		}
	}

	if err := message.PublishMessage(nc, "done", ""); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	for ; count < max; count++ {
		if err := message.PublishMessage(nc, "hello", "Transistor"); err != nil {
			panic(err)
		}
	}
}
