package chain

import (
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

type Log struct {
	BlockHash        string
	Address          string
	LogIndex         hexutil.Uint64
	Data             string
	TransactionHash  string
	TransactionIndex hexutil.Uint64
	Removed          bool
	Topics           []string
	BlockNumber      hexutil.Uint64
}
