package transponder

type ConnectionStatus string

const (
	InvalidConnectionStatus ConnectionStatus = "invalid"
	PendingConnectionStatus ConnectionStatus = "pending"
	ActiveConnectionStatus  ConnectionStatus = "active"
	ErrorConnectionStatus   ConnectionStatus = "error"
)

var ConnectionStatus_index = map[ConnectionStatus]int32{
	"invalid": 0,
	"pending": 1,
	"active":  2,
	"error":   3,
}
var ConnectionStatus_value = map[int32]ConnectionStatus{
	0: "invalid",
	1: "pending",
	2: "active",
	3: "error",
}

type PutConnectionResponseBody struct {
	// Code
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Properties
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type ErrorResponseBody struct {
	// Errors
	Errors []*Error `json:"errors,omitempty" yaml:"errors,omitempty"`

	// Meta
	Meta map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

type Error struct {
	// Code an application specific error code expressed as a string value
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Detail a human readable explanation specific to this occurrence of the problem
	// Like title this field s value can be localized
	Detail string `json:"detail,omitempty" yaml:"detail,omitempty"`

	// Id a unique identifier for this particular occurrence of the problem
	ID string `json:"id,omitempty" yaml:"id,omitempty"`

	// Status the HTTP status code applicable to this problem expressed as a string
	// value
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Title a short human readable summary of the problem that SHOULD NOT change from
	// occurrence to occurrence of the problem except for purposes of localization
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

type FindConnectionsResponseBody struct {
	// Data
	Data []*Connection `json:"data,omitempty" yaml:"data,omitempty"`
}

type Connection struct {
	// Code
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Platform
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`

	// Properties
	Properties []*ConfigurationProperty `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Provider
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`

	// Status
	Status ConnectionStatus `json:"status,omitempty" yaml:"status,omitempty"`

	// Technology
	Technology string `json:"technology,omitempty" yaml:"technology,omitempty"`
}

type ConfigurationProperty struct {
	// Default
	Default string `json:"default,omitempty" yaml:"default,omitempty"`

	// Description
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Required
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Sensitive
	Sensitive bool `json:"sensitive,omitempty" yaml:"sensitive,omitempty"`
}
