package message

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

const (
	natsAddr = "192.168.2.251"
	natsPort = "4223"
)

var (
	Nc *nats.Conn
)

func init() {
	natsUrl := fmt.Sprintf("%s:%s", natsAddr, natsPort)

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}

	Nc = nc
}
