package gobuoy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	WaterQuality string   `xml:"waterquality,attr"`
}

// func (s Station) String() string {
// 	return fmt.Sprintf("(Station id=%s, name=%s, lat=%f, lon=%f)",
// 		s.ID, s.Name, s.Lat, s.Lon,
// 	)
// }

func GetActiveStations() string {
	url := fmt.Sprintf("%s.%s", ActiveStations, "xml")

	response := request(url)

	var activeStations Stations
	if err := xml.Unmarshal(response, &activeStations); err != nil {
		panic(err)
	}
	activeStationsJson, _ := json.Marshal(activeStations)

	return string(activeStationsJson)
}
