package main

import (
	"{{.CommonImportPath}}"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) {{.CurrentPath.ServiceMethod}}(w rest.ResponseWriter, r *rest.Request) {
	{{ range $parameter := .CurrentPath.Parameters }}
		{{if eq $parameter.In "path"}}
		// Get path parameter.
		// {{$parameter.Name}} := r.PathParam("{{$parameter.Name}}")
		{{end}}
	{{ end }}

	{{ $gen_param_values := "false" }}

	{{ range $parameter := .CurrentPath.Parameters }}
		{{if eq $parameter.In "query"}}
			{{if eq $gen_param_values "false"}}
			// Getting query parameters.
			// paramValues := r.URL.Query()
			{{$gen_param_values := "true"}}
			{{end}}
		// {{$parameter.Name}} := paramValues.Get("{{$parameter.Name}}")
		{{end}}
	{{ end }}

	{{ range $response := .CurrentPath.Responses }}
	{{$response.Name}} := {{$response.Type}}{}
	{{ end }}
	
	// Business rules here :-)

	{{ range $response := .CurrentPath.Responses }}
	w.WriteJson({{$response.Name}})
	{{ end }}
	
}
