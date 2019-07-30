package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Token is a struct that contains all the metadata of a name
type Token struct {
	Name           string         `json:"name"`
	Symbol         string         `json:"symbol"`
	OriginalSymbol string         `json:"original_symbol"`
	Mintable       bool           `json:"mintable"`
	TotalSupply    sdk.Coins      `json:"total_supply"`
	Owner          sdk.AccAddress `json:"owner"`
}

// NewToken Returns a new Token with the minprice as the price
func NewToken() Token {
	return Token{}
}

// implement fmt.Stringer
func (t Token) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
	Name: %s
	Symbol: %s
	OriginalSymbol: %s
	TotalSupply: %s`, t.Owner, t.Name, t.Symbol, t.OriginalSymbol, t.TotalSupply))
}
