package apps

import (
	"encoding/json"
	"time"
)

type Error struct {
	Code    string
	Message string
	Status  int64
}

type ErrorResponse struct {
	Error *Error
}

type AppType string

const (
	AppTypeSystem   AppType = "system"
	AppTypeInternal AppType = "internal"
	AppTypeCustom   AppType = "custom"
)

var AppCapabilityValues = map[string]int32{
	"UNSPECIFIED_APP_CAPABILITY": 0,
	"permitPropagation":          1,
	"permitIngestion":            2,
	"rightPropagation":           3,
	"rightIngestion":             4,
	"tagOrchestration":           5,
}

var AppContactTypeValues = map[string]int32{
	"UNSPECIFIED_APP_CONTACT_TYPE": 0,
	"technical":                    1,
	"marketing":                    2,
	"finance":                      3,
	"security":                     4,
}

var AppMarketplaceCategoryValues = map[string]int32{
	"UNSPECIFIED_APP_MARKETPLACE_CATEGORY": 0,
	"privacy":                              1,
	"asset":                                2,
}

var AppTypeValues = map[AppType]int32{
	"UnspecifiedAppType": 0,
	"SystemAppType":      1,
	"InternalAppType":    2,
	"CustomAppType":      3,
}

var AppAllowedInstancesValues = map[string]int32{
	"UnspecifiedAppAllowedInstances": 0,
	"MultipleAppAllowedInstances":    1,
	"SingleAppAllowedInstances":      2,
}

var DataRoleValues = map[string]int32{
	"DATA_ROLE_UNSPECIFIED": 0,
	"processor":             1,
	"controller":            2,
	"jointController":       3,
	"jointProcessor":        4,
	"coController":          5,
}

var DataSubjectRoleValues = map[string]int32{
	"DATA_SUBJECT_ROLE_UNSPECIFIED": 0,
	"customer":                      1,
	"employee":                      2,
}

var CookieDurationValues = map[string]int32{
	"COOKIE_DURATION_UNSPECIFIED": 0,
	"session":                     1,
	"persistent":                  2,
}

var CookieProvenanceValues = map[string]int32{
	"COOKIE_TYPE_UNSPECIFIED": 0,
	"firstParty":              1,
	"thirdParty":              2,
}

var CookieCategoryValues = map[string]int32{
	"COOKIE_CATEGORY_UNSPECIFIED": 0,
	"strictlyNecessary":           1,
	"functional":                  2,
	"performance":                 3,
	"marketing":                   4,
}

var ParentClosePolicyValues = map[string]int32{
	"terminate":     0,
	"requestCancel": 1,
	"abandon":       2,
}

type SelectData struct {
	// These are the values that will be selected on this field
	Values []*SelectDataValue `yaml:"values,omitempty" json:"values,omitempty"`
	// Enter a JSON Array to use. It should be formatted as an array of objects with named properties
	Json string `yaml:"json,omitempty" json:"json,omitempty"`
	// Enter a url with a data source in JSON Array format. This can be used to populate a Select list with external JSON values
	Url string `yaml:"url,omitempty" json:"url,omitempty"`
	// Key of the label in json object
	LabelKey string `yaml:"labelKey,omitempty" json:"labelKey,omitempty"`
	// Key of the value in json object
	ValueKey string `yaml:"valueKey,omitempty" json:"valueKey,omitempty"`
}

type ConditionForDisplay struct {
	// Name of the field to condition upon
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// Value of named field for this form element to display
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
}

type SelectDataValue struct {
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
	Label string `yaml:"label,omitempty" json:"label,omitempty"`
}

type FormComponent struct {
	// The key field is where the data will be saved to. This must be unique per field.
	// For example, if key = 'customers' then the value of the field will be saved in data.customers
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// The name or title for this component
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	// The description that will appear before this component
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	// The type property will be used to select which component to render on the frontend. It cannot be an existing field type
	Type            string `yaml:"type,omitempty" json:"type,omitempty"`
	ShowOnView      bool   `yaml:"showOnView,omitempty" json:"showOnView,omitempty"`
	ShowValueOnEdit bool   `yaml:"showValueOnEdit,omitempty" json:"showValueOnEdit,omitempty"`
	Editable        bool   `yaml:"editable,omitempty" json:"editable,omitempty"`
	// If true, the field will be required to have a value
	Required  bool   `yaml:"required,omitempty" json:"required,omitempty"`
	MinLength int64  `yaml:"minLength,omitempty" json:"minLength,omitempty"`
	MaxLength int64  `yaml:"maxLength,omitempty" json:"maxLength,omitempty"`
	Pattern   string `yaml:"pattern,omitempty" json:"pattern,omitempty"`
	// The placeholder text that will appear when this field is empty
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`
	// Default will be the default value for this field, before user interaction. Having a default value will override the placeholder text
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
	// type == "string"
	// This setting will enable spell check on the field if
	Spellcheck bool `yaml:"spellcheck,omitempty" json:"spellcheck,omitempty"`
	// type == "array"
	// If true, multiple values can be added in this field.
	// The values will appear as an array in the API and an “Add Another” button will be visible on the field allowing the creation of additional fields for this component
	Multiple bool `yaml:"multiple,omitempty" json:"multiple,omitempty"`
	// type == "array"
	// Data is the definition of how data is provided for the dropdown
	Data *SelectData `yaml:"data,omitempty" json:"data,omitempty"`
	// The condition for display determines whether this component should be displayed
	ConditionForDisplay []*ConditionForDisplay `yaml:"conditionForDisplay,omitempty" json:"conditionForDisplay,omitempty"`
}

type Webhook struct {
	URL           string   `yaml:"url" json:"url,omitempty"`
	Secret        string   `yaml:"secret" json:"secret,omitempty"`
	Authorization string   `yaml:"authorization,omitempty" json:"authorization,omitempty"`
	Events        []string `yaml:"events,omitempty" json:"events,omitempty"`
	TLS           TLS      `yaml:"tls,omitempty" json:"tls,omitempty"`
	MaxQPS        int32    `yaml:"maxQPS,omitempty" json:"maxQPS,omitempty"`
}

// TODO - add rest of TLS properties with proper secrets handling
type TLS struct {
	Insecure bool `yaml:"insecure,omitempty" json:"insecure,omitempty"`
}

type IdentitySpaceFilter struct {
	Type   string `yaml:"type" json:"type,omitempty"`
	Format string `yaml:"format" json:"format,omitempty"`
}

type IdentitySpaceVariableJWT struct {
	Location string `yaml:"location" json:"location"`
	Key      string `yaml:"key" json:"key"`
}

type IdentitySpaceVariable struct {
	Name     string                    `yaml:"name" json:"name"`
	Location string                    `yaml:"location" json:"location"`
	JWT      *IdentitySpaceVariableJWT `yaml:"jwt,omitempty" json:"jwt,omitempty"`
}

type IdentitySpace struct {
	Managed     bool                   `yaml:"managed" json:"managed"`
	Code        string                 `yaml:"code,omitempty" json:"code,omitempty"`
	Name        string                 `yaml:"name,omitempty" json:"name,omitempty"`
	Description string                 `yaml:"description,omitempty" json:"description,omitempty"`
	Filters     []*IdentitySpaceFilter `yaml:"filters,omitempty" json:"filters,omitempty"`
	Type        string                 `yaml:"type,omitempty" json:"type,omitempty"`
	Format      string                 `yaml:"format,omitempty" json:"format,omitempty"`
	Variable    *IdentitySpaceVariable `yaml:"variable,omitempty" json:"variable,omitempty"`
}

type AppContact struct {
	ContactType string `yaml:"type,omitempty" json:"type,omitempty"`
	Email       string `yaml:"email,omitempty" json:"email,omitempty"`
}

type AppConfigContact struct {
	ContactType string `yaml:"type,omitempty" json:"type,omitempty"`
	Email       string `yaml:"email,omitempty" json:"email,omitempty"`
}

type AppConfigImage struct {
	Title  string `yaml:"title,omitempty" json:"title,omitempty"`
	Link   string `yaml:"link,omitempty" json:"link,omitempty"`
	Width  int32  `yaml:"width,omitempty" json:"width,omitempty"`
	Height int32  `yaml:"height,omitempty" json:"height,omitempty"`
}

type AppConfigActivityStep struct {
	ID          string                        `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                        `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                        `yaml:"description,omitempty" json:"description,omitempty"`
	Step        string                        `yaml:"step,omitempty" json:"step,omitempty"`
	Fn          string                        `yaml:"fn,omitempty" json:"fn,omitempty"`
	Next        string                        `yaml:"next,omitempty" json:"next,omitempty"`
	Options     *AppConfigActivityStepOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Params      *json.RawMessage              `yaml:"params,omitempty" json:"params,omitempty"`
}

type WorkflowStepChildWorkflowOptions struct {
	TaskQueue                string            `yaml:"taskQueue,omitempty" json:"task_queue,omitempty"`
	WorkflowExecutionTimeout int64             `yaml:"workflowExecutionTimeout,omitempty" json:"workflow_execution_timeout,omitempty"`
	WorkflowRunTimeout       int64             `yaml:"workflowRunTimeout,omitempty" json:"workflow_run_timeout,omitempty"`
	WorkflowTaskTimeout      int64             `yaml:"workflowTaskTimeout,omitempty" json:"workflow_task_timeout,omitempty"`
	WaitForCancellation      bool              `yaml:"waitForCancellation,omitempty" json:"wait_for_cancellation,omitempty"`
	ParentClosePolicy        int64             `yaml:"parentClosePolicy,omitempty" json:"parent_close_policy,omitempty"`
	Memo                     map[string]string `yaml:"memo,omitempty" json:"memo,omitempty"`
	SearchAttributes         map[string]string `yaml:"searchAttributes,omitempty" json:"search_attributes,omitempty"`
	RetryPolicy              *RetryPolicy      `yaml:"retryPolicy,omitempty" json:"retry_policy,omitempty"`
}

type AppConfigChildWorkflowStep struct {
	ID          string                           `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                           `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                           `yaml:"description,omitempty" json:"description,omitempty"`
	Step        string                           `yaml:"step,omitempty" json:"step,omitempty"`
	Fn          string                           `yaml:"fn,omitempty" json:"fn,omitempty"`
	Next        string                           `yaml:"next,omitempty" json:"next,omitempty"`
	Options     WorkflowStepChildWorkflowOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Params      *json.RawMessage                 `yaml:"params,omitempty" json:"params,omitempty"`
}

type AppConfigGatewayStep struct {
	ID          string            `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string            `yaml:"code,omitempty" json:"code,omitempty"`
	Description string            `yaml:"description,omitempty" json:"description,omitempty"`
	Mode        string            `yaml:"mode,omitempty" json:"mode,omitempty"`
	Next        []*StepTransition `yaml:"next,omitempty" json:"next,omitempty"`
}

