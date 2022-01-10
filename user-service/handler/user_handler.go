package handler

import (
	"comm/auth"
	"comm/db"
	"comm/errors"
	"comm/logger"
	"context"
	"proto/user"
	"regexp"
	"strings"
	"time"
	"user-service/model"

	"github.com/jinzhu/copier"
)

var (
	emailFormat = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

func (h *Handler) Create(ctx context.Context, req *user.CreateRequest, rsp *user.CreateResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Create", acc.Name)
	}

	if !emailFormat.MatchString(req.Email) {
		return errors.BadRequest("email has wrong format")
	}
	if len(req.Password) < 8 {
		return errors.InternalServerError("password is less than 8 characters")
	}
	req.Username = strings.ToLower(req.Username)
	req.Email = strings.ToLower(req.Email)

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	where, user := model.User{Username: req.Username}, model.User{}
	err = h.QueryUserDetailDB(ctx, session, &where, &user)
	if err != nil && !errors.Is(err, errors.RecordNotFound) {
		return err
	}
	if user.Id > 0 {
		return errors.BadRequest("username already exists")
	}

	where = model.User{Email: req.Email}
	err = h.QueryUserDetailDB(ctx, session, &where, &user)
	if err != nil && !errors.Is(err, errors.RecordNotFound) {
		return err
	}
	if user.Id > 0 {
		return errors.BadRequest("email already exists")
	}

	mu := model.User{
		Username: req.Username,
		Email:    req.Email,
		Profile:  req.Profile,
		Created:  uint32(time.Now().Unix()),
		Updated:  uint32(time.Now().Unix()),
	}
	mu.GenerateFromPassword(req.Password)
	err = h.InsertUserDB(ctx, session, &mu)
	if err != nil {
		return err
	}

	copier.Copy(&rsp.Account, &acc)
	return nil
}

func (h *Handler) Read(ctx context.Context, req *user.ReadRequest, rsp *user.ReadResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Read", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}
	switch {
	case req.Id > 0:
		where, user := model.User{Id: req.Id}, model.User{}
		err := h.QueryUserDetailDB(ctx, session, &where, &user)
		if err != nil {
			return err
		}
		copier.Copy(&rsp.Account, &user)
		return nil
	case req.Username != "" || req.Email != "":
		where, user := model.User{Username: req.Username}, model.User{Email: req.Email}
		err := h.QueryUserDetailDB(ctx, session, &where, &user)
		if err != nil {
			return err
		}
		copier.Copy(&rsp.Account, &user)
		return nil
	}
	return nil
}

func (h *Handler) Update(ctx context.Context, req *user.UpdateRequest, rsp *user.UpdateResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Update", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}
	return h.UpdateUserDB(ctx, session, &model.User{
		Id:       req.Id,
		Username: strings.ToLower(req.Username),
		Email:    strings.ToLower(req.Email),
		Profile:  req.Profile,
	})
}

func (h *Handler) Delete(ctx context.Context, req *user.DeleteRequest, rsp *user.DeleteResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Update", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	return h.DeleteUserDB(ctx, session, &model.User{Id: req.Id})
}

// func (h *Handler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest, rsp *pb.UpdatePasswordResponse) error {
// 	usr, err := s.domain.Read(ctx, req.UserId)
// 	if err != nil {
// 		return errors.InternalServerError("user.updatepassword", err.Error())
// 	}
// 	if req.NewPassword != req.ConfirmPassword {
// 		return errors.InternalServerError("user.updatepassword", "Passwords don't match")
// 	}

// 	salt, hashed, err := s.domain.SaltAndPassword(ctx, usr.Id)
// 	if err != nil {
// 		return errors.InternalServerError("user.updatepassword", err.Error())
// 	}

// 	hh, err := base64.StdEncoding.DecodeString(hashed)
// 	if err != nil {
// 		return errors.InternalServerError("user.updatepassword", err.Error())
// 	}

// 	if err := bcrypt.CompareHashAndPassword(hh, []byte(x+salt+req.OldPassword)); err != nil {
// 		return errors.Unauthorized("user.updatepassword", err.Error())
// 	}

// 	salt = random(16)
// 	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.NewPassword), 10)
// 	if err != nil {
// 		return errors.InternalServerError("user.updatepassword", err.Error())
// 	}
// 	pp := base64.StdEncoding.EncodeToString(h)

// 	if err := s.domain.UpdatePassword(ctx, req.UserId, salt, pp); err != nil {
// 		return errors.InternalServerError("user.updatepassword", err.Error())
// 	}
// 	return nil
// }
