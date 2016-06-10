package main

import (
	"github.com/verdverm/frisby"
)

const (
	MYSERVICE_URL = "http://localhost:8090/" // TODO: put in config.yaml
)

func main() {
	frisby.Global.PrintProgressName = true

	{{ range $method := .Methods }}New{{$method.TestType}}().RunAllTests()
	{{ end }}

	frisby.Global.PrintReport()
}
