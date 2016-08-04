package config

import (
    "encoding/json"
    "os"
)

type Configuration struct {
    Port int
    DBServer string
    DBPort int
    DBUsername string
    DBPassword string
    AdjuvantServer string
    AdjuvantInternalNumberAPI string
    AdjuvantAuthKey string
}

func ReadFromFile(path string) (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
	  return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
	  return configuration, err
	}
	return configuration, nil
}
