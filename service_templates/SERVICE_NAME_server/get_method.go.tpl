package main

import (
	"{{.CommonImportPath}}"

	"github.com/ant0ine/go-json-rest/rest"
	{{ range $import := .CurrentMethod.Imports }}
	"{{$import}}"{{ end }}
)

func (s *Service) {{.CurrentMethod.ServiceMethod}}(w rest.ResponseWriter, r *rest.Request) {
	{{ range $parameter := .CurrentMethod.Parameters }}
		{{if eq $parameter.In "path"}}
		// Get path parameter.
		// {{$parameter.Name}} := r.PathParam("{{$parameter.Name}}")
		{{end}}
	{{ end }}
	
	{{ range $parameter := .CurrentMethod.Parameters }}
		{{if eq $parameter.In "body"}}
		// Get body parameters.

	{{$parameter.Name}} := {{$parameter.Type}}{}

	if err := r.DecodeJsonPayload(&{{$parameter.Name}}); err != nil {
		rest.Error(
			w,
			fmt.Sprintf("When decoding json: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}
		{{end}}
	{{ end }}
	


	{{ $gen_param_values := "false" }}

	{{ range $parameter := .CurrentMethod.Parameters }}
		{{if eq $parameter.In "query"}}
			{{if eq $gen_param_values "false"}}
			// Getting query parameters.
			// paramValues := r.URL.Query()
			{{$gen_param_values := "true"}}
			{{end}}
		// {{$parameter.Name}} := paramValues.Get("{{$parameter.Name}}")
		{{end}}
	{{ end }}

	{{ range $response := .CurrentMethod.Responses }}
		{{if ne $response.Name ""}}
			{{$response.Name}} := {{$response.Type}}{}
		{{end}}
	{{ end }}
	
	// Business rules here :-)

	{{ range $response := .CurrentMethod.Responses }}
		{{if ne $response.Name ""}}
			w.WriteJson({{$response.Name}})
		{{end}}
	{{ end }}
}
