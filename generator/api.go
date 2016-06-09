package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"

	"github.com/alexgarzao/ms-gen/swaggerparser"
)

type (
	Api struct {
		Filename            string
		OutputDir           string
		ServiceName         string
		FriendlyServiceName string
		CommonImportPath    string
		Methods             []*Method
		Definitions         []Definition
		CurrentMethod       *Method
	}

	Definition struct {
		Name       string
		Properties []*Property
	}
)

func NewApi(filename string, outputDir string) (api *Api) {
	api = new(Api)
	api.Filename = filename
	api.OutputDir = outputDir

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
	swagger := new(swaggerparser.Swagger)

	if err := yaml.Unmarshal(text, swagger); err != nil {
		return err
	}

	api.ServiceName = "myservice"
	api.FriendlyServiceName = swagger.Info.Title

	api.Methods = api.fillMethods(swagger.Paths)

	api.Definitions = api.fillDefinitions(swagger.Definitions)

	commonImportPath, err := GetCommonImportPath(api.OutputDir, api.ServiceName)
	if err != nil {
		return err
	}

	api.CommonImportPath = commonImportPath

	return nil
}

// Fill methods.
func (api *Api) fillMethods(pathDefinitions map[string]*swaggerparser.Path) []*Method {
	var methods []*Method
	for k, v := range pathDefinitions {
		if v.Get != nil {
			methods = append(methods, NewMethod(api.ServiceName, k, "Get", v.Get))
		}

		if v.Post != nil {
			methods = append(methods, NewMethod(api.ServiceName, k, "Post", v.Post))
		}

		if v.Put != nil {
			methods = append(methods, NewMethod(api.ServiceName, k, "Put", v.Put))
		}
	}

	return methods
}

// Fill definitions.
func (api *Api) fillDefinitions(apiDefinitions map[string]*swaggerparser.JSONSchema) []Definition {
	var definitions []Definition

	for apiDefinitionKey, apiDefinitionValue := range apiDefinitions {
		definition := Definition{}
		definition.Name = apiDefinitionKey
		for propertyKey, propertyValue := range apiDefinitionValue.Properties {
			property := NewProperty(propertyKey, propertyValue)
			definition.Properties = append(definition.Properties, property)
		}
		definitions = append(definitions, definition)
	}

	return definitions
}
