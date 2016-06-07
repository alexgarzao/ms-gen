package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"strings"
)

type (
	Api struct {
		Filename            string
		ServiceName         string
		FriendlyServiceName string
		Paths               []ApiPath
		Definitions         []Definition
		CurrentPath         ApiPath
	}

	ApiPath struct {
		MethodType         string
		PathWithParameters string
		ServiceMethod      string
		CodeFilename       string
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
	api.FriendlyServiceName = swagger.Info.Title

	api.Paths = api.fillPaths(swagger.Paths)

	api.Definitions = api.fillDefinitions(swagger.Definitions)

	return nil
}

// Fill paths.
func (api *Api) fillPaths(pathDefinitions map[string]*Path) []ApiPath {
	var paths []ApiPath
	for k, v := range pathDefinitions {
		var methodType = "UNDEFINED_METHOD_TYPE"
		var serviceMethod = "UNDEFINED_SERVICE_METHOD"
		if v.Get != nil {
			methodType = "Get"
			serviceMethod = v.Get.OperationID
		}
		path := ApiPath{
			MethodType:         methodType,
			PathWithParameters: k,
			ServiceMethod:      strings.Title(serviceMethod),
			CodeFilename:       "service_" + k[1:] + ".go",
		}
		paths = append(paths, path)
	}

	return paths

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
