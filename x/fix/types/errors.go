package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Error Code Enums

const (
	// DID
	errInvalidDidDocument uint32 = iota + 101
	errAccountIsTaken
	errNotDIDCreator
	errDIDIsTaken
	errCompanyNameIsTaken
	errWebsiteIstaken
	errNotAccountCreator

	// session
	errSessionNameFound
	errSessionSameDID
	errEmptySession
	errWrongSession
	errIncorrectDID
	errSessionIsAccepted
	errSessionIsNotLoggedIn

	// Quote
	errQuoteIsEmpty
	errQuoteReqIDIsTaken
	errQuoteReqIDIsEmpty
	errQuoteSymbolIsEmpty
	errQuoteSideIsEmpty
	errQuoteBidPxIsEmpty
	errQuoteCurrencyIsEmpty
	errQuoteTypeIsEmpty
	errMICInQuoteRquestIsNotFound
	errQuoteRequestIsRejected
	errQuoteAcknowledgementCreatorIsWrong
	errQuoteSession
	errQuoteRequestRejectCreatorIsWrong
	errQuoteRequestIsAcknowledged
	errQuoteRequestRejectReasonIsEmpty

	// Market Data
	errSubscriptionRequestTypeIsEmpty
	errMarketDepthIsEmpty
	errMdUpdateTypeIsEmpty
	errNoRelatedSymIsEmpty
	errSymbolIsEmpty
	errMdReqIDIsNotFound
	errMarketDataSnapShotFullRefreshCreatorIsWrong
	errMarketDataRequestIsRejected
	errMdReqIDIsEmpty
	errNoMDEntriesIsEmpty
	errMdUpdateActionIsEmpty
	errMdEntryTypeIsEmpty
	errMdEntryPxIsEmpty
	errMdEntrySizeIsEmpty
	errMarketDataSession
	errMarketDataRequestRejectCreatorIsWrong
	errMarketDataRequestIsAcknowlodged
	errMdReqRejReasonIsEmpty

	// Orders
	errOrdersExecutionReportIsNotFound
	errOrderIDIsNotFound
	errExecIDIsNotFound
	errOrderIsNotFound
	errOrderIsExecutedAlready
	errOrderEmptyField

	// Trade Capture
	errTradeReportIDIsEmpty
	errTradeReportTransTypeIsEmpty
	errTrdTypeIsEmpty
	errTradeCaptureSideIsEmpty
	errTradeCaptureOrderQtyIsEmpty
	errTradeCaptureLastQtyIsEmpty
	errTradeCaptureLastPxIsEmpty
	errTradeCaptureGrossTradeAmtIsEmpty
	errTradeCaptureSymbolIsEmpty
	errTradeCaptureSecurityIDIsEmpty
	errTradeCaptureTradeDateIsEmpty
	errTradeCaptureTransactTimeIsEmpty
	errTradeCaptureIsNotFound
	errTradeCaptureMismatchField
	errExecTypeIsEmpty
	errTradeReportStatusIsEmpty
	errTradeCaptureReportIsRejected
	errTradeCaptureReportIsAcknowledged
	errTradeCaptureSession
	errTradeReportRejectReasonIsEmpty

	// security
	errSecurityReqIDIsEmpty
	errSecurityRequestTypeIsEmpty
	errSecuritySymbolIsEmpty
	errSecurityExchangeIsEmpty
	errSecuritySession
	errSecurityIsNotFound
	errSecurityDefinitionRequestIsRejected
	errSecurityDefinitionRequestIsAcknowledged
	errSecurityMismatchField
	errSecurityTypeIsEmpty
	errSecurityCurrencyIsEmpty
	errSecurityMinPriceIncrementIsEmpty
	errSecurityRequestResultIsEmpty
	errSecurityRequestErrorIsEmpty

	// Order Mass Status
	errOrderMassStatusSession
	errOrderMassStatusMismatchField
	errOrderMassStatusEmptyField
	errOrderMassStatusIsNotFound
	errOrderMassStatusCreatorIsWrong
	errOrderMassStatusRequestIsAcknowledged
	errOrderMassStatusRequestIsRejected

	// Trading Session
	errTradingSessionEmptyField
	errTradingSessionSession
	errTradingSessionIsNotFound
	errTradingSessionStatusRequestIsRejected
	errTradingSessionStatusRequestIsAcknowledged
	errTradingSessionMismatchField

	// Trading Session List
	errTradingSessionListEmptyField
	errTradingSessionListSession
	errTradingSessionListIsNotFound
	errTradingSessionListRequestIsRejected
	errTradingSessionListRequestIsAcknowledged
)

