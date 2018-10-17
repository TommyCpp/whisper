package config

import (
	"encoding/json"
	"os"
)

var Config *Configuration

type Configuration struct {
	//model.Singleton
	Port              int    `json:"port"`
	MiddlewareSize    int    `json:"middlewareSize"`
	DatabaseDriveName string `json:"database.drive"`
	DatabaseURLName   string `json:"database.url"`
}

func (c *Configuration) Init() {
	c.Port = 8086
	c.MiddlewareSize = 2
	c.DatabaseDriveName = "mysql"
	c.DatabaseURLName = "whisper:123456@127.0.0.1:3306/whisper"
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Port:              8086,
		MiddlewareSize:    2,
		DatabaseDriveName: "mysql",
		DatabaseURLName:   "whisper:123456@127.0.0.1:3306/whisper",
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
//	configuration := new(Configuration)
//	configuration.Init()
//
//	ReadConfig("./config.json", configuration)
//	fmt.Print(configuration.MiddlewareSize)
//}
