package token

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// SetToken Sets the entire Token metadata struct for a symbol
func (k Keeper) SetToken(ctx sdk.Context, symbol string, token Token) {
	if token.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(symbol), k.cdc.MustMarshalBinaryBare(token))
}

// GetToken Gets the entire Token metadata struct for a symbol
func (k Keeper) GetToken(ctx sdk.Context, symbol string) Token {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(symbol)) {
		return NewToken()
	}
	bz := store.Get([]byte(symbol))
	var token Token
	k.cdc.MustUnmarshalBinaryBare(bz, &token)
	return token
}

// ResolveSymbol - returns the string that the symbol resolves to
func (k Keeper) ResolveSymbol(ctx sdk.Context, symbol string) string {
	return k.GetToken(ctx, symbol)
}

// SetName - sets the value string that a symbol resolves to
func (k Keeper) SetName(ctx sdk.Context, symbol string, value string) {
	token := k.GetToken(ctx, symbol)
	token.Value = value
	k.SetToken(ctx, symbol, token)
}

// GetName - gets the value string that a symbol resolves to
func (k Keeper) GetName(ctx sdk.Context, symbol string, value string) {
	token := k.GetToken(ctx, symbol)
	token.Name = value
	k.SetToken(ctx, symbol, token)
}

// HasOwner - returns whether or not the symbol already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, symbol string) bool {
	return !k.GetToken(ctx, symbol).Owner.Empty()
}

// GetOwner - get the current owner of a symbol
func (k Keeper) GetOwner(ctx sdk.Context, symbol string) sdk.AccAddress {
	return k.GetToken(ctx, symbol).Owner
}

// SetOwner - sets the current owner of a symbol
func (k Keeper) SetOwner(ctx sdk.Context, symbol string, owner sdk.AccAddress) {
	token := k.GetToken(ctx, symbol)
	token.Owner = owner
	k.SetToken(ctx, symbol, token)
}

// GetTotalSupply - gets the current price of a symbol.
func (k Keeper) GetTotalSupply(ctx sdk.Context, symbol string) sdk.Coins {
	return k.GetToken(ctx, symbol).TotalSupply
}

// SetTotalSupply - sets the current price of a symbol
func (k Keeper) SetTotalSupply(ctx sdk.Context, symbol string, totalSupply sdk.Coins) {
	token := k.GetToken(ctx, symbol)
	token.TotalSupply = totalSupply
	k.SetToken(ctx, symbol, token)
}

// GetSymbol - gets the token's symbol
func (k Keeper) GetSymbol(ctx sdk.Context, symbol string) string {
	return k.GetToken(ctx, symbol).Symbol
}

// SetSmybol - sets the token's symbol
func (k Keeper) SetSmybol(ctx sdk.Context, symbol string, value string) {
	token := k.GetToken(ctx, symbol)
	token.Symbol = value
	k.SetToken(ctx, symbol, token)
}

// GetOriginalSymbol - gets the token's symbol
func (k Keeper) GetOriginalSymbol(ctx sdk.Context, symbol string) string {
	return k.GetToken(ctx, symbol).OriginalSymbol
}

// SetOriginalSmybol - sets the token's symbol
func (k Keeper) SetOriginalSmybol(ctx sdk.Context, symbol string, value string) {
	token := k.GetToken(ctx, symbol)
	token.OriginalSymbol = value
	k.SetToken(ctx, symbol, token)
}

// GetMintable - gets the token's symbol
func (k Keeper) GetMintable(ctx sdk.Context, symbol string) bool {
	return k.GetToken(ctx, symbol).OriginalSymbol
}

// SetMintable - sets the token's symbol
func (k Keeper) SetMintable(ctx sdk.Context, symbol string, mintable bool) {
	token := k.GetToken(ctx, symbol)
	token.Mintable = mintable
	k.SetToken(ctx, symbol, token)
}

// GetTokensIterator Get an iterator over all tokens in which the keys are the symbols and the values are the whois
func (k Keeper) GetTokensIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// NewToken - sets the value string that a symbol resolves to
func (k Keeper) NewToken(ctx sdk.Context, owner sdk.AccAddress, symbol string, name string, totalSupply sdk.Coins, mintable bool) {
	token := k.GetToken(ctx, symbol)

	if !token.Owner.Empty() {
		return //token already exists, cant issue
	}

	token.Owner = owner
	token.Name = name
	token.Symbol = symbol
	token.TotalSupply = totalSupply
	token.Mintable = mintable

	k.SetToken(ctx, symbol, token)
}

// NewKeeper creates new instances of the token Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}
