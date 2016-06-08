package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"strings"
)

type (
	Api struct {
		Filename            string
		OutputDir           string
		ServiceName         string
		FriendlyServiceName string
		CommonImportPath    string
		Paths               []ApiPath
		Definitions         []Definition
		CurrentPath         ApiPath
	}

	ApiPath struct {
		MethodType         string
		PathWithParameters string
		ServiceMethod      string
		CodeFilename       string
		Responses          []ApiResponse
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

	ApiResponse struct {
		ResultCode  string
		Description string
		Ref         string
		Name        string
		Type        string
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
	swagger := new(Swagger)

	if err := yaml.Unmarshal(text, swagger); err != nil {
		return err
	}

	api.ServiceName = "myservice"
	api.FriendlyServiceName = swagger.Info.Title

	api.Paths = api.fillPaths(swagger.Paths)

	api.Definitions = api.fillDefinitions(swagger.Definitions)

	commonImportPath, err := GetCommonImportPath(api.OutputDir, api.ServiceName)
	if err != nil {
		return err
	}

	api.CommonImportPath = commonImportPath

	return nil
}

// Fill paths.
func (api *Api) fillPaths(pathDefinitions map[string]*Path) []ApiPath {
	var paths []ApiPath
	for k, v := range pathDefinitions {
		var methodType = "UNDEFINED_METHOD_TYPE"
		var serviceMethod = "UNDEFINED_SERVICE_METHOD"
		var operation *Operation
		if v.Get != nil {
			methodType = "Get"
			operation = v.Get
		}

		serviceMethod = operation.OperationID

		path := ApiPath{
			MethodType:         methodType,
			PathWithParameters: k,
			ServiceMethod:      strings.Title(serviceMethod),
			CodeFilename:       "service_" + k[1:] + ".go",
		}

		path.Responses = api.fillResponses(operation.Responses)

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
				Type:     ToGolangType(string(propertyValue.Type), ""),
				JsonName: propertyKey,
			}
			definition.Properties = append(definition.Properties, property)
		}
		definitions = append(definitions, definition)
	}

	return definitions
}

// Fill responses.
func (api *Api) fillResponses(apiResponses map[string]*Response) []ApiResponse {
	var responses []ApiResponse

	for apiResponseKey, apiResponseValue := range apiResponses {
		response := ApiResponse{}
		response.ResultCode = apiResponseKey
		completeRef := apiResponseValue.Schema.Ref // "#/definitions/GetMethod1Response"
		response.Ref = completeRef[strings.LastIndex(completeRef, "/")+1:]

		// Help fields.
		response.Name = strings.ToLower(string(response.Ref[0])) + response.Ref[1:]
		response.Type = api.ServiceName + "_common." + response.Ref
		responses = append(responses, response)
	}

	return responses
}
