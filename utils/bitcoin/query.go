package bitcoin

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jim380/Re/utils/bitcoin/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/patrickmn/go-cache"
)

const (
	bitcoinTxsAPI = "http://localhost:5001/bitcoin/txs?limit=15"
)

func fetchBitcoinTxs() (*types.TransactionResponse, error) {
	resp, err := http.Get(bitcoinTxsAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactions []types.Transaction
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&transactions)
	if err != nil {
		return nil, err
	}

	// Create TransactionResponse object and set the transactions field
	response := &types.TransactionResponse{
		Transactions: transactions,
	}

	return response, nil
}

func GetCachedBitcoinTransactions(cacheKey string) (*types.TransactionResponse, error) {
	data, found := constants.C.Get(cacheKey)
	if found {
		return data.(*types.TransactionResponse), nil
	}

	newData, err := fetchBitcoinTxs()
	if err != nil {
		return nil, err
	}

	constants.C.Set(cacheKey, newData, cache.DefaultExpiration)
	return newData, err
}

func RefetchBitcoinTxsDataPeriodically(cacheKey string) {
	ticker := time.NewTicker(constants.FetchEvery)
	defer ticker.Stop()

	for range ticker.C {
		data, err := fetchBitcoinTxs()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			continue
		}

		constants.C.Set(cacheKey, data, cache.DefaultExpiration)
	}
}
