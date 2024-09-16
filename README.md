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

## Documentation

For detailed documentation on all available methods and options, please refer to the [GoDoc](https://pkg.go.dev/github.com/jdotcurs/pirateweather-go).

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



