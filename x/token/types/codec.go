package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc moduleCdc
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueToken{}, "token/IssueToken", nil)
}
