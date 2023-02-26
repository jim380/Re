package keeper

import (
	"crypto/rand"
	"math/big"
)

const (
	sessionNameChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// randomly generate unique session name
func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(sessionNameChars))))
		if err != nil {
			return "", err
		}
		b[i] = sessionNameChars[n.Int64()]
	}
	return string(b), nil
}

func calculateChecksum(msg string) int64 {
	var sum int
	for _, r := range msg {
		sum += int(r)
	}
	sum %= 256
	return int64(sum)
}
