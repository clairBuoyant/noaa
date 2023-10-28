package noaa

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
