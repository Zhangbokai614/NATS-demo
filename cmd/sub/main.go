package main

import (
	"fmt"
	"nats-test/message"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc := message.Nc
	defer nc.Close()

	if err := message.AsyncSubDrain(nc, "hello"); err != nil {
		panic(err)
	}

	if _, err := nc.Subscribe("done", func(msg *nats.Msg) {
		fmt.Println("start close")
		nc.Close()
		fmt.Println("conn close")

	}); err != nil {
		panic(err)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("is closed:", nc.IsClosed())
	fmt.Println("is draining:", nc.IsDraining())
}
