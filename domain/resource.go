package domain

import "time"

// Resource struct
type Resource struct {
	ID           string                 `json:"id" yaml:"id"`
	ProviderType string                 `json:"provider_type" yaml:"provider_type"`
	ProviderURN  string                 `json:"provider_urn" yaml:"provider_urn"`
	Type         string                 `json:"type" yaml:"type"`
	URN          string                 `json:"urn" yaml:"urn"`
	Name         string                 `json:"name" yaml:"name"`
	Details      map[string]interface{} `json:"details" yaml:"details"`
	Labels       map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	IsDeleted    bool                   `json:"is_deleted,omitempty" yaml:"is_deleted,omitempty"`
}
