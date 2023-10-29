package ndbc

import (
	"time"
)

// Realtime Datasets
//
//	There are nine different data sources:
//	  - data_spec     Raw Spectral Wave Data
//	  - ocean         Oceanographic Data
//	  - spec          Spectral Wave Summary Data
//	  - supl          Supplemental Measurements Data
//	  - swdir         Spectral Wave Data (alpha1)
//	  - swdir2        Spectral Wave Data (alpha2)
//	  - swr1          Spectral Wave Data (r1)
//	  - swr2          Spectral Wave Data (r2)
//	  - txt           Standard Meteorological Data
const (
	DATASPEC = "data_spec"
	OCEAN    = "ocean"
	SPEC     = "spec"
	SUPL     = "supl"
	SWDIR    = "swdir"
	SWDIR2   = "swdir2"
	SWR1     = "swr1"
	SWR2     = "swr2"
	TXT      = "txt"
)

type ObservationDate struct {
	Year   int `csv:"YY"`
	Month  int `csv:"MM"`
	Day    int `csv:"DD"`
	Hour   int `csv:"hh"`
	Minute int `csv:"mm"`
}

type MetObs struct {
	Data []MeteorologicalObservation `csv:",inline"`
}

type MeteorologicalObservation struct {
	// Datetime            time.Time `json:"datetime"`
	Datetime            ObservationDate `csv:",inline"`
	WindDirection       int16           `csv:"WDIR" json:"wind_direction"`
	WindSpeed           float32         `csv:"WSPD" json:"wind_speed"`
	WindGust            float32         `csv:"GST" json:"wind_gust"`
	WaveHeight          float32         `csv:"WVHT" json:"wave_height"`
	DominantWavePeriod  float32         `csv:"DPD" json:"dominant_wave_period"`
	AverageWavePeriod   float32         `csv:"APD" json:"average_wave_period"`
	WaveDirection       int16           `csv:"MWD" json:"wave_direction"`
	SeaLevelPressure    float32         `csv:"PRES" json:"sea_level_pressure"`
	PressureTendency    float32         `csv:"PTDY" json:"pressure_tendency"`
	AirTemperature      float32         `csv:"ATMP" json:"air_temperature"`
	WaterTemperature    float32         `csv:"WTMP" json:"water_temperature"`
	DewpointTemperature float32         `csv:"DEWP" json:"dewpoint_temperature"`
	Visibility          float32         `csv:"VIS" json:"visibility"`
	Tide                float32         `csv:"TIDE" json:"tide"`
}

type WaveSummaryObservation struct {
	Datetime              time.Time `json:"datetime"`
	SignificantWaveHeight float32   `json:"significant_wave_height"`
	SwellHeight           float32   `json:"swell_height"`
	SwellPeriod           float32   `json:"swell_period"`
	WindWaveHeight        float32   `json:"wind_wave_height"`
	WindWavePeriod        float32   `json:"wind_wave_period"`
	SwellDirection        string    `json:"swell_direction"`
	WindWaveDirection     float32   `json:"wind_wave_direction"`
	Steepness             string    `json:"steepness"`
	AverageWavePeriod     float32   `json:"average_wave_period"`
	DominantWaveDirection int16     `json:"dominant_wave_direction"`
}

// TODO: handle time.Time object from multi-col and standardize parsing
// Encodes time.Time object to ISODateString
// func (o MeteorologicalObservation) MarshalJSON() ([]byte, error) {
// 	// TODO: refactor
// 	metObs := struct {
// 		Datetime            string  `json:"datetime"`
// 		WindDirection       int16   `json:"wind_direction"`
// 		WindSpeed           float32 `json:"wind_speed"`
// 		WindGust            float32 `json:"wind_gust"`
// 		WaveHeight          float32 `json:"wave_height"`
// 		DominantWavePeriod  float32 `json:"dominant_wave_period"`
// 		AverageWavePeriod   float32 `json:"average_wave_period"`
// 		WaveDirection       int16   `json:"wave_direction"`
// 		SeaLevelPressure    float32 `json:"sea_level_pressure"`
// 		PressureTendency    float32 `json:"pressure_tendency"`
// 		AirTemperature      float32 `json:"air_temperature"`
// 		WaterTemperature    float32 `json:"water_temperature"`
// 		DewpointTemperature float32 `json:"dewpoint_temperature"`
// 		Visibility          float32 `json:"visibility"`
// 		Tide                float32 `json:"tide"`
// 	}{
// 		Datetime:            o.Datetime.Format((time.RFC3339)),
// 		WindDirection:       o.WindDirection,
// 		WindSpeed:           o.WindSpeed,
// 		WindGust:            o.WindGust,
// 		WaveHeight:          o.WaveHeight,
// 		DominantWavePeriod:  o.DominantWavePeriod,
// 		AverageWavePeriod:   o.AverageWavePeriod,
// 		WaveDirection:       o.WaveDirection,
// 		SeaLevelPressure:    o.SeaLevelPressure,
// 		PressureTendency:    o.PressureTendency,
// 		AirTemperature:      o.AirTemperature,
// 		WaterTemperature:    o.WaterTemperature,
// 		DewpointTemperature: o.DewpointTemperature,
// 		Visibility:          o.Visibility,
// 		Tide:                o.Tide,
// 	}
// 	return json.Marshal(metObs)
// }
