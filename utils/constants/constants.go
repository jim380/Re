package constants

import "time"

const (
	NameChars           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SessionNameLength   = 10
	QuoteReqIDLength    = 15
	QuoteID             = 15
	MdReqID             = 15
	TradeReportID       = 15
	SecurityReqID       = 15
	MassStatusReqID     = 15
	ClOrdID             = 15
	TradSesReqID        = 15
	SecurityStatusReqID = 15
	SecurityResponseID  = 15

	LogonRequestStatus = "logon-request"
	LoggedInStatus     = "loggedIn"
	RejectedStatus     = "rejected"
)

var SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")
