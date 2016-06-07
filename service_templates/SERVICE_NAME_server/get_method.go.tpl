package main

import (
	"github.com/alexgarzao/ms-gen/generator/code/{{.ServiceName}}_common"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) {{.CurrentPath.ServiceMethod}}(w rest.ResponseWriter, r *rest.Request) {
//	par := r.PathParam("{{/*.ParameterName*/}}")

	response := {{.ServiceName}}_common.{{.CurrentPath.ServiceMethod}}Response{}
	
	// Business rules here :-)

	w.WriteJson(response)
}
