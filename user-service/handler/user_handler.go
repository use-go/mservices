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

	users := []model.User{}
	err = h.QueryUserDB(ctx, session, &model.User{Username: req.Username}, &users)
	if err != nil && !errors.Is(err, errors.RecordNotFound) {
		return err
	}
	if len(users) > 0 {
		return errors.BadRequest("username already exists")
	}

	err = h.QueryUserDB(ctx, session, &model.User{Email: req.Email}, &users)
	if err != nil && !errors.Is(err, errors.RecordNotFound) {
		return err
	}
	if len(users) > 0 {
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
	rsp.Id = mu.Id
	return nil
}

func (h *Handler) Read(ctx context.Context, req *user.ReadRequest, rsp *user.ReadResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Read", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}
	switch {
	case req.Id > 0:
		where, user := model.User{Id: req.Id}, model.User{}
		err := h.QueryUserDetailDB(ctx, session, &where, &user)
		if err != nil {
			return err
		}
		copier.Copy(rsp, &user)
		return nil
	case req.Username != "" || req.Email != "":
		where, user := model.User{Username: req.Username}, model.User{Email: req.Email}
		err := h.QueryUserDetailDB(ctx, session, &where, &user)
		if err != nil {
			return err
		}
		copier.Copy(&rsp, &user)
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
		return errors.InternalServerError("init db failed %v", err)
	}

	err = h.UpdateUserDB(ctx, session, &model.User{
		Id:       req.Id,
		Username: strings.ToLower(req.Username),
		Email:    strings.ToLower(req.Email),
		Profile:  req.Profile,
	})
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err)
	}
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *user.DeleteRequest, rsp *user.DeleteResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	err = h.DeleteUserDB(ctx, session, &model.User{Id: req.Id})
	if err != nil {
		return errors.InternalServerError("DeleteUserDB failed %v", err.Error())
	}
	return nil
}

func (h *Handler) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, rsp *user.UpdatePasswordResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdatePassword", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Id: req.Id}, &usr)
	if err != nil {
		return errors.InternalServerError("QueryUserDetailDB %v", err.Error())
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.InternalServerError("passwords don't match")
	}

	if err := usr.CompareHashAndPassword(req.OldPassword); err != nil {
		return errors.InternalServerError("unauthorized")
	}

	if _, err = usr.GenerateFromPassword(req.NewPassword); err != nil {
		return errors.InternalServerError("generate password failed")
	}

	err = h.UpdateUserDB(ctx, session, &usr)
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err.Error())
	}
	return nil
}

func (h *Handler) List(ctx context.Context, req *user.ListRequest, rsp *user.ListResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do List", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}
	session = db.SetLimit(ctx, session, req)

	var list []model.User
	err = h.QueryUserDB(ctx, session, &model.User{}, &list)
	if err != nil {
		return errors.InternalServerError("QueryUserDB failed %v", err.Error())
	}
	return nil
}

func (h *Handler) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest, rsp *user.ResetPasswordResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do ResetPassword", acc.Name)
	}

	if len(req.Email) == 0 {
		return errors.BadRequest("missing email")
	}
	if len(req.Code) == 0 {
		return errors.BadRequest("missing code")
	}
	if len(req.ConfirmPassword) == 0 {
		return errors.BadRequest("missing confirm password")
	}
	if len(req.NewPassword) == 0 {
		return errors.BadRequest("missing new password")
	}
	if req.ConfirmPassword != req.NewPassword {
		return errors.BadRequest("passwords do not match")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &usr)
	if err != nil {
		return errors.InternalServerError("QueryUserDetailDB %v", err.Error())
	}

	// check if a request was made to reset the password, we should have saved it
	// code, err := s.domain.ReadPasswordResetCode(ctx, usr.Id, req.Code)
	// if err != nil {
	// 	return err
	// }

	// validate the code, e.g its an OTP token and hasn't expired
	// resp, err := s.Otp.Validate(ctx, &otp.ValidateRequest{
	// 	Id:   req.Email,
	// 	Code: req.Code,
	// })
	// if err != nil {
	// 	return err
	// }

	// check if the code is actually valid
	// if !resp.Success {
	// 	return errors.BadRequest("user.resetpassword", "invalid code")
	// }

	if _, err = usr.GenerateFromPassword(req.NewPassword); err != nil {
		return errors.InternalServerError("generate password failed")
	}

	err = h.UpdateUserDB(ctx, session, &usr)
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err.Error())
	}
	return nil
}

func (h *Handler) SendPasswordResetEmail(ctx context.Context, req *user.SendPasswordResetEmailRequest, rsp *user.SendPasswordResetEmailResponse) error {
	return nil
}

func (h *Handler) SendVerificationEmail(ctx context.Context, req *user.SendVerificationEmailRequest, rsp *user.SendVerificationEmailResponse) error {
	return nil
}

func (h *Handler) VerifyEmail(ctx context.Context, req *user.VerifyEmailRequest, rsp *user.VerifyEmailResponse) error {
	if len(req.Email) == 0 {
		return errors.BadRequest("missing email")
	}
	if len(req.Token) == 0 {
		return errors.BadRequest("missing token")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	var usr model.User
	// err = h.QueryUserDetailDB(ctx, session, &model.User{Id: req.Id}, &usr)
	// if err != nil {
	// 	return errors.InternalServerError("QueryUserDetailDB %v", err.Error())
	// }

	usr.Verified = true
	err = h.UpdateUserDB(ctx, session, &usr)
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err.Error())
	}
	return nil
}
