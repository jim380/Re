package bitcoin

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	bitcoinTxsAPI = "http://localhost:5001/bitcoin/txs?limit=15"
	cacheTTL      = 5 * time.Second
	fetchEvery    = 5 * time.Second
	CacheKey      = "bitcoin"
)

type Input struct {
	PrevAddresses []string `json:"prev_addresses"`
	PrevPosition  int      `json:"prev_position"`
	PrevTxHash    string   `json:"prev_tx_hash"`
	PrevType      string   `json:"prev_type"`
	PrevValue     int      `json:"prev_value"`
	Sequence      int      `json:"sequence"`
}

type Output struct {
	Addresses         []string `json:"addresses"`
	Value             int      `json:"value"`
	Type              string   `json:"type"`
	SpentByTx         string   `json:"spent_by_tx"`
	SpentByTxPosition int      `json:"spent_by_tx_position"`
}

type Transaction struct {
	BlockHeight   int      `json:"block_height"`
	Hash          string   `json:"hash"`
	BlockTime     string   `json:"block_time"`
	Fee           int      `json:"fee"`
	IsDoubleSpend bool     `json:"is_double_spend"`
	OutputsCount  int      `json:"outputs_count"`
	OutputsValue  int      `json:"outputs_value"`
	Inputs        []Input  `json:"inputs"`
	Outputs       []Output `json:"outputs"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}

func fetchBitcoinTxs() (*TransactionResponse, error) {
	resp, err := http.Get(bitcoinTxsAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactions TransactionResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &TransactionResponse{}, nil
}
