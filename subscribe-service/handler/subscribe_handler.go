package handler

import (
	"comm/auth"
	"comm/broker"
	"comm/errors"
	"comm/logger"
	"context"
	"path"
	"proto/subscribe"
)

func (h *Handler) Publish(ctx context.Context, req *subscribe.PublishRequest, rsp *subscribe.PublishResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Call", acc.Name)
	}
	if len(req.Topic) == 0 {
		return errors.BadRequest("topic is blank")
	}

	logger.Infof(ctx, "Publishing to %v\n", req.Topic)

	topic := path.Join("event", req.Topic)
	err := broker.Publish(topic, &broker.Message{Body: req.Message})
	if err != nil {
		return errors.InternalServerError("Failed to publish %v", err.Error())
	}
	return nil
}

// Subscribe defined TODO, setting timeout should be careful
func (h *Handler) Subscribe(ctx context.Context, req *subscribe.SubscribeRequest, stream subscribe.Subscribe_SubscribeStream) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Subscribe", acc.Name)
	}
	if len(req.Topic) == 0 {
		return errors.BadRequest("topic is blank")
	}

	logger.Infof(ctx, "Subscribing to %v\n", req.Topic)

	topic := path.Join("event", req.Topic)
	ch := make(chan *broker.Message, 100)
	_, err := broker.Subscribe(topic, func(msg *broker.Message) error {
		ch <- msg
		return nil
	})
	if err != nil {
		return errors.InternalServerError("Failed to subscribe %v", err.Error())
	}
	for m := range ch {
		if err := stream.Send(&subscribe.SubscribeResponse{
			Topic:   req.Topic,
			Message: m.Body,
		}); err != nil {
			logger.Infof(ctx, "Failed to Send %v", err)
		}
	}
	return nil
}
