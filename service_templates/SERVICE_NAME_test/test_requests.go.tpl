package main

import (
	common 	"{{.CommonImportPath}}"

	"fmt"

	gojson "encoding/json"

	"github.com/bitly/go-simplejson"
	"github.com/verdverm/frisby"
)


{{/* Build test functions for GET operations */}}
{{ range $method := .Methods }}
	{{ if eq $method.MethodType "Get" }}
		func SendTestValid{{$method.ServiceMethod}}(testId string, parameter string, request interface{}, expectedStatusCode int, result interface{}) {
			SendTestValidGetRequest(testId, "{{$method.Path}}", parameter, request, expectedStatusCode, result)
		}
		
		func SendTestInvalid{{$method.ServiceMethod}}(testId string, parameter string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
			SendTestInvalidGetRequest(testId, "{{$method.Path}}", parameter, request, expectedStatusCode, expectedErrorMessage)
		}
	{{ end }}
{{ end }}

{{/* Build test functions for PUT operations */}}
{{ range $method := .Methods }}
	{{ if eq $method.MethodType "Put" }}
		func SendTestValid{{$method.ServiceMethod}}(testId string, parameter string, request interface{}, expectedStatusCode int, result interface{}) {
			SendTestValidPutRequest(testId, "{{$method.Path}}", parameter, request, expectedStatusCode, result)
		}
		
		func SendTestInvalid{{$method.ServiceMethod}}(testId string, parameter string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
			SendTestInvalidPutRequest(testId, "{{$method.Path}}", parameter, request, expectedStatusCode, expectedErrorMessage)
		}
	{{ end }}
{{ end }}

{{/* Build test functions for POST operations */}}
{{ range $method := .Methods }}
	{{ if eq $method.MethodType "Post" }}
		func SendTestValid{{$method.ServiceMethod}}(testId string, request interface{}, expectedStatusCode int, result interface{}) {
			SendTestValidPostRequest(testId, "{{$method.Path}}", request, expectedStatusCode, result)
		}
		
		func SendTestInvalid{{$method.ServiceMethod}}(testId string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
			SendTestInvalidPostRequest(testId, "{{$method.Path}}", request, expectedStatusCode, expectedErrorMessage)
		}
	{{ end }}
{{ end }}

// Generic functions

//
// Tools to send request using test system.
//
func SendTestValidPutRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, serviceResult interface{}) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	if parameter != "" {
		parameter = "/" + parameter
	}

	f := frisby.Create(testId).
		Put(MYSERVICE_URL + uri + parameter).
		SetJson(request).
		Send()

	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode)

	if serviceResult != nil {
		f.AfterContent(func(F *frisby.Frisby, content []byte, inputErr error) {
			if inputErr != nil {
				F.AddError(inputErr.Error())
				return
			}

			if err := gojson.Unmarshal(content, serviceResult); err != nil {
				F.AddError(err.Error())
				return
			}
		})
	}
}

func SendTestValidPostRequest(testId string, uri string, request interface{}, expectedStatusCode int, serviceResult interface{}) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	f := frisby.Create(testId).
		Post(MYSERVICE_URL + uri).
		SetJson(request).
		Send()
	
	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode)

	if serviceResult != nil {
		f.AfterContent(func(F *frisby.Frisby, content []byte, inputErr error) {
			if inputErr != nil {
				F.AddError(inputErr.Error())
				return
			}

			if err := gojson.Unmarshal(content, serviceResult); err != nil {
				F.AddError(err.Error())
				return
			}
		})
	}
}

func SendTestValidGetRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, serviceResult interface{}) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."
	
	if parameter != "" {
		parameter = "/" + parameter
	}

	f := frisby.Create(testId).
		Get(MYSERVICE_URL + uri + parameter).
		Send()
	
	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode)

	if serviceResult != nil {
		f.AfterContent(func(F *frisby.Frisby, content []byte, inputErr error) {
			if inputErr != nil {
				F.AddError(inputErr.Error())
				return
			}

			if err := gojson.Unmarshal(content, serviceResult); err != nil {
				F.AddError(err.Error())
				return
			}
		})
	}
}

func SendTestInvalidPutRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	if parameter != "" {
		parameter = "/" + parameter
	}

	f := frisby.Create(testId).
		Put(MYSERVICE_URL+uri+parameter).
		SetJson(request).
		Send()

	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode).
		ExpectJson("Error", expectedErrorMessage).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			errorMessage, _ := json.Get("Error").String()
			if errorMessage != expectedErrorMessage {
				F.AddError(fmt.Sprintf("Value of error [%s] differs from expected [%s]", errorMessage, expectedErrorMessage))
			}
		})
}

func SendTestInvalidPostRequest(testId string, uri string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	f := frisby.Create(testId).
		Post(MYSERVICE_URL + uri).
		SetJson(request).
		Send()

	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			errorMessage, _ := json.Get("Error").String()
			if errorMessage != expectedErrorMessage {
				F.AddError(fmt.Sprintf("Value of error [%s] differs from expected [%s]", errorMessage, expectedErrorMessage))
			}
		})
}

func SendTestInvalidGetRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	if parameter != "" {
		parameter = "/" + parameter
	}

	f := frisby.Create(testId).
		Get(MYSERVICE_URL + uri + parameter).
		Send()

	if f.Error() != nil {
		f.AddError(fmt.Sprintf("%s", f.Error()))
		return
	}

	f.ExpectStatus(expectedStatusCode).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			errorMessage, _ := json.Get("Error").String()
			if errorMessage != expectedErrorMessage {
				F.AddError(fmt.Sprintf("Value of error [%s] differs from expected [%s]", errorMessage, expectedErrorMessage))
			}
		})
}