// x/fix module sentinel errors
var (
	// DID
	ErrInvalidDidDocument = sdkerrors.Register(ModuleName, errInvalidDidDocument, "Invalid DID Document")
	ErrAccountIsTaken     = sdkerrors.Register(ModuleName, errAccountIsTaken, "Account Address is used already")
	ErrNotDIDCreator      = sdkerrors.Register(ModuleName, errNotDIDCreator, "Account creator must be the same with DID creator")
	ErrDIDIsTaken         = sdkerrors.Register(ModuleName, errDIDIsTaken, "DID exists with an account already")
	ErrCompanyNameIsTaken = sdkerrors.Register(ModuleName, errCompanyNameIsTaken, "Company Name exists with an account already")
	ErrWebsiteIstaken     = sdkerrors.Register(ModuleName, errWebsiteIstaken, "Website is used already")
	ErrNotAccountCreator  = sdkerrors.Register(ModuleName, errNotAccountCreator, "Incorrect Account Owner")

	// Session
	ErrSessionNameFound     = sdkerrors.Register(ModuleName, errSessionNameFound, "Session Name exists")
	ErrSessionSameDID       = sdkerrors.Register(ModuleName, errSessionSameDID, "Session can not use same DID for senderCompID and targetCompID")
	ErrEmptySession         = sdkerrors.Register(ModuleName, errEmptySession, "Session does not Exist")
	ErrWrongSession         = sdkerrors.Register(ModuleName, errWrongSession, "The Session provided does not tally with account")
	ErrIncorrectDID         = sdkerrors.Register(ModuleName, errIncorrectDID, "senderCompID and targetCompID does not match in session")
	ErrSessionIsAccepted    = sdkerrors.Register(ModuleName, errSessionIsAccepted, "session accepted already")
	ErrSessionIsNotLoggedIn = sdkerrors.Register(ModuleName, errSessionIsNotLoggedIn, "There is no active session with this sessionID")

	// Quote
	ErrQuoteIsEmpty                       = sdkerrors.Register(ModuleName, errQuoteIsEmpty, "Quote Is Empty")
	ErrQuoteReqIDIsTaken                  = sdkerrors.Register(ModuleName, errQuoteReqIDIsTaken, "QuoteReqID exists already for a given Quote Request")
	ErrQuoteReqIDIsEmpty                  = sdkerrors.Register(ModuleName, errQuoteReqIDIsEmpty, "QuoteReqID is empty")
	ErrQuoteSymbolIsEmpty                 = sdkerrors.Register(ModuleName, errQuoteSymbolIsEmpty, "Quote Symbol is empty")
	ErrQuoteSideIsEmpty                   = sdkerrors.Register(ModuleName, errQuoteSideIsEmpty, "Quote Side is empty")
	ErrQuoteBidPxIsEmpty                  = sdkerrors.Register(ModuleName, errQuoteBidPxIsEmpty, "Quote BidPx is empty")
	ErrQuoteCurrencyIsEmpty               = sdkerrors.Register(ModuleName, errQuoteCurrencyIsEmpty, "Quote Currency is empty")
	ErrQuoteTypeIsEmpty                   = sdkerrors.Register(ModuleName, errQuoteTypeIsEmpty, "QuoteType is empty")
	ErrMICInQuoteRquestIsNotFound         = sdkerrors.Register(ModuleName, errMICInQuoteRquestIsNotFound, "MIC used in Quote is not registered")
	ErrQuoteRequestIsRejected             = sdkerrors.Register(ModuleName, errQuoteRequestIsRejected, "Quote Request has been Rejected")
	ErrQuoteAcknowledgementCreatorIsWrong = sdkerrors.Register(ModuleName, errQuoteAcknowledgementCreatorIsWrong, "This Account Address is not allowed to Acknowledge this Quote Request")
	ErrQuoteSession                       = sdkerrors.Register(ModuleName, errQuoteSession, "The sessionID is not recognized by the Quote Request system")
	ErrQuoteRequestRejectCreatorIsWrong   = sdkerrors.Register(ModuleName, errQuoteRequestRejectCreatorIsWrong, "This Account Address is not allowed to Reject this Quote Request")
	ErrQuoteRequestIsAcknowledged         = sdkerrors.Register(ModuleName, errQuoteRequestIsAcknowledged, "Quote Request has been Acknowledged")
	ErrQuoteRequestRejectReasonIsEmpty    = sdkerrors.Register(ModuleName, errQuoteRequestRejectReasonIsEmpty, "Quote Request Reject Reason is empty")

	// Market Data
	ErrSubscriptionRequestTypeIsEmpty              = sdkerrors.Register(ModuleName, errSubscriptionRequestTypeIsEmpty, "SubscriptionRequestType is empty")
	ErrMarketDepthIsEmpty                          = sdkerrors.Register(ModuleName, errMarketDepthIsEmpty, "MarketDepth is empty")
	ErrMdUpdateTypeIsEmpty                         = sdkerrors.Register(ModuleName, errMdUpdateTypeIsEmpty, "MdUpdateType is empty")
	ErrNoRelatedSymIsEmpty                         = sdkerrors.Register(ModuleName, errNoRelatedSymIsEmpty, "NoRelatedSym is empty")
	ErrSymbolIsEmpty                               = sdkerrors.Register(ModuleName, errSymbolIsEmpty, "Symbol is empty")
	ErrMdReqIDIsNotFound                           = sdkerrors.Register(ModuleName, errMdReqIDIsNotFound, "MdReqID is not found")
	ErrMarketDataSnapShotFullRefreshCreatorIsWrong = sdkerrors.Register(ModuleName, errMarketDataSnapShotFullRefreshCreatorIsWrong, "This account address is not allowed to respond to the market data request")
	ErrMarketDataRequestIsRejected                 = sdkerrors.Register(ModuleName, errMarketDataRequestIsRejected, "Market Data Request has been rejected")
	ErrMdReqIDIsEmpty                              = sdkerrors.Register(ModuleName, errMdReqIDIsEmpty, "MdReqID is empty")
	ErrNoMDEntriesIsEmpty                          = sdkerrors.Register(ModuleName, errNoMDEntriesIsEmpty, "NoMDEntries is empty")
	ErrMdUpdateActionIsEmpty                       = sdkerrors.Register(ModuleName, errMdUpdateActionIsEmpty, "MdUpdateAction is empty")
	ErrMdEntryTypeIsEmpty                          = sdkerrors.Register(ModuleName, errMdEntryTypeIsEmpty, "MdEntryType is empty")
	ErrMdEntryPxIsEmpty                            = sdkerrors.Register(ModuleName, errMdEntryPxIsEmpty, "MdEntryPx is empty")
	ErrMdEntrySizeIsEmpty                          = sdkerrors.Register(ModuleName, errMdEntrySizeIsEmpty, "MdEntrySize is empty")
	ErrMarketDataSession                           = sdkerrors.Register(ModuleName, errMarketDataSession, "The sessionID is not recognized by the Market Data Request system")
	ErrMarketDataRequestRejectCreatorIsWrong       = sdkerrors.Register(ModuleName, errMarketDataRequestRejectCreatorIsWrong, "This account address is not allowed to respond to the market data request reject")
	ErrMarketDataRequestIsAcknowlodged             = sdkerrors.Register(ModuleName, errMarketDataRequestIsAcknowlodged, "Market Data Request has been acknowledged")
	ErrMdReqRejReasonIsEmpty                       = sdkerrors.Register(ModuleName, errMdReqRejReasonIsEmpty, "MdReqRejReason is empty")

	// Orders
	ErrOrdersExecutionReportIsNotFound = sdkerrors.Register(ModuleName, errOrdersExecutionReportIsNotFound, "Order Execution Report is not found")
	ErrOrderIDIsNotFound               = sdkerrors.Register(ModuleName, errOrderIDIsNotFound, "OrderID is not found")
	ErrExecIDIsNotFound                = sdkerrors.Register(ModuleName, errExecIDIsNotFound, "ExecID is not found")
	ErrOrderIsNotFound                 = sdkerrors.Register(ModuleName, errOrderIsNotFound, "Order is not found")
	ErrOrderIsExecutedAlready          = sdkerrors.Register(ModuleName, errOrderIsExecutedAlready, "Order is executed already")
	ErrOrderEmptyField                 = sdkerrors.Register(ModuleName, errOrderEmptyField, "This field can not be left empty")

	// Trade Capture
	ErrTradeReportIDIsEmpty             = sdkerrors.Register(ModuleName, errTradeReportIDIsEmpty, "TradeReportID is not empty")
	ErrTradeReportTransTypeIsEmpty      = sdkerrors.Register(ModuleName, errTradeReportTransTypeIsEmpty, "TradeReportTransType is not Empty")
	ErrTrdTypeIsEmpty                   = sdkerrors.Register(ModuleName, errTrdTypeIsEmpty, "TrdType is not Empty")
	ErrTradeCaptureSideIsEmpty          = sdkerrors.Register(ModuleName, errTradeCaptureSideIsEmpty, "Trade Capture Side is not Empty")
	ErrTradeCaptureOrderQtyIsEmpty      = sdkerrors.Register(ModuleName, errTradeCaptureOrderQtyIsEmpty, "Trade Capture OrderQty is not Empty")
	ErrTradeCaptureLastQtyIsEmpty       = sdkerrors.Register(ModuleName, errTradeCaptureLastQtyIsEmpty, "Trade Capture LastQty is not Empty")
	ErrTradeCaptureLastPxIsEmpty        = sdkerrors.Register(ModuleName, errTradeCaptureLastPxIsEmpty, "Trade Capture LastPx is not Empty")
	ErrTradeCaptureGrossTradeAmtIsEmpty = sdkerrors.Register(ModuleName, errTradeCaptureGrossTradeAmtIsEmpty, "Trade Capture GrossTradeAmt is not Empty")
	ErrTradeCaptureSymbolIsEmpty        = sdkerrors.Register(ModuleName, errTradeCaptureSymbolIsEmpty, "Trade Capture Symbol is not Empty")
	ErrTradeCaptureSecurityIDIsEmpty    = sdkerrors.Register(ModuleName, errTradeCaptureSecurityIDIsEmpty, "Trade Capture SecurityID is not Empty")
	ErrTradeCaptureTradeDateIsEmpty     = sdkerrors.Register(ModuleName, errTradeCaptureTradeDateIsEmpty, "Trade Capture TradeDate is not Empty")
	ErrTradeCaptureTransactTimeIsEmpty  = sdkerrors.Register(ModuleName, errTradeCaptureTransactTimeIsEmpty, "Trade Capture TransactTime is not Empty")
	ErrTradeCaptureIsNotFound           = sdkerrors.Register(ModuleName, errTradeCaptureIsNotFound, "Trade Capture Report is not found")
	ErrTradeCaptureMismatchField        = sdkerrors.Register(ModuleName, errTradeCaptureMismatchField, "This value does not match the value from trade capture report")
	ErrExecTypeIsEmpty                  = sdkerrors.Register(ModuleName, errExecTypeIsEmpty, "ExecType is empty")
	ErrTradeReportStatusIsEmpty         = sdkerrors.Register(ModuleName, errTradeReportStatusIsEmpty, "TradeReportStatus is empty")
	ErrTradeCaptureReportIsRejected     = sdkerrors.Register(ModuleName, errTradeCaptureReportIsRejected, "Trade Capture Report has been rejected")
	ErrTradeCaptureReportIsAcknowledged = sdkerrors.Register(ModuleName, errTradeCaptureReportIsAcknowledged, "Trade Capture Report has been acknowledged")
	ErrTradeCaptureSession              = sdkerrors.Register(ModuleName, errTradeCaptureSession, "The sessionID is not recognized by the Trade Capture Report system")
	ErrTradeReportRejectReasonIsEmpty   = sdkerrors.Register(ModuleName, errTradeReportRejectReasonIsEmpty, "TradeReportRejectReason is empty")

	// security
	ErrSecurityReqIDIsEmpty                    = sdkerrors.Register(ModuleName, errSecurityReqIDIsEmpty, "SecurityReqID is empty")
	ErrSecurityRequestTypeIsEmpty              = sdkerrors.Register(ModuleName, errSecurityRequestTypeIsEmpty, "SecurityRequestType is empty")
	ErrSecuritySymbolIsEmpty                   = sdkerrors.Register(ModuleName, errSecuritySymbolIsEmpty, "Symbol is empty")
	ErrSecurityExchangeIsEmpty                 = sdkerrors.Register(ModuleName, errSecurityExchangeIsEmpty, "SecurityExchange is empty")
	ErrSecuritySession                         = sdkerrors.Register(ModuleName, errSecuritySession, "The sessionID is not recognized by the Security Definition Request system")
	ErrSecurityIsNotFound                      = sdkerrors.Register(ModuleName, errSecurityIsNotFound, "Security Definition Request is not found")
	ErrSecurityDefinitionRequestIsRejected     = sdkerrors.Register(ModuleName, errSecurityDefinitionRequestIsRejected, "Security Definition Request has been rejected")
	ErrSecurityDefinitionRequestIsAcknowledged = sdkerrors.Register(ModuleName, errSecurityDefinitionRequestIsAcknowledged, "Security Definition Request has been acknowledged")
	ErrSecurityMismatchField                   = sdkerrors.Register(ModuleName, errSecurityMismatchField, "This value does not match the value from security definition request")
	ErrSecurityTypeIsEmpty                     = sdkerrors.Register(ModuleName, errSecurityTypeIsEmpty, "SecurityType is empty")
	ErrSecurityCurrencyIsEmpty                 = sdkerrors.Register(ModuleName, errSecurityCurrencyIsEmpty, "Currency is empty")
	ErrSecurityMinPriceIncrementIsEmpty        = sdkerrors.Register(ModuleName, errSecurityMinPriceIncrementIsEmpty, "MinPriceIncrement is empty")
	ErrSecurityRequestResultIsEmpty            = sdkerrors.Register(ModuleName, errSecurityRequestResultIsEmpty, "SecurityRequestResult is empty")
	ErrSecurityRequestErrorIsEmpty             = sdkerrors.Register(ModuleName, errSecurityRequestErrorIsEmpty, "SecurityRequestError is empty")

	// Order Mass Status
	ErrOrderMassStatusSession               = sdkerrors.Register(ModuleName, errOrderMassStatusSession, "The sessionID is not recognized by the Order Mass Status Request system")
	ErrOrderMassStatusMismatchField         = sdkerrors.Register(ModuleName, errOrderMassStatusMismatchField, "This value does not match the value from the Order")
	ErrOrderMassStatusEmptyField            = sdkerrors.Register(ModuleName, errOrderMassStatusEmptyField, "This field can not be left empty")
	ErrOrderMassStatusIsNotFound            = sdkerrors.Register(ModuleName, errOrderMassStatusIsNotFound, "Order mass status is not found")
	ErrOrderMassStatusCreatorIsWrong        = sdkerrors.Register(ModuleName, errOrderMassStatusCreatorIsWrong, "This account address can not respond to order mass status request")
	ErrOrderMassStatusRequestIsAcknowledged = sdkerrors.Register(ModuleName, errOrderMassStatusRequestIsAcknowledged, "Order Mass Status Request has been responded to")
	ErrOrderMassStatusRequestIsRejected     = sdkerrors.Register(ModuleName, errOrderMassStatusRequestIsRejected, "Order Mass Status Request has been rejected")

	// Trading Session
	ErrTradingSessionEmptyField                  = sdkerrors.Register(ModuleName, errTradingSessionEmptyField, "This field can not be left empty")
	ErrTradingSessionSession                     = sdkerrors.Register(ModuleName, errTradingSessionSession, "The sessionID is not recognized by the Trading Session Status Request system")
	ErrTradingSessionIsNotFound                  = sdkerrors.Register(ModuleName, errTradingSessionIsNotFound, "Trading Session Status Request is not found")
	ErrTradingSessionStatusRequestIsRejected     = sdkerrors.Register(ModuleName, errTradingSessionStatusRequestIsRejected, "Trading Session Status Request has been rejected")
	ErrTradingSessionStatusRequestIsAcknowledged = sdkerrors.Register(ModuleName, errTradingSessionStatusRequestIsAcknowledged, "Trading Session Status Request has been acknowledged")
	ErrTradingSessionMismatchField               = sdkerrors.Register(ModuleName, errTradingSessionMismatchField, "This value does not match the value from Trade Session Status Request")

	// Trading Session List
	ErrTradingSessionListEmptyField            = sdkerrors.Register(ModuleName, errTradingSessionListEmptyField, "This field can not be left empty")
	ErrTradingSessionListSession               = sdkerrors.Register(ModuleName, errTradingSessionListSession, "The sessionID is not recognized by the Trading Session List Request system")
	ErrTradingSessionListIsNotFound            = sdkerrors.Register(ModuleName, errTradingSessionListIsNotFound, "Trading Session List Request is not found")
	ErrTradingSessionListRequestIsRejected     = sdkerrors.Register(ModuleName, errTradingSessionListRequestIsRejected, "Trading Session List Request has been rejected")
	ErrTradingSessionListRequestIsAcknowledged = sdkerrors.Register(ModuleName, errTradingSessionListRequestIsAcknowledged, "Trading Session List Request has been acknowledged")
)
