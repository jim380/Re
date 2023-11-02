package fixoperations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	fixTypes "github.com/jim380/Re/x/fix/types"
)

type ExecutionReportData struct {
	SessionData  SessionData
	ClOrdID      string
	OrderID      string
	ExecID       string
	OrdStatus    string
	ExecTyp      string
	Symbol       string
	Side         int64
	OrderQty     string
	Price        string
	TimeInForce  int64
	LastPx       int64
	LastQty      int64
	LeavesQty    int64
	CumQty       int64
	AvgPx        int64
	Text         string
	TransactTime string
}

// ExecutionReportOperation creates and sets the order execution report based on the provided data
func ExecutionReportOperation(ctx sdk.Context, k fixKeeper.Keeper, executionReportData ExecutionReportData) {
	orderExecutionReport := &fixTypes.OrdersExecutionReport{
		SessionID: executionReportData.SessionData.SessionID,
		Header: &fixTypes.Header{
			BeginString:  "FIX4.4",
			MsgType:      executionReportData.SessionData.MsgType,
			SenderCompID: executionReportData.SessionData.SenderCompID,
			TargetCompID: executionReportData.SessionData.TargetCompID,
			MsgSeqNum:    executionReportData.SessionData.MsgSeqNum,
			SendingTime:  executionReportData.SessionData.SendingTime,
			ChainID:      executionReportData.SessionData.ChainID,
		},
		ClOrdID:      executionReportData.ClOrdID,
		OrderID:      executionReportData.OrderID,
		ExecID:       executionReportData.ExecID,
		OrdStatus:    executionReportData.OrdStatus,
		ExecType:     executionReportData.ExecTyp,
		Symbol:       executionReportData.Symbol,
		Side:         executionReportData.Side,
		OrderQty:     executionReportData.OrderQty,
		Price:        executionReportData.Price,
		TimeInForce:  executionReportData.TimeInForce,
		LastPx:       executionReportData.LastPx,
		LastQty:      executionReportData.LastQty,
		LeavesQty:    executionReportData.LeavesQty,
		CumQty:       executionReportData.CumQty,
		AvgPx:        executionReportData.AvgPx,
		Text:         executionReportData.Text,
		TransactTime: executionReportData.TransactTime,
		Trailer: &fixTypes.Trailer{
			CheckSum: 0,
		},
	}

	k.SetOrdersExecutionReport(ctx, executionReportData.ClOrdID, *orderExecutionReport)
}
