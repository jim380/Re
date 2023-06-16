package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jim380/Re/x/fix/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateAccount())
	cmd.AddCommand(CmdUpdateAccount())
	cmd.AddCommand(CmdDeleteAccount())
	cmd.AddCommand(CmdLogonInitiator())
	cmd.AddCommand(CmdLogonAcceptor())
	cmd.AddCommand(CmdTerminateLogon())
	cmd.AddCommand(CmdLogonReject())
	cmd.AddCommand(CmdLogoutInitiator())
	cmd.AddCommand(CmdLogoutAcceptor())
	cmd.AddCommand(CmdNewOrderSingle())
	cmd.AddCommand(CmdOrderCancelRequest())
	cmd.AddCommand(CmdOrderCancelReject())
	cmd.AddCommand(CmdOrderExecutionReport())
	cmd.AddCommand(CmdQuoteRequest())
	cmd.AddCommand(CmdQuoteAcknowledgement())
	cmd.AddCommand(CmdQuoteRequestReject())
	cmd.AddCommand(CmdTradeCaptureReport())
	cmd.AddCommand(CmdTradeCaptureReportAcknowledgement())
	cmd.AddCommand(CmdTradeCaptureReportRejection())
	cmd.AddCommand(CmdMarketDataRequest())
	cmd.AddCommand(CmdMarketDataSnapshotFullRefresh())
	cmd.AddCommand(CmdMarketDataIncremental())
	cmd.AddCommand(CmdMarketDataRequestReject())
	cmd.AddCommand(CmdSecurityDefinitionRequest())
	cmd.AddCommand(CmdSecurityDefinition())
	cmd.AddCommand(CmdSecurityDefinitionRequestReject())
	cmd.AddCommand(CmdOrderMassStatusRequest())
	cmd.AddCommand(CmdOrderMassStatusReport())
	cmd.AddCommand(CmdOrderMassStatusRequestReject())
	cmd.AddCommand(CmdTradingSessionStatusRequest())
	cmd.AddCommand(CmdTradingSessionStatus())
	cmd.AddCommand(CmdTradingSessionStatusRequestReject())
	cmd.AddCommand(CmdTradingSessionListRequest())
	cmd.AddCommand(CmdTradingSessionListResponse())
	cmd.AddCommand(CmdTradingSessionListRequestReject())
	cmd.AddCommand(CmdSecurityListRequest())
	cmd.AddCommand(CmdSecurityListResponse())
	cmd.AddCommand(CmdSecurityListRequestReject())
	cmd.AddCommand(CmdSecurityStatusRequest())
	cmd.AddCommand(CmdSecurityStatusResponse())
	cmd.AddCommand(CmdSecurityStatusRequestReject())
	cmd.AddCommand(CmdSecurityTypesRequest())
	cmd.AddCommand(CmdSecurityTypesResponse())
	cmd.AddCommand(CmdSecurityTypesRequestReject())
	// this line is used by starport scaffolding # 1

	return cmd
}
