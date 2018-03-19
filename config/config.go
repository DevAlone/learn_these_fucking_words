package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var Settings = make(map[string]interface{})

func updateFromFile(config map[string]interface{}, filename string) error {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.New("file: " + err.Error())
	}

	var result map[string]interface{}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	for key, val := range result {
		config[key] = val
	}

	return nil
}

func init() {
	err := updateFromFile(Settings, "config/default_settings.json")

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat("config/settings.json"); err == nil {
		err = updateFromFile(Settings, "config/settings.json")

		if err != nil {
			panic(err)
		}
	}
}