type AppConfigActivityOptions struct {
	TaskQueue                string `yaml:"task_queue,omitempty" json:"task_queue,omitempty"`
	WorkflowExecutionTimeout int64  `yaml:"workflow_execution_timeout,omitempty" json:"workflow_execution_timeout,omitempty"`
	WorkflowRunTimeout       int64  `yaml:"workflow_run_timeout,omitempty" json:"workflow_run_timeout,omitempty"`
	WorkflowTaskTimeout      int64  `yaml:"workflow_task_timeout,omitempty" json:"workflow_task_timeout,omitempty"`
	WaitForCancellation      bool   `yaml:"wait_for_cancellation,omitempty" json:"wait_for_cancellation,omitempty"`
	ScheduleToCloseTimeout   int64  `yaml:"schedule_to_close_timeout,omitempty" json:"schedule_to_close_timeout,omitempty"`
	ScheduleToStartTimeout   int64  `yaml:"schedule_to_start_timeout,omitempty" json:"schedule_to_start_timeout,omitempty"`
	StartToCloseTimeout      int64  `yaml:"start_to_close_timeout,omitempty" json:"start_to_close_timeout,omitempty"`
	HeartbeatTimeout         int64  `yaml:"heartbeat_timeout,omitempty" json:"heartbeat_timeout,omitempty"`
}

type ActivityOptions struct {
	TaskQueue              string       `json:"task_queue,omitempty"`
	ScheduleToCloseTimeout int64        `json:"schedule_to_close_timeout,omitempty"`
	ScheduleToStartTimeout int64        `json:"schedule_to_start_timeout,omitempty"`
	StartToCloseTimeout    int64        `json:"start_to_close_timeout,omitempty"`
	HeartbeatTimeout       int64        `json:"heartbeat_timeout,omitempty"`
	WaitForCancellation    bool         `json:"wait_for_cancellation,omitempty"`
	RetryPolicy            *RetryPolicy `json:"retry_policy,omitempty"`
}

type AppConfigActivityStepOptions struct {
	TaskQueue              string `yaml:"taskQueue,omitempty" json:"task_queue,omitempty"`
	ScheduleToCloseTimeout int64  `yaml:"scheduleToCloseTimeout,omitempty" json:"schedule_to_close_timeout,omitempty"`
	ScheduleToStartTimeout int64  `yaml:"scheduleToStartTimeout,omitempty" json:"schedule_to_start_timeout,omitempty"`
	StartToCloseTimeout    int64  `yaml:"startToCloseTimeout,omitempty" json:"start_to_close_timeout,omitempty"`
	HeartbeatTimeout       int64  `yaml:"heartbeatTimeout,omitempty" json:"heartbeat_timeout,omitempty"`
	WaitForCancellation    bool   `yaml:"waitForCancellation,omitempty" json:"wait_for_cancellation,omitempty"`
	RetryPolicy            string `yaml:"retryPolicy,omitempty" json:"retry_policy,omitempty"`
}

type ParameterDefinition struct {
	Code    string `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Type    string `yaml:"type,omitempty" json:"type,omitempty"`
	Default string `yaml:"default,omitempty" json:"default_value,omitempty"`
}

type AppConfigWorkflowActivityDefinition struct {
	Code    string                    `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string                    `yaml:"name,omitempty" json:"name,omitempty"`
	Icon    string                    `yaml:"icon,omitempty" json:"icon,omitempty"`
	Fn      string                    `yaml:"fn,omitempty" json:"fn,omitempty"`
	Options *AppConfigActivityOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Params  []*ParameterDefinition    `yaml:"params,omitempty" json:"params,omitempty"`
	Outputs []*ParameterDefinition    `yaml:"outputs,omitempty" json:"outputs,omitempty"`
	Config  *json.RawMessage          `yaml:"config,omitempty" json:"config,omitempty"`
}

type IconDefinition struct {
	SVG  string `json:"SVG,omitempty"`
	URL  string `json:"URL,omitempty"`
	ETag string `json:"ETag,omitempty"`
}

type ActivityDefinition struct {
	Options              *ActivityOptions       `json:"options,omitempty"`
	Params               []*ParameterDefinition `json:"params,omitempty"`
	Outputs              []*ParameterDefinition `json:"outputs,omitempty"`
	Config               *json.RawMessage       `json:"config,omitempty"`
	TemporalFunctionName string                 `json:"temporalFunctionName,omitempty"`
}

type ChildWorkflowOptions struct {
	TaskQueue                string            `json:"task_queue,omitempty"`
	WorkflowExecutionTimeout int64             `json:"workflow_execution_timeout,omitempty"`
	WorkflowRunTimeout       int64             `json:"workflow_run_timeout,omitempty"`
	WorkflowTaskTimeout      int64             `json:"workflow_task_timeout,omitempty"`
	WaitForCancellation      bool              `json:"wait_for_cancellation,omitempty"`
	Memo                     map[string]string `json:"memo,omitempty"`
	SearchAttributes         map[string]string `json:"search_attributes,omitempty"`
	RetryPolicy              *RetryPolicy      `json:"retry_policy,omitempty"`
}

type ChildWorkflowDefinition struct {
	Options              *ChildWorkflowOptions  `json:"options,omitempty"`
	Params               []*ParameterDefinition `json:"params,omitempty"`
	Outputs              []*ParameterDefinition `json:"outputs,omitempty"`
	Config               *json.RawMessage       `json:"config,omitempty"`
	TemporalFunctionName string                 `json:"temporalFunctionName,omitempty"`
}

