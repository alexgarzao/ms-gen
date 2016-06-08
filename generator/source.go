package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"
)

type Source struct {
	api  *Api
	tmpl *template.Template
}

func NewSource(api *Api, templateFilename string) *Source {
	source := new(Source)

	source.api = api
	source.tmpl = template.Must(template.ParseGlob(templateFilename))

	return source
}

func (s *Source) SaveToFile(templateFilename string) error {

	baseSourceDir := s.api.OutputDir + "/"

	// Replace tokens in filename.
	t := template.Must(template.New("template_filename").Parse(templateFilename))

	buffFilename := bytes.NewBufferString("")
	t.Execute(buffFilename, s.api)

	filename := baseSourceDir + buffFilename.String()

	log.Printf("Generating %s", filename)

	if err := CreateBasePath(filename); err != nil {
		return errors.New(fmt.Sprintf("creating base for %s: %s", filename, err))
	}

	f, err := os.Create(filename)
	if err != nil {
		return errors.New(fmt.Sprintf("create file: %s", err))
	}

	defer f.Close()

	err = s.tmpl.Execute(f, s.api)
	if err != nil {
		return errors.New(fmt.Sprintf("template execution: %s", err))
	}

	return nil
}
