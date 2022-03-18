package chain

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

type TransactionReceipt struct {
	TransactionHash   string
	TransactionIndex  hexutil.Uint64
	BlockHash         string
	BlockNumber       hexutil.Uint64
	From              types.Address
	To                *types.Address
	CumulativeGasUsed hexutil.Uint64
	GasUsed           hexutil.Uint64
	ContractAddress   *types.Address
	Logs              []*Log
	Status            hexutil.Uint64
}
