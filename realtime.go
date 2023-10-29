package noaa

// TODO: move to ndbc package

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/clairBuoyant/noaa/ndbc"
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

type MeteorologicalObservation struct {
	Datetime            time.Time `csv:"datetime" json:"datetime"`
	WindDirection       int16     `csv:"wind_direction" json:"wind_direction"`
	WindSpeed           float32   `csv:"wind_speed" json:"wind_speed"`
	WindGust            float32   `csv:"wind_gust" json:"wind_gust"`
	WaveHeight          float32   `csv:"wave_height" json:"wave_height"`
	DominantWavePeriod  float32   `csv:"dominant_wave_period" json:"dominant_wave_period"`
	AverageWavePeriod   float32   `csv:"average_wave_period" json:"average_wave_period"`
	WaveDirection       int16     `csv:"wave_direction" json:"wave_direction"`
	SeaLevelPressure    float32   `csv:"sea_level_pressure" json:"sea_level_pressure"`
	PressureTendency    float32   `csv:"pressure_tendency" json:"pressure_tendency"`
	AirTemperature      float32   `csv:"air_temperature" json:"air_temperature"`
	WaterTemperature    float32   `csv:"water_temperature" json:"water_temperature"`
	DewpointTemperature float32   `csv:"dewpoint_temperature" json:"dewpoint_temperature"`
	Visibility          float32   `csv:"visibility" json:"visibility"`
	Tide                float32   `csv:"tide" json:"tide"`
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

// Encodes time.Time object to ISODateString
func (o MeteorologicalObservation) MarshalJSON() ([]byte, error) {
	// TODO: refactor
	metObs := struct {
		Datetime            string  `json:"datetime"`
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
	}{
		Datetime:            o.Datetime.Format((time.RFC3339)),
		WindDirection:       o.WindDirection,
		WindSpeed:           o.WindSpeed,
		WindGust:            o.WindGust,
		WaveHeight:          o.WaveHeight,
		DominantWavePeriod:  o.DominantWavePeriod,
		AverageWavePeriod:   o.AverageWavePeriod,
		WaveDirection:       o.WaveDirection,
		SeaLevelPressure:    o.SeaLevelPressure,
		PressureTendency:    o.PressureTendency,
		AirTemperature:      o.AirTemperature,
		WaterTemperature:    o.WaterTemperature,
		DewpointTemperature: o.DewpointTemperature,
		Visibility:          o.Visibility,
		Tide:                o.Tide,
	}
	return json.Marshal(metObs)
}

// func (mo MeteorologicalObservation) String() string {
// 	return fmt.Sprintf("(MeteorologicalObservation datetime=%s, wind_direction=%v, wind_speed=%v, wind_gust=%v)",
// 		mo.Datetime, mo.WindDirection, mo.WindSpeed, mo.WindGust,
// 	)
// }

// TODO: add support for other realtime datasets and different return signatures (json, struct)
func GetRealtime(stationId string) string {
	url := fmt.Sprintf("%s/%s.%s", ndbc.Realtime, stationId, TXT)

	data := realtimeMeteorological(url)

	jsonData, _ := json.Marshal(data)

	return string(jsonData)
}

func request(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func realtimeMeteorological(url string) []MeteorologicalObservation {
	body := request(url)

	r := csv.NewReader(strings.NewReader(string(body)))
	r.FieldsPerRecord = 0
	r.TrimLeadingSpace = true

	var mos []MeteorologicalObservation
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		mos = append(mos, parseMeteorologicalObservation(record))
	}
	return mos
}

func parseMeteorologicalObservation(record []string) MeteorologicalObservation {
	mo := MeteorologicalObservation{}

	row := record[0]
	trimmed := strings.TrimSpace(row)
	singleSpacePattern := regexp.MustCompile(`\s+`)
	rowValues := strings.Split(singleSpacePattern.ReplaceAllString(trimmed, " "), " ")

	year, _ := strconv.ParseInt(rowValues[0], 10, 16)
	month, _ := strconv.ParseInt(rowValues[1], 10, 16)
	day, _ := strconv.ParseInt(rowValues[2], 10, 16)
	hour, _ := strconv.ParseInt(rowValues[3], 10, 16)
	minute, _ := strconv.ParseInt(rowValues[4], 10, 16)

	mo.Datetime = time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), 0, 0, time.UTC)
	mo.WindDirection = parseInt16(rowValues[5])
	mo.WindSpeed = parseFloat32(rowValues[6])
	mo.WindGust = parseFloat32(rowValues[7])
	mo.WaveHeight = parseFloat32(rowValues[8])
	mo.DominantWavePeriod = parseFloat32(rowValues[9])
	mo.AverageWavePeriod = parseFloat32(rowValues[10])
	mo.WaveDirection = parseInt16(rowValues[11])
	mo.SeaLevelPressure = parseFloat32(rowValues[12])
	mo.AirTemperature = parseFloat32(rowValues[13])
	mo.WaterTemperature = parseFloat32(rowValues[14])
	mo.DewpointTemperature = parseFloat32(rowValues[15])
	mo.Visibility = parseFloat32(rowValues[16])
	mo.PressureTendency = parseFloat32(rowValues[17])
	mo.Tide = parseFloat32(rowValues[18])

	return mo
}

func parseInt16(value string) int16 {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

func parseFloat32(value string) float32 {
	f64, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0.0
	}
	return float32(f64)
}
