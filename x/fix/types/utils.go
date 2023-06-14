package types

import (
	"crypto/rand"
	"math/big"
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

	LoggedInStatus = "loggedIn"
)

// randomly generate unique string
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(NameChars))))
		if err != nil {
			return "", err
		}
		b[i] = NameChars[n.Int64()]
	}
	return string(b), nil
}

// calcualtes the messgae length
func BodyLength(msgLength string) int64 {
	logonMsg := msgLength
	length := len(logonMsg)
	return int64(length)
}

// calculates length of checkSum in standard Trailer
func CalculateChecksum(msg string) int64 {
	var sum int
	for _, r := range msg {
		sum += int(r)
	}
	sum %= 256
	return int64(sum)
}
