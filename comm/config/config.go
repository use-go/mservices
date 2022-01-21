package config

import (
	"comm/service"

	"github.com/2637309949/micro/v3/service/config"
)

func Get(path string, options ...config.Option) (config.Value, error) {
	return config.Get(service.GetName()+"."+path, options...)
}

func Set(path string, val interface{}, options ...config.Option) error {
	return config.Set(service.GetName()+"."+path, val, options...)
}

func Delete(path string, options ...config.Option) error {
	return config.Delete(service.GetName()+"."+path, options...)
}
