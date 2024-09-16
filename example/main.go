package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/geocoding"
	"github.com/jdotcurs/pirateweather-go/pkg/models"
	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/jdotcurs/pirateweather-go/pkg/utils"
)

func main() {
	apiKey := os.Getenv("PIRATE_WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("PIRATE_WEATHER_API_KEY environment variable is not set")
	}

	client := pirateweather.NewClient(apiKey)

	// Forecast example
	fmt.Println("Current Forecast:")
	forecast, err := client.Forecast(45.42, -75.69,
		pirateweather.WithUnits("si"),
		pirateweather.WithExclude([]string{"minutely"}),
		pirateweather.WithExtend("hourly"),
		pirateweather.WithVersion(2),
	)
	if err != nil {
		log.Fatalf("Error getting forecast: %v", err)
	}

	printForecastData(forecast)

	// Time Machine example
	fmt.Println("\nTime Machine (Yesterday's Weather):")
	pastTimestamp := time.Now().AddDate(0, 0, -1) // Yesterday
	timeMachine, err := client.TimeMachine(45.42, -75.69, pastTimestamp, pirateweather.WithUnits("si"))
	if err != nil {
		log.Fatalf("Error getting time machine data: %v", err)
	}

	printForecastData(timeMachine)

	// Demonstrate unit conversion
	tempF, err := utils.ConvertTemperature(forecast.Currently.Temperature, "C", "F")
	if err != nil {
		log.Fatalf("Error converting temperature: %v", err)
	}
	fmt.Printf("Current temperature: %.2f°C (%.2f°F)\n", forecast.Currently.Temperature, tempF)

	windSpeedKmh, err := utils.ConvertUnit(forecast.Currently.WindSpeed, "m/s", "km/h")
	if err != nil {
		log.Fatalf("Error converting wind speed: %v", err)
	}
	fmt.Printf("Current wind speed: %.2f m/s (%.2f km/h)\n", forecast.Currently.WindSpeed, windSpeedKmh)

	// Demonstrate different unit systems
	fmt.Println("\nForecast with different unit systems:")
	unitSystems := []string{"si", "us", "uk", "ca"}
	for _, units := range unitSystems {
		forecast, err := client.Forecast(45.42, -75.69, pirateweather.WithUnits(units))
		if err != nil {
			log.Fatalf("Error getting forecast with %s units: %v", units, err)
		}
		fmt.Printf("%s units - Temperature: %.2f, Wind Speed: %.2f\n", units, forecast.Currently.Temperature, forecast.Currently.WindSpeed)
	}
}

func printForecastData(forecast *models.ForecastResponse) {
	fmt.Printf("Location: %.4f, %.4f\n", forecast.Latitude, forecast.Longitude)

	// Get address information
	geocodeResult, err := geocoding.ReverseGeocode(forecast.Latitude, forecast.Longitude)
	if err != nil {
		fmt.Printf("Error getting address information: %v\n", err)
	} else {
		fmt.Printf("Address: %s\n", geocodeResult.DisplayName)
		fmt.Printf("City: %s\n", geocodeResult.Address.City)
		fmt.Printf("State: %s\n", geocodeResult.Address.State)
		fmt.Printf("Country: %s\n", geocodeResult.Address.Country)
	}

	fmt.Printf("Timezone: %s\n", forecast.Timezone)
	fmt.Printf("Time: %s\n", utils.FormatTime(forecast.Currently.Time))
	fmt.Printf("Temperature: %.2f°C\n", forecast.Currently.Temperature)
	fmt.Printf("Feels like: %.2f°C\n", forecast.Currently.ApparentTemperature)
	fmt.Printf("Humidity: %.2f%%\n", forecast.Currently.Humidity*100)
	fmt.Printf("Wind Speed: %.2f m/s\n", forecast.Currently.WindSpeed)
	fmt.Printf("Wind Direction: %.2f°\n", forecast.Currently.WindBearing)
	fmt.Printf("Cloud Cover: %.2f%%\n", forecast.Currently.CloudCover*100)
	fmt.Printf("UV Index: %.1f\n", forecast.Currently.UVIndex)
	fmt.Printf("Visibility: %.2f km\n", forecast.Currently.Visibility)
	fmt.Printf("Fire Index: %.2f\n", forecast.Currently.FireIndex)
	fmt.Printf("Smoke: %.2f\n", forecast.Currently.Smoke)

	if len(forecast.Alerts) > 0 {
		fmt.Println("\nWeather Alerts:")
		for _, alert := range forecast.Alerts {
			fmt.Printf("- %s: %s\n", alert.Title, alert.Description)
		}
	}

	if forecast.Hourly != nil && len(forecast.Hourly.Data) > 0 {
		fmt.Printf("\nHourly forecast available for the next %d hours\n", len(forecast.Hourly.Data))
	}

	if forecast.Daily != nil && len(forecast.Daily.Data) > 0 {
		fmt.Printf("\nDaily forecast available for the next %d days\n", len(forecast.Daily.Data))
	}
}
