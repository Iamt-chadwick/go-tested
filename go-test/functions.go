package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type responseJSON struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type appConfig struct {
	Port  string
	Host  string
}
func sendSimpleResponse(w http.ResponseWriter, status bool, message string) {
	json.NewEncoder(w).Encode(
		responseJSON{
			Status:  status,
			Message: message,
		})
}

func readConf(conf string) appConfig {
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		panic(err)
	}
	obj := appConfig{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}
	return obj
}