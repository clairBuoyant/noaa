package noaa

import "encoding/xml"

type Dwml struct {
	XMLName xml.Name `xml:"dwml" json:"dwml,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Head    struct {
		Text    string `xml:",chardata" json:"text,omitempty"`
		Product struct {
			SrsName         string `xml:"srsName,attr" json:"srsname,omitempty"`
			ConciseName     string `xml:"concise-name,attr" json:"concise-name,omitempty"`
			OperationalMode string `xml:"operational-mode,attr" json:"operational-mode,omitempty"`
			Title           string `xml:"title"`
			Field           string `xml:"field"`
			Category        string `xml:"category"`
			CreationDate    struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				RefreshFrequency string `xml:"refresh-frequency,attr" json:"refresh-frequency,omitempty"`
			} `xml:"creation-date" json:"creation-date,omitempty"`
		} `xml:"product" json:"product,omitempty"`
		Source struct {
			Text             string `xml:",chardata" json:"text,omitempty"`
			MoreInformation  string `xml:"more-information"`
			ProductionCenter struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				SubCenter string `xml:"sub-center"`
			} `xml:"production-center" json:"production-center,omitempty"`
			Disclaimer string `xml:"disclaimer"`
		} `xml:"source" json:"source,omitempty"`
	} `xml:"head" json:"head,omitempty"`
	Data struct {
		Text     string `xml:",chardata" json:"text,omitempty"`
		Location struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			LocationKey string `xml:"location-key"`
			Point       struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Latitude  string `xml:"latitude,attr" json:"latitude,omitempty"`
				Longitude string `xml:"longitude,attr" json:"longitude,omitempty"`
			} `xml:"point" json:"point,omitempty"`
		} `xml:"location" json:"location,omitempty"`
		MoreWeatherInformation struct {
			Text               string `xml:",chardata" json:"text,omitempty"`
			ApplicableLocation string `xml:"applicable-location,attr" json:"applicable-location,omitempty"`
		} `xml:"moreWeatherInformation" json:"moreweatherinformation,omitempty"`
		TimeLayout []struct {
			Text           string   `xml:",chardata" json:"text,omitempty"`
			TimeCoordinate string   `xml:"time-coordinate,attr" json:"time-coordinate,omitempty"`
			Summarization  string   `xml:"summarization,attr" json:"summarization,omitempty"`
			LayoutKey      string   `xml:"layout-key"`
			StartValidTime []string `xml:"start-valid-time"`
			EndValidTime   []string `xml:"end-valid-time"`
		} `xml:"time-layout" json:"time-layout,omitempty"`
		Parameters struct {
			Text               string `xml:",chardata" json:"text,omitempty"`
			ApplicableLocation string `xml:"applicable-location,attr" json:"applicable-location,omitempty"`
			Temperature        []struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Type       string `xml:"type,attr" json:"type,omitempty"`
				Units      string `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string `xml:"name"`
				Value      []struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
				} `xml:"value" json:"value,omitempty"`
			} `xml:"temperature" json:"temperature,omitempty"`
			Precipitation []struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Type       string `xml:"type,attr" json:"type,omitempty"`
				Units      string `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string `xml:"name"`
				Value      []struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
				} `xml:"value" json:"value,omitempty"`
			} `xml:"precipitation" json:"precipitation,omitempty"`
			ProbabilityOfPrecipitation struct {
				Text       string   `xml:",chardata" json:"text,omitempty"`
				Type       string   `xml:"type,attr" json:"type,omitempty"`
				Units      string   `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string   `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string   `xml:"name"`
				Value      []string `xml:"value"`
			} `xml:"probability-of-precipitation" json:"probability-of-precipitation,omitempty"`
			ConvectiveHazard []struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Outlook struct {
					Text       string   `xml:",chardata" json:"text,omitempty"`
					TimeLayout string   `xml:"time-layout,attr" json:"time-layout,omitempty"`
					Name       string   `xml:"name"`
					Value      []string `xml:"value"`
				} `xml:"outlook" json:"outlook,omitempty"`
				SevereComponent struct {
					Text       string   `xml:",chardata" json:"text,omitempty"`
					Type       string   `xml:"type,attr" json:"type,omitempty"`
					Units      string   `xml:"units,attr" json:"units,omitempty"`
					TimeLayout string   `xml:"time-layout,attr" json:"time-layout,omitempty"`
					Name       string   `xml:"name"`
					Value      []string `xml:"value"`
				} `xml:"severe-component" json:"severe-component,omitempty"`
			} `xml:"convective-hazard" json:"convective-hazard,omitempty"`
			WindSpeed []struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Type       string `xml:"type,attr" json:"type,omitempty"`
				Units      string `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string `xml:"name"`
				Value      []struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
				} `xml:"value" json:"value,omitempty"`
			} `xml:"wind-speed" json:"wind-speed,omitempty"`
			Direction struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Type       string `xml:"type,attr" json:"type,omitempty"`
				Units      string `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string `xml:"name"`
				Value      []struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
				} `xml:"value" json:"value,omitempty"`
			} `xml:"direction" json:"direction,omitempty"`
			CloudAmount struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Type       string `xml:"type,attr" json:"type,omitempty"`
				Units      string `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string `xml:"name"`
				Value      []struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
					Nil  string `xml:"nil,attr" json:"nil,omitempty"`
				} `xml:"value" json:"value,omitempty"`
			} `xml:"cloud-amount" json:"cloud-amount,omitempty"`
			Humidity struct {
				Text       string   `xml:",chardata" json:"text,omitempty"`
				Type       string   `xml:"type,attr" json:"type,omitempty"`
				Units      string   `xml:"units,attr" json:"units,omitempty"`
				TimeLayout string   `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string   `xml:"name"`
				Value      []string `xml:"value"`
			} `xml:"humidity" json:"humidity,omitempty"`
			Weather struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				TimeLayout        string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name              string `xml:"name"`
				WeatherConditions []struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Value struct {
						Text        string `xml:",chardata" json:"text,omitempty"`
						Coverage    string `xml:"coverage,attr" json:"coverage,omitempty"`
						Intensity   string `xml:"intensity,attr" json:"intensity,omitempty"`
						WeatherType string `xml:"weather-type,attr" json:"weather-type,omitempty"`
						Qualifier   string `xml:"qualifier,attr" json:"qualifier,omitempty"`
						Visibility  struct {
							Text string `xml:",chardata" json:"text,omitempty"`
							Nil  string `xml:"nil,attr" json:"nil,omitempty"`
						} `xml:"visibility" json:"visibility,omitempty"`
					} `xml:"value" json:"value,omitempty"`
				} `xml:"weather-conditions" json:"weather-conditions,omitempty"`
			} `xml:"weather" json:"weather,omitempty"`
			ConditionsIcon struct {
				Text       string   `xml:",chardata" json:"text,omitempty"`
				Type       string   `xml:"type,attr" json:"type,omitempty"`
				TimeLayout string   `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name       string   `xml:"name"`
				IconLink   []string `xml:"icon-link"`
			} `xml:"conditions-icon" json:"conditions-icon,omitempty"`
			Hazards struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				TimeLayout       string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Name             string `xml:"name"`
				HazardConditions []struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					Hazard struct {
						Text          string `xml:",chardata" json:"text,omitempty"`
						HazardCode    string `xml:"hazardCode,attr" json:"hazardcode,omitempty"`
						Phenomena     string `xml:"phenomena,attr" json:"phenomena,omitempty"`
						Significance  string `xml:"significance,attr" json:"significance,omitempty"`
						HazardType    string `xml:"hazardType,attr" json:"hazardtype,omitempty"`
						HazardTextURL string `xml:"hazardTextURL"`
					} `xml:"hazard" json:"hazard,omitempty"`
				} `xml:"hazard-conditions" json:"hazard-conditions,omitempty"`
			} `xml:"hazards" json:"hazards,omitempty"`
			WaterState struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				TimeLayout string `xml:"time-layout,attr" json:"time-layout,omitempty"`
				Waves      struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Type  string `xml:"type,attr" json:"type,omitempty"`
					Units string `xml:"units,attr" json:"units,omitempty"`
					Name  string `xml:"name"`
					Value []struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Nil  string `xml:"nil,attr" json:"nil,omitempty"`
					} `xml:"value" json:"value,omitempty"`
				} `xml:"waves" json:"waves,omitempty"`
			} `xml:"water-state" json:"water-state,omitempty"`
		} `xml:"parameters" json:"parameters,omitempty"`
	} `xml:"data" json:"data,omitempty"`
}
