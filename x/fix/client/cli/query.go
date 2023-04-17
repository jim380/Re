package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jim380/Re/x/fix/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group fix queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdShowAccount())
	cmd.AddCommand(CmdListSessions())
	cmd.AddCommand(CmdShowSessions())
	cmd.AddCommand(CmdListSessionReject())
	cmd.AddCommand(CmdShowSessionReject())
	cmd.AddCommand(CmdListSessionLogout())
	cmd.AddCommand(CmdShowSessionLogout())
	cmd.AddCommand(CmdListOrders())
	cmd.AddCommand(CmdShowOrders())
	cmd.AddCommand(CmdListOrdersCancelRequest())
	cmd.AddCommand(CmdShowOrdersCancelRequest())
	cmd.AddCommand(CmdListOrdersCancelReject())
	cmd.AddCommand(CmdShowOrdersCancelReject())
	cmd.AddCommand(CmdListOrdersExecutionReport())
	cmd.AddCommand(CmdShowOrdersExecutionReport())
	cmd.AddCommand(CmdListQuote())
	cmd.AddCommand(CmdShowQuote())
	cmd.AddCommand(CmdShowQuotesBySessionID())
	cmd.AddCommand(CmdListMarketData())
	cmd.AddCommand(CmdShowMarketData())
	// this line is used by starport scaffolding # 1

	return cmd
}
