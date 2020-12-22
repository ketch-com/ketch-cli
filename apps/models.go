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
	"permit_propagation":         1,
	"permit_ingestion":           2,
	"rights_propagation":         3,
	"rights_ingestion":           4,
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

type TextField struct {
	// Default value will be the value for this field, before user interaction. Having a default value will override the placeholder text
	DefaultValue string `yaml:"default_value,omitempty" json:"defaultValue,omitempty"`
	// The placeholder text that will appear when this field is empty
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`
	// This setting will enable spell check on the field.
	Spellcheck bool `yaml:"spellcheck,omitempty" json:"spellcheck,omitempty"`
}

type Select struct {
	// The placeholder text that will appear before an option is selected
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`

	// Default value will be the value for this field, before user interaction. Having a default value will override the placeholder text
	DefaultValue string `yaml:"default_value,omitempty" json:"defaultValue,omitempty"`

	// If true, multiple values can be added in this field.
	// The values will appear as an array in the API and an “Add Another” button will be visible on the field allowing the creation of additional fields for this component
	Multiple bool `yaml:"multiple,omitempty" json:"multiple,omitempty"`

	// Data is the definition of how data is provided for the dropdown
	Data *SelectData `yaml:"data,omitempty" json:"data,omitempty"`
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
	Key string `yaml:"key,omitempty" json:"key,omitempty"`
	// The name or title for this component
	Label string `yaml:"label,omitempty" json:"label,omitempty"`
	// The type property will be used to select which component to render on the frontend. It cannot be an existing field type
	Type string `yaml:"type,omitempty" json:"type,omitempty"`
	// If true, the field will be required to have a value
	Required  bool       `yaml:"required,omitempty" json:"required,omitempty"`
	MinLength int64      `yaml:"minLength,omitempty" json:"minLength,omitempty"`
	MaxLength int64      `yaml:"maxLength,omitempty" json:"maxLength,omitempty"`
	Pattern   string     `yaml:"pattern,omitempty" json:"pattern,omitempty"`
	TextField *TextField `yaml:"text_field,omitempty" json:"textField,omitempty"`
	Select    *Select    `yaml:"select,omitempty" json:"select,omitempty"`
}

