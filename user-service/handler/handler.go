package handler

import (
	"comm/errors"
	"comm/service"
	"comm/store"
	"context"
	"crypto/rand"
	"fmt"
	"proto/email"
	"regexp"
	"strings"
	"time"
)

var (
	alphanum    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	emailFormat = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Handler struct {
	CacheService store.Cache
	EmailService email.EmailService
}

type verificationToken struct {
	UserID uint32 `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type passwordResetCode struct {
	Expires time.Time `json:"expires"`
	UserID  uint32    `json:"userId"`
	Code    string    `json:"code"`
}

func (h *Handler) ReadToken(ctx context.Context, token string) (string, error) {
	if len(token) == 0 {
		return "", errors.BadRequest(service.GetName(), "missing token")
	}
	key := generateVerificationTokenStoreKey(token)
	vt := verificationToken{}
	err := h.CacheService.Get(ctx, key, &vt)
	if err != nil {
		return "", err
	}

	// pass back account id
	return vt.Email, nil
}

// ReadPasswordResetCode returns the user reset code
func (h *Handler) ReadPasswordResetCode(ctx context.Context, userId uint32, code string) (*passwordResetCode, error) {
	resetCode := &passwordResetCode{}
	err := h.CacheService.Get(ctx, generatePasswordResetCodeStoreKey(ctx, userId, code), &resetCode)
	if err != nil {
		return nil, err
	}

	// check the expiry
	if resetCode.Expires.Before(time.Now()) {
		return nil, errors.InternalServerError(service.GetName(), "password reset code expired")
	}
	return resetCode, nil
}

func (h *Handler) SavePasswordResetCode(ctx context.Context, userId uint32, code string, expiry time.Duration) (*passwordResetCode, error) {
	pwcode := passwordResetCode{
		Expires: time.Now().Add(expiry),
		UserID:  userId,
		Code:    code,
	}

	err := h.CacheService.Set(ctx, generatePasswordResetCodeStoreKey(ctx, userId, code), &pwcode, expiry)
	if err != nil {
		return nil, err
	}
	return &pwcode, err
}

func (h *Handler) DeletePasswordResetCode(ctx context.Context, userId uint32, code string) error {
	return h.CacheService.Delete(ctx, generatePasswordResetCodeStoreKey(ctx, userId, code))
}

// CreateToken returns the created and saved token
func (h *Handler) CreateToken(ctx context.Context, email, token string) error {
	tk := verificationToken{
		Email: email,
		Token: token,
	}

	var expiry int64 = 1800 // 1800 secs = 30 min
	return h.CacheService.Set(ctx, generateVerificationTokenStoreKey(token), &tk, time.Duration(expiry)*time.Second)
}

func generateVerificationTokenStoreKey(token string) string {
	return fmt.Sprintf("user/verification-token/%s", token)
}

func generatePasswordResetCodeStoreKey(ctx context.Context, userId uint32, code string) string {
	return fmt.Sprintf("%vpassword-reset-codes/%v-%v", getStoreKeyPrefix(ctx), userId, code)
}

func getStoreKeyPrefix(ctx context.Context) string {
	tenantId := "micro"
	return getStoreKeyPrefixForTenent(tenantId)
}

func getStoreKeyPrefixForTenent(tenantId string) string {
	tid := strings.Replace(strings.Replace(tenantId, "/", "_", -1), "-", "_", -1)

	return fmt.Sprintf("user/%s/", tid)
}

// random generate i length alphanum string
func random(i int) string {
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
	return "ughwhy?!!!"
}
