package cosmostxs

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	cosmosHubTxsAPI = "http://localhost:5001/cosmos/txs?limit=20"
	cacheTTL        = 5 * time.Second
	fetchEvery      = 5 * time.Second
	CacheKey        = "cosmosHub"
)

var c = cache.New(cacheTTL, 10*time.Minute)

type Transaction struct {
	TxHash   string    `json:"txHash"`
	Messages []Message `json:"messages"`
	Memo     string    `json:"memo"`
	Result   int       `json:"result"`
	Raw_log  string    `json:"raw_log"`
	Fee      []Coin    `json:"fee"`
	Height   int64     `json:"height"`
	Time     string    `json:"time"`
}

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

func fetchTxs() ([]Transaction, error) {
	resp, err := http.Get(cosmosHubTxsAPI)
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
	ticker := time.NewTicker(fetchEvery)
	defer ticker.Stop()

	for range ticker.C {
		data, err := fetchTxs()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			continue
		}

		c.Set(cacheKey, data, cache.DefaultExpiration)
	}
}
