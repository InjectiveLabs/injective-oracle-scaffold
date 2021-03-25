package cli

import (
	"github.com/spf13/cobra"

	"github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/modules/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the parent command for all modules/oracle CLi query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand()
	return cmd
}
