package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	. "config"
)

var client = &http.Client{
	Timeout: time.Duration(Settings.HttpClientTimeout) * time.Second,
}

func GetJson(url string) (interface{}, error) {
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
