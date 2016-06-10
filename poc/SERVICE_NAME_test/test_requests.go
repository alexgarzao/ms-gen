package main

import (
	common "github.com/alexgarzao/ms-gen/poc/SERVICE_NAME_common"

	"fmt"

	gojson "encoding/json"

	"github.com/bitly/go-simplejson"
	"github.com/verdverm/frisby"
)

func SendTestValidRequest1(testId string, parameter string, request interface{}) {
	SendTestValidGetRequest(testId, "method_1/", parameter, request, 200, nil)
}

func SendTestInvalidRequest1(testId string, parameter string, request interface{}, expectedErrorMessage string) {
	SendTestInvalidGetRequest(testId, "method_1/", parameter, request, 400, expectedErrorMessage)
}

// Generic functions

//
// Tools to send request using test system.
//
func SendTestValidPutRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, serviceResult interface{}) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	f := frisby.Create(testId).
		Put(MYSERVICE_URL + uri + parameter).
		SetJson(request).
		Send().
		ExpectStatus(expectedStatusCode)

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
		Send().
		ExpectStatus(expectedStatusCode)

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

	f := frisby.Create(testId).
		Get(MYSERVICE_URL + uri + parameter).
		Send().
		ExpectStatus(expectedStatusCode)

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

	frisby.Create(testId).
		Put(MYSERVICE_URL+uri+parameter).
		SetJson(request).
		Send().
		ExpectStatus(expectedStatusCode).
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

	frisby.Create(testId).
		Post(MYSERVICE_URL + uri).
		SetJson(request).
		Send().
		ExpectStatus(expectedStatusCode).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			errorMessage, _ := json.Get("Error").String()
			if errorMessage != expectedErrorMessage {
				F.AddError(fmt.Sprintf("Value of error [%s] differs from expected [%s]", errorMessage, expectedErrorMessage))
			}
		})
}

func SendTestInvalidGetRequest(testId string, uri string, parameter string, request interface{}, expectedStatusCode int, expectedErrorMessage string) {
	testId = "Test:" + common.GetFuncName(3)[4:] + ": " + testId + "."

	frisby.Create(testId).
		Get(MYSERVICE_URL + uri + parameter).
		Send().
		ExpectStatus(expectedStatusCode).
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			errorMessage, _ := json.Get("Error").String()
			if errorMessage != expectedErrorMessage {
				F.AddError(fmt.Sprintf("Value of error [%s] differs from expected [%s]", errorMessage, expectedErrorMessage))
			}
		})
}
