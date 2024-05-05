package util

import (
	"fmt"
	"io"
	"net/http"
)

type Osm struct {
}

func Request(path string, queryArr []string) error {
	baseUrl := "https://api.openstreetmap.org"

	query := ""

	url := baseUrl + path + "?" + query
	res, err := http.Get(url)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	fmt.Println("Response from API:", string(body))

	return nil
}
