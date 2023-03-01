package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) OrderExecutionReport(goCtx context.Context, msg *types.MsgOrderExecutionReport) (*types.MsgOrderExecutionReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message err
	//Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator
	//checks for the fields in the FIX Protocol

	orderExecutionReport := types.OrdersExecutionReport{
		SessionID:    msg.SessionID,
		Header:       session.LogonAcceptor.Header,
		ClOrdID:      msg.ClOrdID,
		OrderID:      msg.OrderID,
		ExecID:       msg.ExecID,
		OrdStatus:    msg.OrdStatus,
		ExecType:     msg.ExecType,
		Symbol:       msg.Symbol,
		Side:         msg.Side,
		OrderQty:     msg.OrderQty,
		Price:        msg.Price,
		TimeInForce:  msg.TimeInForce,
		LastPx:       msg.LastPx,
		LastQty:      msg.LastQty,
		LeavesQty:    msg.LeavesQty,
		CumQty:       msg.CumQty,
		AvgPx:        msg.AvgPx,
		Text:         msg.Text,
		TransactTime: msg.TransactTime,
		Trailer:      session.LogonAcceptor.Trailer,
		Creator:      msg.Creator,
	}

	orderExecutionReport.Header.MsgType = "8"
	orderExecutionReport.TransactTime = time.Now().UTC().Format("20060102-15:04:05.000")

	//set order execution report to store
	k.SetOrdersExecutionReport(ctx, msg.SessionID, orderExecutionReport)

	return &types.MsgOrderExecutionReportResponse{}, nil
}
