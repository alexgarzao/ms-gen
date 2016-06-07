package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Starting ms-gen...\n")

	api := NewApi("../generator_tests/api_sample_1.yaml")

	source := NewSource(api, "../service_templates/Makefile.tpl")
	source.SaveToFile("./code/Makefile")

	fmt.Printf("Finishing ms-gen...\n")
}
