package ndbc

import (
	"encoding/xml"
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
