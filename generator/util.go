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
		if chr >= 'A' && chr <= 'Z' {
			if result != "" {
				result += "_"
			}
			result += string(chr + ('a' - 'A'))
		} else {
			result += string(chr)
		}
	}

	return result
}
