package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"strings"

	"github.com/alexgarzao/ms-gen/swaggerparser"
)

type (
	Api struct {
		Filename            string
		OutputDir           string
		ServiceName         string
		FriendlyServiceName string
		CommonImportPath    string
		Paths               []Path
		Definitions         []Definition
		CurrentPath         Path
	}

	Path struct { // TODO: In really, its a method.
		MethodType         string
		PathWithParameters string
		ServiceMethod      string
		CodeFilename       string
		Parameters         []Parameter
		Responses          []Response
		Imports            []string
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

	Parameter struct {
		Name        string
		In          string
		Description string
		Required    bool
		Type        string
		Format      string
	}

	Response struct {
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
	swagger := new(swaggerparser.Swagger)

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
func (api *Api) fillPaths(pathDefinitions map[string]*swaggerparser.Path) []Path {
	var paths []Path
	for k, v := range pathDefinitions {
		if v.Get != nil {
			paths = append(paths, api.newMethod(k, "Get", v.Get))
		}

		if v.Post != nil {
			paths = append(paths, api.newMethod(k, "Post", v.Post))
		}

		if v.Put != nil {
			paths = append(paths, api.newMethod(k, "Put", v.Put))
		}
	}

	return paths
}

func (api *Api) newMethod(pathWithParameters string, methodType string, operation *swaggerparser.Operation) Path {

	serviceMethod := operation.OperationID
	pathWithoutParameter := GetPathWithoutParameter(pathWithParameters)

	path := Path{
		MethodType:    methodType,
		ServiceMethod: strings.Title(serviceMethod),
		CodeFilename:  "service_" + CamelToSnake(operation.OperationID) + ".go",
	}

	path.Parameters = api.fillPathParameters(operation.Parameters)
	path.Responses = api.fillResponses(operation.Responses)

	pathParamName := api.getPathParamName(operation.Parameters)

	normalizedPath := pathWithoutParameter
	if pathParamName != "" {
		normalizedPath += "/:" + pathParamName
	}

	if api.getBodyParamName(operation.Parameters) != "" {
		path.Imports = append(path.Imports, "fmt")
		path.Imports = append(path.Imports, "net/http")

	}

	path.PathWithParameters = normalizedPath

	return path
}

// Fill path parameters.
func (api *Api) fillPathParameters(swgParameters []*swaggerparser.Parameter) []Parameter {
	var parameters []Parameter

	for _, swgParameter := range swgParameters {
		parameter := Parameter{
			Name:        swgParameter.Name,
			In:          swgParameter.In,
			Description: swgParameter.Description,
			Required:    swgParameter.Required,
			Format:      swgParameter.Format,
		}

		if swgParameter.Schema != nil {
			completeRef := swgParameter.Schema.Ref // "#/definitions/GetMethod1Response"
			ref := completeRef[strings.LastIndex(completeRef, "/")+1:]
			parameter.Type = api.ServiceName + "_common." + ref
		} else if swgParameter.Type != "" {
			parameter.Type = api.ServiceName + "_common." + swgParameter.Type
		}

		parameters = append(parameters, parameter)
	}

	return parameters
}

// Get the first path parameter name.
func (api *Api) getPathParamName(swgParameters []*swaggerparser.Parameter) string {
	for _, swgParameter := range swgParameters {
		if swgParameter.In == "path" {
			return swgParameter.Name
		}
	}

	return ""
}

// Get the first body parameter name.
func (api *Api) getBodyParamName(swgParameters []*swaggerparser.Parameter) string {
	for _, swgParameter := range swgParameters {
		if swgParameter.In == "body" {
			return swgParameter.Name
		}
	}

	return ""
}

// Fill definitions.
func (api *Api) fillDefinitions(apiDefinitions map[string]*swaggerparser.JSONSchema) []Definition {
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
func (api *Api) fillResponses(apiResponses map[string]*swaggerparser.Response) []Response {
	var responses []Response

	for apiResponseKey, apiResponseValue := range apiResponses {
		response := Response{}
		response.ResultCode = apiResponseKey

		if apiResponseValue.Schema != nil {
			completeRef := apiResponseValue.Schema.Ref // "#/definitions/GetMethod1Response"
			response.Ref = completeRef[strings.LastIndex(completeRef, "/")+1:]

			// Help fields.
			response.Name = strings.ToLower(string(response.Ref[0])) + response.Ref[1:]
			response.Type = api.ServiceName + "_common." + response.Ref
		}
		responses = append(responses, response)
	}

	return responses
}
