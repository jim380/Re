package types

const (
	// ModuleName defines the module name
	ModuleName = "fix"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fix"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	AccountKey      = []byte{0x00}
	AccountCountKey = []byte{0x01}

	SessionsKey      = []byte{0x02}
	SessionsCountKey = []byte{0x03}

	SessionRejectKey      = []byte{0x04}
	SessionRejectCountKey = []byte{0x05}

	SessionLogoutKey      = []byte{0x06}
	SessionLogoutCountKey = []byte{0x07}

	OrdersKey      = []byte{0x08}
	OrdersCountKey = []byte{0x09}

	OrdersCancelRequestKey      = []byte{0x010}
	OrdersCancelRequestCountKey = []byte{0x011}

	OrdersCancelRejectKey      = []byte{0x012}
	OrdersCancelRejectCountKey = []byte{0x012}

	OrdersExecutionReportKey      = []byte{0x013}
	OrdersExecutionReportCountKey = []byte{0x014}

	QuoteKey = []byte{0x015}

	MarketDataKey      = []byte{0x016}
	MarketDataCountKey = []byte{0x017}

	TradeCaptureKey      = []byte{0x018}
	TradeCaptureCountKey = []byte{0x019}

	SecurityKey      = []byte{0x020}
	SecurityCountKey = []byte{0x021}

	OrderMassStatusKey      = []byte{0x022}
	OrderMassStatusCountKey = []byte{0x023}

	TradingSessionKey      = []byte{0x024}
	TradingSessionCountKey = []byte{0x025}
)

func GetAccountKey() []byte {
	return AccountKey
}

func GetAccountCountKey() []byte {
	return AccountCountKey
}

func GetSessionsKey() []byte {
	return SessionsKey
}

func GetSessionsCountKey() []byte {
	return SessionsCountKey
}

func GetSessionRejectKey() []byte {
	return SessionRejectKey
}

func GetSessionRejectCountKey() []byte {
	return SessionRejectCountKey
}

func GetSessionLogoutKey() []byte {
	return SessionLogoutKey
}

func GetSessionLogoutCountKey() []byte {
	return SessionLogoutCountKey
}

func GetOrdersKey() []byte {
	return OrdersKey
}

func GetOrdersCountKey() []byte {
	return OrdersCountKey
}

func GetOrdersCancelRequestKey() []byte {
	return OrdersCancelRequestKey
}

func GetOrdersCancelRequestCountKey() []byte {
	return OrdersCancelRequestCountKey
}

func GetOrdersCancelRejectKey() []byte {
	return OrdersCancelRejectKey
}

func GetOrdersCancelRejectCountKey() []byte {
	return OrdersCancelRejectCountKey
}

func GetOrdersExecutionReportKey() []byte {
	return OrdersExecutionReportKey
}

func GetOrdersExecutionReportCountKey() []byte {
	return OrdersExecutionReportCountKey
}

func GetQuoteKey() []byte {
	return QuoteKey
}

func GetTradeCaptureKey() []byte {
	return TradeCaptureKey
}

func GetTradeCaptureCountKey() []byte {
	return TradeCaptureCountKey
}

func GetMarketDataKey() []byte {
	return MarketDataKey
}

func GetMarketDataCountKey() []byte {
	return MarketDataCountKey
}

func GetSecurityKey() []byte {
	return SecurityKey
}

func GetSecurityCountKey() []byte {
	return SecurityCountKey
}

func GetOrderMassStatusKey() []byte {
	return OrderMassStatusKey
}

func GetOrderMassStatusCountKey() []byte {
	return OrderMassStatusCountKey
}

func GetTradingSessionKey() []byte {
	return TradingSessionKey
}

func GetTradingSessionCountKey() []byte {
	return TradingSessionCountKey
}

const (
	TradingSessionListKey      = "TradingSessionList/value/"
	TradingSessionListCountKey = "TradingSessionList/count/"
)
