package main

import (
	"fmt"
	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"log"
	"sync"
	"time"
)

func main() {

	appInsightsClient := appinsights.NewTelemetryClient("5aec4d47-203e-480b-b7d4-da020d468e92")

	var urls = GetDataSource()

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
		// Ten second timeout for retries.

		// If we got here, then all telemetry was submitted
		// successfully, and we can proceed to exiting.
	case <-time.After(30 * time.Second):
		// Thirty second absolute timeout.  This covers any
		// previous telemetry submission that may not have
		// completed before Close was called.

		// There are a number of reasons we could have
		// reached here.  We gave it a go, but telemetry
		// submission failed somewhere.  Perhaps old events
		// were still retrying, or perhaps we're throttled.
		// Either way, we don't want to wait around for it
		// to complete, so let's just exit.
	}
}
