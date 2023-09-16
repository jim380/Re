package cosmostxs

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/cosmos_txs/types"
	"github.com/patrickmn/go-cache"
)

const (
	cosmosHubTxsAPI = "http://localhost:5001/cosmos/txs?limit=20"
)

var c = cache.New(constants.CacheTTL, 10*time.Minute)

func fetchCosmosTxs() (*types.TransactionResponse, error) {
	resp, err := http.Get(cosmosHubTxsAPI)
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

func GetCachedCosmosTransactions(cacheKey string) (*types.TransactionResponse, error) {
	data, found := c.Get(cacheKey)
	if found {
		return data.(*types.TransactionResponse), nil
	}

	newData, err := fetchCosmosTxs()
	if err != nil {
		return nil, err
	}

	c.Set(cacheKey, newData, cache.DefaultExpiration)
	return newData, err
}

func RefetchCosmosTxsDataPeriodically(cacheKey string) {
	ticker := time.NewTicker(constants.FetchEvery)
	defer ticker.Stop()

	for range ticker.C {
		data, err := fetchCosmosTxs()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			continue
		}

		c.Set(cacheKey, data, cache.DefaultExpiration)
	}
}
