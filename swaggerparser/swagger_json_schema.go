package swaggerparser

import (
	"strings"
)

type (
	// JSONSchema represents an instance of a JSON schema.
	// See http://json-schema.org/documentation.html
	JSONSchema struct {
		Schema string `yaml:"$schema,omitempty"`
		// Core schema
		ID           string                 `yaml:"id,omitempty"`
		Title        string                 `yaml:"title,omitempty"`
		Type         JSONType               `yaml:"type,omitempty"`
		Items        *JSONSchema            `yaml:"items,omitempty"`
		Properties   map[string]*JSONSchema `yaml:"properties,omitempty"`
		Definitions  map[string]*JSONSchema `yaml:"definitions,omitempty"`
		Description  string                 `yaml:"description,omitempty"`
		DefaultValue interface{}            `yaml:"default,omitempty"`
		Example      interface{}            `yaml:"example,omitempty"`

		// Hyper schema
		Media     *JSONMedia  `yaml:"media,omitempty"`
		ReadOnly  bool        `yaml:"readOnly,omitempty"`
		PathStart string      `yaml:"pathStart,omitempty"`
		Links     []*JSONLink `yaml:"links,omitempty"`
		Ref       string      `yaml:"$ref,omitempty"`

		// Validation
		Enum                 []interface{} `yaml:"enum,omitempty"`
		Format               string        `yaml:"format,omitempty"`
		Pattern              string        `yaml:"pattern,omitempty"`
		Minimum              float64       `yaml:"minimum,omitempty"`
		Maximum              float64       `yaml:"maximum,omitempty"`
		MinLength            int           `yaml:"minLength,omitempty"`
		MaxLength            int           `yaml:"maxLength,omitempty"`
		Required             []string      `yaml:"required,omitempty"`
		AdditionalProperties bool          `yaml:"additionalProperties,omitempty"`

		// Union
		AnyOf []*JSONSchema `yaml:"anyOf,omitempty"`
	}

	// JSONType is the JSON type enum.
	JSONType string

	// JSONMedia represents a "media" field in a JSON hyper schema.
	JSONMedia struct {
		BinaryEncoding string `yaml:"binaryEncoding,omitempty"`
		Type           string `yaml:"type,omitempty"`
	}

	// JSONLink represents a "link" field in a JSON hyper schema.
	JSONLink struct {
		Title        string      `yaml:"title,omitempty"`
		Description  string      `yaml:"description,omitempty"`
		Rel          string      `yaml:"rel,omitempty"`
		Href         string      `yaml:"href,omitempty"`
		Method       string      `yaml:"method,omitempty"`
		Schema       *JSONSchema `yaml:"schema,omitempty"`
		TargetSchema *JSONSchema `yaml:"targetSchema,omitempty"`
		MediaType    string      `yaml:"mediaType,omitempty"`
		EncType      string      `yaml:"encType,omitempty"`
	}
)

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
func (schema *JSONSchema) ToGolangType() string {
	if schema.Ref != "" {
		completeRef := schema.Ref // "#/definitions/GetMethod1Response"
		return completeRef[strings.LastIndex(completeRef, "/")+1:]
	}

	if string(schema.Type) == "array" {
		if schema.Items == nil {
			return "arrayundefinedtype"
		}

		completeRef := schema.Items.Ref // "#/definitions/GetMethod1Response"
		return "[]" + completeRef[strings.LastIndex(completeRef, "/")+1:]
	}

	goType := map[string]string{
		"integer|int32":    "int32",
		"integer|int64":    "int64",
		"integer|":         "int64",
		"number|float":     "float32",
		"number|double":    "float64",
		"number|":          "uint",
		"string|":          "string",
		"string|byte":      "undefined",
		"string|binary":    "undefined",
		"boolean|":         "bool",
		"string|date":      "time.Time",
		"string|date-time": "time.Time",
		"string|password":  "undefined",
	}

	swaggerType := string(schema.Type)
	swaggerFormat := schema.Format

	result, ok := goType[swaggerType+"|"+swaggerFormat]
	if ok == false {
		result = "invalidgolangtype"
	}

	return result
}
