package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey k
const RouterKey = ModuleName // this was defined in your key.go file

// MsgIssueToken defines a SetName message
type MsgIssueToken struct {
	Name           string         `json:"name"`
	Symbol         string         `json:"symbol"`
	Mintable       bool           `json:"mintable"`
	TotalSupply    sdk.Coins      `json:"total_supply"`
	Owner          sdk.AccAddress `json:"owner"`
}

// NewMsgIssueToken is a constructor function for MsgIssueToken
func NewMsgIssueToken(name string, symbol string, mintable bool, totalSupply sdk.Coins, owner sdk.AccAddress) MsgIssueToken {
	return MsgIssueToken{
		Name:           name,
		Symbol:         symbol,
		Mintable:       mintable,
		TotalSupply:    totalSupply,
		Owner:          owner,
	}
}

// Route should return the name of the module
func (msg MsgIssueToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgIssueToken) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgIssueToken) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Symbol) == 0  || len(msg.TotalSupply) == 0 {
		return sdk.ErrUnknownRequest("Name, Symbol, TotalSuppky cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgIssueToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgIssueToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}