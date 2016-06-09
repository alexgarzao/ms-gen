package main

import (
	"strings"

	"github.com/alexgarzao/ms-gen/swaggerparser"
)

type Method struct {
	MethodType         string
	PathWithParameters string
	ServiceMethod      string
	CodeFilename       string
	Parameters         []Parameter
	Responses          []Response
	Imports            []string
}

func NewMethod(serviceName string, pathWithParameters string, methodType string, operation *swaggerparser.Operation) *Method {

	serviceMethod := operation.OperationID
	pathWithoutParameter := GetPathWithoutParameter(pathWithParameters)

	method := &Method{
		MethodType:    methodType,
		ServiceMethod: strings.Title(serviceMethod),
		CodeFilename:  "service_" + CamelToSnake(operation.OperationID) + ".go",
	}

	method.Parameters = method.fillPathParameters(serviceName, operation.Parameters)
	method.Responses = method.fillResponses(serviceName, operation.Responses)

	pathParamName := operation.GetFirstPathParamName()

	normalizedPath := pathWithoutParameter
	if pathParamName != "" {
		normalizedPath += "/:" + pathParamName
	}

	if operation.GetBodyParamName() != "" {
		method.Imports = append(method.Imports, "fmt")
		method.Imports = append(method.Imports, "net/http")

	}

	method.PathWithParameters = normalizedPath

	return method
}

// Fill path parameters.
func (method *Method) fillPathParameters(serviceName string, swgParameters []*swaggerparser.Parameter) []Parameter {
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
			parameter.Type = serviceName + "_common." + ref
		} else if swgParameter.Type != "" {
			parameter.Type = serviceName + "_common." + swgParameter.Type
		}

		parameters = append(parameters, parameter)
	}

	return parameters
}

// Fill responses.
func (method *Method) fillResponses(serviceName string, apiResponses map[string]*swaggerparser.Response) []Response {
	var responses []Response

	for apiResponseKey, apiResponseValue := range apiResponses {
		response := Response{}
		response.ResultCode = apiResponseKey

		if apiResponseValue.Schema != nil {
			completeRef := apiResponseValue.Schema.Ref // "#/definitions/GetMethod1Response"
			response.Ref = completeRef[strings.LastIndex(completeRef, "/")+1:]

			// Help fields.
			response.Name = strings.ToLower(string(response.Ref[0])) + response.Ref[1:]
			response.Type = serviceName + "_common." + response.Ref
		}
		responses = append(responses, response)
	}

	return responses
}
