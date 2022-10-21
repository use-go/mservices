package api

import (
	"net/http"
	"strconv"
)

func MustCookie(r *http.Request, name string) string {
	cke, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return cke.Value
}

type Uri struct {
	User   string `json:"user"`
	PassWd string `json:"passwd"`
	Host   string `json:"host"`
	Port   uint32 `json:"port"`
	DB     string `json:"db"`
}

func (u *Uri) Unmarshal(r *http.Request) *Uri {
	u.User = MustCookie(r, "user")
	u.PassWd = MustCookie(r, "passwd")
	u.Host = MustCookie(r, "host")
	pt, _ := strconv.ParseUint(MustCookie(r, "port"), 10, 64)
	u.Port = uint32(pt)
	u.DB = MustCookie(r, "db")
	return u
}
