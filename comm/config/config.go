package config

import (
	"github.com/2637309949/micro/v3/service/config"
)

func Get(path string, options ...config.Option) (config.Value, error) {
	return config.Get(path, options...)
}

func Set(path string, val interface{}, options ...config.Option) error {
	return config.Set(path, val, options...)
}

func Delete(path string, options ...config.Option) error {
	return config.Delete(path, options...)
}
