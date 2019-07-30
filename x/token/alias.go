package token

import (
	"github.com/antonnell/sdk-application-tutorial/x/token/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgIssueToken = types.NewMsgIssueToken
	NewToken         = types.NewToken
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	MsgIssueToken = types.MsgIssueToken
	Token         = types.Token
)
