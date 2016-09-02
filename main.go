package main

import (
	"flag"
	"sync"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	CUrl     = flag.String("url", "http://thanhson1085.github.io", "That you want to crawl")
	NReq     = flag.Int("req", 1, "The number of requests to crawl")
)

var wg sync.WaitGroup

func main() {
	// Parse the command-line flags.
	flag.Parse()

	wg.Add(*NReq)

	// Start the dispatcher.
	StartDispatcher(*NWorkers)

	for i := 0; i < *NReq; i++ {
		Collector(*CUrl)
	}

	wg.Wait()
}
