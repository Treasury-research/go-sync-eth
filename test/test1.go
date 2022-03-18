package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fun1(i)
	}

	wg.Wait()
	log.Println("split")
	main()
}

func fun1(i int) {
	log.Println("fun1 s%", i)
	wg.Done()
}
