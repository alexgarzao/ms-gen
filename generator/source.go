package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"
)

type Source struct {
	tmpl *template.Template
	api  *Api
}

type (
	TemplateData struct {
		ServiceName string
		MethodList  []string
		Definitions []Definition
	}

	Definition struct {
		Name       string
		Properties []Property
	}

	Property struct {
		Name     string
		Type     string
		JsonName string
	}
)

func NewSource(api *Api, templateFilename string) *Source {
	source := new(Source)

	source.api = api
	source.tmpl = template.Must(template.ParseGlob(templateFilename))

	return source
}

func (s *Source) SaveToFile(templateFilename string) error {

	templateData := TemplateData{
		ServiceName: "myservice",
	}

	// Fill method list.
	for k := range s.api.swagger.Paths {
		templateData.MethodList = append(templateData.MethodList, "service_"+k[1:])
	}

	templateData.Definitions = s.fillDefinitions(s.api.swagger.Definitions)

	// Replace tokens in filename.
	t := template.Must(template.New("template_filename").Parse(templateFilename))

	buffFilename := bytes.NewBufferString("")
	t.Execute(buffFilename, templateData)

	filename := buffFilename.String()

	if err := CreateBasePath(filename); err != nil {
		log.Fatalf("creating base for %s: %s", filename, err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("create file: %s", err)
	}

	defer f.Close()

	err = s.tmpl.Execute(f, templateData)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return nil
}

// Fill definitions.
func (s *Source) fillDefinitions(apiDefinitions map[string]*JSONSchema) []Definition {
	var definitions []Definition

	for apiDefinitionKey, apiDefinitionValue := range apiDefinitions {
		definition := Definition{}
		definition.Name = apiDefinitionKey
		for propertyKey, propertyValue := range apiDefinitionValue.Properties {
			property := Property{
				Name:     strings.Title(propertyKey),
				Type:     string(propertyValue.Type),
				JsonName: propertyKey,
			}
			definition.Properties = append(definition.Properties, property)
		}
		definitions = append(definitions, definition)
	}

	return definitions
}
