# Pirate Weather Go SDK

## Overview

Pirate Weather Go SDK is a robust, feature-rich Go client library for the Pirate Weather API. This SDK provides a seamless way to interact with Pirate Weather's forecasting services, offering the same functionality as the Python implementation while showcasing best practices in Go development.

## Why This Project?

1. **Open-Source Weather Data**: Pirate Weather provides transparent, open-source weather forecasting, allowing users to understand the origin and processing of their weather data.
2. **Dark Sky API Replacement**: With the shutdown of the Dark Sky API, this SDK offers a drop-in compatible solution for legacy services and new projects alike.
3. **Go Expertise Showcase**: This project demonstrates advanced Go techniques such as concurrent API handling, robust error management, and comprehensive testing.

## Features

- Full implementation of Pirate Weather API endpoints
- Support for both forecast and time machine requests
- Customizable units (SI, Imperial, etc.)
- Extensive error handling and retry logic
- Rate limiting to respect API constraints
- Comprehensive test suite with mocking
- Detailed documentation and examples

## Installation

To install the Pirate Weather Go SDK, use the following command:


```bash
go get github.com/jdotcurs/pirateweather-go
```


## Quick Start

Here's a simple example to get the current weather forecast:

```go
package main
import (
"fmt"
"log"
"os"
"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
)
func main() {
apiKey := os.Getenv("PIRATE_WEATHER_API_KEY")
if apiKey == "" {
log.Fatal("PIRATE_WEATHER_API_KEY environment variable is not set")
}
client := pirateweather.NewClient(apiKey)
forecast, err := client.Forecast(45.42, -75.69,
pirateweather.WithUnits("si"),
pirateweather.WithExclude([]string{"minutely"}),
)
if err != nil {
log.Fatalf("Error getting forecast: %v", err)
}
fmt.Printf("Current temperature: %.2f°C\n", forecast.Currently.Temperature)
}
```

This example demonstrates how to initialize the client, make a forecast request, and handle the response.


## API Key

To use this SDK, you need to obtain an API key from [Pirate Weather](https://pirateweather.net/). After signing up, subscribe to the forecast API. It may take up to 20 minutes for your key to become active.

**Important**: Keep your API key secret and never commit it to version control.


## Advanced Usage

### Time Machine Requests

To get historical weather data:

```go
pastTimestamp := time.Now().AddDate(0, 0, -1) // Yesterday
timeMachine, err := client.TimeMachine(45.42, -75.69, pastTimestamp, pirateweather.WithUnits("si"))
if err != nil {
log.Fatalf("Error getting time machine data: %v", err)
}
fmt.Printf("Temperature yesterday: %.2f°C\n", timeMachine.Currently.Temperature)
```


### Customizing Requests

The SDK supports various options to customize your requests:

```go
forecast, err := client.Forecast(45.42, -75.69,
pirateweather.WithUnits("si"),
pirateweather.WithExclude([]string{"minutely"}),
pirateweather.WithExtend("hourly"),
pirateweather.WithVersion(2),
)
```

### Time Machine Requests with Different Times

You can request weather data for a specific time in the past or future:

```go
pastTime := time.Now().AddDate(-1, 0, 0) // One year ago
forecast, err := client.TimeMachine(45.42, -75.69, pastTime,
    pirateweather.WithUnits("si"),
    pirateweather.WithExclude([]string{"minutely"}),
)
```

### Using Different Unit Systems

The SDK supports four unit systems: si (SI units), us (Imperial units), uk (UK units), and ca (Canada units). Here's how to use them:

```go
// SI units (default)
forecast, err := client.Forecast(45.42, -75.69, pirateweather.WithUnits("si"))

// US units
forecast, err := client.Forecast(45.42, -75.69, pirateweather.WithUnits("us"))

// UK units
forecast, err := client.Forecast(45.42, -75.69, pirateweather.WithUnits("uk"))

// Canada units
forecast, err := client.Forecast(45.42, -75.69, pirateweather.WithUnits("ca"))
```

### Excluding Data Blocks
You can exclude specific data blocks to reduce the amount of data returned:

```go
forecast, err := client.Forecast(45.42, -75.69,
    pirateweather.WithExclude([]string{"minutely", "hourly", "daily", "alerts"}),
)
```

### Using API Version 2

Version 2 of the API includes additional fields such as fireIndex, smoke, dawnTime, and duskTime:

```go
forecast, err := client.Forecast(45.42, -75.69,
    pirateweather.WithVersion(2),
)
```

### Handling Rate Limits

The SDK automatically handles rate limiting. If you exceed the rate limit, the Forecast and TimeMachine methods will return an error:

```go
forecast, err := client.Forecast(45.42, -75.69)
if err != nil {
    if strings.Contains(err.Error(), "rate limit exceeded") {
        // Handle rate limit error
    } else {
        // Handle other errors
    }
}
```

### Error Handling

The SDK includes retry logic for transient errors. If an API request fails after multiple retries, an error will be returned:

```go
forecast, err := client.Forecast(45.42, -75.69)
if err != nil {
    if strings.Contains(err.Error(), "API request failed after") {
        // Handle API failure error
    } else {
        // Handle other errors
    }
}
```

## Go Techniques Showcase

### Concurrent API Handling
- Utilizes Go's powerful concurrency primitives for efficient API requests
- Implements a custom rate limiter to manage concurrent requests while respecting API limits

### Robust Error Management
- Custom error types for different scenarios (APIError, RateLimitError, JSONError)
- Comprehensive error handling with detailed error messages

### Comprehensive Testing
- Extensive unit tests for all major components (see client_test.go and forecast_test.go)
- Mock client implementation for testing API interactions without network calls
- Table-driven tests for thorough coverage of different scenarios

### Flexible API Design
- Functional options pattern for customizable API requests (see WithUnits, WithExclude, WithExtend, WithVersion)
- Chainable method calls for intuitive SDK usage

### Type Safety and Generics
- Leverages Go's type system for compile-time safety
- Utilizes generics for reusable, type-safe utility functions (e.g., ConvertTemperature and ConvertUnit in utils.go)

### Modular Architecture
- Well-organized package structure for easy maintenance and scalability
- Clear separation of concerns between client, models, and utilities

## Contributing

Contributions to the Pirate Weather Go SDK are welcome! Please follow these steps:

1. Fork the repository
2. Create a new branch for your feature
3. Write tests for your new feature
4. Implement your feature
5. Run the test suite
6. Create a pull request

Please ensure your code adheres to the existing style and passes all tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Pirate Weather team for providing the API
- Inspired by the Dark Sky API and its community

## Support

If you find this SDK helpful, consider supporting the Pirate Weather project. Donations help keep the API free and allow for more frequent weather data updates.

## Disclaimer

This SDK is not officially associated with Pirate Weather. It's an independent, open-source project aimed at providing a Go interface to the Pirate Weather API.
