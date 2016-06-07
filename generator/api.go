package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"strings"
)

type (
	Api struct {
		Filename    string
		ServiceName string
		MethodList  []string
		Definitions []Definition
	}

	Definition struct {
		Name       string
		Properties []Property
	}

	Property struct {
		Name     string
		Type     string
		JsonName string
	}
)

func NewApi(filename string) (api *Api) {
	api = new(Api)
	api.Filename = filename

	return api
}

func (api *Api) LoadFromSwagger() error {
	// Reading from swagger file.
	config, err := ioutil.ReadFile(api.Filename)
	if err != nil {
		return err
	}

	if err := api.parser(config); err != nil {
		return err
	}

	return nil
}

func (api *Api) parser(text []byte) error {
	swagger := new(Swagger)

	if err := yaml.Unmarshal(text, swagger); err != nil {
		return err
	}

	api.ServiceName = "myservice"

	// Fill method list.
	for k := range swagger.Paths {
		api.MethodList = append(api.MethodList, "service_"+k[1:])
	}

	api.Definitions = api.fillDefinitions(swagger.Definitions)

	return nil
}

// Fill definitions.
func (api *Api) fillDefinitions(apiDefinitions map[string]*JSONSchema) []Definition {
	var definitions []Definition

	for apiDefinitionKey, apiDefinitionValue := range apiDefinitions {
		definition := Definition{}
		definition.Name = apiDefinitionKey
		for propertyKey, propertyValue := range apiDefinitionValue.Properties {
			property := Property{
				Name:     strings.Title(propertyKey),
				Type:     string(propertyValue.Type),
				JsonName: propertyKey,
			}
			definition.Properties = append(definition.Properties, property)
		}
		definitions = append(definitions, definition)
	}

	return definitions
}
