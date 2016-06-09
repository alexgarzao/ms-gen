package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"

	"strings"
)

type Property struct {
	Name     string
	Type     string
	JsonName string
}

func NewProperty(name string, schema *swaggerparser.JSONSchema) *Property {
	return &Property{
		Name:     strings.Title(name),
		Type:     ToGolangType(string(schema.Type), ""),
		JsonName: name,
	}
}