type WorkflowActivityDefinition struct {
	Code     string                   `json:"code,omitempty"`
	Name     string                   `json:"name,omitempty"`
	Icon     *IconDefinition          `json:"icon,omitempty"`
	Activity *ActivityDefinition      `json:"activity,omitempty"`
	Workflow *ChildWorkflowDefinition `json:"workflow,omitempty"`
}

type PurposeTemplate struct {
	Code                       string            `yaml:"code,omitempty" json:"code,omitempty"`
	Name                       string            `yaml:"name,omitempty" json:"name,omitempty"`
	Description                string            `yaml:"description,omitempty" json:"description,omitempty"`
	TcfID                      int               `yaml:"tcfId,omitempty" json:"tcfId,omitempty"`
	TcfType                    string            `yaml:"tcfType,omitempty" json:"tcfType,omitempty"`
	Editable                   bool              `yaml:"editable,omitempty" json:"editable,omitempty"`
	LegalBasisRestriction      string            `yaml:"legalBasisRestriction,omitempty" json:"legalBasisRestriction,omitempty"`
	Required                   bool              `yaml:"required,omitempty" json:"required,omitempty"`
	DisplayName                string            `yaml:"displayName,omitempty" json:"displayName,omitempty"`
	DisplayDescription         string            `yaml:"displayDescription,omitempty" json:"displayDescription,omitempty"`
	LegalBasis                 map[string]string `yaml:"legalBasis,omitempty" json:"legalBasis,omitempty"`
	Cookies                    []*Cookie         `yaml:"cookies,omitempty" json:"cookies,omitempty"`
	CanonicalPurposes          []string          `yaml:"canonicalPurposes,omitempty" json:"canonicalPurposes,omitempty"`
	Translations               map[string]string `yaml:"translations,omitempty" json:"translations,omitempty"`
	DataSubjectRole            int32             `yaml:"dataSubjectRole,omitempty" json:"dataSubjectRole,omitempty"`
	DataRole                   int32             `yaml:"dataRole,omitempty" json:"dataRole,omitempty"`
	PurposeTemplateCollections []string          `yaml:"collections,omitempty" json:"collections,omitempty"`
}

