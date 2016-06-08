package main

import (
	"flag"
	"log"
)

var outputDir string

func main() {
	const (
		defaultOutputDir = "code"
		outputDirUsage   = "Directory where to put the generated source"
	)

	log.Printf("Starting ms-gen...\n")

	flag.StringVar(&outputDir, "output", defaultOutputDir, outputDirUsage)
	flag.StringVar(&outputDir, "o", defaultOutputDir, outputDirUsage+" (shorthand)")

	flag.Parse()

	log.Printf("Configuration: output=%s", outputDir)

	api := NewApi("../generator_tests/api_sample_1.yaml", outputDir)
	if err := api.LoadFromSwagger(); err != nil {
		log.Fatalf("When loading swagger file %s: %s", api.Filename, err)
	}

	projectSource := NewSource(api)
	if err := projectSource.Build(); err != nil {
		log.Fatalf("When build source files: %s", err)
	}

	log.Printf("Finishing ms-gen...\n")
}
