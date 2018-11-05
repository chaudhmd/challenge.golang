package main

import (
	"log"
	"time"
	"math/rand"
	"sync"
)

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	// create a ticker with an interval of one second, tickers are used when you want to do something repeatedly at regular intervals
	ticker := time.NewTicker(time.Second)

	// create a new instance of a sync.WaitGroup, named wg
	var wg sync.WaitGroup

	for inx := 0; inx < 10; inx++ {

		// add another go routine to WaitGroup in order to wait for it to finish executing before method problem2 finishes its execution
		wg.Add(1)

		// call the method printRandom2 as a concurrent go routine, pass it the reference for WaitGroup, and ticker
		go printRandom2(inx, &wg, ticker)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	// wait for all the go routines to finish
	wg.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

// add the following parameters to the function printRandom1
// 1. the pointer to WaitGroup instance, wg *sync.WaitGroup
// 2. the pointer to ticker
func printRandom2(slot int, wg *sync.WaitGroup, ticker *time.Ticker) {

	// schedule the call to Done to tell problem2 we are done
	// defer is used to call the wg.Done() method at the end of the enclosing function, i.e., printRandom2
	defer wg.Done();

	for inx := 0; inx < 10; inx++ {

		// receive from the ticker channel, this operation will block until a message is received, a message is received every second
		<-ticker.C

		// print the random number
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}
}