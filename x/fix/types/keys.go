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
	SessionsKey      = []byte{0x01}
	SessionsCountKey = []byte{0x02}

	SessionRejectKey      = []byte{0x0A}
	SessionRejectCountKey = []byte{0x0B}

	SessionLogoutKey      = []byte{0x05}
	SessionLogoutCountKey = []byte{0x06}

	OrdersKey      = []byte{0x07}
	OrdersCountKey = []byte{0x08}

	OrdersCancelRequestKey      = []byte{0x09}
	OrdersCancelRequestCountKey = []byte{0x010}

	OrdersCancelRejectKey      = []byte{0x011}
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

	TradingSessionListKey      = []byte{0x026}
	TradingSessionListCountKey = []byte{0x027}

	SecurityListKey      = []byte{0x028}
	SecurityListCountKey = []byte{0x029}

	SecurityStatusKey      = []byte{0x030}
	SecurityStatusCountKey = []byte{0x031}

	SecurityTypesKey      = []byte{0x032}
	SecurityTypesCountKey = []byte{0x033}

	AccountRegistrationKey      = []byte{0x034}
	AccountRegistrationCountKey = []byte{0x035}
)

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

func GetTradingSessionListKey() []byte {
	return TradingSessionListKey
}

func GetTradingSessionListCountKey() []byte {
	return TradingSessionListCountKey
}

func GetSecurityListKey() []byte {
	return SecurityListKey
}

func GetSecurityListCountKey() []byte {
	return SecurityListCountKey
}

func GetSecurityStatusKey() []byte {
	return SecurityStatusKey
}

func GetSecurityStatusCountKey() []byte {
	return SecurityStatusCountKey
}

func GetSecurityTypesKey() []byte {
	return SecurityTypesKey
}

func GetSecurityTypesCountKey() []byte {
	return SecurityTypesCountKey
}

func GetAccountRegistrationKey() []byte {
	return AccountRegistrationKey
}

func GetAccountRegistrationCountKey() []byte {
	return AccountRegistrationCountKey
}
