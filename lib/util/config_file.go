package util

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig(configs interface{}) error {
	configsJSON, err := ioutil.ReadFile("./config.json")

	if err == nil {
		err := json.Unmarshal(configsJSON, &configs)
		if err != nil {
			return err
		}
	}
	return nil
}
