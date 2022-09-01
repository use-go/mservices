package errors

import (
	"fmt"

	"github.com/2637309949/micro/v3/service/errors"
)

const (
	StatusRecordNotFound = 100
)

func RecordNotFound(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   StatusRecordNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: "RecordNotFound",
	}
}
