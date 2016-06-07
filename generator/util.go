package main

import (
	"os"
	"path/filepath"

	"errors"
	"fmt"
)

func CreateBasePath(filename string) error {
	basePath := filepath.Dir(filename)

	if err := os.MkdirAll(basePath, 0700); err != nil {
		return errors.New(fmt.Sprintf("Creating base path %s: %s", basePath, err))
	}

	return nil
}
