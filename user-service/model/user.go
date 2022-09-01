package model

import (
	"comm/errors"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

var (
	x        = "sdfvw12d3s"
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func random(i int) string {
	bytes := make([]byte, i)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// UnmarshalUser defined TODO
func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal defined TODO
func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// User defined TODO
type User struct {
	// unique account id
	Id uint32 `gorm:"column:id;primary_key" json:"id"`
	// alphanumeric username
	Username string `gorm:"column:username" json:"username,omitempty"`
	// alphanumeric salt
	Salt string `gorm:"column:salt" json:"-"`
	// an salt password
	Password string `gorm:"column:password" json:"-"`
	// an email address
	Email string `gorm:"column:email" json:"email,omitempty"`
	// unix timestamp
	Created uint32 `gorm:"column:created" json:"created,omitempty"`
	// unix timestamp
	Updated uint32 `gorm:"column:updated" json:"updated,omitempty"`
	// if the account is verified
	Verified bool `gorm:"column:verified" json:"verified,omitempty"`
	// date of verification
	VerificationDate uint32 `gorm:"column:verification_date" json:"verification_date,omitempty"`
	// Store any custom data you want about your users in this fields.
	Profile map[string]string `gorm:"column:name" json:"profile,omitempty"`
}

func (h *User) GenerateFromPassword(password string) (string, error) {
	salt := random(16)
	hb, err := bcrypt.GenerateFromPassword([]byte(x+salt+password), 10)
	if err != nil {
		return "", errors.InternalServerError(err.Error())
	}
	hashedPassword := base64.StdEncoding.EncodeToString(hb)
	h.Password = hashedPassword
	h.Salt = salt
	return hashedPassword, nil
}

func (h *User) CompareHashAndPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(h.Password), []byte(x+h.Salt+password)); err != nil {
		return errors.Unauthorized(err.Error())
	}
	return nil
}

// TableName sets the insert table name for this struct type
func (h *User) TableName() string {
	return "user"
}
