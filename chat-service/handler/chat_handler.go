package handler

import (
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"context"
	"path"
	"time"

	"proto/chat"

	"github.com/2637309949/micro/v3/service/events"
	"github.com/2637309949/micro/v3/service/store"
	"github.com/google/uuid"
)

const (
	chatStoreKeyPrefix    = "chats/"
	chatEventKeyPrefix    = "chats/"
	messageStoreKeyPrefix = "messages/"
)

func (h *Handler) Create(ctx context.Context, req *chat.CreateRequest, rsp *chat.CreateResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// generate a unique id for the chat
	groupId := uuid.New().String()

	// create a new group
	group := &chat.Group{
		Id:          groupId,
		Name:        req.Name,
		Description: req.Description,
		UserIds:     req.UserIds,
		Private:     req.Private,
		CreatedAt:   time.Now().Format(time.RFC3339Nano),
	}

	// key to lookup the chat in the store using, e.g. "chat/usera-userb-userc"
	key := path.Join(chatStoreKeyPrefix, groupId)

	// create a new record for the group
	rec := store.NewRecord(key, group)

	// write a record for the new group
	if err := store.Write(rec); err != nil {
		logger.Errorf(ctx, "Error writing to the store. Key: %v. Error: %v", key, err)
		return errors.InternalServerError("chat.new", "error creating chat group")
	}

	// return the group
	rsp.Group = group

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *chat.DeleteRequest, rsp *chat.DeleteResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.delete", "missing group id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	recs, err := store.Read(key, store.ReadLimit(1))
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.delete", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.delete", "error reading chat group")
	}

	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.delete", "error reading chat group")
	}
	// set response
	rsp.Group = group

	// delete the group
	if err := store.Delete(key); err != nil {
		return errors.InternalServerError("chat.delete", "error deleting chat group")
	}

	// get all messages
	// TODO: paginate the list
	key = path.Join(messageStoreKeyPrefix, req.GroupId)
	srecs, err := store.List(store.ListPrefix(key))
	if err != nil {
		return errors.InternalServerError("chat.delete", "failed to list messages")
	}

	// delete all the messages
	for _, rec := range srecs {
		if err := store.Delete(rec); err != nil {
			return errors.InternalServerError("chat.delete", "failed to list messages")
		}
	}

	// TODO: notify users of the event that the group is deleted

	return nil
}

func (h *Handler) List(ctx context.Context, req *chat.ListRequest, rsp *chat.ListResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	key := chatStoreKeyPrefix + "/"

	// read all the groups from the store for the user
	recs, err := store.Read(key, store.ReadPrefix())
	if err != nil {
		return errors.InternalServerError("chat.list", "error listing chat groups")
	}

	// list all the groups
	for _, rec := range recs {
		group := new(chat.Group)
		err := rec.Decode(group)
		if err != nil {
			continue
		}

		if len(req.UserId) == 0 {
			rsp.Groups = append(rsp.Groups, group)
			continue
		}

		// check if there's a user id match
		for _, user := range group.UserIds {
			if user == req.UserId {
				rsp.Groups = append(rsp.Groups, group)
				break
			}
		}
	}

	return nil
}

