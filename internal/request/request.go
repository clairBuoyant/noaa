package request

import (
	"io"
	"log"
	"net/http"
)

func Request(url string) []byte {
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
