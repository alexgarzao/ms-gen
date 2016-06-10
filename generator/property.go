package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"

	"strings"
)

type Property struct {
	Name            string
	Type            string
	JsonName        string
	JsonValidations string
	Required        bool
}

func NewProperty(name string, schema *swaggerparser.JSONSchema, required bool) *Property {

	jsonValidations := ""
	if required {
		jsonValidations = " valid:\"Required\""
	}

	return &Property{
		Name:            strings.Title(name),
		Type:            schema.ToGolangType(),
		JsonName:        name,
		JsonValidations: jsonValidations,
		Required:        required,
	}
}
