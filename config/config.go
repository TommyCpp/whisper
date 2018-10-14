package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port           int `json:"port"`
	MiddlewareSize int `json:"middlewareSize"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Port:           8086,
		MiddlewareSize: 2,
	}
}

func ReadConfig(filename string, configuration *Configuration) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(configuration)
	if err != nil {
		return err
	}
	return nil
}

//func main() {
//	configuration := NewConfiguration()
//	ReadConfig("./config.json", configuration)
//	fmt.Print(configuration.MiddlewareSize)
//}
