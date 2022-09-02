package auth

import (
	"comm/errors"
	"comm/logger"
	"comm/service"
	"context"
	"sync"

	"github.com/2637309949/micro/v3/service/auth"
)

var (
	account *auth.Account
	once    sync.Once
)

// GenerateAccount defined todo
func GenerateAccount(ctx context.Context) (*auth.Account, error) {
	once.Do(func() {
		accName := service.GetName() + "-latest"
		opts := []auth.GenerateOption{
			auth.WithIssuer("micro"),
			auth.WithScopes("service"),
			auth.WithType("service"),
		}
		acc, err := auth.Generate(accName, opts...)
		if err != nil {
			logger.Errorf(ctx, err.Error())
			return
		}
		account = acc
	})
	if account != nil {
		return account, nil
	}
	return nil, errors.InternalServerError(service.GetName(), "Generate fail")
}

func GenerateSecret(ctx context.Context) (string, error) {
	acc, err := GenerateAccount(ctx)
	if err != nil {
		return "", err
	}
	return acc.Secret, nil
}

func MustSecret(ctx context.Context) string {
	secret, _ := GenerateSecret(ctx)
	return secret
}
