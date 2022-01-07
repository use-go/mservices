package auth

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/util/auth/token"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type authClaims struct {
	Type     string            `json:"type"`
	Scopes   []string          `json:"scopes"`
	Metadata map[string]string `json:"metadata"`
	Name     string            `json:"name"`

	jwt.StandardClaims
}

type JWTAccessGenerate struct {
	opts token.Options
}

// NewJWTAccessGenerate create to generate the jwt access token instance
func NewJWTAccessGenerate(opts ...token.Option) oauth2.AccessGenerate {
	return &JWTAccessGenerate{
		opts: token.NewOptions(opts...),
	}
}

func (j *JWTAccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	var priv []byte
	if strings.HasPrefix(j.opts.PrivateKey, "-----BEGIN RSA PRIVATE KEY-----") {
		priv = []byte(j.opts.PrivateKey)
	} else {
		var err error
		priv, err = base64.StdEncoding.DecodeString(j.opts.PrivateKey)
		if err != nil {
			return "", "", err
		}
	}

	// parse the private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return "", "", token.ErrEncodingToken
	}

	// mock user
	acc := &auth.Account{
		ID:   data.UserID,
		Type: "user",
	}
	// backwards compatibility
	name := acc.Name
	if name == "" {
		name = acc.ID
	}

	// generate the JWT
	expiry := time.Now().Add(data.TokenInfo.GetAccessExpiresIn())
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaims{
		Type: acc.Type, Scopes: acc.Scopes, Metadata: acc.Metadata, Name: name,
		StandardClaims: jwt.StandardClaims{
			Subject:   acc.ID,
			Issuer:    acc.Issuer,
			ExpiresAt: expiry.Unix(),
		},
	})
	tok, err := t.SignedString(key)
	if err != nil {
		return "", "", err
	}
	refresh := ""

	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(tok)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return tok, refresh, nil
}
