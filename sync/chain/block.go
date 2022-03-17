package chain

import "github.com/chain5j/chain5j-pkg/util/hexutil"

type BlockHeader struct {
	Number     hexutil.Uint64 `json:"number"`
	Hash       string
	ParentHash string
	Coinbase   string
	Size       hexutil.Uint64
	Timestamp  hexutil.Uint64
	GasLimit   hexutil.Uint64
	GasUsed    hexutil.Uint64
}

type Block struct {
	BlockHeader
	Transactions []*Transaction
}

// BlockTxHashes block with txHashes
type BlockTxHashes struct {
	BlockHeader
	Transactions []string
}
