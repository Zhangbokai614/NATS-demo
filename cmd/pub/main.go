package main

import (
	"fmt"
	"log"
	"nats-test/message"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc := message.Nc
	defer nc.Close()

	tryCount := 1000

	wgCount := tryCount * 2
	wg := sync.WaitGroup{}
	wg.Add(wgCount)

	replyMap := make(map[string]int)

	if _, err := nc.Subscribe("reply", func(msg *nats.Msg) {
		key := string(msg.Data)
		if _, exists := replyMap[key]; !exists {
			replyMap[key] = 1
		} else {
			replyMap[key] += 1
		}

		wg.Done()
	}); err != nil {
		panic(err)
	}

	for i := 0; i < tryCount; i++ {
		err := message.PublishMessage(nc, "hello", "Transistor")
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Wait()

	for key := range replyMap {
		fmt.Println(key, ":", replyMap[key])
	}
}
