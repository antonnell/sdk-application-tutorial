package token

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	TokenRecords []Token `json:"token_records"`
}

func NewGenesisState(tokenRecords []Token) GenesisState {
	return GenesisState{TokenRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.TokenRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid TokenRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Name == "" {
			return fmt.Errorf("Invalid TokenRecord: Name: %s. Error: Missing Name", record.Name)
		}
		if record.TotalSupply == nil {
			return fmt.Errorf("Invalid TokenRecord: TotalSupply: %s. Error: Missing TotalSupply", record.TotalSupply)
		}
		if record.Symbol == nil {
			return fmt.Errorf("Invalid TokenRecord: Symbol: %s. Error: Missing Symbol", record.Symbol)
		}
		if record.Mintable == nil {
			return fmt.Errorf("Invalid TokenRecord: Mintable: %s. Error: Missing Mintable", record.Mintable)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		TokenRecords: []Token{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.TokenRecords {
		keeper.SetToken(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Token
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		var token Token
		token = k.GetToken(ctx, name)
		records = append(records, token)
	}
	return GenesisState{TokenRecords: records}
}
