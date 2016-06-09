package main

import (
	"strings"

	"github.com/alexgarzao/ms-gen/swaggerparser"

	"github.com/patrickmn/sortutil"
)

type Method struct {
	ServiceName        string
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
		ServiceName:   serviceName,
		MethodType:    methodType,
		ServiceMethod: strings.Title(serviceMethod),
		CodeFilename:  "service_" + CamelToSnake(operation.OperationID) + ".go",
	}

	method.Parameters = method.fillMethodParameters(operation.Parameters)
	method.Responses = method.fillResponses(operation.Responses)

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

// Fill method parameters.
func (method *Method) fillMethodParameters(swgParameters []*swaggerparser.Parameter) []Parameter {
	var parameters []Parameter

	for _, swgParameter := range swgParameters {
		parameters = append(parameters, *NewParameter(method.ServiceName, swgParameter))
	}

	sortutil.AscByField(parameters, "Name")

	return parameters
}

// Fill responses.
func (method *Method) fillResponses(apiResponses map[string]*swaggerparser.Response) []Response {
	var responses []Response

	for apiResponseKey, apiResponseValue := range apiResponses {
		responses = append(responses, *NewResponse(method.ServiceName, apiResponseKey, apiResponseValue))
	}

	sortutil.AscByField(responses, "ResultCode")

	return responses
}

func (method Method) String() string {
	return method.ServiceMethod
}

// Fill methods.
func FillMethods(serviceName string, pathDefinitions map[string]*swaggerparser.Path) []Method {
	var methods []Method
	for k, v := range pathDefinitions {
		if v.Get != nil {
			methods = append(methods, *NewMethod(serviceName, k, "Get", v.Get))
		}

		if v.Post != nil {
			methods = append(methods, *NewMethod(serviceName, k, "Post", v.Post))
		}

		if v.Put != nil {
			methods = append(methods, *NewMethod(serviceName, k, "Put", v.Put))
		}
	}

	sortutil.AscByField(methods, "ServiceMethod")

	return methods
}
