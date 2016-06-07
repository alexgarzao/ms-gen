package {{.ServiceName}}_common

//import (
//	"time"
//)

{{ range $definition := .Definitions }}
type {{$definition.Name}} struct {
	{{ range $property := $definition.Properties }}
	{{$property.Name}} {{$property.Type}} `json:"{{$property.JsonName}}"`
	{{ end }}	
}
{{ end }}
