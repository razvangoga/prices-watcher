package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var urls = GetDataSource()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, pg := range urls {

		go func(pg page) {
			start := time.Now()
			price := GetPrice(pg)
			elapsed := time.Since(start)
			log.Printf("[%s] [%s] - found : %.2f in %s\n", pg.product, pg.site, price, elapsed)
			wg.Done()
		}(pg)
	}

	wg.Wait()
}
