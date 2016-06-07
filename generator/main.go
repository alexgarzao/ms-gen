package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Starting ms-gen...\n")

	api := NewApi("../generator_tests/api_sample_1.yaml")

	source := NewSource(api, "../service_templates/Makefile.tpl")
	source.SaveToFile("./code/Makefile")

	source = NewSource(api, "../service_templates/SERVICE_NAME_common/definitions.go.tpl")
	source.SaveToFile("./code/{{.ServiceName}}_common/definitions.go")

	source = NewSource(api, "../service_templates/SERVICE_NAME_common/requests.go.tpl")
	source.SaveToFile("./code/{{.ServiceName}}_common/requests.go")

	fmt.Printf("Finishing ms-gen...\n")
}
