package main

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"

	"github.com/alexgarzao/ms-gen/swaggerparser"
)

type Api struct {
	Filename            string
	OutputDir           string
	ServiceName         string
	FriendlyServiceName string
	CommonImportPath    string
	Methods             []*Method
	Definitions         []*Definition
	CurrentMethod       *Method
}

func NewApi(filename string, outputDir string) (api *Api) {
	api = new(Api)
	api.Filename = filename
	api.OutputDir = outputDir

	return api
}

func (api *Api) LoadFromSwagger() error {
	// Reading from swagger file.
	config, err := ioutil.ReadFile(api.Filename)
	if err != nil {
		return err
	}

	if err := api.parser(config); err != nil {
		return err
	}

	return nil
}

func (api *Api) parser(text []byte) error {
	swagger := new(swaggerparser.Swagger)

	if err := yaml.Unmarshal(text, swagger); err != nil {
		return err
	}

	api.ServiceName = "myservice"
	api.FriendlyServiceName = swagger.Info.Title

	api.Methods = FillMethods(api.ServiceName, swagger.Paths)

	api.Definitions = FillDefinitions(swagger.Definitions)

	commonImportPath, err := GetCommonImportPath(api.OutputDir, api.ServiceName)
	if err != nil {
		return err
	}

	api.CommonImportPath = commonImportPath

	return nil
}
