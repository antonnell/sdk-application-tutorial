package cli

import (
	"fmt"

	"github.com/antonnell/sdk-application-tutorial/x/token/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetQueryCmd query
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	tokenQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the token module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	tokenQueryCmd.AddCommand(client.GetCommands(
		GetCmdToken(storeKey, cdc),
	)...)
	return tokenQueryCmd
}

// GetCmdToken queries information about a domain
func GetCmdToken(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "token [name]",
		Short: "Query token info of name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/token/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve token - %s \n", name)
				return nil
			}

			var out types.Token
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
