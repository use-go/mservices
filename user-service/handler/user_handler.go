package handler

import (
	"comm/auth"
	"comm/db"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"context"
	"net/url"
	"proto/email"
	"proto/user"
	"strings"
	"time"
	"user-service/model"

	"github.com/jinzhu/copier"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateRequest, rsp *user.CreateResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Create")()

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

	if len(req.Username) < 8 {
		return errors.BadRequest("missing username")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	user := model.User{}
	err = h.QueryUserDetailDB(ctx, session, &model.User{Username: req.Username}, &user)
	timemark.Mark("QueryUserDetailDB")
	if err == nil {
		return errors.BadRequest("username already exists")
	}

	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &user)
	timemark.Mark("QueryUserDetailDB")
	if err == nil {
		return errors.BadRequest("email already exists")
	}

	item := model.User{
		Username: req.Username,
		Email:    req.Email,
		Profile:  req.Profile,
		Created:  uint32(time.Now().Unix()),
		Updated:  uint32(time.Now().Unix()),
	}
	item.GenerateFromPassword(req.Password)
	timemark.Mark("GenerateFromPassword")
	err = h.InsertUserDB(ctx, session, &item)
	if err != nil {
		return err
	}
	rsp.Id = item.Id
	return nil
}

func (h *Handler) Read(ctx context.Context, req *user.ReadRequest, rsp *user.ReadResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Read")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Read", acc.Name)
	}

	if req.Id == 0 && len(req.Username) == 0 && len(req.Email) == 0 {
		return errors.BadRequest("missing id or username or email")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	where, user := model.User{Id: req.Id, Username: req.Username, Email: req.Email}, model.User{}
	err = h.QueryUserDetailDB(ctx, session, &where, &user)
	if err != nil {
		return err
	}
	copier.Copy(rsp, &user)
	return nil
}

func (h *Handler) Update(ctx context.Context, req *user.UpdateRequest, rsp *user.UpdateResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Update")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Update", acc.Name)
	}

	if req.Id == 0 {
		return errors.BadRequest("missing id")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	if len(req.Username) > 0 {
		usr := model.User{}
		session = session.Where("id!=?", req.Id)
		err = h.QueryUserDetailDB(ctx, session, &model.User{Username: req.Username}, &usr)
		timemark.Mark("QueryUserDetailDB")
		if err == nil {
			return errors.BadRequest("username already exists")
		}
	}

	if len(req.Email) > 0 {
		usr := model.User{}
		session = session.Where("id!=?", req.Id)
		err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &usr)
		timemark.Mark("QueryUserDetailDB")
		if err == nil {
			return errors.BadRequest("email already exists")
		}
	}

	err = h.UpdateUserDB(ctx, session, &model.User{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Profile:  req.Profile,
	})
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err)
	}
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *user.DeleteRequest, rsp *user.DeleteResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Delete")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete", acc.Name)
	}

	if req.Id == 0 {
		return errors.BadRequest("missing id")
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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "UpdatePassword")()

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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "List")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do List", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}
	session = db.SetLimit(ctx, session, req)
	session = db.SetOrder(ctx, session, req)

	list := []model.User{}
	where := model.User{
		Username: req.Username,
		Email:    req.Email,
	}
	err = h.QueryUserDB(ctx, session, &where, &list)
	if err != nil {
		return errors.InternalServerError("QueryUserDB failed %v", err.Error())
	}
	return nil
}

// ResetPassword defined todo
func (h *Handler) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest, rsp *user.ResetPasswordResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "ResetPassword")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do ResetPassword", acc.Name)
	}

	if len(req.Email) == 0 {
		logger.Error(ctx, "missing email")
		return errors.BadRequest("missing email")
	}

	if len(req.Code) == 0 {
		logger.Error(ctx, "missing code")
		return errors.BadRequest("missing code")
	}

	if len(req.ConfirmPassword) == 0 {
		logger.Error(ctx, "missing confirm password")
		return errors.BadRequest("missing confirm password")
	}

	if len(req.NewPassword) == 0 {
		logger.Error(ctx, "missing new password")
		return errors.BadRequest("missing new password")
	}

	if req.ConfirmPassword != req.NewPassword {
		logger.Error(ctx, "password do not match")
		return errors.BadRequest("password do not match")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb fail %v", err)
		return err
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &usr)
	if err != nil {
		return err
	}

	// validate the code
	_, err = h.ReadPasswordResetCode(ctx, usr.Id, req.Code)
	if err != nil {
		return err
	}

	// no error means it exists and not expired
	if _, err = usr.GenerateFromPassword(req.NewPassword); err != nil {
		return errors.InternalServerError("generate password failed")
	}

	err = h.UpdateUserDB(ctx, session, &usr)
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err.Error())
	}

	// delete our saved code
	err = h.DeletePasswordResetCode(ctx, usr.Id, req.Code)
	if err != nil {
		return errors.InternalServerError("DeletePasswordResetCode failed %v", err.Error())
	}

	return nil
}

