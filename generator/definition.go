package main

import (
	"github.com/alexgarzao/ms-gen/swaggerparser"
	"github.com/patrickmn/sortutil"
)

type Definition struct {
	Name       string
	Properties []Property
}

func NewDefinition(name string, schema *swaggerparser.JSONSchema) *Definition {

	definition := new(Definition)

	definition.Name = name
	for propertyKey, propertyValue := range schema.Properties {
		property := NewProperty(propertyKey, propertyValue)
		definition.Properties = append(definition.Properties, *property)
	}

	sortutil.AscByField(definition.Properties, "Name")

	return definition
}

// Fill definitions.
func FillDefinitions(apiDefinitions map[string]*swaggerparser.JSONSchema) []Definition {
	var definitions []Definition

	for apiDefinitionKey, apiDefinitionValue := range apiDefinitions {
		definition := NewDefinition(apiDefinitionKey, apiDefinitionValue)
		definitions = append(definitions, *definition)
	}

	sortutil.AscByField(definitions, "Name")

	return definitions
}
