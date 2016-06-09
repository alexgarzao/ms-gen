package SERVICE_NAME_common

import (
	"log"
	"regexp"
	"runtime"

	"io/ioutil"

	"errors"

	"bytes"

	"github.com/ant0ine/go-json-rest/rest"
)

var (
	// ErrJsonPayloadEmpty is returned when the JSON payload is empty.
	ErrJsonPayloadEmpty = errors.New("JSON payload is empty")
)

// Regex to extract just the function name (and not the module path)
var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

func GetFuncName(level int) string {
	fnName := "<unknown>"
	pc, _, _, ok := runtime.Caller(level)
	if ok {
		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	}

	return fnName
}

func LogRequest(r *rest.Request) error {
	// Read the content
	var bodyBytes []byte
	if r.Body != nil {
		var err error
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)

	if len(bodyString) == 0 {
		return ErrJsonPayloadEmpty
	}

	log.Printf("Request: ==>%v<==\n", string(bodyString))

	return nil
}

func GetFuncTestName() string {
	funcName := GetFuncName(3)[4:]
	return funcName
}
