package main

import (
	"log"
)

func main() {
	log.Printf("Starting ms-gen...\n")

	api := NewApi("../generator_tests/api_sample_1.yaml")
	if err := api.LoadFromSwagger(); err != nil {
		log.Fatalf("When loading swagger file %s: %s", api.Filename, err)
	}

	source := NewSource(api, "../service_templates/Makefile.tpl")
	if err := source.SaveToFile("./code/Makefile"); err != nil {
		log.Fatalf("Error when saving Makefile: %s", err)
	}

	source = NewSource(api, "../service_templates/SERVICE_NAME_common/definitions.go.tpl")
	if err := source.SaveToFile("./code/{{.ServiceName}}_common/definitions.go"); err != nil {
		log.Fatalf("Error when saving definitions.go: %s", err)
	}

	source = NewSource(api, "../service_templates/SERVICE_NAME_common/requests.go.tpl")
	if err := source.SaveToFile("./code/{{.ServiceName}}_common/requests.go"); err != nil {
		log.Fatalf("Error when saving requests.go: %s", err)
	}

	source = NewSource(api, "../service_templates/SERVICE_NAME_server/db.go.tpl")
	if err := source.SaveToFile("./code/{{.ServiceName}}_server/db.go"); err != nil {
		log.Fatalf("Error when saving db.go: %s", err)
	}

	source = NewSource(api, "../service_templates/SERVICE_NAME_server/db_models.go.tpl")
	if err := source.SaveToFile("./code/{{.ServiceName}}_server/db_models.go"); err != nil {
		log.Fatalf("Error when saving db_models.go: %s", err)
	}

	source = NewSource(api, "../service_templates/SERVICE_NAME_server/main.go.tpl")
	if err := source.SaveToFile("./code/{{.ServiceName}}_server/main.go"); err != nil {
		log.Fatalf("Error when saving main.go: %s", err)
	}

	log.Printf("Finishing ms-gen...\n")
}
