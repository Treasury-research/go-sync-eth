package chain

import (
	"github.com/Treasury-research/go-sync-eth/pkg/rpc"
	"github.com/Treasury-research/go-sync-eth/pkg/util"
	"github.com/chain5j/chain5j-pkg/math"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"math/big"
	"strings"
)

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

// GetBlockByNumber get block by height
// isFullTx whether return the full transaction
func (eth *Eth) GetBlockByNumber(height uint64, isFullTx bool, result interface{}) error {
	toHex := hexutil.EncodeUint64(height)
	return eth.rpc.Call(&result, eth.clientIdentifier+"_getBlockByNumber", toHex, isFullTx)
}

// GetTransactionReceipt get tx receipt by hash
func (eth *Eth) GetTransactionReceipt(hash string, result interface{}) error {
	return eth.rpc.Call(&result, eth.clientIdentifier+"_getTransactionReceipt", hash)
}

// GetTransactionByHash get transaction by hash
func (eth *Eth) GetTransactionByHash(hash string, result interface{}) error {
	return eth.rpc.Call(&result, eth.clientIdentifier+"_getTransactionByHash", hash)
}

// GetTokenDecimals get token decimals
func (eth *Eth) GetTokenDecimals(contract string) (decimals big.Int, err error) {
	paramsMap := make(map[string]interface{})
	paramsMap["from"] = contract
	paramsMap["to"] = contract
	input := "0x313ce567"
	paramsMap["data"] = input

	extraParam := "latest"
	var result *math.HexOrDecimal256
	err = eth.rpc.Call(&result, eth.clientIdentifier+"_call", paramsMap, extraParam)
	if err != nil {
		return *big.NewInt(0), err
	}
	return big.Int(*result), nil
}

// GetTokenName get token name
func (eth *Eth) GetTokenName(contract string) (name string, err error) {
	paramsMap := make(map[string]interface{})
	paramsMap["from"] = contract
	paramsMap["to"] = contract
	input := "0x06fdde03"
	paramsMap["data"] = input

	extraParam := "latest"
	var result hexutil.Bytes
	err = eth.rpc.Call(&result, eth.clientIdentifier+"_call", paramsMap, extraParam)
	if err != nil {
		return "", err
	}
	resultMap := make(map[int]string)
	resultStr := result.String()
	if strings.HasPrefix(resultStr, "0x") {
		resultStr = resultStr[2:]
	}
	times := len(resultStr) / 64
	for i := 0; i < times; i++ {
		resultMap[i] = util.DecodeStr(resultStr[64*i : 64*(i+1)])
	}

	s := string(util.Hex2Bytes(resultMap[times-1]))
	return s, nil
}

// GetTokenSymbol get token symbol
func (eth *Eth) GetTokenSymbol(contract string) (symbol string, err error) {
	paramsMap := make(map[string]interface{})
	paramsMap["from"] = contract
	paramsMap["to"] = contract
	input := "0x95d89b41"
	paramsMap["data"] = input

	extraParam := "latest"
	var result hexutil.Bytes
	err = eth.rpc.Call(&result, eth.clientIdentifier+"_call", paramsMap, extraParam)
	if err != nil {
		return "", err
	}
	resultMap := make(map[int]string)
	resultStr := result.String()

	if strings.HasPrefix(resultStr, "0x") {
		resultStr = resultStr[2:]
	}
	times := len(resultStr) / 64
	for i := 0; i < times; i++ {
		resultMap[i] = util.DecodeStr(resultStr[64*i : 64*(i+1)])
	}

	s := string(util.Hex2Bytes(resultMap[times-1]))
	return s, nil
}
