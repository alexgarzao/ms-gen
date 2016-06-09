package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"

	"strings"
)

type Response struct {
	ResultCode  string
	Description string
	Ref         string
	Name        string
	Type        string
}

func NewResponse(serviceName string, resultCode string, swgResponse *swaggerparser.Response) *Response {
	response := new(Response)
	response.ResultCode = resultCode

	if swgResponse.Schema != nil {
		completeRef := swgResponse.Schema.Ref // "#/definitions/GetMethod1Response"
		response.Ref = completeRef[strings.LastIndex(completeRef, "/")+1:]

		// Help fields.
		response.Name = strings.ToLower(string(response.Ref[0])) + response.Ref[1:]
		response.Type = serviceName + "_common." + response.Ref
	}

	return response
}
