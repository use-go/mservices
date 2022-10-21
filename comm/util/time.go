package util

import "time"

const (
	UnixZero = 0
)

func Unix(t time.Time) int64 {
	localtime := t.In(time.Local)
	sec := localtime.Unix()
	if sec < 0 {
		sec = UnixZero
	}
	return sec
}

func Time(sec int64) (t time.Time) {
	if sec > 0 {
		t = time.Unix(sec, 0)
	}
	return
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseDate(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", str, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}

func Parse(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}
