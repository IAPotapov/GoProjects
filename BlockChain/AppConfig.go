package main

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Text       string
	JobNumber  int
	Complexity int
}

func (con *AppConfig) Save() (err error) {
	var file *os.File
	file, err = os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(con)
	}
	return
}

func (con *AppConfig) Load() (err error) {
	var file *os.File
	file, err = os.Open("config.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(con)
	}
	return
}
