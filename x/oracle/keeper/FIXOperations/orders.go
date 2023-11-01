package fixoperations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	fixTypes "github.com/jim380/Re/x/fix/types"
)

type OrdersData struct {
	SessionData  SessionData
	ClOrdID      string
	Symbol       string
	Side         int64
	OrderQty     string
	OrdType      int64
	Price        string
	TimeInForce  int64
	Text         string
	TransactTime string
}

// OrdersOperations creates and sets orders based on the provided data
func OrdersOperations(ctx sdk.Context, k fixKeeper.Keeper, ordersData OrdersData) {
	order := &fixTypes.Orders{
		SessionID: ordersData.SessionData.SessionID,
		Header: &fixTypes.Header{
			BeginString:  "FIX4.4",
			MsgType:      ordersData.SessionData.MsgType,
			SenderCompID: ordersData.SessionData.SenderCompID,
			TargetCompID: ordersData.SessionData.TargetCompID,
			MsgSeqNum:    ordersData.SessionData.MsgSeqNum,
			SendingTime:  ordersData.SessionData.SendingTime,
			ChainID:      ordersData.SessionData.ChainID,
		},
		ClOrdID:      ordersData.ClOrdID,
		Symbol:       ordersData.Symbol,
		Side:         ordersData.Side,
		OrderQty:     ordersData.OrderQty,
		OrdType:      ordersData.OrdType,
		Price:        ordersData.Price,
		TimeInForce:  ordersData.TimeInForce,
		Text:         ordersData.Text,
		TransactTime: ordersData.TransactTime,
		Trailer: &fixTypes.Trailer{
			CheckSum: 0,
		},
	}

	k.SetOrders(ctx, ordersData.ClOrdID, *order)
}
