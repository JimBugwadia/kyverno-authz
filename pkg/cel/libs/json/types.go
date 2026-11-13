package json

import (
	"cel.dev/cel-go/common/types"
)

var JsonType = types.NewOpaqueType("json.Json")

type JsonImpl interface {
	Unmarshal([]byte) (any, error)
}

type Json struct {
	JsonImpl
}
