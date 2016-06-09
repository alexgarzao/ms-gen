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
	api *Api
}

func NewSource(api *Api) *Source {
	source := new(Source)

	source.api = api

	return source
}

func (s *Source) Build() error {
	if err := s.saveToFile("../service_templates/Makefile.tpl", "Makefile"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving Makefile: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_common/definitions.go.tpl", "{{.ServiceName}}_common/definitions.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving definitions.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_common/requests.go.tpl", "{{.ServiceName}}_common/requests.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving requests.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_common/utils.go.tpl", "{{.ServiceName}}_common/utils.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving utils.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/db.go.tpl", "{{.ServiceName}}_server/db.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving db.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/db_models.go.tpl", "{{.ServiceName}}_server/db_models.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving db_models.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/main.go.tpl", "{{.ServiceName}}_server/main.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving main.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/service.go.tpl", "{{.ServiceName}}_server/service.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving service.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/request_validation.go.tpl", "{{.ServiceName}}_server/request_validation.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving request_validation.go: %s", err))
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/request_validation_test.go.tpl", "{{.ServiceName}}_server/request_validation_test.go"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving request_validation_test.go: %s", err))
	}

	for _, method := range s.api.Methods {
		s.api.CurrentMethod = method
		if err := s.saveToFile("../service_templates/SERVICE_NAME_server/get_method.go.tpl", "{{.ServiceName}}_server/"+method.CodeFilename); err != nil {
			return errors.New(fmt.Sprintf("Error when saving %s: %s", method.CodeFilename, err))
		}
	}

	if err := s.saveToFile("../service_templates/SERVICE_NAME_server/SERVICE_NAME_config.yaml.example.tpl", "{{.ServiceName}}_server/{{.ServiceName}}_config.yaml.example"); err != nil {
		return errors.New(fmt.Sprintf("Error when saving SERVICE_NAME_config.yaml.example: %s", err))
	}

	if err := s.saveTestFiles(); err != nil {
		return errors.New(fmt.Sprintf("Error when saving test files: %s", err))
	}

	return nil
}

func (s *Source) saveTestFiles() error {
	testFiles := []string{
		"tests",
		"test_requests",
		"test_behaviour_1",
	}

	for _, file := range testFiles {
		tplFile := fmt.Sprintf("../service_templates/SERVICE_NAME_test/%s.go.tpl", file)
		outputFile := fmt.Sprintf("{{.ServiceName}}_test/%s.go", file)

		if err := s.saveToFile(tplFile, outputFile); err != nil {
			return errors.New(fmt.Sprintf("Error when saving %s.go: %s", file, err))
		}
	}

	return nil
}

func (s *Source) saveToFile(templateFilename string, outputFilename string) error {

	tmpl := template.Must(template.ParseGlob(templateFilename))

	baseSourceDir := s.api.OutputDir + "/"

	// Replace tokens in filename.
	t := template.Must(template.New("output_filename").Parse(outputFilename))

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

	err = tmpl.Execute(f, s.api)
	if err != nil {
		return errors.New(fmt.Sprintf("template execution: %s", err))
	}

	return nil
}
