package main

import (
	"os"
	"path/filepath"

	"errors"
	"fmt"
	"strings"
)

func CreateBasePath(filename string) error {
	basePath := filepath.Dir(filename)

	if err := os.MkdirAll(basePath, 0700); err != nil {
		return errors.New(fmt.Sprintf("Creating base path %s: %s", basePath, err))
	}

	return nil
}

func GetImportBasePath(absolutePath string) (string, error) {
	n := strings.Index(absolutePath, "/github.com/")
	if n == -1 {
		return "", errors.New(fmt.Sprintf("Path github.com not found in %s", absolutePath))
	}

	return absolutePath[n+1:], nil
}

func GetCommonImportPath(outputDir string, serviceName string) (string, error) {
	baseSourceDir := outputDir + "/"

	absPath, err := filepath.Abs(baseSourceDir)
	if err != nil {
		return "", errors.New(fmt.Sprintf("When get abs(%s): %s", baseSourceDir, err))
	}

	importBasePath, err := GetImportBasePath(absPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("When getting import path from %s: %s", absPath, err))
	}

	commonImportPath := importBasePath + "/" + serviceName + "_common"

	return commonImportPath, nil
}

// Conversion based on this table: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
//
//	Common Name		type			format		Comments
//	-------------------------------------------------
//	integer			integer		int32		signed 32 bits
//	long				integer		int64		signed 64 bits
//	float			number		float
//	double			number		double
//	string			string
//	byte				string		byte			base64 encoded characters
//	binary			string		binary		any sequence of octets
//	boolean			boolean
//	date				string		date			As defined by full-date - RFC3339
//	dateTime			string		date-time	As defined by date-time - RFC3339
//	password			string		password		Used to hint UIs the input needs to be obscured.
//
func ToGolangType(swaggerType string, swaggerFormat string) string {
	goType := map[string]string{
		"integer|int32":    "int32",
		"integer|int64":    "int64",
		"integer|":         "int64",
		"number|float":     "float32",
		"number|double":    "float64",
		"number|":          "float64",
		"string|":          "string",
		"string|byte":      "undefined",
		"string|binary":    "undefined",
		"boolean|":         "bool",
		"string|date":      "time.Time",
		"string|date-time": "time.Time",
		"string|password":  "undefined",
	}

	result, ok := goType[swaggerType+"|"+swaggerFormat]
	if ok == false {
		result = "error"
	}

	return result
}

func GetPathWithoutParameter(path string) string {
	// Input example: "/get_method_4/{par1}"
	// Result: "/get_method_4"

	n := strings.Index(path, "/{")
	if n == -1 {
		return path
	}

	result := path[:n]

	return result
}

func CamelToSnake(text string) string {
	result := ""

	for _, chr := range text {
		if chr >= 'A' && chr <= 'Z' && result != "" {
			result += "_"
			result += string(chr + ('a' - 'A'))
		} else {
			result += string(chr)
		}
	}

	return result
}
