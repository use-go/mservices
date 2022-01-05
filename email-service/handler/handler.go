package handler

import (
	"github.com/Teamwork/spamc"
)

type SendGridConf struct {
	Key       string `json:"key"`
	EmailFrom string `json:"email_from"`
	PoolName  string `json:"ip_pool_name"`
}

type SpamcConf struct {
	SpamdAddress string `json:"spamd_address"`
}

type Sent struct {
	UserID        string
	SendgridMsgID string
}

type Handler struct {
	Config SendGridConf
	Spamc  *spamc.Client
}
