package chain

import "github.com/Treasury-research/go-sync-eth/pkg/rpc"

type Eth struct {
	rpc              rpc.JsonRpc
	clientIdentifier string
	chainId          int64
	isEip155         bool
}

// NewETH new eth struct
func NewETH(host string, clientIdentifier string, chainId int64, isEip155 bool) (*Eth, error) {
	rpc, err := rpc.NewRpc(host)
	if err != nil {
		return nil, err
	}
	return &Eth{
		rpc:              rpc,
		clientIdentifier: clientIdentifier,
		chainId:          chainId,
		isEip155:         isEip155,
	}, nil
}

// GetLatestBlock get latest block
func (eth *Eth) GetLatestBlock(isFullTx bool, result interface{}) error {
	return eth.rpc.Call(&result, eth.clientIdentifier+"_getBlockByNumber", "latest", isFullTx)
}
