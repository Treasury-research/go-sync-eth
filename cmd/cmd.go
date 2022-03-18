package cmd

import (
	"log"
	"sync"
	_ "sync"
)

var wg sync.WaitGroup

func Execute() {
	sync, err := NewSync()

	if err != nil {
		log.Fatalln("sync.Start err", "err", err)
	}

	err = sync.Start()

}
