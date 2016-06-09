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
	Parameters         []*Parameter
	Responses          []*Response
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
func (method *Method) fillPathParameters(serviceName string, swgParameters []*swaggerparser.Parameter) []*Parameter {
	var parameters []*Parameter

	for _, swgParameter := range swgParameters {
		parameters = append(parameters, NewParameter(serviceName, swgParameter))
	}

	return parameters
}

// Fill responses.
func (method *Method) fillResponses(serviceName string, apiResponses map[string]*swaggerparser.Response) []*Response {
	var responses []*Response

	for apiResponseKey, apiResponseValue := range apiResponses {
		responses = append(responses, NewResponse(serviceName, apiResponseKey, apiResponseValue))
	}

	return responses
}

// Fill methods.
func FillMethods(serviceName string, pathDefinitions map[string]*swaggerparser.Path) []*Method {
	var methods []*Method
	for k, v := range pathDefinitions {
		if v.Get != nil {
			methods = append(methods, NewMethod(serviceName, k, "Get", v.Get))
		}

		if v.Post != nil {
			methods = append(methods, NewMethod(serviceName, k, "Post", v.Post))
		}

		if v.Put != nil {
			methods = append(methods, NewMethod(serviceName, k, "Put", v.Put))
		}
	}

	return methods
}
