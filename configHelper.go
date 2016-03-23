package webutils

import (
	"encoding/json"
	"os"
)

func ReadConfig(configpath string, config interface{}) error {
	file, _ := os.Open(configpath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}
