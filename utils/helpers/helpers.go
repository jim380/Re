package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/jim380/Re/utils/constants"
)

// randomly generate unique string
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(constants.NameChars))))
		if err != nil {
			return "", err
		}
		b[i] = constants.NameChars[n.Int64()]
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

// AbbrTxMessage abbreviates a message or a list of messages by:
// 1. Removing any prefix up to and including the last "." or "/" character.
// 2. Removing any "Msg" substring.
// If provided an array of strings, the function will also aggregate duplicates
// and provide a count of each unique abbreviated message in the format: "message×count".
// For example:
// Given ["/cosmos.gov.v1beta1.MsgVote", "/cosmos.bank.v1beta1.MsgSend", "/cosmos.bank.v1beta1.MsgSend"],
// the function will return: "Vote, Send×2".
// If provided a single string, it will simply return the abbreviated form.
// For instance:
// Given "/cosmos.bank.v1beta1.MsgSend", the function will return: "Send".
func AbbrTxMessage(msg interface{}) string {
	switch v := msg.(type) {
	case []string:
		sum := make(map[string]int)
		for _, item := range v {
			m := AbbrTxMessage(item)
			sum[m]++
		}

		var output []string
		for k, s := range sum {
			if s > 1 {
				output = append(output, fmt.Sprintf("%s×%d", k, s))
			} else {
				output = append(output, k)
			}
		}
		return strings.Join(output, ", ")

	case string:
		return abbreviate(v)

	default:
		return ""
	}
}

func abbreviate(str string) string {
	lastDot := strings.LastIndex(str, ".")
	if lastDot != -1 {
		str = str[lastDot+1:]
	}
	lastSlash := strings.LastIndex(str, "/")
	if lastSlash != -1 {
		str = str[lastSlash+1:]
	}
	return strings.ReplaceAll(str, "Msg", "")
}
