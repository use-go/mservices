package broker

import (
	"github.com/2637309949/micro/v3/service/broker"
)

type (
	Message = broker.Message
	Handler = broker.Handler
)

func Publish(topic string, m *Message) error {
	return broker.Publish(topic, m)
}

func Subscribe(topic string, h Handler) (broker.Subscriber, error) {
	return broker.Subscribe(topic, h)
}
