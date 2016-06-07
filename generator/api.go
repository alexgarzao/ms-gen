package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"

	"log"
)

type Api struct {
	swagger *Swagger
}

func NewApi(filename string) (api *Api) {
	// Reading from file.
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	api = apiParser(config)

	return
}

func apiParser(data []byte) (api *Api) {
	api = new(Api)
	api.swagger = new(Swagger)

	err := yaml.Unmarshal(data, api.swagger)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}
