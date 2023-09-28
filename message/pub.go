package message

import (
	"time"

	"github.com/nats-io/nats.go"
)

func PublishMessage(nc *nats.Conn, subject, msg string) error {
	if err := nc.Publish(subject, []byte(msg)); err != nil {
		return err
	}

	return nil
}

func RequestMessage(nc *nats.Conn, subject, msg string) (string, error) {
	replyMsg, err := nc.Request(subject, []byte(msg), 5*time.Second)
	if err != nil {
		return "", err
	}

	return string(replyMsg.Data), nil
}

func JsonEncoderMessage(ec *nats.EncodedConn, subject string, msg interface{}) error {
	if err := ec.Publish(subject, msg); err != nil {
		return err
	}

	return nil
}
