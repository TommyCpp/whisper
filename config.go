package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port           int `json:"port"`
	MiddlewareSize int `json:"middleware.size"` //todo:test
}

func ReadConfig(filename string, configuration Configuration) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}
}
