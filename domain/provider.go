package domain

import (
	"net/http"
	"time"
)

const (
	// ProviderTypeBigQuery is the type name for BigQuery provider
	ProviderTypeBigQuery = "google_bigquery"
)

// RoleConfig is the configuration to define a role and mapping the permissions in the provider
type RoleConfig struct {
	ID          string        `json:"id" yaml:"id" validate:"required"`
	Name        string        `json:"name" yaml:"name" validate:"required"`
	Description string        `json:"description,omitempty" yaml:"description"`
	Permissions []interface{} `json:"permissions" yaml:"permissions" validate:"required"`
}

// PolicyConfig is the configuration that defines which policy is being used in the provider
type PolicyConfig struct {
	ID      string `json:"id" yaml:"id" validate:"required"`
	Version int    `json:"version" yaml:"version" validate:"required"`
}

// ResourceConfig is the configuration for a resource type within a provider
type ResourceConfig struct {
	Type   string        `json:"type" yaml:"type" validate:"required"`
	Policy *PolicyConfig `json:"policy" yaml:"policy"`
	Roles  []*RoleConfig `json:"roles" yaml:"roles" validate:"required"`
}

// AppealConfig is the policy configuration of the appeal
type AppealConfig struct {
	AllowPermanentAccess         bool   `json:"allow_permanent_access" yaml:"allow_permanent_access"`
	AllowActiveAccessExtensionIn string `json:"allow_active_access_extension_in" yaml:"allow_active_access_extension_in" validate:"required"`
}

// ProviderConfig is the configuration for a data provider
type ProviderConfig struct {
	Type        string                 `json:"type" yaml:"type" validate:"required,eq=google_bigquery"`
	URN         string                 `json:"urn" yaml:"urn" validate:"required"`
	Labels      map[string]interface{} `json:"labels" yaml:"labels"`
	Credentials interface{}            `json:"credentials,omitempty" yaml:"credentials" validate:"required"`
	Appeal      *AppealConfig          `json:"appeal" yaml:"appeal" validate:"required"`
	Resources   []*ResourceConfig      `json:"resources" yaml:"resources" validate:"required"`
}

// Provider domain structure
type Provider struct {
	ID        uint            `json:"id"`
	Type      string          `json:"type"`
	URN       string          `json:"urn"`
	Config    *ProviderConfig `json:"config"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// ProviderRepository interface
type ProviderRepository interface {
	Create(*Provider) error
	Update(*Provider) error
	Find() ([]*Provider, error)
	GetOne(uint) (*Provider, error)
	Delete(uint) error
}

// ProviderService interface
type ProviderService interface {
	Create(*Provider) error
	Find() ([]*Provider, error)
	Update(*Provider) error
	FetchResources() error
}

// ProviderHandler interface
type ProviderHandler interface {
	Create(http.ResponseWriter, *http.Request)
}

// ProviderInterface abstracts guardian communicates with external data providers
type ProviderInterface interface {
	GetType() string
	CreateConfig(*ProviderConfig) error
	GetResources(pc *ProviderConfig) ([]*Resource, error)
}
