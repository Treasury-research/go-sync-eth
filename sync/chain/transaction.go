package chain

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

type Transaction struct {
	Hash             string
	Nonce            hexutil.Uint64
	BlockHash        string
	BlockNumber      hexutil.Uint64
	TransactionIndex hexutil.Uint64
	From             types.Address
	To               *types.Address
	Value            *hexutil.Big
	GasPrice         *hexutil.Big
	GasLimit         hexutil.Uint64
	Input            hexutil.Bytes
}
