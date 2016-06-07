package main

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"io/ioutil"

	"text/template"

	"log"
	//	"path/filepath"
	"os"
)

var apiSample1 = `
swagger: '2.0'

info:
  version: "0.0.1"
  title: ms-gen sample 1
  description: GET method, without parameter, and returning one value.

paths:

  /get_method_1:
    get:
      description: |
        GET method, without parameter, and returning one value.
      responses:
        200:
          description: Successful response
          schema:
            $ref: "#/definitions/GetMethod1Response"

definitions:

  GetMethod1Response:
    type: object
    properties:
      fieldName1:
        type: string
        description: Field 1
`

func main() {
	fmt.Printf("Starting ms-gen...\n")

	// Reading from file.
	filename := "../generator_tests/api_sample_1.yaml"
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	api := ApiParser(config)

	source := NewSource(&api, "../service_templates/Makefile.tpl")

	source.SaveToFile("code/Makefile")

	fmt.Printf("Finishing ms-gen...\n")
}

func ApiParserTest() {
	// Reading from string.
	api1 := ApiParser([]byte(apiSample1))

	fmt.Printf("api: %v\n", api1)

	fmt.Printf("paths: %v\n", api1.Paths["/get_method_1"])
	fmt.Printf("\tget: %v\n", *(api1.Paths["/get_method_1"].Get))

	fmt.Printf("definitions: %v\n", api1.Definitions["GetMethod1Response"])
	fmt.Printf("\tdefinitions.Properties: %v\n", api1.Definitions["GetMethod1Response"].Properties)

	// Reading from file.
	filename := "../generator_tests/api_sample_1.yaml"
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	api2 := ApiParser(source)

	fmt.Printf("api: %v\n", api2)

	fmt.Printf("paths: %v\n", api2.Paths["/get_method_1"])
	fmt.Printf("\tget: %v\n", *(api2.Paths["/get_method_1"].Get))

	fmt.Printf("definitions: %v\n", api2.Definitions["GetMethod1Response"])
	fmt.Printf("\tdefinitions.Properties: %v\n", api2.Definitions["GetMethod1Response"].Properties)
}

func ApiParser(data []byte) (api Swagger) {
	api = Swagger{}

	err := yaml.Unmarshal(data, &api)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}
