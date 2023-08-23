package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jim380/Re/utils/helpers"
	"github.com/patrickmn/go-cache"
)

const (
	cosmosHubTxsAPI = "http://localhost:5001/cosmos/txs?limit=10"
	cacheTTL        = 5 * time.Second
	fetchEvery      = 5 * time.Second
)

var c = cache.New(cacheTTL, 10*time.Minute)

type Transaction struct {
	TxHash   string    `json:"txHash"`
	Messages []Message `json:"messages"`
	Result   int       `json:"result"`
	Fee      []Coin    `json:"fee"`
	Height   int       `json:"height"`
	Time     string    `json:"time"`
}

type Message struct {
	Type        string          `json:"@type"`
	FromAddress string          `json:"from_address,omitempty"`
	ToAddress   string          `json:"to_address,omitempty"`
	Amount      json.RawMessage `json:"amount,omitempty"`
	Token       Coin            `json:"token,omitempty"`
	ProposalID  string          `json:"proposal_id,omitempty"`
	Voter       string          `json:"voter,omitempty"`
	Option      string          `json:"option,omitempty"`
	Delegator   string          `json:"delegator_address,omitempty"`
	Validator   string          `json:"validator_address,omitempty"`
	Sender      string          `json:"sender,omitempty"`
	Receiver    string          `json:"receiver,omitempty"`
	Msgs        []Msg           `json:"msgs,omitempty"`
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

func main() {
	t, _ := fetchTxs()
	for _, a := range t {
		for _, d := range a.Messages {
			//fmt.Println(d.Type)
			fmt.Println(helpers.AbbrTxMessage(d.Type))
		}
		//fmt.Println(helpers.AbbrTxMessage(a))
	}
}
