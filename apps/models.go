package apps

import (
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

var AppCapabilityValues = map[string]int32{
	"UNSPECIFIED_APP_CAPABILITY": 0,
	"permitPropagation":          1,
	"permitIngestion":            2,
	"rightPropagation":           3,
	"rightIngestion":             4,
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

var AppTypeValues = map[string]int32{
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

type SelectData struct {
	// These are the values that will be selected on this field
	Values []*SelectDataValue `yaml:"values,omitempty" json:"values,omitempty"`
	// Enter a JSON Array to use. It should be formatted as an array of objects with named properties
	Json string `yaml:"json,omitempty" json:"json,omitempty"`
	// Enter a url with a data source in JSON Array format. This can be used to populate a Select list with external JSON values
	Url string `yaml:"url,omitempty" json:"url,omitempty"`
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

type WorkflowRetryPolicy struct {
	InitialInterval          int64    `yaml:"initialInterval,omitempty" json:"initialInterval,omitempty"`
	BackoffCoefficient       float64  `yaml:"backoffCoefficient,omitempty" json:"backoffCoefficient,omitempty"`
	MaximumInterval          int64    `yaml:"maximumInterval,omitempty" json:"maximumInterval,omitempty"`
	MaximumAttempts          int64    `yaml:"maximumAttempts,omitempty" json:"maximumAttempts,omitempty"`
	NonRetryableErrorReasons []string `yaml:"nonRetryableErrorReasons,omitempty" json:"nonRetryableErrorReasons,omitempty"`
}

type WorkflowStepActivityOptions struct {
	TaskQueue              string `yaml:"taskQueue,omitempty" json:"taskQueue,omitempty"`
	ScheduleToCloseTimeout int64  `yaml:"scheduleToCloseTimeout,omitempty" json:"scheduleToCloseTimeout,omitempty"`
	ScheduleToStartTimeout int64  `yaml:"scheduleToStartTimeout,omitempty" json:"scheduleToStartTimeout,omitempty"`
	StartToCloseTimeout    int64  `yaml:"startToCloseTimeout,omitempty" json:"startToCloseTimeout,omitempty"`
	HeartbeatTimeout       int64  `yaml:"heartbeatTimeout,omitempty" json:"heartbeatTimeout,omitempty"`
	WaitForCancellation    bool   `yaml:"waitForCancellation,omitempty" json:"waitForCancellation,omitempty"`
	RetryPolicy            string `yaml:"retryPolicy,omitempty" json:"retryPolicy,omitempty"`
}

type WorkflowStepActivity struct {
	ID          string                      `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                      `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                      `yaml:"description,omitempty" json:"description,omitempty"`
	Step        string                      `yaml:"step,omitempty" json:"step,omitempty"`
	Fn          string                      `yaml:"fn,omitempty" json:"fn,omitempty"`
	Next        string                      `yaml:"next,omitempty" json:"next,omitempty"`
	Options     WorkflowStepActivityOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Params      map[string]interface{}      `yaml:"params,omitempty" json:"params,omitempty"`
}

type WorkflowStepChildWorkflowOptions struct {
	TaskQueue                string                 `yaml:"taskQueue,omitempty" json:"taskQueue,omitempty"`
	WorkflowExecutionTimeout int64                  `yaml:"workflowExecutionTimeout,omitempty" json:"workflowExecutionTimeout,omitempty"`
	WorkflowRunTimeout       int64                  `yaml:"workflowRunTimeout,omitempty" json:"workflowRunTimeout,omitempty"`
	WorkflowTaskTimeout      int64                  `yaml:"workflowTaskTimeout,omitempty" json:"workflowTaskTimeout,omitempty"`
	WaitForCancellation      bool                   `yaml:"waitForCancellation,omitempty" json:"waitForCancellation,omitempty"`
	ParentClosePolicy        string                 `yaml:"parentClosePolicy,omitempty" json:"parentClosePolicy,omitempty"`
	Memo                     map[string]interface{} `yaml:"memo,omitempty" json:"memo,omitempty"`
	SearchAttributes         map[string]interface{} `yaml:"searchAttributes,omitempty" json:"searchAttributes,omitempty"`
	RetryPolicy              WorkflowRetryPolicy    `yaml:"retryPolicy,omitempty" json:"retryPolicy,omitempty"`
}

type WorkflowStepChildWorkflow struct {
	ID          string                           `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                           `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                           `yaml:"description,omitempty" json:"description,omitempty"`
	Step        string                           `yaml:"step,omitempty" json:"step,omitempty"`
	Fn          string                           `yaml:"fn,omitempty" json:"fn,omitempty"`
	Next        string                           `yaml:"next,omitempty" json:"next,omitempty"`
	Options     WorkflowStepChildWorkflowOptions `yaml:"options,omitempty" json:"options,omitempty"`
	Params      map[string]interface{}           `yaml:"params,omitempty" json:"params,omitempty"`
}

type WorkflowStepGatewayNext struct {
	ID       string `yaml:"id,omitempty" json:"id,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	Variable string `yaml:"variable,omitempty" json:"variable,omitempty"`
	Operator string `yaml:"operator,omitempty" json:"operator,omitempty"`
	Operand  string `yaml:"operand,omitempty" json:"operand,omitempty"`
}

type WorkflowStepGateway struct {
	ID          string                    `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                    `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                    `yaml:"description,omitempty" json:"description,omitempty"`
	Mode        string                    `yaml:"mode,omitempty" json:"mode,omitempty"`
	Next        []WorkflowStepGatewayNext `yaml:"next,omitempty" json:"next,omitempty"`
}

type WorkflowStepStart struct {
	ID          string                 `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                 `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                 `yaml:"description,omitempty" json:"description,omitempty"`
	Params      map[string]interface{} `yaml:"params,omitempty" json:"params,omitempty"`
}

type WorkflowStepFinish struct {
	ID          string                 `yaml:"id,omitempty" json:"id,omitempty"`
	Code        string                 `yaml:"code,omitempty" json:"code,omitempty"`
	Description string                 `yaml:"description,omitempty" json:"description,omitempty"`
	Next        string                 `yaml:"next,omitempty" json:"next,omitempty"`
	Params      map[string]interface{} `yaml:"params,omitempty" json:"params,omitempty"`
}

type WorkflowStep struct {
	Activity      WorkflowStepActivity      `yaml:"activity,omitempty" json:"activity,omitempty"`
	ChildWorkflow WorkflowStepChildWorkflow `yaml:"childWorkflow,omitempty" json:"childWorkflow,omitempty"`
	Gateway       WorkflowStepGateway       `yaml:"gateway,omitempty" json:"gateway,omitempty"`
	Start         WorkflowStepStart         `yaml:"start,omitempty" json:"start,omitempty"`
	Finish        WorkflowStepFinish        `yaml:"finish,omitempty" json:"finish,omitempty"`
}

type WorkflowOptions struct {
	TaskQueue                string              `yaml:"taskQueue,omitempty" json:"taskQueue,omitempty"`
	WorkflowExecutionTimeout int64               `yaml:"workflowExecutionTimeout,omitempty" json:"workflowExecutionTimeout,omitempty"`
	WorkflowRunTimeout       int64               `yaml:"workflowRunTimeout,omitempty" json:"workflowRunTimeout,omitempty"`
	WorkflowTaskTimeout      int64               `yaml:"workflowTaskTimeout,omitempty" json:"workflowTaskTimeout,omitempty"`
	Memo                     map[string]string   `yaml:"memo,omitempty" json:"memo,omitempty"`
	SearchAttributes         map[string]string   `yaml:"searchAttributes,omitempty" json:"searchAttributes,omitempty"`
	RetryPolicy              WorkflowRetryPolicy `yaml:"retryPolicy,omitempty" json:"retryPolicy,omitempty"`
}

type Workflow struct {
	Code     string
	Name     string
	Readonly bool
	Options  *WorkflowOptions
	Steps    []WorkflowStep `yaml:"steps,omitempty" json:"steps,omitempty"`
}

type ActivityOptions struct {
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

type ActivityParam struct {
	Code    string `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Type    string `yaml:"type,omitempty" json:"type,omitempty"`
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
}

type ActivityOutput struct {
	Code    string `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Type    string `yaml:"type,omitempty" json:"type,omitempty"`
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
}

type Activity struct {
	Code    string                 `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string                 `yaml:"name,omitempty" json:"name,omitempty"`
	Icon    string                 `yaml:"icon,omitempty" json:"icon,omitempty"`
	Fn      string                 `yaml:"fn,omitempty" json:"fn,omitempty"`
	Options ActivityOptions        `yaml:"options,omitempty" json:"options,omitempty"`
	Params  []*ActivityParam       `yaml:"params,omitempty" json:"params,omitempty"`
	Outputs []*ActivityOutput      `yaml:"outputs,omitempty" json:"outputs,omitempty"`
	Config  map[string]interface{} `yaml:"config,omitempty" json:"config,omitempty"`
}

type ChildWorkflow struct {
	Code    string           `yaml:"code,omitempty" json:"code,omitempty"`
	Name    string           `yaml:"name,omitempty" json:"name,omitempty"`
	Icon    string           `yaml:"icon,omitempty" json:"icon,omitempty"`
	Fn      string           `yaml:"fn,omitempty" json:"fn,omitempty"`
	Options ActivityOptions  `yaml:"options,omitempty" json:"options,omitempty"`
	Params  []*ActivityParam `yaml:"params,omitempty" json:"params,omitempty"`
}

type PurposeTemplate struct {
	Code                  string `yaml:"code,omitempty" json:"code,omitempty"`
	Name                  string `yaml:"name,omitempty" json:"name,omitempty"`
	Description           string `yaml:"description,omitempty" json:"description,omitempty"`
	TcfID                 int    `yaml:"tcfId,omitempty" json:"tcfId,omitempty"`
	TcfType               string `yaml:"tcfType,omitempty" json:"tcfType,omitempty"`
	Editable              bool   `yaml:"editable,omitempty" json:"editable,omitempty"`
	LegalBasisRestriction string `yaml:"legalBasisRestriction,omitempty" json:"legalBasisRestriction,omitempty"`
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

type Purpose struct {
	Code                  string             `json:"code,omitempty"`
	Name                  string             `json:"name,omitempty"`
	Description           string             `json:"description,omitempty"`
	TcfID                 int                `json:"tcfId,omitempty"`
	TcfType               string             `json:"tcfType,omitempty"`
	Editable              bool               `json:"editable,omitempty"`
	LegalBasisRestriction string             `json:"legalBasisRestriction,omitempty"`
	DisplayName           string             `json:"displayName,omitempty"`
	DisplayDescription    string             `json:"displayDescription,omitempty"`
	ProcessingPurpose     string             `json:"processingPurpose,omitempty"`
	LegalBasis            map[string]string  `json:"legalBasis,omitempty"`
	Cookies               []*AppConfigCookie `json:"cookies,omitempty"`
	CanonicalPurposes     []string           `json:"canonicalPurposes,omitempty"`
	Translations          map[string]string  `json:"translations,omitempty"`
	DataSubjectRole       int32              `json:"dataSubjectRole,omitempty"`
	DataRole              int32              `json:"dataRole,omitempty"`
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
	Version           int64               `json:"version,omitempty"`
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

type PublishAppConfig struct {
	Code                   string                   `yaml:"code,omitempty" json:"code,omitempty"`
	OrgCode                string                   `yaml:"org,omitempty" json:"org,omitempty"`
	Name                   string                   `yaml:"name,omitempty" json:"name,omitempty"`
	Version                string                   `yaml:"version,omitempty" json:"version,omitempty"`
	Depends                string                   `yaml:"depends,omitempty" json:"depends,omitempty"`
	Provides               []string                 `yaml:"provides,omitempty" json:"provides,omitempty"`
	Type                   string                   `yaml:"type,omitempty" json:"type,omitempty"`
	AutoUpgrade            bool                     `yaml:"autoUpgrade,omitempty" json:"autoUpgrade,omitempty"`
	Instances              string                   `yaml:"instances,omitempty" json:"instances,omitempty"`
	PrimaryCategory        string                   `yaml:"primaryCategory,omitempty" json:"primaryCategory,omitempty"`
	SecondaryCategory      string                   `yaml:"secondaryCategory,omitempty" json:"secondaryCategory,omitempty"`
	Rules                  map[string]string        `yaml:"rules" json:"rules,omitempty"`
	Capabilities           []string                 `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
	SupportedLanguages     []string                 `yaml:"supportedLanguages,omitempty" json:"supportedLanguages,omitempty"`
	SupportedPurposes      []string                 `yaml:"supportedPurposes,omitempty" json:"supportedPurposes,omitempty"`
	SupportedRights        []string                 `yaml:"supportedRights,omitempty" json:"supportedRights,omitempty"`
	ShortDescription       string                   `yaml:"shortDescription,omitempty" json:"shortDescription,omitempty"`
	DetailedDescription    string                   `yaml:"detailedDescription,omitempty" json:"detailedDescription,omitempty"`
	PermissionNote         string                   `yaml:"permissionNote,omitempty" json:"permissionNote,omitempty"`
	Permissions            []string                 `yaml:"permissions,omitempty" json:"permissions,omitempty"`
	SetupUrl               string                   `yaml:"setupURL,omitempty" json:"setupURL,omitempty"`
	HomepageUrl            string                   `yaml:"homepageURL,omitempty" json:"homepageURL,omitempty"`
	CustomerSupportUrl     string                   `yaml:"customerSupportURL,omitempty" json:"customerSupportURL,omitempty"`
	PrivacyPolicyUrl       string                   `yaml:"privacyPolicyURL,omitempty" json:"privacyPolicyURL,omitempty"`
	StatusUrl              string                   `yaml:"statusURL,omitempty" json:"statusURL,omitempty"`
	TosUrl                 string                   `yaml:"tosURL,omitempty" json:"tosURL,omitempty"`
	DocUrl                 string                   `yaml:"docURL,omitempty" json:"docURL,omitempty"`
	Logo                   *AppConfigImage          `yaml:"logo,omitempty" json:"logo,omitempty"`
	Previews               []*AppConfigImage        `yaml:"previews,omitempty" json:"previews,omitempty"`
	Contacts               []*AppConfigContact      `yaml:"contacts,omitempty" json:"contacts,omitempty"`
	ExpireUserTokens       bool                     `yaml:"expireUserTokens,omitempty" json:"expireUserTokens,omitempty"`
	RefreshInterval        string                   `yaml:"refreshInterval,omitempty" json:"refreshInterval,omitempty"`
	RequestUserAuth        bool                     `yaml:"requestUserAuth,omitempty" json:"requestUserAuth,omitempty"`
	UserAuthCallbackUrl    string                   `yaml:"userAuthCallbackURL,omitempty" json:"userAuthCallbackURL,omitempty"`
	RedirectOnUpdate       bool                     `yaml:"redirectOnUpdate,omitempty" json:"redirectOnUpdate,omitempty"`
	Webhook                *Webhook                 `yaml:"webhook,omitempty" json:"webhook,omitempty"`
	Form                   []*FormComponent         `yaml:"form,omitempty" json:"form,omitempty"`
	IdentitySpaces         []*IdentitySpace         `yaml:"identitySpaces,omitempty" json:"identitySpaces,omitempty"`
	Workflow               []*Workflow              `yaml:"workflow,flow,omitempty" json:"workflow,omitempty"`
	Activities             []*Activity              `yaml:"activities,flow,omitempty" json:"activities,omitempty"`
	ChildWorkflows         []*ChildWorkflow         `yaml:"childWorkflows,flow,omitempty" json:"childWorkflows,omitempty"`
	PurposeTemplates       []*PurposeTemplate       `yaml:"purposeTemplates,flow,omitempty" json:"purposeTemplates,omitempty"`
	Purposes               []*AppConfigPurpose      `yaml:"purposes,flow,omitempty" json:"purposes,omitempty"`
	LegalBasisRestrictions []*LegalBasisRestriction `yaml:"legalBasisRestrictions,flow,omitempty" json:"legalBasisRestrictions,omitempty"`
	PolicyScopes           []*AppConfigPolicyScope  `yaml:"policyScopes,flow,omitempty" json:"policyScopes,omitempty"`
	LegalBases             []*AppConfigLegalBasis   `yaml:"legalBases,flow,omitempty" json:"legalBases,omitempty"`
	Themes                 []*Theme                 `yaml:"themes,flow,omitempty" json:"themes,omitempty"`
}

type App struct {
	ID                  string            `json:"id,omitempty"`
	Code                string            `json:"code,omitempty"`
	OrgCode             string            `yaml:"orgCode" json:"orgCode,omitempty"`
	Name                string            `json:"name,omitempty"`
	Version             string            `json:"version,omitempty"`
	Depends             string            `json:"depends,omitempty"`
	Provides            []string          `json:"provides,omitempty"`
	Type                int32             `json:"type,omitempty"`
	AutoUpgrade         bool              `json:"autoUpgrade,omitempty"`
	Instances           int32             `json:"instances,omitempty"`
	Rules               map[string]string `yaml:",flow" json:"rules,omitempty"`
	Capabilities        []int32           `yaml:",flow"`
	PermissionNote      string            `json:"permissionNode,omitempty"`
	Permissions         []string          `yaml:",flow"`
	SetupUrl            string            `json:"setupURL,omitempty"`
	HomepageUrl         string            `json:"homepageURL,omitempty"`
	ExpireUserTokens    bool              `json:"expireUserTokens,omitempty"`
	RefreshInterval     time.Duration     `yaml:"refreshInterval" json:"refreshInterval,omitempty"`
	RequestUserAuth     bool              `json:"requestUserAuth,omitempty"`
	UserAuthCallbackUrl string            `json:"userAuthCallbackURL,omitempty"`
	RedirectOnUpdate    bool              `json:"redirectOnUpdate,omitempty"`
	WebhookId           string            `json:"webhookID,omitempty"`
	Form                []*FormComponent  `yaml:",flow"`
	IdentitySpaces      []*IdentitySpace  `yaml:",flow" json:"identitySpaces,omitempty"`
	Purposes            []string          `yaml:",flow"`
	Rights              []string          `yaml:",flow" json:"rights,omitempty"`
	Readme              string            `json:"readme,omitempty"`
	// TODO: Is it also should be sent using createApp?
	//Workflow               []*Workflow              `json:"workflow,omitempty"`
	//Activities             []*Activity              `json:"activities,omitempty"`
	//ChildWorkflows         []*ChildWorkflow         `json:"childWorkflows,omitempty"`
	//PurposeTemplates       []*PurposeTemplate       `json:"purposeTemplates,omitempty"`
	//LegalBasisRestrictions []*LegalBasisRestriction `json:"legalBasisRestrictions,omitempty"`
	//PolicyScopes           []*PolicyScope           `json:"policyScopes,omitempty"`
	//LegalBases             []*LegalBasis            `json:"legalBases,omitempty"`
	//Themes                 []*Theme                 `json:"themes,omitempty"`
}

type AppMarketplaceEntry struct {
	AppID               string `json:"appID,omitempty"`
	Version             string
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

type WebhookResponse struct {
	Webhook *Webhook
}

type PutAppResponse struct {
	App *App
}

func NewApp(p PublishAppConfig) (*App, error) {
	var appCapabilities []int32
	for _, capability := range p.Capabilities {
		if appCapability, ok := AppCapabilityValues[capability]; ok {
			appCapabilities = append(appCapabilities, appCapability)
		}
	}

	refreshInterval, err := time.ParseDuration(p.RefreshInterval)
	if err != nil {
		return nil, err
	}

	return &App{
		Code:                p.Code,
		OrgCode:             p.OrgCode,
		Name:                p.Name,
		Version:             p.Version,
		Depends:             p.Depends,
		Provides:            p.Provides,
		Type:                AppTypeValues[p.Type],
		AutoUpgrade:         p.AutoUpgrade,
		Instances:           AppAllowedInstancesValues[p.Instances],
		Rules:               p.Rules,
		Capabilities:        appCapabilities,
		PermissionNote:      p.PermissionNote,
		Permissions:         p.Permissions,
		SetupUrl:            p.SetupUrl,
		HomepageUrl:         p.HomepageUrl,
		ExpireUserTokens:    p.ExpireUserTokens,
		RefreshInterval:     refreshInterval,
		RequestUserAuth:     p.RequestUserAuth,
		UserAuthCallbackUrl: p.UserAuthCallbackUrl,
		RedirectOnUpdate:    p.RedirectOnUpdate,
		Form:                p.Form,
		IdentitySpaces:      p.IdentitySpaces,
		Purposes:            p.SupportedPurposes,
		Rights:              p.SupportedRights,
	}, nil
}

func NewAppMarketplaceEntry(p PublishAppConfig) *AppMarketplaceEntry {
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
		//TODO: Previews:            p.Previews,
	}
}
