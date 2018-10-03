package main

import (
	"fmt"
	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	appInsightsClient := appinsights.NewTelemetryClient(os.Getenv("APPSETTING_ApplicationInsightsInstrumentationKey"))

	var dataSource IDataSource

	if accountName != "" {
		dataSource = AzureBlobDataSource { accountKey:accountKey, accountName:accountName}
	} else {
		dataSource = StaticDataSource { }
	}

	var urls = dataSource.get()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, pg := range urls {

		go func(pg page) {
			start := time.Now()
			price := GetPrice(pg)
			elapsed := time.Since(start)

			logMessage := fmt.Sprintf("[%s] [%s] - found : %.2f in %s\n", pg.Product, pg.Site, price, elapsed)
			log.Print(logMessage)

			event := appinsights.NewEventTelemetry(fmt.Sprintf("[%s] [%s]", pg.Site, pg.Product))
			event.Properties["product"] = pg.Product
			event.Properties["site"] = pg.Site
			event.Properties["price"] = fmt.Sprintf("%v", price)
			event.Properties["elapsed"] = fmt.Sprintf("%v", elapsed)

			appInsightsClient.Track(event)

			wg.Done()
		}(pg)
	}

	wg.Wait()
	select {
	case <-appInsightsClient.Channel().Close(10 * time.Second):
	case <-time.After(30 * time.Second):
	}
}
