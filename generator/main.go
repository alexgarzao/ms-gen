package main

import (
	"flag"
	"log"
)

var (
	outputDir string
	apiFile   string
)

func main() {
	const (
		// Output dir flag.
		defaultOutputDir = "code"
		outputDirUsage   = "Directory where to put the generated source"

		// API file flag.
		apiFileUsage = "File that defines the API"
	)

	log.Printf("Starting ms-gen...\n")

	flag.StringVar(&outputDir, "output", defaultOutputDir, outputDirUsage)
	flag.StringVar(&outputDir, "o", defaultOutputDir, outputDirUsage+" (shorthand)")

	flag.StringVar(&apiFile, "api", "", apiFileUsage)
	flag.StringVar(&apiFile, "a", "", apiFileUsage+" (shorthand)")

	flag.Parse()

	log.Printf("Configuration: output='%s' api='%s'", outputDir, apiFile)

	if apiFile == "" {
		log.Fatalf("Parameter 'api' must be informed")
	}

	api := NewApi(apiFile, outputDir)
	if err := api.LoadFromSwagger(); err != nil {
		log.Fatalf("When loading swagger file %s: %s", api.Filename, err)
	}

	projectSource := NewSource(api)
	if err := projectSource.Build(); err != nil {
		log.Fatalf("When build source files: %s", err)
	}

	log.Printf("Finishing ms-gen...\n")
}
