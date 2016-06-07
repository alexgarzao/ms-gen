package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type Source struct {
	tmpl *template.Template
	api  *Api
}

func NewSource(api *Api, templateFilename string) *Source {
	source := new(Source)

	source.api = api
	source.tmpl = template.Must(template.ParseGlob(templateFilename))

	return source
}

func (s *Source) SaveToFile(templateFilename string) error {
	type TemplateData struct {
		ServiceName string
		MethodList  []string
	}

	templateData := TemplateData{
		ServiceName: "myservice",
	}

	for k := range s.api.swagger.Paths {
		templateData.MethodList = append(templateData.MethodList, "service_"+k[1:])
	}

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
