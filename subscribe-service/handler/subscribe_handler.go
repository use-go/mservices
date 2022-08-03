package handler

import (
	"comm/auth"
	"comm/broker"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"context"
	"path"
	"proto/subscribe"
)

func (h *Handler) Publish(ctx context.Context, req *subscribe.PublishRequest, rsp *subscribe.PublishResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Publish")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Call", acc.Name)
	}
	if len(req.Topic) == 0 {
		return errors.BadRequest("topic is blank")
	}

	logger.Infof(ctx, "Publishing to %v\n", req.Topic)

	topic := path.Join("event", req.Topic)
	err = broker.Publish(topic, &broker.Message{Body: req.Message})
	timemark.Mark("Publish")
	if err != nil {
		return errors.InternalServerError("Failed to publish %v", err.Error())
	}
	return nil
}

// Subscribe defined TODO, setting timeout should be careful
func (h *Handler) Subscribe(ctx context.Context, req *subscribe.SubscribeRequest, stream subscribe.Subscribe_SubscribeStream) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Subscribe")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Subscribe", acc.Name)
	}
	if len(req.Topic) == 0 {
		return errors.BadRequest("topic is blank")
	}

	logger.Infof(ctx, "Subscribing to %v\n", req.Topic)

	topic := path.Join("event", req.Topic)
	sub := make(chan *broker.Message, 100)
	bs, err := broker.Subscribe(topic, func(msg *broker.Message) error {
		sub <- msg
		return nil
	})
	defer bs.Unsubscribe()
	if err != nil {
		return errors.InternalServerError("Failed to subscribe %v", err.Error())
	}
	for msg := range sub {
		err := stream.Send(&subscribe.SubscribeResponse{
			Topic:   req.Topic,
			Message: msg.Body,
		})
		if err != nil {
			logger.Infof(ctx, "Failed to Send %v", err)
		}
	}
	return nil
}
