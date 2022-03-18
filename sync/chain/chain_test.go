package chain

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

// var eth, _ = NewETH("https://speedy-nodes-nyc.moralis.io/bebf0541f444d9229b16b6b8/polygon/mainnet", "eth", 137, true)

var eth, _ = NewETH("https://polygon-mainnet.g.alchemy.com/v2/gg3wYYReOFEBsOYmNYtVK2xFeynOauhv", "eth", 137, true)

func TestEth_GetLatestBlock(t *testing.T) {
	var blockInfo Block
	eth.GetLatestBlock(true, &blockInfo)
	spew.Dump(blockInfo)
	var block2 BlockTxHashes
	eth.GetLatestBlock(false, &block2)
	spew.Dump(block2)
}

func TestEth_GetBlockByNumber(t *testing.T) {
	var blockInfo Block
	eth.GetBlockByNumber(26032924, true, &blockInfo)
	spew.Dump(blockInfo)
	var block2 BlockTxHashes
	eth.GetBlockByNumber(26032924, false, &block2)
	spew.Dump(block2)
}

func TestEth_GetTokenDecimals(t *testing.T) {
	decimals, err := eth.GetTokenDecimals("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	if err != nil {
		panic(err)
	}
	fmt.Println(decimals)
}

func TestEth_GetTokenName(t *testing.T) {
	name, err := eth.GetTokenName("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}

func TestEth_GetTokenSymbol(t *testing.T) {
	name, err := eth.GetTokenSymbol("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}
