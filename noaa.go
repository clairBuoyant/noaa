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
		var mo MeteorologicalObservation
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// TODO: refactor parsing approach
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

		windDirValue, _ := strconv.ParseFloat(rowValues[5], 32)
		mo.WindDirection = int16(windDirValue)

		windSpeedValue, _ := strconv.ParseFloat(rowValues[6], 32)
		mo.WindSpeed = float32(windSpeedValue)

		windGustValue, _ := strconv.ParseFloat(rowValues[7], 32)
		mo.WindGust = float32(windGustValue)

		waveHeightValue, _ := strconv.ParseFloat(rowValues[8], 32)
		mo.WaveHeight = float32(waveHeightValue)

		dominantWavePeriodValue, _ := strconv.ParseFloat(rowValues[9], 32)
		mo.DominantWavePeriod = float32(dominantWavePeriodValue)

		averageWavePeriodValue, _ := strconv.ParseFloat(rowValues[10], 32)
		mo.AverageWavePeriod = float32(averageWavePeriodValue)

		waveDirectionValue, _ := strconv.ParseFloat(rowValues[11], 32)
		mo.WaveDirection = int16(waveDirectionValue)

		seaLevelPresValue, _ := strconv.ParseFloat(rowValues[12], 32)
		mo.SeaLevelPressure = float32(seaLevelPresValue)

		airTempValue, _ := strconv.ParseFloat(rowValues[13], 32)
		mo.AirTemperature = float32(airTempValue)

		waterTempValue, _ := strconv.ParseFloat(rowValues[14], 32)
		mo.WaterTemperature = float32(waterTempValue)

		dewpointTempValue, _ := strconv.ParseFloat(rowValues[15], 32)
		mo.DewpointTemperature = float32(dewpointTempValue)

		visibilityValue, _ := strconv.ParseFloat(rowValues[16], 32)
		mo.Visibility = float32(visibilityValue)

		pressureTendencyVal, _ := strconv.ParseFloat(rowValues[17], 32)
		mo.PressureTendency = float32(pressureTendencyVal)

		tideVal, _ := strconv.ParseFloat(rowValues[18], 32)
		mo.Tide = float32(tideVal)

		mos = append(mos, mo)
	}
	return mos
}
