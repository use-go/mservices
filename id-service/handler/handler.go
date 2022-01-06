package handler

import (
	"github.com/mattheath/kala/bigflake"
	"github.com/mattheath/kala/snowflake"
)

type Handler struct {
	Snowflake *snowflake.Snowflake
	Bigflake  *bigflake.Bigflake
}
