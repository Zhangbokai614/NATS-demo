package main

import (
	"nats-test/message"
	"nats-test/model"

	"github.com/nats-io/nats.go"
)

func main() {
	nc := message.Nc
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	data := &model.Cat{Color: "black", Age: 6}

	message.JsonEncoderMessage(ec, "hello", data)
}