type WebHook struct {
	URL           string   `yaml:"url" json:"url,omitempty"`
	Secret        string   `yaml:"secret" json:"secret,omitempty"`
	Authorization string   `yaml:"authorization,omitempty" json:"authorization,omitempty"`
	Events        []string `yaml:"events,omitempty" json:"events,omitempty"`
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

type WorkflowOptions struct {
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

type WorkflowConfig struct {
	Code     string
	Name     string
	Icon     string
	Type     string
	Function string `yaml:"fn" json:"fn"`
	Options  *WorkflowOptions
}

type PublishAppConfig struct {
	ID                  string              `yaml:"id,omitempty" json:"id,omitempty"`
	Name                string              `yaml:"name,omitempty" json:"name,omitempty"`
	OrgCode             string              `yaml:"org,omitempty" json:"org,omitempty"`
	Version             string              `yaml:"version,omitempty" json:"version,omitempty"`
	PrimaryCategory     string              `yaml:"primary_category,omitempty" json:"primary_category,omitempty"`
	SecondaryCategory   string              `yaml:"secondary_category,omitempty" json:"secondary_category,omitempty"`
	Rules               map[string]string   `yaml:"rules" json:"rules,omitempty"`
	Capabilities        []string            `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
	SupportedLanguages  []string            `yaml:"supported_languages,omitempty" json:"supported_languages,omitempty"`
	Purposes            []string            `yaml:"purposes,omitempty" json:"purposes,omitempty"`
	Rights              []string            `yaml:"rights,omitempty" json:"rights,omitempty"`
	IdentitySpaces      []*IdentitySpace    `yaml:"identity_spaces,omitempty" json:"identity_spaces,omitempty"`
	ShortDescription    string              `yaml:"short_description,omitempty" json:"short_description,omitempty"`
	DetailedDescription string              `yaml:"detailed_description,omitempty" json:"detailed_description,omitempty"`
	PermissionNote      string              `yaml:"permission_note,omitempty" json:"permission_note,omitempty"`
	Permissions         []string            `yaml:"permissions,omitempty" json:"permissions,omitempty"`
	SetupUrl            string              `yaml:"setup_url,omitempty" json:"setup_url,omitempty"`
	HomepageUrl         string              `yaml:"homepage_url,omitempty" json:"homepage_url,omitempty"`
	CustomerSupportUrl  string              `yaml:"customer_support_url,omitempty" json:"customer_support_url,omitempty"`
	PrivacyPolicyUrl    string              `yaml:"privacy_policy_url,omitempty" json:"privacy_policy_url,omitempty"`
	StatusUrl           string              `yaml:"status_url,omitempty" json:"status_url,omitempty"`
	TosUrl              string              `yaml:"tos_url,omitempty" json:"tos_url,omitempty"`
	DocUrl              string              `yaml:"doc_url,omitempty" json:"doc_url,omitempty"`
	Logo                *AppConfigImage     `yaml:"logo,omitempty" json:"logo,omitempty"`
	Previews            []*AppConfigImage   `yaml:"previews,omitempty" json:"previews,omitempty"`
	Contacts            []*AppConfigContact `yaml:"contacts,omitempty" json:"contacts,omitempty"`
	ExpireUserTokens    bool                `yaml:"expire_user_tokens,omitempty" json:"expire_user_tokens,omitempty"`
	RefreshInterval     string              `yaml:"refresh_interval,omitempty" json:"refresh_interval,omitempty"`
	RequestUserAuth     bool                `yaml:"request_user_auth,omitempty" json:"request_user_auth,omitempty"`
	UserAuthCallbackUrl string              `yaml:"user_auth_callback_url,omitempty" json:"user_auth_callback_url,omitempty"`
	RedirectOnUpdate    bool                `yaml:"redirect_on_update,omitempty" json:"redirect_on_update,omitempty"`
	Webhook             *WebHook            `yaml:"webhook,omitempty" json:"webhook,omitempty"`
	Workflow            []*WorkflowConfig   `yaml:"workflow,flow,omitempty" json:"workflow,omitempty"`
	Form                []*FormComponent    `yaml:"form,omitempty" json:"form,omitempty"`
}

type App struct {
	ID                  string            `json:"ID,omitempty"`
	OrgCode             string            `yaml:"orgCode" json:"orgCode,omitempty"`
	Code                string            `json:"code,omitempty"`
	Name                string            `json:"name,omitempty"`
	Version             string            `json:"version,omitempty"`
	Readme              string            `json:"readme,omitempty"`
	HomepageUrl         string            `json:"homepage_url,omitempty"`
	UserAuthCallbackUrl string            `json:"user_auth_callback_url,omitempty"`
	ExpireUserTokens    bool              `json:"expire_user_tokens,omitempty"`
	RequestUserAuth     bool              `json:"request_user_auth,omitempty"`
	SetupUrl            string            `json:"setup_url,omitempty"`
	RedirectOnUpdate    bool              `json:"redirect_on_update,omitempty"`
	WebhookId           string            `json:"webhook_id,omitempty"`
	Capabilities        []int32           `yaml:",flow"`
	Permissions         []string          `yaml:",flow"`
	PermissionNote      string            `json:"permission_node,omitempty"`
	Purposes            []string          `yaml:",flow"`
	Form                []*FormComponent  `yaml:",flow"`
	IdentitySpaces      []*IdentitySpace  `yaml:",flow" json:"identity_spaces,omitempty"`
	Rights              []string          `yaml:",flow" json:"rights,omitempty"`
	Rules               map[string]string `yaml:",flow" json:"rules,omitempty"`
	RefreshInterval     time.Duration     `yaml:"refresh_interval" json:"refresh_interval,omitempty"`
}

type AppMarketplaceEntry struct {
	AppID               string `json:"appID,omitempty"`
	Version             string
	Contacts            []*AppContact `yaml:",flow"`
	ShortDescription    string        `json:"short_description,omitempty"`
	PrimaryCategory     int32         `yaml:",inline" json:"primary_category,omitempty"`
	SecondaryCategory   int32         `yaml:",inline" json:"secondary_category,omitempty"`
	SupportedLanguages  []string      `yaml:",flow" json:"supported_languages,omitempty"`
	CustomerSupportUrl  string        `json:"customer_support_url,omitempty"`
	PrivacyPolicyUrl    string        `json:"privacy_policy_url,omitempty"`
	StatusUrl           string        `json:"status_url,omitempty"`
	TosUrl              string        `json:"tos_url,omitempty"`
	DocUrl              string        `json:"doc_url,omitempty"`
	Logo                Image         `yaml:",flow"`
	IntroDescription    string        `json:"intro_description,omitempty"`
	DetailedDescription string        `json:"detailed_description,omitempty"`
	Previews            []*Image      `yaml:",flow" json:"previews,omitempty"`
}

type Image struct {
	Data   []byte `yaml:"data" json:"data,omitempty"`
	Title  string `yaml:"title" json:"title,omitempty"`
	Url    string `yaml:"url" json:"url,omitempty"`
	Width  int32  `yaml:"width" json:"width,omitempty"`
	Height int32  `yaml:"height" json:"height,omitempty"`
}

type WebhookResponse struct {
	Webhook *WebHook
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
		ID:                  p.ID,
		Name:                p.Name,
		OrgCode:             p.OrgCode,
		Version:             p.Version,
		HomepageUrl:         p.HomepageUrl,
		UserAuthCallbackUrl: p.UserAuthCallbackUrl,
		ExpireUserTokens:    p.ExpireUserTokens,
		RequestUserAuth:     p.RequestUserAuth,
		SetupUrl:            p.SetupUrl,
		RedirectOnUpdate:    p.RedirectOnUpdate,
		Capabilities:        appCapabilities,
		Permissions:         p.Permissions,
		PermissionNote:      p.PermissionNote,
		Purposes:            p.Purposes,
		Form:                p.Form,
		IdentitySpaces:      p.IdentitySpaces,
		Rights:              p.Rights,
		Rules:               p.Rules,
		RefreshInterval:     refreshInterval,
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
