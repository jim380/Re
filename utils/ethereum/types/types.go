package types

type Transaction struct {
	BlockNumber      int64  `json:"blockNumber"`
	Hash             string `json:"hash"`
	Timestamp        int64  `json:"timestamp"`
	TransactionIndex int    `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Nonce            int    `json:"nonce"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
