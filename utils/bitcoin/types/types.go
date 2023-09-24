package types

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
	IsCoinbase    bool     `json:"is_coinbase"`
	IsDoubleSpend bool     `json:"is_double_spend"`
	OutputsCount  int      `json:"outputs_count"`
	OutputsValue  int      `json:"outputs_value"`
	WitnessHash   string   `json:"witness_hash"`
	Inputs        []Input  `json:"inputs"`
	Outputs       []Output `json:"outputs"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
