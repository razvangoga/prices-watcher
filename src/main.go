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

	appInsightsClient := appinsights.NewTelemetryClient(os.Getenv("APPSETTING_ApplicationInsightsInstrumentationKey"))

	var urls = getPagesToWatch()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, pg := range urls {

		go func(pg page) {
			start := time.Now()
			price := GetPrice(pg)
			elapsed := time.Since(start)

			logMessage := fmt.Sprintf("[%s] [%s] - found : %.2f in %s\n", pg.product, pg.site, price, elapsed)
			log.Print(logMessage)

			event := appinsights.NewEventTelemetry(fmt.Sprintf("[%s] [%s]", pg.site, pg.product))
			event.Properties["product"] = pg.product
			event.Properties["site"] = pg.site
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
