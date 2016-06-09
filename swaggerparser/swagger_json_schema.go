package swaggerparser

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
