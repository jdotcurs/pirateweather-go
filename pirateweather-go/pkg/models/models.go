package models

import "time"

// ForecastResponse represents the response from the Pirate Weather API forecast endpoint
type ForecastResponse struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Timezone  string     `json:"timezone"`
	Offset    float64    `json:"offset"`
	Elevation float64    `json:"elevation"`
	Currently *DataPoint `json:"currently"`
	Minutely  *DataBlock `json:"minutely"`
	Hourly    *DataBlock `json:"hourly"`
	Daily     *DataBlock `json:"daily"`
	Alerts    []Alert    `json:"alerts"`
	Flags     *Flags     `json:"flags"`
	SourceIDX *SourceIDX `json:"sourceIDX,omitempty"`
}

// DataPoint represents a single weather data point
type DataPoint struct {
	Time                        int64   `json:"time"`
	Summary                     string  `json:"summary"`
	Icon                        string  `json:"icon"`
	NearestStormDistance        float64 `json:"nearestStormDistance,omitempty"`
	NearestStormBearing         float64 `json:"nearestStormBearing,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity"`
	PrecipProbability           float64 `json:"precipProbability"`
	PrecipIntensityError        float64 `json:"precipIntensityError"`
	PrecipType                  string  `json:"precipType"`
	Temperature                 float64 `json:"temperature"`
	ApparentTemperature         float64 `json:"apparentTemperature"`
	DewPoint                    float64 `json:"dewPoint"`
	Humidity                    float64 `json:"humidity"`
	Pressure                    float64 `json:"pressure"`
	WindSpeed                   float64 `json:"windSpeed"`
	WindGust                    float64 `json:"windGust"`
	WindBearing                 float64 `json:"windBearing"`
	CloudCover                  float64 `json:"cloudCover"`
	UVIndex                     float64 `json:"uvIndex"`
	Visibility                  float64 `json:"visibility"`
	Ozone                       float64 `json:"ozone"`
	PrecipAccumulation          float64 `json:"precipAccumulation"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         int64   `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          int64   `json:"temperatureLowTime"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime"`
	MoonPhase                   float64 `json:"moonPhase"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime"`
	SunriseTime                 int64   `json:"sunriseTime"`
	SunsetTime                  int64   `json:"sunsetTime"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          int64   `json:"temperatureMinTime"`
	TemperatureMax              float64 `json:"temperatureMax"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime"`
	Smoke                       float64 `json:"smoke,omitempty"`
	FireIndex                   float64 `json:"fireIndex,omitempty"`
	LiquidAccumulation          float64 `json:"liquidAccumulation,omitempty"`
	SnowAccumulation            float64 `json:"snowAccumulation,omitempty"`
	IceAccumulation             float64 `json:"iceAccumulation,omitempty"`
	DawnTime                    int64   `json:"dawnTime,omitempty"`
	DuskTime                    int64   `json:"duskTime,omitempty"`
}

// DataBlock represents a block of weather data points
type DataBlock struct {
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	Data    []DataPoint `json:"data"`
}

// Alert represents a weather alert
type Alert struct {
	Title       string    `json:"title"`
	Regions     []string  `json:"regions"`
	Severity    string    `json:"severity"`
	Time        time.Time `json:"time"`
	Expires     time.Time `json:"expires"`
	Description string    `json:"description"`
	URI         string    `json:"uri"`
}

// Flags represents additional metadata about the forecast
type Flags struct {
	Sources        []string          `json:"sources"`
	SourceTimes    map[string]string `json:"sourceTimes"`
	NearestStation float64           `json:"nearest-station"`
	Units          string            `json:"units"`
	Version        string            `json:"version"`
}

type SourceIDX struct {
	X         int     `json:"x"`
	Y         int     `json:"y"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
