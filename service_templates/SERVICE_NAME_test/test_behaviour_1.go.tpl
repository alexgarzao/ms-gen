package main

import (
	"log"

	common 	"{{.CommonImportPath}}"
)

type {{$.CurrentMethod.TestType}} struct {
}

func New{{$.CurrentMethod.TestType}}() *{{$.CurrentMethod.TestType}} {
	return &{{$.CurrentMethod.TestType}}{}
}

func (t *{{$.CurrentMethod.TestType}}) RunAllTests() {
	log.Println("Checking behaviour 1")

	t.test1() // Remember to use great names :-)
	t.test2()
	//	...
	//	t.testN()
}

func (t *{{$.CurrentMethod.TestType}}) test1() {
	// Action 1
	//     For example, send a valid request to method 1, method 2, ...
	//     Sure, you can send invalid requests to test desired behaviours...
	// Action 2
	// ...
	// Action N

	{{ range $parameter := .CurrentMethod.Parameters }}
		{{if eq $parameter.In "body"}}
			// Body parameter.
			// {{$parameter.Name}} := {{$parameter.Type}}{
				// FieldName1: "xxx",
				// FieldName2: "yyy",
			// }
		{{end}}
		{{if eq $parameter.In "path"}}
			// Path parameter.
			// {{$parameter.Name}} := ""
		{{end}}
		{{if eq $parameter.In "query"}}
			// Query parameter.
			// {{$parameter.Name}} := ""
		{{end}}
	{{ end }}

	request := common.ServiceRequest1{
		FieldName1: "xxx",
		FieldName2: "yyy",
	}

	// Be sure about the expected status code bellow... Is 200?
	{{ if ne $.CurrentMethod.MethodType "Post" }}
	SendTestValid{{$.CurrentMethod.ServiceMethod}}("Request XXX with valid infos", "parameter_value", request, 200, nil)
	{{else}}
	SendTestValid{{$.CurrentMethod.ServiceMethod}}("Request XXX with valid infos", request, 200, nil)
	{{end}}
}

func (T *{{$.CurrentMethod.TestType}}) test2() {
	request := common.ServiceRequest1{
		FieldName1: "invalid",
		FieldName2: "yyy",
	}

	// Be sure about the expected status code bellow... Is 400?
	{{ if ne $.CurrentMethod.MethodType "Post" }}
	SendTestInvalid{{$.CurrentMethod.ServiceMethod}}("Request XXX with invalid infos", "parameter_value", request, 400, "Expected message error")
	{{else}}
	SendTestInvalid{{$.CurrentMethod.ServiceMethod}}("Request XXX with invalid infos", request, 400, "Expected message error")
	{{end}}
}
