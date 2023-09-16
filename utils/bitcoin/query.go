package bitcoin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jim380/Re/utils/bitcoin/types"
)

const (
	bitcoinTxsAPI = "http://localhost:5001/bitcoin/txs?limit=15"
	cacheTTL      = 5 * time.Second
	fetchEvery    = 5 * time.Second
	CacheKey      = "bitcoin"
)


func fetchBitcoinTxs() (*types.TransactionResponse, error) {
	resp, err := http.Get(bitcoinTxsAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactions types.TransactionResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &types.TransactionResponse{}, nil
}
