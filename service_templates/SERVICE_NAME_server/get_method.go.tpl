package main

import (
	"{{.CommonImportPath}}"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) {{.CurrentPath.ServiceMethod}}(w rest.ResponseWriter, r *rest.Request) {
	{{if .CurrentPath.Parameter.Name}}
	// Get path parameter.
	// {{.CurrentPath.Parameter.Name}} := r.PathParam("{{.CurrentPath.Parameter.Name}}")
	{{end}}

	{{ range $response := .CurrentPath.Responses }}
	{{$response.Name}} := {{$response.Type}}{}
	{{ end }}
	
	// Business rules here :-)

	{{ range $response := .CurrentPath.Responses }}
	w.WriteJson({{$response.Name}})
	{{ end }}
	
}
