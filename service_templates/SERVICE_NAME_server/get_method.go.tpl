package main

import (
	"{{.CommonImportPath}}"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) {{.CurrentPath.ServiceMethod}}(w rest.ResponseWriter, r *rest.Request) {
	{{ range $response := .CurrentPath.Responses }}
	{{$response.Name}} := {{$response.Type}}{}
	{{ end }}
	
	// Business rules here :-)

	{{ range $response := .CurrentPath.Responses }}
	w.WriteJson({{$response.Name}})
	{{ end }}
	
}
