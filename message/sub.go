package message

import (
	"fmt"
	"nats-test/model"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func SyncSub(nc *nats.Conn, subject string) error {
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		return err
	}

	msg, err := sub.NextMsg(5 * time.Second)
	if err != nil {
		return err
	}

	fmt.Println(string(msg.Data))

	return nil
}

func SyncSubAutoUnsubscribe(nc *nats.Conn, subject string) error {
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		return err
	}

	sub.AutoUnsubscribe(6)

	for {
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			return err
		}

		fmt.Println(string(msg.Data))
	}
}

func AsyncSub(nc *nats.Conn, subject string) error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))

		wg.Done()
	}); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func JsonEncoderAsyncSub(ec *nats.EncodedConn, subject string, payload model.Payload) error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := ec.Subscribe(subject, func(msg map[string]interface{}) {
		payload.Bind(msg)

		wg.Done()
	}); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func AsyncSubDrain(nc *nats.Conn, subject string) error {
	if _, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("work start", "is closed:", nc.IsClosed(), "is draining:", nc.IsDraining())
		time.Sleep(2 * time.Second)
		fmt.Println(string(msg.Data))
	}); err != nil {
		return err
	}

	return nil
}

func AsyncSubReply(nc *nats.Conn, wg *sync.WaitGroup, subject, reply string) error {
	if _, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("subMsg:", string(msg.Data))

		if replyErr := nc.Publish(msg.Reply, []byte(reply)); replyErr != nil {
			fmt.Println(replyErr)
		}
	}); err != nil {
		wg.Done()
		return err
	}

	return nil
}

func AsyncQueueSub(nc *nats.Conn, wg *sync.WaitGroup, subject, queue, key string) error {
	if _, err := nc.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		fmt.Println("subMsg:", string(msg.Data))

		if err := nc.Publish("reply", []byte(key)); err != nil {
			wg.Done()
		}
	}); err != nil {
		wg.Done()

		return err
	}

	return nil
}
