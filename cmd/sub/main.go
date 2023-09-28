package main

import (
	"fmt"
	"log"
	"nats-test/message"
	"sync"
)

func main() {
	nc := message.Nc
	defer nc.Close()

	wc := 3
	wg := sync.WaitGroup{}

	for i := 0; i < wc; i++ {
		keyA := fmt.Sprintf("hello-queue.A.A-%d", i)
		keyB := fmt.Sprintf("Hello-queue.*.A-%d", i)

		wg.Add(1)
		go func() {
			if err := message.AsyncQueueSub(nc, &wg, "hello", "queue.A.A", keyA); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("sub: %s \n", keyA)
		}()

		wg.Add(1)
		go func() {
			if err := message.AsyncQueueSub(nc, &wg, "hello", "queue.*.A", keyB); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("sub: %s \n", keyB)
		}()
	}

	wg.Wait()
}
