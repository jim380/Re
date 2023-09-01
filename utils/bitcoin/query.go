package bitcoin

type Input struct {
	PrevAddresses []string `json:"prev_addresses"`
	PrevPosition  int      `json:"prev_position"`
	PrevTxHash    string   `json:"prev_tx_hash"`
	PrevType      string   `json:"prev_type"`
	PrevValue     int64    `json:"prev_value"`
	Sequence      uint32   `json:"sequence"`
}

type Output struct {
	Addresses         []string `json:"addresses"`
	Value             int64    `json:"value"`
	Type              string   `json:"type"`
	SpentByTx         string   `json:"spent_by_tx"`
	SpentByTxPosition int      `json:"spent_by_tx_position"`
}

type Transaction struct {
	BlockHeight   int      `json:"block_height"`
	BlockHash     string   `json:"block_hash"`
	BlockTime     int64    `json:"block_time"`
	CreatedAt     int64    `json:"created_at"`
	Confirmations int      `json:"confirmations"`
	Fee           int64    `json:"fee"`
	Hash          string   `json:"hash"`
	InputsCount   int      `json:"inputs_count"`
	InputsValue   int64    `json:"inputs_value"`
	IsCoinbase    bool     `json:"is_coinbase"`
	IsDoubleSpend bool     `json:"is_double_spend"`
	IsSwTx        bool     `json:"is_sw_tx"`
	LockTime      int      `json:"lock_time"`
	OutputsCount  int      `json:"outputs_count"`
	OutputsValue  int64    `json:"outputs_value"`
	Sigops        int      `json:"sigops"`
	Size          int      `json:"size"`
	Version       int      `json:"version"`
	Vsize         int      `json:"vsize"`
	Weight        int      `json:"weight"`
	WitnessHash   string   `json:"witness_hash"`
	Inputs        []Input  `json:"inputs"`
	Outputs       []Output `json:"outputs"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"list"`
}
