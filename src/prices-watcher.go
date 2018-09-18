package main

import (
	"log"
)

func main() {

	var urls = GetDataSource()

	for _, page := range urls {
		log.Printf("Fetching price for [%s] [%s]\n", page.product, page.site)
		log.Printf("Found : %.2f\n", GetPrice(page))
	}
}