// SendPasswordResetEmail defined todo
func (h *Handler) SendPasswordResetEmail(ctx context.Context, req *user.SendPasswordResetEmailRequest, rsp *user.SendPasswordResetEmailResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "SendPasswordResetEmail")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do SendPasswordResetEmail", acc.Name)
	}

	if len(req.Email) == 0 {
		return errors.BadRequest("user.sendpasswordresetemail", "missing email")
	}
	if len(req.Subject) == 0 {
		return errors.BadRequest("user.sendpasswordresetemail", "missing subject")
	}
	if len(req.TextContent) == 0 {
		return errors.BadRequest("user.sendpasswordresetemail", "missing textContent")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb fail %v", err)
		return err
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &usr)
	if err != nil {
		return err
	}

	var expiry int64 = 1800 // 1800 secs = 30 min
	if req.Expiration > 0 {
		expiry = req.Expiration
	}

	if err != nil {
		return err
	}
	code := random(8)

	// save the password reset code
	_, err = h.SavePasswordResetCode(ctx, usr.Id, code, time.Duration(expiry)*time.Second)
	if err != nil {
		return err
	}

	// set the code in the text content
	req.TextContent = strings.Replace(req.TextContent, "$code", code, -1)
	sendRequest := email.SendRequest{
		To:       req.Email,
		Subject:  req.Subject,
		TextBody: req.TextContent,
	}
	_, err = h.EmailService.Send(ctx, &sendRequest)
	if err != nil {
		return err
	}
	return nil
}

// SendVerificationEmail defined todo
func (h *Handler) SendVerificationEmail(ctx context.Context, req *user.SendVerificationEmailRequest, rsp *user.SendVerificationEmailResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "SendVerificationEmail")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do SendVerificationEmail", acc.Name)
	}

	if len(req.Email) == 0 {
		return errors.BadRequest("user.sendverificationemail", "missing email")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb fail %v", err)
		return err
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: req.Email}, &usr)
	if err != nil {
		return err
	}

	// generate random token
	token := random(256)

	// generate/save a token for verification
	err = h.CreateToken(ctx, req.Email, token)
	if err != nil {
		return err
	}

	// set the code in the text content
	uri := "http://127.0.0.1:8080/xxx"
	query := "?token=" + token + "&redirectUrl=" + url.QueryEscape(req.RedirectUrl) + "&failureRedirectUrl=" + url.QueryEscape(req.FailureRedirectUrl)
	// set the text content
	req.TextContent = strings.Replace(req.TextContent, "$micro_verification_link", uri+query, -1)
	sendRequest := email.SendRequest{
		To:       req.Email,
		Subject:  req.Subject,
		TextBody: req.TextContent,
	}
	_, err = h.EmailService.Send(ctx, &sendRequest)
	if err != nil {
		return err
	}
	return nil
}

// VerifyEmail deifned todo
func (h *Handler) VerifyEmail(ctx context.Context, req *user.VerifyEmailRequest, rsp *user.VerifyEmailResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "SendPasswordResetEmail")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do VerifyEmail", acc.Name)
	}

	if len(req.Token) == 0 {
		return errors.BadRequest("missing token")
	}

	// check the token exists
	email, err := h.ReadToken(ctx, req.Token)
	if err != nil {
		logger.Error(ctx, "Failed to read token: %v", err)
		return err
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db failed %v", err)
	}

	var usr model.User
	err = h.QueryUserDetailDB(ctx, session, &model.User{Email: email}, &usr)
	if err != nil {
		return errors.InternalServerError("QueryUserDetailDB %v", err.Error())
	}

	t := time.Now().Unix()
	usr.Verified = true
	usr.Updated = uint32(t)
	usr.VerificationDate = uint32(t)
	err = h.UpdateUserDB(ctx, session, &usr)
	if err != nil {
		return errors.InternalServerError("UpdateUserDB failed %v", err.Error())
	}
	return nil
}

// ValidPassword defined todo
func (h *Handler) ValidPassword(ctx context.Context, req *user.ValidPasswordRequest, rsp *user.ValidPasswordResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "ValidPassword")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do ValidPassword", acc.Name)
	}

	if req.Id == 0 {
		return errors.BadRequest("missing id")
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

	err = usr.CompareHashAndPassword(req.Password)
	if err != nil {
		return errors.InternalServerError("unauthorized")
	}
	return nil
}