// History returns the historical messages in a chat
func (h *Handler) History(ctx context.Context, req *chat.HistoryRequest, rsp *chat.HistoryResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.history", "missing group id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	if _, err := store.Read(key); err == store.ErrNotFound {
		return errors.BadRequest("chat.history", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.history", "error reading chat group")
	}

	// lookup the messages
	key = path.Join(messageStoreKeyPrefix, req.GroupId)
	recs, err := store.Read(key+"/", store.ReadPrefix())
	if err != nil {
		logger.Errorf(ctx, "Error reading messages the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.history", "failed to read messages")
	}

	for _, rec := range recs {
		msg := new(chat.Message)
		err := rec.Decode(msg)
		if err != nil {
			return errors.InternalServerError("chat.history", "failed to decode message")
		}
		rsp.Messages = append(rsp.Messages, msg)
	}

	return nil
}

func (h *Handler) Invite(ctx context.Context, req *chat.InviteRequest, rsp *chat.InviteResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.invite", "missing group id")
	}

	if len(req.UserId) == 0 {
		return errors.BadRequest("chat.invite", "missing user id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.invite", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.invite", "error reading chat group")
	}

	// check the user is in the group
	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.invite", "Error reading group")
	}

	var exists bool

	// check the user is in the group
	for _, user := range group.UserIds {
		if user == req.UserId {
			exists = true
			break
		}
	}

	// TODO: send join message
	if !exists {
		group.UserIds = append(group.UserIds, req.UserId)
		// write the record
		rec := store.NewRecord(key, group)
		if err := store.Write(rec); err != nil {
			return errors.InternalServerError("chat.invite", "Error adding user to group")
		}
	}

	rsp.Group = group

	return nil
}

// Send a single message to the chat, designed for ease of use via the API / CLI
func (h *Handler) Send(ctx context.Context, req *chat.SendRequest, rsp *chat.SendResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.send", "missing group id")
	}
	if len(req.UserId) == 0 {
		return errors.BadRequest("chat.send", "missing user id")
	}
	if len(req.Text) == 0 {
		return errors.BadRequest("chat.send", "missing text")
	}

	// check the group exists
	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat group from the store to ensure it's valid
	recs, err := store.Read(key, store.ReadLimit(1))
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.send", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.send", "error reading chat group")
	}

	// decode the group
	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.send", "error reading chat group")
	}

	var exists bool

	// check the user is in the group
	for _, user := range group.UserIds {
		if user == req.UserId {
			exists = true
			break
		}
	}

	if !exists {
		return errors.BadRequest("chat.send", "user is not in the group")
	}

	// construct the message
	msg := &chat.Message{
		Id:      uuid.New().String(),
		Client:  req.Client,
		GroupId: req.GroupId,
		UserId:  req.UserId,
		Subject: req.Subject,
		Text:    req.Text,
		SentAt:  time.Now().Format(time.RFC3339Nano),
	}

	// default the client id if not provided
	if len(msg.Client) == 0 {
		msg.Client = uuid.New().String()
	}

	// create the message
	if err := h.createMessage(msg); err != nil {
		return err
	}

	// return the response
	rsp.Message = msg

	return nil
}

func (h *Handler) Join(ctx context.Context, req *chat.JoinRequest, stream chat.Chat_JoinStream) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.send", "missing group id")
	}
	if len(req.UserId) == 0 {
		return errors.BadRequest("chat.send", "missing user id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	recs, err := store.Read(key, store.ReadLimit(1))
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.join", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.join", "Error reading group")
	}

	// check the user is in the group
	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.join", "Error reading group")
	}

	var exists bool

	// check the user is in the group
	for _, user := range group.UserIds {
		if user == req.UserId {
			exists = true
			break
		}
	}

	// TODO: send join message
	if !exists {
		group.UserIds = append(group.UserIds, req.UserId)
		// write the record
		rec := store.NewRecord(key, group)
		if err := store.Write(rec); err != nil {
			return errors.InternalServerError("chat.join", "Error adding user to group")
		}
	}

	// create a channel to send errors on, because the subscriber / publisher will run in seperate go-
	// routines, they need a way of returning errors to the client
	errChan := make(chan error)

	eventKey := path.Join(chatEventKeyPrefix, req.GroupId)

	// create an event stream to consume messages posted by other users into the chat. we'll use the
	// user id as a queue to ensure each user recieves the message
	evStream, err := events.Consume(eventKey, events.WithGroup(req.UserId), events.WithContext(ctx))
	if err != nil {
		logger.Errorf(ctx, "Error streaming events. Group ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.join", "Error joining the group")
	}

	for {
		select {
		case <-ctx.Done():
			// the context has been cancelled or timed out, stop subscribing to new messages
			return nil
		case ev := <-evStream:
			// recieved a message, unmarshal it into a message struct. if an error occurs log it and
			// cancel the context
			var msg chat.Message
			if err := ev.Unmarshal(&msg); err != nil {
				logger.Errorf(ctx, "Error unmarshaling message. Group ID: %v. Error: %v", req.GroupId, err)
				errChan <- err
				return nil
			}

			// ignore any messages published by the current user
			if msg.UserId == req.UserId {
				continue
			}

			// publish the message to the stream
			if err := stream.Send(&chat.JoinResponse{Message: &msg}); err != nil {
				logger.Errorf(ctx, "Error sending message to stream. ChatID: %v. Message ID: %v. Error: %v", msg.GroupId, msg.Id, err)
				errChan <- err
				return nil
			}
		}
	}

	return nil
}

