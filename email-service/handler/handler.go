package handler

import (
	"github.com/Teamwork/spamc"
)

type Smtp struct {
	Addr     string `json:"addr"`
	UserName string `json:"username"`
	Identity string `json:"identity"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

type SpamcConf struct {
	SpamdAddress string `json:"spamd_address"`
}

type Sent struct {
	UserID        string
	SendgridMsgID string
}

type Handler struct {
	Smtp  *Smtp
	Spamc *spamc.Client
}
