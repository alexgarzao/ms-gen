package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"
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

	parameter.Type = "common." + swgParameter.ToGolangType()

	return parameter
}
