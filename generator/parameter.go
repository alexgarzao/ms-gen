package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"

	"strings"
)

type Parameter struct {
	Name        string
	In          string
	Description string
	Required    bool
	Type        string
	Format      string
}

func NewParameter(serviceName string, swgParameter *swaggerparser.Parameter) *Parameter {
	parameter := &Parameter{
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

	return parameter
}
