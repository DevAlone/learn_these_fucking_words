package config

import (
	"encoding/json"
	"os"
)

var Settings struct {
	ListenAddress                      string
	Debug                              bool
	RegisterToken                      string
	RegisterForbiddenUsernames         []string
	RegisterUsernameMinLength          int
	RegisterUsernameMaxLength          int
	RegisterPasswordMinLength          int
	RegisterPasswordMaxLength          int
	MaxWordLength                      int
	Database                           map[string]string
	PixabayApiKey                      string
	HttpClientTimeout                  int32
	MemorizationsUpdateTimeDelta       uint32
	MemorizationFullForgettingInDays   float64
	MemorizationMinimumForgettingSpeed float64
	LearningNextShowMinTime            int64
	LearningNextShowMaxTime            int64
}

// var Settings = make(map[string]interface{})

func updateSettingsFromFile(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Settings)

	return err

	/*bytes, err := ioutil.ReadFile(filename)

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

	return nil*/
}

func init() {
	err := updateSettingsFromFile("config/default_settings.json")

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat("config/settings.json"); err == nil {
		err = updateSettingsFromFile("config/settings.json")

		if err != nil {
			panic(err)
		}
	}
}
