//
//
package main

import (
	"github.com/tempodb/tempodb-go"
	"github.com/wolfeidau/timeseries/config"
	"log"
)

func main() {
	settings := config.LoadSettings()

	// create a client for tempodb.
	client := tempodb.NewClient(settings.ApiKey, settings.ApiSecret)

	// create the time series
	created, err := client.CreateSeries("my-series3")

	if err != nil {
		log.Fatal("Error occured creating timeseries %s", err)
	}

	log.Printf("timeseries %+v\n", created)

}