func (h *Handler) Kick(ctx context.Context, req *chat.KickRequest, rsp *chat.KickResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.kick", "missing group id")
	}
	if len(req.UserId) == 0 {
		return errors.BadRequest("chat.kick", "missing user id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	recs, err := store.Read(key, store.ReadLimit(1))
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.kick", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Chat ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.kick", "Error reading group")
	}

	// check the user is in the group
	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.kick", "Error reading group")
	}

	var users []string

	// check the user is in the group
	for _, user := range group.UserIds {
		if user == req.UserId {
			continue
		}
		users = append(users, user)
	}

	group.UserIds = users

	rec := store.NewRecord(key, group)
	if err := store.Write(rec); err != nil {
		return errors.InternalServerError("chat.kick", "Error leaveing from group")
	}

	// TODO: send leave message
	// TODO: disconnect the actual event consumption
	rsp.Group = group

	return nil
}
func (h *Handler) Leave(ctx context.Context, req *chat.LeaveRequest, rsp *chat.LeaveResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	// validate the request
	if len(req.GroupId) == 0 {
		return errors.BadRequest("chat.leave", "missing group id")
	}
	if len(req.UserId) == 0 {
		return errors.BadRequest("chat.leave", "missing user id")
	}

	key := path.Join(chatStoreKeyPrefix, req.GroupId)

	// lookup the chat from the store to ensure it's valid
	recs, err := store.Read(key, store.ReadLimit(1))
	if err == store.ErrNotFound {
		return errors.BadRequest("chat.leave", "group not found")
	} else if err != nil {
		logger.Errorf(ctx, "Error reading from the store. Chat ID: %v. Error: %v", req.GroupId, err)
		return errors.InternalServerError("chat.leave", "Error reading group")
	}

	// check the user is in the group
	group := new(chat.Group)
	err = recs[0].Decode(group)
	if err != nil {
		return errors.InternalServerError("chat.leave", "Error reading group")
	}

	var users []string

	// check the user is in the group
	for _, user := range group.UserIds {
		if user == req.UserId {
			continue
		}
		users = append(users, user)
	}

	group.UserIds = users

	rec := store.NewRecord(key, group)
	if err := store.Write(rec); err != nil {
		return errors.InternalServerError("chat.leave", "Error leaveing from group")
	}

	// TODO: send leave message
	// TODO: disconnect the actual event consumption
	rsp.Group = group

	return nil
}

// createMessage is a helper function which creates a message in the event stream. It handles the
// logic for ensuring client id is unique.
func (h *Handler) createMessage(msg *chat.Message) error {
	storekey := path.Join(messageStoreKeyPrefix, msg.GroupId, msg.Id)
	eventKey := path.Join(chatEventKeyPrefix, msg.GroupId)

	// send the message to the event stream
	if err := events.Publish(eventKey, msg); err != nil {
		return err
	}

	// create a new record
	rec := store.NewRecord(storekey, msg)

	// record the messages client id
	return store.Write(rec)
}
