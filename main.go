//
// Simple test application which uses the tempodb golang library to
// upload a entry ever 500ms for a period of time.
//
package main

import (
	"fmt"
	"github.com/tempodb/tempodb-go"
	"github.com/wolfeidau/timeseries/config"
	"log"
	"time"
)

func main() {
	settings := config.LoadSettings()

	// create a client for tempodb.
	client := tempodb.NewClient(settings.ApiKey, settings.ApiSecret)

	// create the time series
	//	created, err := client.CreateSeries("my-series3")

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("sending data point at", t)

			datapoints := []*tempodb.DataPoint{
				&tempodb.DataPoint{
					Ts: time.Now().UTC(),
					V:  1.23,
				},
			}

			err := client.WriteKey("1_D90343085872CA9C_687_0_2000", datapoints)

			if err != nil {
				log.Fatal("Error occured creating timeseries %s", err)
				return
			}

		}
	}()

	fmt.Println("Writing data for 60 seconds")
	time.Sleep(time.Second * 60)
}
