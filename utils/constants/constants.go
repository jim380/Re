package constants

import (
	"time"

	"github.com/patrickmn/go-cache"
)

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
	OracleID            = 15

	LogonRequestStatus = "logon-request"
	LoggedInStatus     = "loggedIn"
	RejectedStatus     = "rejected"

	// cache
	CacheTTL          = 5 * time.Second
	FetchEvery        = 5 * time.Second
	CosmosHubCacheKey = "cosmosHub"
)

var (
	SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")
	CreatedAt   = time.Now().UTC().Format("20060102-15:04:05.000")
)

// cache
var (
	C = cache.New(CacheTTL, 10*time.Minute)
)
