package types

import "encoding/json"

type Message struct {
	Type             string          `json:"@type"`
	FromAddress      string          `json:"from_address,omitempty"`
	ToAddress        string          `json:"to_address,omitempty"`
	Amount           json.RawMessage `json:"amount,omitempty"`
	Token            Coin            `json:"token,omitempty"`
	ProposalID       string          `json:"proposal_id,omitempty"`
	Voter            string          `json:"voter,omitempty"`
	Option           string          `json:"option,omitempty"`
	DelegatorAddress string          `json:"delegator_address,omitempty"`
	ValidatorAddress string          `json:"validator_address,omitempty"`
	Sender           string          `json:"sender,omitempty"`
	Receiver         string          `json:"receiver,omitempty"`
	Msgs             []Msg           `json:"msgs,omitempty"`
}

type Msg struct {
	Type      string `json:"@type"`
	Delegator string `json:"delegator_address,omitempty"`
	Validator string `json:"validator_address,omitempty"`
	Amount    Coin   `json:"amount,omitempty"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Transaction struct {
	TxHash   string    `json:"txHash"`
	Messages []Message `json:"messages"`
	Memo     string    `json:"memo"`
	Result   int       `json:"result"`
	RawLog   string    `json:"raw_log"`
	Fee      []Coin    `json:"fee"`
	Height   int64     `json:"height"`
	Time     string    `json:"time"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
