package noaa

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/clairBuoyant/noaa/ndbc"
	"github.com/jszwec/csvutil"
)

type client struct {
	client *http.Client
	addr   string
}

func New(uri string) client {
	return client{http.DefaultClient, strings.TrimSuffix(uri, "/")}
}

func (c *client) get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(resp)
		// return c.decodeError(resp)
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "xml") {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(body, result); err != nil {
			panic(err)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO: make this concise approach work, if possible
	// err = csvutil.Unmarshal(body, result)
	// if err != nil {
	// 	return err
	// }

	csvReader := csv.NewReader(strings.NewReader(string(body)))
	csvReader.Comma = ' '
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = 19
	csvReader.TrimLeadingSpace = true

	// in real application this should be done once in init function.
	userHeader, err := csvutil.Header(ndbc.MeteorologicalObservation{}, "csv")
	if err != nil {
		fmt.Println("NewDecoder")
		log.Fatal(err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	if err != nil {
		fmt.Println("NewDecoder")
		log.Fatal(err)
	}
	dec.Map = func(field, column string, v any) string {
		if field == "MM" {
			return "0"
		}
		return field
	}

	// var mos []ndbc.MeteorologicalObservation
	for {
		if err := dec.Decode(result); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) ActiveStations() (ndbc.Stations, error) {
	url := ndbc.ActiveStations

	var s ndbc.Stations
	err := c.get(context.Background(), url, &s)

	return s, err
}

// TODO: work with []ndbc.MeteorologicalObservation
func (c *client) Realtime(id string) string {
	return GetRealtime(id)
}