type PurposeTemplateCollection struct {
	Code string `yaml:"code,omitempty" json:"code,omitempty"`
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

type Vendor struct {
	Id                         string   `yaml:"id,omitempty" json:"id,omitempty"`
	Name                       string   `yaml:"name,omitempty" json:"name,omitempty"`
	Purposes                   []string `yaml:"purposes,omitempty" json:"purposes,omitempty"`
	SpecialPurposes            []string `yaml:"specialPurposes,omitempty" json:"specialPurposes,omitempty"`
	Features                   []string `yaml:"features,omitempty" json:"features,omitempty"`
	SpecialFeatures            []string `yaml:"specialFeatures,omitempty" json:"specialFeatures,omitempty"`
	PolicyUrl                  string   `yaml:"policyUrl,omitempty" json:"policyUrl,omitempty"`
	LegIntPurposes             []string `yaml:"legIntPurposes,omitempty" json:"legIntPurposes,omitempty"`
	FlexiblePurposes           []string `yaml:"flexiblePurposes,omitempty" json:"flexiblePurposes,omitempty"`
	UsesCookies                bool     `yaml:"usesCookies,omitempty" json:"usesCookies,omitempty"`
	CookieMaxAgeSeconds        int64    `yaml:"cookieMaxAgeSeconds,omitempty" json:"cookieMaxAgeSeconds,omitempty"`
	CookieRefresh              bool     `yaml:"cookieRefresh,omitempty" json:"cookieRefresh,omitempty"`
	UsesNonCookieAccess        bool     `yaml:"usesNonCookieAccess,omitempty" json:"usesNonCookieAccess,omitempty"`
	DeviceStorageDisclosureUrl string   `yaml:"deviceStorageDisclosureUrl,omitempty" json:"deviceStorageDisclosureUrl,omitempty"`
}

type Tcf struct {
	Vendor                  Vendor `yaml:"vendor,omitempty" json:"vendor,omitempty"`
	GvlSpecificationVersion string `yaml:"gvlSpecificationVersion,omitempty" json:"gvlSpecificationVersion,omitempty"`
	VendorListVersion       string `yaml:"vendorListVersion,omitempty" json:"vendorListVersion,omitempty"`
	TcfPolicyVersion        string `yaml:"tcfPolicyVersion,omitempty" json:"tcfPolicyVersion,omitempty"`
}

type AppConfigCookie struct {
	Code            string `yaml:"code,omitempty" json:"code,omitempty"`
	Name            string `yaml:"name,omitempty" json:"name,omitempty"`
	Description     string `yaml:"description,omitempty" json:"description,omitempty"`
	Host            string `yaml:"host,omitempty" json:"host,omitempty"`
	Duration        string `yaml:"duration,omitempty" json:"duration,omitempty"`
	Provenance      string `yaml:"provenance,omitempty" json:"provenance,omitempty"`
	Category        string `yaml:"category,omitempty" json:"category,omitempty"`
	ServiceProvider string `yaml:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
}

type Cookie struct {
	Code            string `json:"code,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Host            string `json:"host,omitempty"`
	Duration        int32  `json:"duration,omitempty"`
	Provenance      int32  `json:"provenance,omitempty"`
	Category        int32  `json:"category,omitempty"`
	ServiceProvider string `json:"serviceProvider,omitempty"`
	AppCode         string `json:"appCode,omitempty"`
}

type AppConfigPurpose struct {
	Code                  string             `yaml:"code,omitempty" json:"code,omitempty"`
	Name                  string             `yaml:"name,omitempty" json:"name,omitempty"`
	Description           string             `yaml:"description,omitempty" json:"description,omitempty"`
	TcfID                 int                `yaml:"tcfId,omitempty" json:"tcfId,omitempty"`
	TcfType               string             `yaml:"tcfType,omitempty" json:"tcfType,omitempty"`
	Editable              bool               `yaml:"editable,omitempty" json:"editable,omitempty"`
	LegalBasisRestriction string             `yaml:"legalBasisRestriction,omitempty" json:"legalBasisRestriction,omitempty"`
	DisplayName           string             `yaml:"displayName,omitempty" json:"displayName,omitempty"`
	DisplayDescription    string             `yaml:"displayDescription,omitempty" json:"displayDescription,omitempty"`
	ProcessingPurpose     string             `yaml:"processingPurpose,omitempty" json:"processingPurpose,omitempty"`
	LegalBasis            map[string]string  `yaml:"legalBasis,omitempty" json:"legalBasis,omitempty"`
	Cookies               []*AppConfigCookie `yaml:"cookies,omitempty" json:"cookies,omitempty"`
	CanonicalPurposes     []string           `yaml:"canonicalPurposes,omitempty" json:"canonicalPurposes,omitempty"`
	Translations          map[string]string  `yaml:"translations,omitempty" json:"translations,omitempty"`
	DataSubjectRole       string             `yaml:"dataSubjectRole,omitempty" json:"dataSubjectRole,omitempty"`
	DataRole              string             `yaml:"dataRole,omitempty" json:"dataRole,omitempty"`
}

type AppConfigPurposeTemplate struct {
	Code                  string             `yaml:"code,omitempty" json:"code,omitempty"`
	Name                  string             `yaml:"name,omitempty" json:"name,omitempty"`
	Description           string             `yaml:"description,omitempty" json:"description,omitempty"`
	TcfID                 int                `yaml:"tcfId,omitempty" json:"tcfId,omitempty"`
	TcfType               string             `yaml:"tcfType,omitempty" json:"tcfType,omitempty"`
	Editable              bool               `yaml:"editable,omitempty" json:"editable,omitempty"`
	LegalBasisRestriction string             `yaml:"legalBasisRestriction,omitempty" json:"legalBasisRestriction,omitempty"`
	DisplayName           string             `yaml:"displayName,omitempty" json:"displayName,omitempty"`
	DisplayDescription    string             `yaml:"displayDescription,omitempty" json:"displayDescription,omitempty"`
	ProcessingPurpose     string             `yaml:"processingPurpose,omitempty" json:"processingPurpose,omitempty"`
	LegalBasis            map[string]string  `yaml:"legalBasis,omitempty" json:"legalBasis,omitempty"`
	Cookies               []*AppConfigCookie `yaml:"cookies,omitempty" json:"cookies,omitempty"`
	CanonicalPurposes     []string           `yaml:"canonicalPurposes,omitempty" json:"canonicalPurposes,omitempty"`
	Translations          map[string]string  `yaml:"translations,omitempty" json:"translations,omitempty"`
	DataSubjectRole       string             `yaml:"dataSubjectRole,omitempty" json:"dataSubjectRole,omitempty"`
	DataRole              string             `yaml:"dataRole,omitempty" json:"dataRole,omitempty"`
}

type Purpose struct {
	Code                  string            `json:"code,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Description           string            `json:"description,omitempty"`
	TcfID                 int               `json:"tcfId,omitempty"`
	TcfType               string            `json:"tcfType,omitempty"`
	Editable              bool              `json:"editable,omitempty"`
	LegalBasisRestriction string            `json:"legalBasisRestriction,omitempty"`
	DisplayName           string            `json:"displayName,omitempty"`
	DisplayDescription    string            `json:"displayDescription,omitempty"`
	ProcessingPurpose     string            `json:"processingPurpose,omitempty"`
	LegalBasis            map[string]string `json:"legalBasis,omitempty"`
	Cookies               []*Cookie         `json:"cookies,omitempty"`
	CanonicalPurposes     []string          `json:"canonicalPurposes,omitempty"`
	Translations          map[string]string `json:"translations,omitempty"`
	DataSubjectRole       int32             `json:"dataSubjectRole,omitempty"`
	DataRole              int32             `json:"dataRole,omitempty"`
}

type RightTranslation struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type CanonicalRight struct {
	Code        string `yaml:"code,omitempty" json:"code,omitempty"`
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type Right struct {
	Code            string                       `yaml:"code,omitempty" json:"code,omitempty"`
	Name            string                       `yaml:"name,omitempty" json:"name,omitempty"`
	Description     string                       `yaml:"description,omitempty" json:"description,omitempty"`
	Translations    map[string]*RightTranslation `yaml:"translations,omitempty" json:"translations,omitempty"`
	CanonicalRights []*CanonicalRight            `yaml:"canonicalRights,omitempty" json:"canonicalRights,omitempty"`
}

type Regulation struct {
	Code              string              `yaml:"code,omitempty" json:"code,omitempty"`
	Name              string              `yaml:"name,omitempty" json:"name,omitempty"`
	Description       string              `yaml:"description,omitempty" json:"description,omitempty"`
	RightsFulfillment []*RightFulfillment `yaml:"rightsFulfillment,omitempty" json:"rightsFulfillment,omitempty"`
}

type LegalBasisRestriction struct {
	Regulation string   `yaml:"regulation,omitempty" json:"regulation,omitempty"`
	LegalBasis []string `yaml:"legalBasis,omitempty" json:"legalBasis,omitempty"`
}

type AppConfigPolicyScope struct {
	Code        string           `yaml:"code,omitempty" json:"code,omitempty"`
	Name        string           `yaml:"name,omitempty" json:"name,omitempty"`
	Description string           `yaml:"description,omitempty" json:"description,omitempty"`
	Regions     []string         `yaml:"regions,omitempty" json:"regions,omitempty"`
	Regulations []string         `yaml:"regulations,omitempty" json:"regulations,omitempty"`
	Fulfillment map[string]int64 `yaml:"fulfillment,omitempty" json:"fulfillment,omitempty"`
}

type RightFulfillment struct {
	RightCode   string `json:"rightCode,omitempty"`
	Fulfillment int64  `json:"fulfillment,omitempty"`
}

type PolicyScope struct {
	Code              string              `json:"code,omitempty"`
	Name              string              `json:"name,omitempty"`
	Description       string              `json:"description,omitempty"`
	RegionCodes       []string            `json:"regionCodes,omitempty"`
	RegulationCodes   []string            `json:"regulationCodes,omitempty"`
	RightsFulfillment []*RightFulfillment `json:"rightsFulfillment,omitempty"`
}

type AppConfigLegalBasis struct {
	Code                  string `yaml:"code,omitempty" json:"code,omitempty"`
	Name                  string `yaml:"name,omitempty" json:"name,omitempty"`
	Description           string `yaml:"description,omitempty" json:"description,omitempty"`
	RequiresOptIn         bool   `yaml:"requiresOptIn,omitempty" json:"requiresOptIn,omitempty"`
	AllowOptOut           bool   `yaml:"allowOptOut,omitempty" json:"allowOptOut,omitempty"`
	RequiresPrivacyPolicy bool   `yaml:"requiresPrivacyPolicy,omitempty" json:"requiresPrivacyPolicy,omitempty"`
}

type LegalBasis struct {
	Code                  string   `json:"code,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Description           string   `json:"description,omitempty"`
	RequiresOptIn         bool     `json:"requiresOptIn,omitempty"`
	AllowOptOut           bool     `json:"allowOptOut,omitempty"`
	RequiresPrivacyPolicy bool     `json:"requiresPrivacyPolicy,omitempty"`
	RegulationCodes       []string `json:"regulationCodes,omitempty"`
}

type Theme struct {
	Code                  string `yaml:"code,omitempty" json:"code,omitempty"`
	Name                  string `yaml:"name,omitempty" json:"name,omitempty"`
	Description           string `yaml:"description,omitempty" json:"description,omitempty"`
	BannerBackgroundColor string `yaml:"bannerBackgroundColor,omitempty" json:"bannerBackgroundColor,omitempty"`
	LightboxRibbonColor   string `yaml:"lightboxRibbonColor,omitempty" json:"lightboxRibbonColor,omitempty"`
	FormHeaderColor       string `yaml:"formHeaderColor,omitempty" json:"formHeaderColor,omitempty"`
	StatusColor           string `yaml:"statusColor,omitempty" json:"statusColor,omitempty"`
	HighlightColor        string `yaml:"highlightColor,omitempty" json:"highlightColor,omitempty"`
	FeedbackColor         string `yaml:"feedbackColor,omitempty" json:"feedbackColor,omitempty"`
}

type RetryPolicy struct {
	InitialInterval          int64    `yaml:"initialInterval,omitempty" json:"initial_interval,omitempty"`
	BackoffCoefficient       float64  `yaml:"backoffCoefficient,omitempty" json:"backoff_coefficient,omitempty"`
	MaximumInterval          int64    `yaml:"maximumInterval,omitempty" json:"maximum_interval,omitempty"`
	MaximumAttempts          int64    `yaml:"maximumAttempts,omitempty" json:"maximum_attempts,omitempty"`
	NonRetryableErrorReasons []string `yaml:"nonRetryableErrorReasons,omitempty" json:"non_retriable_error_reasons,omitempty"`
}

type WorkflowOptions struct {
	TaskQueue                string            `yaml:"taskQueue,omitempty" json:"task_queue,omitempty"`
	WorkflowExecutionTimeout int64             `yaml:"workflowExecutionTimeout,omitempty" json:"workflow_execution_timeout,omitempty"`
	WorkflowRunTimeout       int64             `yaml:"workflowRunTimeout,omitempty" json:"workflow_run_timeout,omitempty"`
	WorkflowTaskTimeout      int64             `yaml:"workflowTaskTimeout,omitempty" json:"workflow_task_timeout,omitempty"`
	Memo                     map[string]string `yaml:"memo,omitempty" json:"memo,omitempty"`
	SearchAttributes         map[string]string `yaml:"searchAttributes,omitempty" json:"search_attributes,omitempty"`
	RetryPolicy              *RetryPolicy      `yaml:"retryPolicy,omitempty" json:"retry_policy,omitempty"`
}

type AppConfigStep struct {
	Activity      *AppConfigActivityStep      `yaml:"activity,omitempty" json:"activity,omitempty"`
	ChildWorkflow *AppConfigChildWorkflowStep `yaml:"childWorkflow,omitempty" json:"childWorkflow,omitempty"`
	Gateway       *AppConfigGatewayStep       `yaml:"gateway,omitempty" json:"gateway,omitempty"`
	Start         *AppConfigStartStep         `yaml:"start,omitempty" json:"start,omitempty"`
	Finish        *AppConfigFinishStep        `yaml:"finish,omitempty" json:"finish,omitempty"`
}

type AppConfigStartStep struct {
	ID          string           `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string           `yaml:"code,omitempty" json:"code,omitempty"`
	Description string           `yaml:"description,omitempty" json:"description,omitempty"`
	Next        string           `yaml:"next,omitempty" json:"next,omitempty"`
	Params      *json.RawMessage `yaml:"params,omitempty" json:"params,omitempty"`
}

type AppConfigFinishStep struct {
	ID          string           `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string           `yaml:"code,omitempty" json:"code,omitempty"`
	Description string           `yaml:"description,omitempty" json:"description,omitempty"`
	Params      *json.RawMessage `yaml:"params,omitempty" json:"params,omitempty"`
}

type StepTransition struct {
	ID       string `yaml:"id,omitempty" json:"id,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	Variable string `yaml:"variable,omitempty" json:"variable,omitempty"`
	Operator string `yaml:"operator,omitempty" json:"operator,omitempty"`
	Operand  string `yaml:"operand,omitempty" json:"operand,omitempty"`
}

type AppConfigWorkflowDefinition struct {
	Code     string           `yaml:"code,omitempty" json:"code,omitempty"`
	Name     string           `yaml:"name,omitempty" json:"name,omitempty"`
	Readonly bool             `yaml:"readonly,omitempty" json:"readonly,omitempty"`
	Options  *WorkflowOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Steps    []*AppConfigStep `yaml:"steps,omitempty" json:"steps,omitempty"`
}

type ActivityStep struct {
	Code                 string           `json:"code,omitempty"`
	Options              *ActivityOptions `json:"options,omitempty"`
	Params               *json.RawMessage `json:"params,omitempty"`
	Transition           string           `json:"transition,omitempty"`
	TemporalFunctionName string           `json:"temporalFunctionName,omitempty"`
}

type ChildWorkflowStep struct {
	Code                 string                `json:"code,omitempty"`
	Options              *ChildWorkflowOptions `json:"options,omitempty"`
	Params               *json.RawMessage      `json:"params,omitempty"`
	Transition           string                `json:"transition,omitempty"`
	TemporalFunctionName string                `json:"temporalFunctionName,omitempty"`
}

var GatewayStepModeValues = map[string]int32{
	"invalid": 0,
	"split":   1,
	"join":    2,
	"single":  3,
	"multi":   4,
}

type GatewayStep struct {
	Mode        int32             `json:"mode,omitempty"`
	Transitions []*StepTransition `json:"transitions,omitempty"`
}

type StartStep struct {
	Transition string           `json:"transition,omitempty"`
	Params     *json.RawMessage `json:"params,omitempty"`
}

type FinishStep struct {
	Params *json.RawMessage `json:"params,omitempty"`
}

type Step struct {
	ID          string             `json:"ID,omitempty"`
	Code        string             `json:"code,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Activity    *ActivityStep      `json:"activity,omitempty"`
	Workflow    *ChildWorkflowStep `json:"workflow,omitempty"`
	Gateway     *GatewayStep       `json:"gateway,omitempty"`
	Start       *StartStep         `json:"start,omitempty"`
	Finish      *FinishStep        `json:"finish,omitempty"`
}

type WorkflowDefinition struct {
	Code     string           `yaml:"code,omitempty" json:"code,omitempty"`
	Name     string           `yaml:"name,omitempty" json:"name,omitempty"`
	Readonly bool             `yaml:"readonly,omitempty" json:"readonly,omitempty"`
	Options  *WorkflowOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Steps    []*Step          `yaml:"steps,omitempty" json:"steps,omitempty"`
}

type ManifestInputs struct {
	ID                         string                                 `yaml:"ID,omitempty" json:"ID,omitempty"`
	Code                       string                                 `yaml:"code,omitempty" json:"code,omitempty"`
	OrgCode                    string                                 `yaml:"org,omitempty" json:"org,omitempty"`
	Name                       string                                 `yaml:"name,omitempty" json:"name,omitempty"`
	Version                    string                                 `yaml:"version,omitempty" json:"version,omitempty"`
	Depends                    string                                 `yaml:"depends,omitempty" json:"depends,omitempty"`
	Provides                   []string                               `yaml:"provides,omitempty" json:"provides,omitempty"`
	Type                       AppType                                `yaml:"type,omitempty" json:"type,omitempty"`
	AutoUpgrade                bool                                   `yaml:"autoUpgrade,omitempty" json:"autoUpgrade,omitempty"`
	Instances                  string                                 `yaml:"instances,omitempty" json:"instances,omitempty"`
	PrimaryCategory            string                                 `yaml:"primaryCategory,omitempty" json:"primaryCategory,omitempty"`
	SecondaryCategory          string                                 `yaml:"secondaryCategory,omitempty" json:"secondaryCategory,omitempty"`
	Rules                      map[string]string                      `yaml:"rules" json:"rules,omitempty"`
	Capabilities               []string                               `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
	SupportedLanguages         []string                               `yaml:"supportedLanguages,omitempty" json:"supportedLanguages,omitempty"`
	SupportedPurposes          []string                               `yaml:"supportedPurposes,omitempty" json:"supportedPurposes,omitempty"`
	SupportedRights            []string                               `yaml:"supportedRights,omitempty" json:"supportedRights,omitempty"`
	ShortDescription           string                                 `yaml:"shortDescription,omitempty" json:"shortDescription,omitempty"`
	DetailedDescription        string                                 `yaml:"detailedDescription,omitempty" json:"detailedDescription,omitempty"`
	PermissionNote             string                                 `yaml:"permissionNote,omitempty" json:"permissionNote,omitempty"`
	Permissions                []string                               `yaml:"permissions,omitempty" json:"permissions,omitempty"`
	InfoUrl                    string                                 `yaml:"infoURL,omitempty" json:"infoURL,omitempty"`
	SetupUrl                   string                                 `yaml:"setupURL,omitempty" json:"setupURL,omitempty"`
	HomepageUrl                string                                 `yaml:"homepageURL,omitempty" json:"homepageURL,omitempty"`
	CustomerSupportUrl         string                                 `yaml:"customerSupportURL,omitempty" json:"customerSupportURL,omitempty"`
	PrivacyPolicyUrl           string                                 `yaml:"privacyPolicyURL,omitempty" json:"privacyPolicyURL,omitempty"`
	StatusUrl                  string                                 `yaml:"statusURL,omitempty" json:"statusURL,omitempty"`
	TosUrl                     string                                 `yaml:"tosURL,omitempty" json:"tosURL,omitempty"`
	DocUrl                     string                                 `yaml:"docURL,omitempty" json:"docURL,omitempty"`
	Logo                       *AppConfigImage                        `yaml:"logo,omitempty" json:"logo,omitempty"`
	Previews                   []*AppConfigImage                      `yaml:"previews,omitempty" json:"previews,omitempty"`
	Contacts                   []*AppConfigContact                    `yaml:"contacts,omitempty" json:"contacts,omitempty"`
	ExpireUserTokens           bool                                   `yaml:"expireUserTokens,omitempty" json:"expireUserTokens,omitempty"`
	RefreshInterval            string                                 `yaml:"refreshInterval,omitempty" json:"refreshInterval,omitempty"`
	RequestUserAuth            bool                                   `yaml:"requestUserAuth,omitempty" json:"requestUserAuth,omitempty"`
	UserAuthCallbackUrl        string                                 `yaml:"userAuthCallbackURL,omitempty" json:"userAuthCallbackURL,omitempty"`
	RedirectOnUpdate           bool                                   `yaml:"redirectOnUpdate,omitempty" json:"redirectOnUpdate,omitempty"`
	Webhook                    *Webhook                               `yaml:"webhook,omitempty" json:"webhook,omitempty"`
	FormTitle                  string                                 `yaml:"formTitle,omitempty" json:"formTitle,omitempty"`
	FormSubtitle               string                                 `yaml:"formSubtitle,omitempty" json:"formSubtitle,omitempty"`
	Form                       []*FormComponent                       `yaml:"form,omitempty" json:"form,omitempty"`
	IdentitySpaces             []*IdentitySpace                       `yaml:"identitySpaces,omitempty" json:"identitySpaces,omitempty"`
	Workflows                  []*AppConfigWorkflowDefinition         `yaml:"workflows,flow,omitempty" json:"workflows,omitempty"`
	Activities                 []*AppConfigWorkflowActivityDefinition `yaml:"activities,flow,omitempty" json:"activities,omitempty"`
	ChildWorkflows             []*AppConfigWorkflowActivityDefinition `yaml:"childWorkflows,flow,omitempty" json:"childWorkflows,omitempty"`
	Tcf                        *Tcf                                   `yaml:"tcf,flow,omitempty" json:"tcf,omitempty"`
	Cookies                    []*AppConfigCookie                     `yaml:"cookies,flow,omitempty" json:"cookies,omitempty"`
	PurposeTemplates           []*AppConfigPurposeTemplate            `yaml:"purposeTemplates,flow,omitempty" json:"purposeTemplates,omitempty"`
	PurposeTemplateCollections []*PurposeTemplateCollection           `yaml:"purposeTemplateCollections,flow,omitempty" json:"purposeTemplateCollections,omitempty"`
	Purposes                   []*AppConfigPurpose                    `yaml:"purposes,flow,omitempty" json:"purposes,omitempty"`
	Rights                     []*Right                               `yaml:"rights,flow" json:"rights,omitempty"`
	Regulations                []*Regulation                          `yaml:"regulations,flow" json:"regulations,omitempty"`
	LegalBasisRestrictions     []*LegalBasisRestriction               `yaml:"legalBasisRestrictions,flow,omitempty" json:"legalBasisRestrictions,omitempty"`
	PolicyScopes               []*AppConfigPolicyScope                `yaml:"policyScopes,flow,omitempty" json:"policyScopes,omitempty"`
	LegalBases                 []*AppConfigLegalBasis                 `yaml:"legalBases,flow,omitempty" json:"legalBases,omitempty"`
	Themes                     []*Theme                               `yaml:"themes,flow,omitempty" json:"themes,omitempty"`
	ResourceTypes              []*ResourceType                        `yaml:"resourceTypes,flow,omitempty" json:"resourceTypes,omitempty"`
}

type App struct {
	ID                         string                        `json:"id,omitempty"`
	Code                       string                        `json:"code,omitempty"`
	OrgCode                    string                        `yaml:"orgCode" json:"orgCode,omitempty"`
	Name                       string                        `json:"name,omitempty"`
	Version                    string                        `json:"version,omitempty"`
	Depends                    string                        `json:"depends,omitempty"`
	Provides                   []string                      `json:"provides,omitempty"`
	Type                       int32                         `json:"type,omitempty"`
	AutoUpgrade                bool                          `json:"autoUpgrade,omitempty"`
	Instances                  int32                         `json:"instances,omitempty"`
	Rules                      map[string]string             `yaml:",flow" json:"rules,omitempty"`
	Capabilities               []int32                       `yaml:",flow" json:"capabilities,omitempty"`
	SupportedLanguages         []string                      `yaml:",flow" json:"supportedLanguages,omitempty"`
	SupportedPurposes          []string                      `yaml:",flow" json:"supportedPurposes,omitempty"`
	SupportedRights            []string                      `yaml:",flow" json:"supportedRights,omitempty"`
	PermissionNote             string                        `json:"permissionNode,omitempty"`
	Permissions                []string                      `yaml:",flow"`
	InfoUrl                    string                        `json:"infoURL,omitempty"`
	SetupUrl                   string                        `json:"setupURL,omitempty"`
	HomepageUrl                string                        `json:"homepageURL,omitempty"`
	ExpireUserTokens           bool                          `json:"expireUserTokens,omitempty"`
	RefreshInterval            int64                         `yaml:"refreshInterval" json:"refreshInterval,omitempty"`
	RequestUserAuth            bool                          `json:"requestUserAuth,omitempty"`
	UserAuthCallbackUrl        string                        `json:"userAuthCallbackURL,omitempty"`
	RedirectOnUpdate           bool                          `json:"redirectOnUpdate,omitempty"`
	WebhookId                  string                        `json:"webhookID,omitempty"`
	Readme                     string                        `json:"readme,omitempty"`
	FormTitle                  string                        `json:"formTitle,omitempty"`
	FormSubtitle               string                        `json:"formSubtitle,omitempty"`
	Form                       []*FormComponent              `yaml:",flow"`
	IdentitySpaces             []*IdentitySpace              `yaml:",flow" json:"identitySpaces,omitempty"`
	Purposes                   []*Purpose                    `yaml:",flow" json:"purposes,omitempty"`
	Workflows                  []*WorkflowDefinition         `json:"workflows,omitempty"`
	Activities                 []*WorkflowActivityDefinition `json:"activities,omitempty"`
	ChildWorkflows             []*WorkflowActivityDefinition `json:"childWorkflows,omitempty"`
	PurposeTemplates           []*PurposeTemplate            `json:"purposeTemplates,omitempty"`
	PurposeTemplateCollections []*PurposeTemplateCollection  `json:"purposeTemplateCollections,omitempty"`
	LegalBasisRestrictions     []*LegalBasisRestriction      `json:"legalBasisRestrictions,omitempty"`
	PolicyScopes               []*PolicyScope                `json:"policyScopes,omitempty"`
	LegalBases                 []*LegalBasis                 `json:"legalBases,omitempty"`
	Themes                     []*Theme                      `json:"themes,omitempty"`
	Rights                     []*Right                      `yaml:",flow" json:"rights,omitempty"`
	Regulations                []*Regulation                 `yaml:",flow" json:"regulations,omitempty"`
	Tcf                        *Tcf                          `json:"tcf,omitempty"`
	EventTypes                 []string                      `json:"eventTypes,omitempty"`
	Cookies                    []*Cookie                     `yaml:"cookies,flow,omitempty" json:"cookies,omitempty"`
	TransponderAppDetailsJSON  []byte                        `json:"transponderAppDetailsJSON,omitempty"`
}

type AppMarketplaceEntry struct {
	AppCode             string        `json:"appCode,omitempty"`
	AppID               string        `json:"appID,omitempty"`
	Version             string        `json:"version,omitempty"`
	Contacts            []*AppContact `yaml:",flow"`
	ShortDescription    string        `json:"shortDescription,omitempty"`
	PrimaryCategory     int32         `yaml:",inline" json:"primaryCategory,omitempty"`
	SecondaryCategory   int32         `yaml:",inline" json:"secondaryCategory,omitempty"`
	SupportedLanguages  []string      `yaml:",flow" json:"supportedLanguages,omitempty"`
	CustomerSupportUrl  string        `json:"customerSupportURL,omitempty"`
	PrivacyPolicyUrl    string        `json:"privacyPolicyURL,omitempty"`
	StatusUrl           string        `json:"statusURL,omitempty"`
	TosUrl              string        `json:"tosURL,omitempty"`
	DocUrl              string        `json:"docURL,omitempty"`
	Logo                Image         `yaml:",flow"`
	IntroDescription    string        `json:"introDescription,omitempty"`
	DetailedDescription string        `json:"detailedDescription,omitempty"`
	Previews            []*Image      `yaml:",flow" json:"previews,omitempty"`
}

type Image struct {
	Data   []byte `yaml:"data" json:"data,omitempty"`
	Title  string `yaml:"title" json:"title,omitempty"`
	Url    string `yaml:"url" json:"url,omitempty"`
	Width  int32  `yaml:"width" json:"width,omitempty"`
	Height int32  `yaml:"height" json:"height,omitempty"`
}

type PublishAppRequest struct {
	AppMarketplaceEntry *AppMarketplaceEntry `json:"marketplaceEntry,omitempty"`
	Webhook             *Webhook             `json:"webhook,omitempty"`
}

type PublishAppResponse struct {
	AppMarketplaceEntry *AppMarketplaceEntry `json:"marketplaceEntry,omitempty"`
}

type WebhookResponse struct {
	Webhook *Webhook
}

type PutAppRequest struct {
	App *App
}

type PutAppResponse struct {
	App *App
}

func NewApp(p ManifestInputs) (*App, error) {
	var appCapabilities []int32
	for _, capability := range p.Capabilities {
		if appCapability, ok := AppCapabilityValues[capability]; ok {
			appCapabilities = append(appCapabilities, appCapability)
		}
	}

	var refreshIntervalHours int64
	if len(p.RefreshInterval) > 0 {
		refreshIntervalNanoseconds, err := time.ParseDuration(p.RefreshInterval)
		if err != nil {
			return nil, err
		}
		refreshIntervalHours = int64(refreshIntervalNanoseconds / time.Hour)
	}

	var policyScopes []*PolicyScope
	for _, policyScope := range p.PolicyScopes {
		var rightsFulfillment []*RightFulfillment
		for rightCode, fulfillment := range policyScope.Fulfillment {
			rightsFulfillment = append(rightsFulfillment, &RightFulfillment{
				RightCode:   rightCode,
				Fulfillment: fulfillment,
			})
		}

		policyScopes = append(policyScopes, &PolicyScope{
			Code:              policyScope.Code,
			Name:              policyScope.Name,
			Description:       policyScope.Description,
			RegionCodes:       policyScope.Regions,
			RegulationCodes:   policyScope.Regulations,
			RightsFulfillment: rightsFulfillment,
		})
	}

	var legalBases []*LegalBasis
	for _, legalBasis := range p.LegalBases {
		legalBases = append(legalBases, &LegalBasis{
			Code:                  legalBasis.Code,
			Name:                  legalBasis.Name,
			Description:           legalBasis.Description,
			RequiresOptIn:         legalBasis.RequiresOptIn,
			AllowOptOut:           legalBasis.AllowOptOut,
			RequiresPrivacyPolicy: legalBasis.RequiresPrivacyPolicy,
		})
	}

	var purposes []*Purpose
	for _, purpose := range p.Purposes {
		var cookies []*Cookie
		for _, cookie := range purpose.Cookies {
			cookies = append(cookies, &Cookie{
				Code:            cookie.Code,
				Name:            cookie.Name,
				Description:     cookie.Description,
				Host:            cookie.Host,
				Duration:        CookieDurationValues[cookie.Duration],
				Provenance:      CookieProvenanceValues[cookie.Provenance],
				Category:        CookieCategoryValues[cookie.Category],
				ServiceProvider: cookie.ServiceProvider,
			})
		}

		purposes = append(purposes, &Purpose{
			Code:                  purpose.Code,
			Name:                  purpose.Name,
			Description:           purpose.Description,
			TcfID:                 purpose.TcfID,
			TcfType:               purpose.TcfType,
			Editable:              purpose.Editable,
			LegalBasisRestriction: purpose.LegalBasisRestriction,
			DisplayName:           purpose.DisplayName,
			DisplayDescription:    purpose.DisplayDescription,
			ProcessingPurpose:     purpose.ProcessingPurpose,
			LegalBasis:            purpose.LegalBasis,
			Cookies:               cookies,
			CanonicalPurposes:     purpose.CanonicalPurposes,
			Translations:          purpose.Translations,
			DataSubjectRole:       DataSubjectRoleValues[purpose.DataSubjectRole],
			DataRole:              DataRoleValues[purpose.DataRole],
		})
	}

	var purposeTemplates []*PurposeTemplate
	for _, purposeTemplate := range p.PurposeTemplates {
		var cookies []*Cookie
		for _, cookie := range purposeTemplate.Cookies {
			cookies = append(cookies, &Cookie{
				Code:            cookie.Code,
				Name:            cookie.Name,
				Description:     cookie.Description,
				Host:            cookie.Host,
				Duration:        CookieDurationValues[cookie.Duration],
				Provenance:      CookieProvenanceValues[cookie.Provenance],
				Category:        CookieCategoryValues[cookie.Category],
				ServiceProvider: cookie.ServiceProvider,
			})
		}

		purposeTemplates = append(purposeTemplates, &PurposeTemplate{
			Code:                  purposeTemplate.Code,
			Name:                  purposeTemplate.Name,
			Description:           purposeTemplate.Description,
			TcfID:                 purposeTemplate.TcfID,
			TcfType:               purposeTemplate.TcfType,
			Editable:              purposeTemplate.Editable,
			LegalBasisRestriction: purposeTemplate.LegalBasisRestriction,
			DisplayName:           purposeTemplate.DisplayName,
			DisplayDescription:    purposeTemplate.DisplayDescription,
			LegalBasis:            purposeTemplate.LegalBasis,
			Cookies:               cookies,
			CanonicalPurposes:     purposeTemplate.CanonicalPurposes,
			Translations:          purposeTemplate.Translations,
			DataSubjectRole:       DataSubjectRoleValues[purposeTemplate.DataSubjectRole],
			DataRole:              DataRoleValues[purposeTemplate.DataRole],
		})
	}

	var workflows []*WorkflowDefinition
	for _, workflow := range p.Workflows {
		var steps []*Step
		for _, appConfigStep := range workflow.Steps {
			step := &Step{}

			if appConfigStep.Start != nil {
				step.ID = appConfigStep.Start.ID
				step.Code = appConfigStep.Start.Code
				step.Description = appConfigStep.Start.Description
				step.Start = &StartStep{
					Transition: appConfigStep.Start.Next,
					Params:     appConfigStep.Start.Params,
				}
			} else if appConfigStep.Finish != nil {
				step.ID = appConfigStep.Finish.ID
				step.Code = appConfigStep.Finish.Code
				step.Description = appConfigStep.Finish.Description
				step.Finish = &FinishStep{
					Params: appConfigStep.Finish.Params,
				}
			} else if appConfigStep.Activity != nil {
				step.ID = appConfigStep.Activity.ID
				step.Code = appConfigStep.Activity.Code
				step.Description = appConfigStep.Activity.Description
				step.Activity = &ActivityStep{
					Code: appConfigStep.Activity.Code,
					Options: &ActivityOptions{
						TaskQueue:              appConfigStep.Activity.Options.TaskQueue,
						ScheduleToCloseTimeout: appConfigStep.Activity.Options.ScheduleToCloseTimeout,
						ScheduleToStartTimeout: appConfigStep.Activity.Options.ScheduleToStartTimeout,
						StartToCloseTimeout:    appConfigStep.Activity.Options.StartToCloseTimeout,
						HeartbeatTimeout:       appConfigStep.Activity.Options.HeartbeatTimeout,
						WaitForCancellation:    appConfigStep.Activity.Options.WaitForCancellation,
					},
					Params:               appConfigStep.Activity.Params,
					Transition:           appConfigStep.Activity.Next,
					TemporalFunctionName: appConfigStep.Activity.Fn,
				}
			} else if appConfigStep.ChildWorkflow != nil {
				step.ID = appConfigStep.ChildWorkflow.ID
				step.Code = appConfigStep.ChildWorkflow.Code
				step.Description = appConfigStep.ChildWorkflow.Description
				step.Workflow = &ChildWorkflowStep{
					Code: appConfigStep.ChildWorkflow.Code,
					Options: &ChildWorkflowOptions{
						TaskQueue:                appConfigStep.ChildWorkflow.Options.TaskQueue,
						WorkflowExecutionTimeout: appConfigStep.ChildWorkflow.Options.WorkflowExecutionTimeout,
						WorkflowRunTimeout:       appConfigStep.ChildWorkflow.Options.WorkflowRunTimeout,
						WorkflowTaskTimeout:      appConfigStep.ChildWorkflow.Options.WorkflowTaskTimeout,
						WaitForCancellation:      appConfigStep.ChildWorkflow.Options.WaitForCancellation,
						Memo:                     appConfigStep.ChildWorkflow.Options.Memo,
						SearchAttributes:         appConfigStep.ChildWorkflow.Options.SearchAttributes,
						RetryPolicy:              appConfigStep.ChildWorkflow.Options.RetryPolicy,
					},
					Params:               appConfigStep.ChildWorkflow.Params,
					Transition:           appConfigStep.ChildWorkflow.Next,
					TemporalFunctionName: appConfigStep.ChildWorkflow.Fn,
				}
			} else if appConfigStep.Gateway != nil {
				step.ID = appConfigStep.Gateway.ID
				step.Code = appConfigStep.Gateway.Code
				step.Description = appConfigStep.Gateway.Description
				step.Gateway = &GatewayStep{
					Mode:        GatewayStepModeValues[appConfigStep.Gateway.Mode],
					Transitions: appConfigStep.Gateway.Next,
				}
			}

			steps = append(steps, step)
		}

		workflows = append(workflows, &WorkflowDefinition{
			Code:     workflow.Code,
			Name:     workflow.Name,
			Readonly: workflow.Readonly,
			Options:  workflow.Options,
			Steps:    steps,
		})
	}

	var activities []*WorkflowActivityDefinition
	for _, activity := range p.Activities {
		// TODO: Upload icon
		icon := &IconDefinition{}

		activities = append(activities, &WorkflowActivityDefinition{
			Code: activity.Code,
			Name: activity.Name,
			Icon: icon,
			Activity: &ActivityDefinition{
				Options: &ActivityOptions{
					TaskQueue:              activity.Options.TaskQueue,
					ScheduleToCloseTimeout: activity.Options.ScheduleToCloseTimeout,
					ScheduleToStartTimeout: activity.Options.ScheduleToStartTimeout,
					StartToCloseTimeout:    activity.Options.StartToCloseTimeout,
					HeartbeatTimeout:       activity.Options.HeartbeatTimeout,
					WaitForCancellation:    activity.Options.WaitForCancellation,
				},
				Params:               activity.Params,
				Outputs:              activity.Outputs,
				Config:               activity.Config,
				TemporalFunctionName: activity.Fn,
			},
		})
	}

	var childWorkflows []*WorkflowActivityDefinition
	for _, childWorkflow := range p.ChildWorkflows {
		// TODO: Upload icon
		icon := &IconDefinition{}

		childWorkflows = append(childWorkflows, &WorkflowActivityDefinition{
			Code: childWorkflow.Code,
			Name: childWorkflow.Name,
			Icon: icon,
			Workflow: &ChildWorkflowDefinition{
				Options: &ChildWorkflowOptions{
					TaskQueue:                childWorkflow.Options.TaskQueue,
					WorkflowExecutionTimeout: childWorkflow.Options.WorkflowExecutionTimeout,
					WorkflowRunTimeout:       childWorkflow.Options.WorkflowRunTimeout,
					WorkflowTaskTimeout:      childWorkflow.Options.WorkflowTaskTimeout,
					WaitForCancellation:      childWorkflow.Options.WaitForCancellation,
				},
				Params:               childWorkflow.Params,
				Outputs:              childWorkflow.Outputs,
				Config:               childWorkflow.Config,
				TemporalFunctionName: childWorkflow.Fn,
			},
		})
	}

	var eventTypes []string
	if p.Webhook != nil {
		eventTypes = p.Webhook.Events
	}

	var cookies []*Cookie
	for _, cookie := range p.Cookies {
		cookies = append(cookies, &Cookie{
			Code:            cookie.Code,
			Name:            cookie.Name,
			Description:     cookie.Description,
			Host:            cookie.Host,
			Duration:        CookieDurationValues[cookie.Duration],
			Provenance:      CookieProvenanceValues[cookie.Provenance],
			Category:        CookieCategoryValues[cookie.Category],
			ServiceProvider: cookie.ServiceProvider,
			AppCode:         p.Code,
		})
	}

	transponderAppDetailsJSON, err := json.Marshal(&TransponderAppDetails{
		ResourceTypes: p.ResourceTypes,
	})
	if err != nil {
		return nil, err
	}

	return &App{
		ID:                         p.ID,
		Code:                       p.Code,
		OrgCode:                    p.OrgCode,
		Name:                       p.Name,
		Version:                    p.Version,
		Depends:                    p.Depends,
		Provides:                   p.Provides,
		Type:                       AppTypeValues[p.Type],
		AutoUpgrade:                p.AutoUpgrade,
		Instances:                  AppAllowedInstancesValues[p.Instances],
		Rules:                      p.Rules,
		Capabilities:               appCapabilities,
		SupportedLanguages:         p.SupportedLanguages,
		SupportedPurposes:          p.SupportedPurposes,
		SupportedRights:            p.SupportedRights,
		PermissionNote:             p.PermissionNote,
		Permissions:                p.Permissions,
		InfoUrl:                    p.InfoUrl,
		SetupUrl:                   p.SetupUrl,
		HomepageUrl:                p.HomepageUrl,
		ExpireUserTokens:           p.ExpireUserTokens,
		RefreshInterval:            refreshIntervalHours,
		RequestUserAuth:            p.RequestUserAuth,
		UserAuthCallbackUrl:        p.UserAuthCallbackUrl,
		RedirectOnUpdate:           p.RedirectOnUpdate,
		FormTitle:                  p.FormTitle,
		FormSubtitle:               p.FormSubtitle,
		Form:                       p.Form,
		IdentitySpaces:             p.IdentitySpaces,
		Purposes:                   purposes,
		Rights:                     p.Rights,
		Regulations:                p.Regulations,
		Tcf:                        p.Tcf,
		Workflows:                  workflows,
		Activities:                 activities,
		ChildWorkflows:             childWorkflows,
		PurposeTemplates:           purposeTemplates,
		PurposeTemplateCollections: p.PurposeTemplateCollections,
		LegalBasisRestrictions:     p.LegalBasisRestrictions,
		PolicyScopes:               policyScopes,
		LegalBases:                 legalBases,
		Themes:                     p.Themes,
		EventTypes:                 eventTypes,
		Cookies:                    cookies,
		TransponderAppDetailsJSON:  transponderAppDetailsJSON,
	}, nil
}

func NewAppMarketplaceEntry(p ManifestInputs) *AppMarketplaceEntry {
	var appCapabilities []int32
	for _, capability := range p.Capabilities {
		if appCapability, ok := AppCapabilityValues[capability]; ok {
			appCapabilities = append(appCapabilities, appCapability)
		}
	}

	var appContacts []*AppContact
	for _, contact := range p.Contacts {
		appContacts = append(appContacts, &AppContact{
			ContactType: contact.ContactType,
			Email:       contact.Email,
		})
	}

	primaryCategory := AppMarketplaceCategoryValues[p.PrimaryCategory]
	secondaryCategory := AppMarketplaceCategoryValues[p.SecondaryCategory]

	return &AppMarketplaceEntry{
		AppCode:             p.Code,
		Version:             p.Version,
		Contacts:            appContacts,
		ShortDescription:    p.ShortDescription,
		PrimaryCategory:     primaryCategory,
		SecondaryCategory:   secondaryCategory,
		SupportedLanguages:  p.SupportedLanguages,
		CustomerSupportUrl:  p.CustomerSupportUrl,
		PrivacyPolicyUrl:    p.PrivacyPolicyUrl,
		StatusUrl:           p.StatusUrl,
		TosUrl:              p.TosUrl,
		DocUrl:              p.DocUrl,
		Logo:                Image{},
		IntroDescription:    p.ShortDescription,
		DetailedDescription: p.DetailedDescription,
	}
}

type TransponderAppDetails struct {
	ResourceTypes []*ResourceType `json:"resourceTypes,omitempty"`
}

type ResourceType struct {
	Code                    string `yaml:"code,omitempty" json:"code,omitempty"`
	Name                    string `yaml:"name,omitempty" json:"name,omitempty"`
	IsScopable              bool   `yaml:"isScopable,omitempty" json:"isScopable,omitempty"`
	HasData                 bool   `yaml:"hasData,omitempty" json:"hasData,omitempty"`
	CanBeLabeled            bool   `yaml:"canBeLabeled,omitempty" json:"canBeLabeled,omitempty"`
	IsDisplayedInNavigation bool   `yaml:"isDisplayedInNavigation,omitempty" json:"isDisplayedInNavigation,omitempty"`
	IsDisplayedInSummary    bool   `yaml:"isDisplayedInSummary,omitempty" json:"isDisplayedInSummary,omitempty"`
}
