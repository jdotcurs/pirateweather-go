package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
)

func main() {
	apiKey := os.Getenv("PIRATE_WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("PIRATE_WEATHER_API_KEY environment variable is not set")
	}

	client := pirateweather.NewClient(apiKey)

	// Forecast example
	forecast, err := client.Forecast(45.42, -75.69,
		pirateweather.WithUnits("si"),
		pirateweather.WithExclude([]string{"minutely"}),
		pirateweather.WithExtend("hourly"),
		pirateweather.WithVersion(2),
	)
	if err != nil {
		log.Fatalf("Error getting forecast: %v", err)
	}

	fmt.Printf("Current temperature in Ottawa: %.2f°C\n", forecast.Currently.Temperature)
	fmt.Printf("Fire index: %.2f\n", forecast.Currently.FireIndex)
	fmt.Printf("Extended hourly forecast available: %v\n", len(forecast.Hourly.Data) > 48)

	// Time Machine example
	pastTimestamp := time.Now().AddDate(0, 0, -1) // Yesterday
	timeMachine, err := client.TimeMachine(45.42, -75.69, pastTimestamp, pirateweather.WithUnits("si"))
	if err != nil {
		log.Fatalf("Error getting time machine data: %v", err)
	}

	fmt.Printf("Temperature in Ottawa yesterday: %.2f°C\n", timeMachine.Currently.Temperature)

	// Print alerts if any
	if len(forecast.Alerts) > 0 {
		fmt.Println("Weather Alerts:")
		for _, alert := range forecast.Alerts {
			fmt.Printf("- %s: %s\n", alert.Title, alert.Description)
		}
	}
}
