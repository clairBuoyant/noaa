package noaa

import (
	"context"
)

// Client options for NOAA NDFD data.
const (
	// NDFD parameters are not summarized.
	NotSummarized = "ndfdXMLclient"

	// NDFD parameters are summarized by 24- or 12-hourly periods.
	SummarizedByDay = "ndfdBrowserClientByDay"
)

func (c *client) GetForecasts() (interface{}, error) {
	// TODO: dynamically set params and client options
	urlWithParams := Forecasts + "?lat=40.369&lon=-73.702&begin=2023-02-01T00:00&end=2023-02-03T23:59&waveh=waveh&wdir=wdir&wspd=wspd&wgust=wgust"

	var s Dwml
	err := c.get(context.Background(), urlWithParams, &s)

	return s, err
}
