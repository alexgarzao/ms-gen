package swaggerparser

type (
	// Based on this definition: https://github.com/goadesign/goa/blob/master/goagen/gen_swagger/swagger.go

	// Swagger represents an instance of a swagger object.
	// See https://swagger.io/specification/
	Swagger struct {
		Swagger             string                         `yaml:"swagger,omitempty"`
		Info                *Info                          `yaml:"info,omitempty"`
		Host                string                         `yaml:"host,omitempty"`
		BasePath            string                         `yaml:"basePath,omitempty"`
		Schemes             []string                       `yaml:"schemes,omitempty"`
		Consumes            []string                       `yaml:"consumes,omitempty"`
		Produces            []string                       `yaml:"produces,omitempty"`
		Paths               map[string]*Path               `yaml:"paths"`
		Definitions         map[string]*JSONSchema         `yaml:"definitions,omitempty"`
		Parameters          map[string]*Parameter          `yaml:"parameters,omitempty"`
		Responses           map[string]*Response           `yaml:"responses,omitempty"`
		SecurityDefinitions map[string]*SecurityDefinition `yaml:"securityDefinitions,omitempty"`
		Tags                []*Tag                         `yaml:"tags,omitempty"`
		ExternalDocs        *ExternalDocs                  `yaml:"externalDocs,omitempty"`
	}

	// Info provides metadata about the API. The metadata can be used by the clients if needed,
	// and can be presented in the Swagger-UI for convenience.
	Info struct {
		Title          string             `yaml:"title,omitempty"`
		Description    string             `yaml:"description,omitempty"`
		TermsOfService string             `yaml:"termsOfService,omitempty"`
		Contact        *ContactDefinition `yaml:"contact,omitempty"`
		License        *LicenseDefinition `yaml:"license,omitempty"`
		Version        string             `yaml:"version"`
	}

	// Path holds the relative paths to the individual endpoints.
	Path struct {
		// Ref allows for an external definition of this path item.
		Ref string `yaml:"$ref,omitempty"`
		// Get defines a GET operation on this path.
		Get *Operation `yaml:"get,omitempty"`
		// Put defines a PUT operation on this path.
		Put *Operation `yaml:"put,omitempty"`
		// Post defines a POST operation on this path.
		Post *Operation `yaml:"post,omitempty"`
		// Delete defines a DELETE operation on this path.
		Delete *Operation `yaml:"delete,omitempty"`
		// Options defines a OPTIONS operation on this path.
		Options *Operation `yaml:"options,omitempty"`
		// Head defines a HEAD operation on this path.
		Head *Operation `yaml:"head,omitempty"`
		// Patch defines a PATCH operation on this path.
		Patch *Operation `yaml:"patch,omitempty"`
		// Parameters is the list of parameters that are applicable for all the operations
		// described under this path.
		Parameters []*Parameter `yaml:"parameters,omitempty"`
	}

	// Operation describes a single API operation on a path.
	Operation struct {
		// Tags is a list of tags for API documentation control. Tags can be used for
		// logical grouping of operations by resources or any other qualifier.
		Tags []string `yaml:"tags,omitempty"`
		// Summary is a short summary of what the operation does. For maximum readability
		// in the swagger-ui, this field should be less than 120 characters.
		Summary string `yaml:"summary,omitempty"`
		// Description is a verbose explanation of the operation behavior.
		// GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		// ExternalDocs points to additional external documentation for this operation.
		ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty"`
		// OperationID is a unique string used to identify the operation.
		OperationID string `yaml:"operationId,omitempty"`
		// Consumes is a list of MIME types the operation can consume.
		Consumes []string `yaml:"consumes,omitempty"`
		// Produces is a list of MIME types the operation can produce.
		Produces []string `yaml:"produces,omitempty"`
		// Parameters is a list of parameters that are applicable for this operation.
		Parameters []*Parameter `yaml:"parameters,omitempty"`
		// Responses is the list of possible responses as they are returned from executing
		// this operation.
		Responses map[string]*Response `yaml:"responses,omitempty"`
		// Schemes is the transfer protocol for the operation.
		Schemes []string `yaml:"schemes,omitempty"`
		// Deprecated declares this operation to be deprecated.
		Deprecated bool `yaml:"deprecated,omitempty"`
		// Secury is a declaration of which security schemes are applied for this operation.
		Security []map[string][]string `yaml:"security,omitempty"`
	}

	// Parameter describes a single operation parameter.
	Parameter struct {
		// Name of the parameter. Parameter names are case sensitive.
		Name string `yaml:"name"`
		// In is the location of the parameter.
		// Possible values are "query", "header", "path", "formData" or "body".
		In string `yaml:"in"`
		// Description is`a brief description of the parameter.
		// GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		// Required determines whether this parameter is mandatory.
		Required bool `yaml:"required"`
		// Schema defining the type used for the body parameter, only if "in" is body
		Schema *JSONSchema `yaml:"schema,omitempty"`

		// properties below only apply if "in" is not body

		//  Type of the parameter. Since the parameter is not located at the request body,
		// it is limited to simple types (that is, not an object).
		Type string `yaml:"type,omitempty"`
		// Format is the extending format for the previously mentioned type.
		Format string `yaml:"format,omitempty"`
		// AllowEmptyValue sets the ability to pass empty-valued parameters.
		// This is valid only for either query or formData parameters and allows you to
		// send a parameter with a name only or an empty value. Default value is false.
		AllowEmptyValue bool `yaml:"allowEmptyValue,omitempty"`
		// Items describes the type of items in the array if type is "array".
		Items *Items `yaml:"items,omitempty"`
		// CollectionFormat determines the format of the array if type array is used.
		// Possible values are csv, ssv, tsv, pipes and multi.
		CollectionFormat string `yaml:"collectionFormat,omitempty"`
		// Default declares the value of the parameter that the server will use if none is
		// provided, for example a "count" to control the number of results per page might
		// default to 100 if not supplied by the client in the request.
		Default          interface{}   `yaml:"default,omitempty"`
		Maximum          float64       `yaml:"maximum,omitempty"`
		ExclusiveMaximum bool          `yaml:"exclusiveMaximum,omitempty"`
		Minimum          float64       `yaml:"minimum,omitempty"`
		ExclusiveMinimum bool          `yaml:"exclusiveMinimum,omitempty"`
		MaxLength        int           `yaml:"maxLength,omitempty"`
		MinLength        int           `yaml:"minLength,omitempty"`
		Pattern          string        `yaml:"pattern,omitempty"`
		MaxItems         int           `yaml:"maxItems,omitempty"`
		MinItems         int           `yaml:"minItems,omitempty"`
		UniqueItems      bool          `yaml:"uniqueItems,omitempty"`
		Enum             []interface{} `yaml:"enum,omitempty"`
		MultipleOf       float64       `yaml:"multipleOf,omitempty"`
	}

	// Response describes an operation response.
	Response struct {
		// Description of the response. GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		// Schema is a definition of the response structure. It can be a primitive,
		// an array or an object. If this field does not exist, it means no content is
		// returned as part of the response. As an extension to the Schema Object, its root
		// type value may also be "file".
		Schema *JSONSchema `yaml:"schema,omitempty"`
		// Headers is a list of headers that are sent with the response.
		Headers map[string]*Header `yaml:"headers,omitempty"`
		// Ref references a global API response.
		// This field is exclusive with the other fields of Response.
		Ref string `yaml:"$ref,omitempty"`
	}

	// Header represents a header parameter.
	Header struct {
		// Description is`a brief description of the parameter.
		// GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		//  Type of the header. it is limited to simple types (that is, not an object).
		Type string `yaml:"type,omitempty"`
		// Format is the extending format for the previously mentioned type.
		Format string `yaml:"format,omitempty"`
		// Items describes the type of items in the array if type is "array".
		Items *Items `yaml:"items,omitempty"`
		// CollectionFormat determines the format of the array if type array is used.
		// Possible values are csv, ssv, tsv, pipes and multi.
		CollectionFormat string `yaml:"collectionFormat,omitempty"`
		// Default declares the value of the parameter that the server will use if none is
		// provided, for example a "count" to control the number of results per page might
		// default to 100 if not supplied by the client in the request.
		Default          interface{}   `yaml:"default,omitempty"`
		Maximum          float64       `yaml:"maximum,omitempty"`
		ExclusiveMaximum bool          `yaml:"exclusiveMaximum,omitempty"`
		Minimum          float64       `yaml:"minimum,omitempty"`
		ExclusiveMinimum bool          `yaml:"exclusiveMinimum,omitempty"`
		MaxLength        int           `yaml:"maxLength,omitempty"`
		MinLength        int           `yaml:"minLength,omitempty"`
		Pattern          string        `yaml:"pattern,omitempty"`
		MaxItems         int           `yaml:"maxItems,omitempty"`
		MinItems         int           `yaml:"minItems,omitempty"`
		UniqueItems      bool          `yaml:"uniqueItems,omitempty"`
		Enum             []interface{} `yaml:"enum,omitempty"`
		MultipleOf       float64       `yaml:"multipleOf,omitempty"`
	}

	// SecurityDefinition allows the definition of a security scheme that can be used by the
	// operations. Supported schemes are basic authentication, an API key (either as a header or
	// as a query parameter) and OAuth2's common flows (implicit, password, application and
	// access code).
	SecurityDefinition struct {
		// Type of the security scheme. Valid values are "basic", "apiKey" or "oauth2".
		Type string `yaml:"type"`
		// Description for security scheme
		Description string `yaml:"description,omitempty"`
		// Name of the header or query parameter to be used when type is "apiKey".
		Name string `yaml:"name,omitempty"`
		// In is the location of the API key when type is "apiKey".
		// Valid values are "query" or "header".
		In string `yaml:"in,omitempty"`
		// Flow is the flow used by the OAuth2 security scheme when type is "oauth2"
		// Valid values are "implicit", "password", "application" or "accessCode".
		Flow string `yaml:"flow,omitempty"`
		// The oauth2 authorization URL to be used for this flow.
		AuthorizationURL string `yaml:"authorizationUrl,omitempty"`
		// TokenURL  is the token URL to be used for this flow.
		TokenURL string `yaml:"tokenUrl,omitempty"`
		// Scopes list the  available scopes for the OAuth2 security scheme.
		Scopes map[string]string `yaml:"scopes,omitempty"`
	}

	// Scope corresponds to an available scope for an OAuth2 security scheme.
	Scope struct {
		// Description for scope
		Description string `yaml:"description,omitempty"`
	}

	// ExternalDocs allows referencing an external resource for extended documentation.
	ExternalDocs struct {
		// Description is a short description of the target documentation.
		// GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		// URL for the target documentation.
		URL string `yaml:"url"`
	}

	// Items is a limited subset of JSON-Schema's items object. It is used by parameter
	// definitions that are not located in "body".
	Items struct {
		//  Type of the items. it is limited to simple types (that is, not an object).
		Type string `yaml:"type,omitempty"`
		// Format is the extending format for the previously mentioned type.
		Format string `yaml:"format,omitempty"`
		// Items describes the type of items in the array if type is "array".
		Items *Items `yaml:"items,omitempty"`
		// CollectionFormat determines the format of the array if type array is used.
		// Possible values are csv, ssv, tsv, pipes and multi.
		CollectionFormat string `yaml:"collectionFormat,omitempty"`
		// Default declares the value of the parameter that the server will use if none is
		// provided, for example a "count" to control the number of results per page might
		// default to 100 if not supplied by the client in the request.
		Default          interface{}   `yaml:"default,omitempty"`
		Maximum          float64       `yaml:"maximum,omitempty"`
		ExclusiveMaximum bool          `yaml:"exclusiveMaximum,omitempty"`
		Minimum          float64       `yaml:"minimum,omitempty"`
		ExclusiveMinimum bool          `yaml:"exclusiveMinimum,omitempty"`
		MaxLength        int           `yaml:"maxLength,omitempty"`
		MinLength        int           `yaml:"minLength,omitempty"`
		Pattern          string        `yaml:"pattern,omitempty"`
		MaxItems         int           `yaml:"maxItems,omitempty"`
		MinItems         int           `yaml:"minItems,omitempty"`
		UniqueItems      bool          `yaml:"uniqueItems,omitempty"`
		Enum             []interface{} `yaml:"enum,omitempty"`
		MultipleOf       float64       `yaml:"multipleOf,omitempty"`
	}

	// Tag allows adding meta data to a single tag that is used by the Operation Object. It is
	// not mandatory to have a Tag Object per tag used there.
	Tag struct {
		// Name of the tag.
		Name string `yaml:"name,omitempty"`
		// Description is a short description of the tag.
		// GFM syntax can be used for rich text representation.
		Description string `yaml:"description,omitempty"`
		// ExternalDocs is additional external documentation for this tag.
		ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty"`
	}

	// ContactDefinition contains the API contact information.
	ContactDefinition struct {
		// Name of the contact person/organization
		Name string `yaml:"name,omitempty"`
		// Email address of the contact person/organization
		Email string `yaml:"email,omitempty"`
		// URL pointing to the contact information
		URL string `yaml:"url,omitempty"`
	}

	// LicenseDefinition contains the license information for the API.
	LicenseDefinition struct {
		// Name of license used for the API
		Name string `yaml:"name,omitempty"`
		// URL to the license used for the API
		URL string `yaml:"url,omitempty"`
	}
)

// Get the first method parameter name.
func (operation *Operation) GetFirstPathParamName() string {
	for _, parameter := range operation.Parameters {
		if parameter.In == "path" {
			return parameter.Name
		}
	}

	return ""
}

// Get the first body parameter name.
func (operation *Operation) GetBodyParamName() string {
	for _, parameter := range operation.Parameters {
		if parameter.In == "body" {
			return parameter.Name
		}
	}

	return ""
}
