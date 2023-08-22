package cosmostxs

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	apiURL     = "https://jsonplaceholder.typicode.com/todos/"
	cacheTTL   = 5 * time.Second
	fetchEvery = 5 * time.Second
)

var c = cache.New(cacheTTL, 10*time.Minute)

type Transaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//Hash  string `json:"hash"`
	// update txs fie
}

func fetchTxs() ([]Transaction, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactions []Transaction
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTxsDataWithCache(cacheKey string) ([]Transaction, error) {
	data, found := c.Get(cacheKey)
	if found {
		return data.([]Transaction), nil
	}

	newData, err := fetchTxs()
	if err != nil {
		return nil, err
	}

	c.Set(cacheKey, newData, cache.DefaultExpiration)
	return newData, err
}

func RefetchTxsDataPeriodically(cacheKey string) {
	for {
		time.Sleep(fetchEvery)

		data, err := fetchTxs()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			continue
		}

		c.Set(cacheKey, data, cache.DefaultExpiration)
	}
}
