package cmd

import (
	"github.com/Treasury-research/go-sync-eth/sync/chain"
	"github.com/davecgh/go-spew/spew"
	"log"
	"sync"
)

type Sync struct {
	dbBlockNumber int64
	eth           *chain.Eth
	wg            sync.WaitGroup
}

// var eth, _ = chain.NewETH("https://polygon-mainnet.g.alchemy.com/v2/gg3wYYReOFEBsOYmNYtVK2xFeynOauhv", "eth", 137, true)

func NewSync() (*Sync, error) {

	var eth, err = chain.NewETH("https://polygon-mainnet.g.alchemy.com/v2/gg3wYYReOFEBsOYmNYtVK2xFeynOauhv", "eth", 137, true)
	if err != nil {
		log.Fatalln("chain NewETH ", "err", err)
	}

	return &Sync{
		eth: eth,
	}, nil
}

func (s *Sync) Start() error {
	log.Println("Sync Start !!!")

	s.sync(26068200)

	return nil
}

func (s *Sync) sync(start int) {
	var batchSize = start + 10
	for i := start; i < batchSize; i++ {
		wg.Add(1)
		go func(i int) {
			var block chain.Block
			err := s.eth.GetBlockByNumber(uint64(i), true, &block)

			if err != nil {
				log.Fatalln("Sync RPC GetBlockByNumber ", "err", err)
			}

			log.Printf("sync blockInfo: %v+\n", spew.Sdump(block))

			if block.Transactions != nil && len(block.Transactions) > 0 {
				for _, tx := range block.Transactions {
					log.Printf("block transactions: %v+\n", spew.Sdump(tx))

					var transactionReceipt chain.TransactionReceipt
					err := s.eth.GetTransactionReceipt(tx.Hash, &transactionReceipt)

					if err != nil {
						log.Fatalln("Sync RPC GetTransactionReceipt ", "err", err)
					}

					log.Printf("block transactionReceipt: %v+\n", spew.Sdump(transactionReceipt))

				}
			}

			wg.Done()
		}(i)

	}
	wg.Wait()
}

func parseTxToLog(tx *chain.Transaction) {

}
