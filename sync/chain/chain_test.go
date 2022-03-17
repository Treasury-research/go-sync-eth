package chain

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestEth_GetLatestBlock(t *testing.T) {
	//eth, _ := NewETH("https://speedy-nodes-nyc.moralis.io/bebf0541f444d9229b16b6b8/eth/mainnet", "eth", 1, true)
	eth, _ := NewETH("https://speedy-nodes-nyc.moralis.io/bebf0541f444d9229b16b6b8/polygon/mainnet", "eth", 137, true)

	var blockInfo Block
	eth.GetLatestBlock(true, &blockInfo)
	spew.Dump(blockInfo)
	var block2 BlockTxHashes
	eth.GetLatestBlock(false, &block2)
	spew.Dump(block2)
}
