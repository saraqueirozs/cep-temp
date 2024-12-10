package model

import "strconv"

type WeatherResponse struct {
	Current struct {
		TemperatureCelsius float64 `json:"temp_c"`
	} `json:"current"`
}

type TemperatureData struct {
	Celsius    Float64Marshal `json:"temp_C"`
	Fahrenheit Float64Marshal `json:"temp_F"`
	Kelvin     Float64Marshal `json:"temp_K"`
}

type Float64Marshal float64

func (f Float64Marshal) MarshalJSON() ([]byte, error) {
	// Convert the float64 value to a string with full precision
	s := strconv.FormatFloat(float64(f), 'f', -1, 64)
	return []byte(s), nil
}
