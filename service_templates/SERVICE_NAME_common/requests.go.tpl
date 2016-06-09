package {{.ServiceName}}_common

import (
	"time"
)

type ServiceRequest1 struct {
	FieldName1 string    `json:"fieldName1"`
	FieldName2 string    `json:"fieldName2"`
	FieldName3 time.Time `json:"fieldName3"`
}

{{ range $definition := .Definitions }}
type {{$definition.Name}} struct {
	{{ range $property := $definition.Properties }}
	{{$property.Name}} {{$property.Type}} `json:"{{$property.JsonName}}"`
	{{ end }}	
}
{{ end }}
