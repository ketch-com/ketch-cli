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

var AppVersionBumpTypeLookup = map[string]AppVersionBumpType{
	"UNSPECIFIED_APP_VERSION_BUMP_TYPE": UnspecifiedAppVersionBumpType,
	"patch":                             PatchAppVersionBumpType,
	"minor":                             MinorAppVersionBumpType,
	"major":                             MajorAppVersionBumpType,
}

type AppVersionBumpType int32

const (
	UnspecifiedAppVersionBumpType AppVersionBumpType = 0
	PatchAppVersionBumpType       AppVersionBumpType = 1
	MinorAppVersionBumpType       AppVersionBumpType = 2
	MajorAppVersionBumpType       AppVersionBumpType = 3
)

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
	Type string `yaml:"type,omitempty" json:"type,omitempty"`
	// If true, the field will be required to have a value
	Required  bool   `yaml:"required,omitempty" json:"required,omitempty"`
	MinLength int64  `yaml:"minLength,omitempty" json:"minLength,omitempty"`
	MaxLength int64  `yaml:"maxLength,omitempty" json:"maxLength,omitempty"`
	Pattern   string `yaml:"pattern,omitempty" json:"pattern,omitempty"`
	// Default will be the default value for this field, before user interaction. Having a default value will override the placeholder text
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
	// The placeholder text that will appear when this field is empty
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`
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

type WorkflowOptions struct {
	TaskQueue                string `yaml:"taskQueue,omitempty" json:"taskQueue,omitempty"`
	WorkflowExecutionTimeout int64  `yaml:"workflowExecutionTimeout,omitempty" json:"workflowExecutionTimeout,omitempty"`
	WorkflowRunTimeout       int64  `yaml:"workflowRunTimeout,omitempty" json:"workflowRunTimeout,omitempty"`
	WorkflowTaskTimeout      int64  `yaml:"workflowTaskTimeout,omitempty" json:"workflowTaskTimeout,omitempty"`
	WaitForCancellation      bool   `yaml:"waitForCancellation,omitempty" json:"waitForCancellation,omitempty"`
	ScheduleToCloseTimeout   int64  `yaml:"scheduleToCloseTimeout,omitempty" json:"scheduleToCloseTimeout,omitempty"`
	ScheduleToStartTimeout   int64  `yaml:"scheduleToStartTimeout,omitempty" json:"scheduleToStartTimeout,omitempty"`
	StartToCloseTimeout      int64  `yaml:"startToCloseTimeout,omitempty" json:"startToCloseTimeout,omitempty"`
	HeartbeatTimeout         int64  `yaml:"heartbeatTimeout,omitempty" json:"heartbeatTimeout,omitempty"`
}

type WorkflowConfig struct {
	Code     string
	Name     string
	Icon     string
	Type     string
	Function string `yaml:"fn" json:"fn"`
	Options  *WorkflowOptions
}

type ManifestInputs struct {
	Code    string `yaml:"code,omitempty" json:"code,omitempty"`
	ID      string `yaml:"id,omitempty" json:"id,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	OrgCode string `yaml:"org,omitempty" json:"org,omitempty"`
	// deprecated - Version string - version assigned by commissary, cannot be dictated by manifest
	VersionBumpType     string              `yaml:"versionBumpType,omitempty" json:"versionBumpType,omitempty"`
	AutoUpgrade         bool                `yaml:"autoUpgrade,omitempty" json:"autoUpgrade,omitempty"`
	PrimaryCategory     string              `yaml:"primaryCategory,omitempty" json:"primaryCategory,omitempty"`
	SecondaryCategory   string              `yaml:"secondaryCategory,omitempty" json:"secondaryCategory,omitempty"`
	Rules               map[string]string   `yaml:"rules" json:"rules,omitempty"`
	Capabilities        []string            `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
	SupportedLanguages  []string            `yaml:"supportedLanguages,omitempty" json:"supportedLanguages,omitempty"`
	Purposes            []string            `yaml:"purposes,omitempty" json:"purposes,omitempty"`
	Rights              []string            `yaml:"rights,omitempty" json:"rights,omitempty"`
	IdentitySpaces      []*IdentitySpace    `yaml:"identitySpaces,omitempty" json:"identitySpaces,omitempty"`
	ShortDescription    string              `yaml:"shortDescription,omitempty" json:"shortDescription,omitempty"`
	DetailedDescription string              `yaml:"detailedDescription,omitempty" json:"detailedDescription,omitempty"`
	PermissionNote      string              `yaml:"permissionNote,omitempty" json:"permissionNote,omitempty"`
	Permissions         []string            `yaml:"permissions,omitempty" json:"permissions,omitempty"`
	SetupUrl            string              `yaml:"setupURL,omitempty" json:"setupURL,omitempty"`
	HomepageUrl         string              `yaml:"homepageURL,omitempty" json:"homepageURL,omitempty"`
	CustomerSupportUrl  string              `yaml:"customerSupportURL,omitempty" json:"customerSupportURL,omitempty"`
	PrivacyPolicyUrl    string              `yaml:"privacyPolicyURL,omitempty" json:"privacyPolicyURL,omitempty"`
	StatusUrl           string              `yaml:"statusURL,omitempty" json:"statusURL,omitempty"`
	TosUrl              string              `yaml:"tosURL,omitempty" json:"tosURL,omitempty"`
	DocUrl              string              `yaml:"docURL,omitempty" json:"docURL,omitempty"`
	Logo                *AppConfigImage     `yaml:"logo,omitempty" json:"logo,omitempty"`
	Previews            []*AppConfigImage   `yaml:"previews,omitempty" json:"previews,omitempty"`
	Contacts            []*AppConfigContact `yaml:"contacts,omitempty" json:"contacts,omitempty"`
	ExpireUserTokens    bool                `yaml:"expireUserTokens,omitempty" json:"expireUserTokens,omitempty"`
	RefreshInterval     string              `yaml:"refreshInterval,omitempty" json:"refreshInterval,omitempty"`
	RequestUserAuth     bool                `yaml:"requestUserAuth,omitempty" json:"requestUserAuth,omitempty"`
	UserAuthCallbackUrl string              `yaml:"userAuthCallbackURL,omitempty" json:"userAuthCallbackURL,omitempty"`
	RedirectOnUpdate    bool                `yaml:"redirectOnUpdate,omitempty" json:"redirectOnUpdate,omitempty"`
	Webhook             *Webhook            `yaml:"webhook,omitempty" json:"webhook,omitempty"`
	Workflow            []*WorkflowConfig   `yaml:"workflow,flow,omitempty" json:"workflow,omitempty"`
	Form                []*FormComponent    `yaml:"form,omitempty" json:"form,omitempty"`
}

type App struct {
	Code                string            `json:"code,omitempty"`
	ID                  string            `json:"id,omitempty"`
	OrgCode             string            `yaml:"orgCode" json:"orgCode,omitempty"`
	Name                string            `json:"name,omitempty"`
	Version             string            `json:"version,omitempty"`
	AutoUpgrade         bool              `json:"autoUpgrade,omitempty"`
	Readme              string            `json:"readme,omitempty"`
	HomepageUrl         string            `json:"homepageURL,omitempty"`
	UserAuthCallbackUrl string            `json:"userAuthCallbackURL,omitempty"`
	ExpireUserTokens    bool              `json:"expireUserTokens,omitempty"`
	RequestUserAuth     bool              `json:"requestUserAuth,omitempty"`
	SetupUrl            string            `json:"setupURL,omitempty"`
	RedirectOnUpdate    bool              `json:"redirectOnUpdate,omitempty"`
	WebhookId           string            `json:"webhookID,omitempty"`
	Capabilities        []int32           `yaml:",flow"`
	Permissions         []string          `yaml:",flow"`
	PermissionNote      string            `json:"permissionNode,omitempty"`
	Purposes            []string          `yaml:",flow"`
	Form                []*FormComponent  `yaml:",flow"`
	IdentitySpaces      []*IdentitySpace  `yaml:",flow" json:"identitySpaces,omitempty"`
	Rights              []string          `yaml:",flow" json:"rights,omitempty"`
	Rules               map[string]string `yaml:",flow" json:"rules,omitempty"`
	RefreshInterval     time.Duration     `yaml:"refreshInterval" json:"refreshInterval,omitempty"`
	EventTypes          []string          `json:"eventTypes,omitempty"`
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
	App             *App
	VersionBumpType AppVersionBumpType
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

	refreshInterval, err := time.ParseDuration(p.RefreshInterval)
	if err != nil {
		return nil, err
	}

	return &App{
		Code:                p.Code,
		ID:                  p.ID,
		OrgCode:             p.OrgCode,
		Name:                p.Name,
		AutoUpgrade:         p.AutoUpgrade,
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
		EventTypes:          p.Webhook.Events,
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
