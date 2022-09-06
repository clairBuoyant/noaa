package gobuoy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Datasets
const (
	DATA_SPEC = "data_spec"
	OCEAN     = "ocean"
	SPEC      = "spec"
	SUPL      = "supl"
	SWDIR     = "swdir"
	SWDIR2    = "swdir2"
	SWR1      = "swr1"
	SWR2      = "swr2"
	TXT       = "txt"
)

// Endpoints
const (
	ActiveStations = "https://www.ndbc.noaa.gov/activestations"
	Realtime       = "https://www.ndbc.noaa.gov/data/realtime2"
)

type Stations struct {
	XMLName  xml.Name  `xml:"stations"`
	Stations []Station `xml:"station"`
}

type Station struct {
	XMLName      xml.Name `xml:"station"`
	ID           string   `xml:"id,attr"`
	Lat          float32  `xml:"lat,attr"`
	Lon          float32  `xml:"lon,attr"`
	Name         string   `xml:"name,attr"`
	Owner        string   `xml:"owner,attr"`
	Type         string   `xml:"type,attr"`
	Met          string   `xml:"met,attr"`
	Currents     string   `xml:"currents,attr"`
	Waterquality string   `xml:"waterquality,attr"`
}

type Observation struct {
	Datetime time.Time `json:"datetime"`
}

type MeteorologicalObservation struct {
	Observation

	WindDirection       int16   `json:"wind_direction"`
	WindSpeed           float32 `json:"wind_speed"`
	WindGust            float32 `json:"wind_gust"`
	WaveHeight          float32 `json:"wave_height"`
	DominantWavePeriod  float32 `json:"dominant_wave_period"`
	AverageWavePeriod   float32 `json:"average_wave_period"`
	WaveDirection       int16   `json:"wave_direction"`
	SeaLevelPressure    float32 `json:"sea_level_pressure"`
	PressureTendency    float32 `json:"pressure_tendency"`
	AirTemperature      float32 `json:"air_temperature"`
	WaterTemperature    float32 `json:"water_temperature"`
	DewpointTemperature float32 `json:"dewpoint_temperature"`
	Visibility          float32 `json:"visibility"`
	Tide                float32 `json:"tide"`
}

type WaveSummaryObservation struct {
	Observation

	SignificantWaveHeight float32 `json:"significant_wave_height"`
	SwellHeight           float32 `json:"swell_height"`
	SwellPeriod           float32 `json:"swell_period"`
	WindWaveHeight        float32 `json:"wind_wave_height"`
	WindWavePeriod        float32 `json:"wind_wave_period"`
	SwellDirection        string  `json:"swell_direction"`
	WindWaveDirection     float32 `json:"wind_wave_direction"`
	Steepness             string  `json:"steepness"`
	AverageWavePeriod     float32 `json:"average_wave_period"`
	DominantWaveDirection int16   `json:"dominant_wave_direction"`
}

// func (s Station) String() string {
// 	return fmt.Sprintf("(Station id=%s, name=%s, lat=%f, lon=%f)",
// 		s.ID, s.Name, s.Lat, s.Lon,
// 	)
// }

func makeRequest(url string) Stations {
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var s Stations
	if err := xml.Unmarshal(body, &s); err != nil {
		panic(err)
	}
	return s
}

func GetActiveStations() string {
	url := fmt.Sprintf("%s.%s", ActiveStations, "xml")

	response := makeRequest(url)

	json, _ := json.Marshal(response)

	return string(json)
}
