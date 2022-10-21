package util

import (
	"fmt"
	"reflect"
	"strings"
)

func IsZero(s interface{}, tags ...string) error {
	v := reflect.Indirect(reflect.ValueOf(s))
	for _, n := range tags {
		for i := 0; i < v.Type().NumField(); i++ {
			if jArr := strings.Split(v.Type().Field(i).Tag.Get("json"), ","); len(jArr) > 0 && jArr[0] == n {
				field := v.Field(i)
				if field.IsZero() {
					return fmt.Errorf("%v is required", n)
				}
			}
		}
	}
	return nil
}
