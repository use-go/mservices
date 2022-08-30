package handler

import (
	"context"
	"path"
	"time"

	"comm/auth"
	"comm/errors"
	"comm/logger"

	"proto/event"

	"github.com/2637309949/micro/v3/service/events"
	"google.golang.org/protobuf/types/known/structpb"
)

// DeleteInfo defined TODO
func (h *Handler) Publish(ctx context.Context, req *event.PublishRequest, rsp *event.PublishResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Publish", acc.Name)
	}

	if len(req.Topic) == 0 {
		return errors.BadRequest("event.publish", "topic is blank")
	}

	// create tenant based topics
	topic := path.Join("event", req.Topic)
	// publish the message
	return events.Publish(topic, req.Message.AsMap())
}

func (h *Handler) Consume(ctx context.Context, req *event.ConsumeRequest, stream event.Event_ConsumeStream) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Consume", acc.Name)
	}

	// create tenant based topics
	topic := path.Join("event", req.Topic)

	logger.Infof(ctx, "Subscribing to %v\n", req.Topic)

	// check if a group os provided
	opts := []events.ConsumeOption{}
	offset := time.Now()
	if len(req.Group) > 0 {
		opts = append(opts, events.WithGroup(req.Group))
	}
	if len(req.Offset) > 0 {
		t, err := time.Parse(time.RFC3339Nano, req.Offset)
		if err == nil {
			offset = t
		}
	}
	opts = append(opts, events.WithOffset(offset))

	sub, err := events.Consume(topic, opts...)
	if err != nil {
		return errors.InternalServerError("event.subscribe", "failed to subscribe to event")
	}

	// range over the messages until the subscriber is closed
	for msg := range sub {
		logger.Info(ctx, "got message, sending")
		// unmarshal the message into a struct
		d := &structpb.Struct{}
		d.UnmarshalJSON(msg.Payload)

		if err := stream.Send(&event.ConsumeResponse{
			Topic:     req.Topic,
			Id:        msg.ID,
			Timestamp: msg.Timestamp.Format(time.RFC3339Nano),
			Message:   d,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) Read(ctx context.Context, req *event.ReadRequest, rsp *event.ReadResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Read", acc.Name)
	}

	// create tenant based topics
	topic := path.Join("event", req.Topic)

	//limit := uint(25)
	//offset := uint(0)
	var opts []events.ReadOption

	if req.Limit > 0 {
		opts = append(opts, events.ReadLimit(uint(req.Limit)))
	}

	if req.Offset > 0 {
		opts = append(opts, events.ReadOffset(uint(req.Offset)))
	}

	logger.Infof(ctx, "Reading %v limit: %v offset: %v\n", req.Topic, req.Limit, req.Offset)

	events, err := events.Read(topic, opts...)
	if err != nil {
		return err
	}

	logger.Infof(ctx, "Events read %v", len(events))

	for _, ev := range events {
		// unmarshal the message into a struct
		d := &structpb.Struct{}
		d.UnmarshalJSON(ev.Payload)

		rsp.Events = append(rsp.Events, &event.Ev{
			Id:        ev.ID,
			Timestamp: ev.Timestamp.Format(time.RFC3339Nano),
			Message:   d,
		})
	}

	return nil
}
