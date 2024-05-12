package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Options struct {
	Method  string      `json:"method"`
	Url     string      `json:"url"`
	Headers http.Header `json:"headers"`
	Body    interface{} `json:"body"`
}

func Get(options *Options) (interface{}, error) {
	client := http.DefaultClient

	var reqBody []byte

	if options.Body != nil {
		reqBody, _ = json.Marshal(options.Body)
	}

	fmt.Println("reqBody", string(reqBody))

	req, err := http.NewRequest(options.Method, options.Url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	req.Header = options.Headers

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var decodedBody interface{}
	err = json.NewDecoder(res.Body).Decode(&decodedBody)

	if err != nil {
		return nil, err
	}

	return decodedBody, nil
}

func Post(options *Options) (interface{}, error) {
	options.Method = "POST"

	return Get(options)
}
