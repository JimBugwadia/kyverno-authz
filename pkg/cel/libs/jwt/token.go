package jwt

import (
	"cel.dev/cel-go/common/types"
	"google.golang.org/protobuf/types/known/structpb"
)

var TokenType = types.NewObjectType("jwt.Token")

type Token struct {
	Claims *structpb.Struct
	Valid  bool
}
