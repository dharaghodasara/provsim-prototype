package core

import (
	"github.com/google/uuid"
)

//ProviderCategory is an enum
type ProviderCategory string

//CapabilityType is an enum
type CapabilityType string

const (
	Service  CapabilityType = "service"
	Security CapabilityType = "security"
)

const (
	Banking   ProviderCategory = "banking"
	Tax       ProviderCategory = "tax"
	CRM       ProviderCategory = "crm"
	eCommerce ProviderCategory = "ecommerce"
	Payment   ProviderCategory = "payment"
	Rideshare ProviderCategory = "ridesharee"
)

//Provider ...
type Provider struct {
	ID               uuid.UUID        `yaml:"id"`
	Name             string           `yaml:"name"`
	BaseURL          string           `yaml:"base-url"`
	LoginURL         string           `yaml:"login-url"`
	ProviderCategory ProviderCategory `yaml:"provider-category"`
	MetaData         Metadata         `yaml:"meta-data"`
	Capabilities     []Capability     `yaml:"capabilities"`
}

// Metadata ...
type Metadata struct {
	//BaseProvider ...
	BaseProvider BaseProvider `yaml:"base-provider"`
}

//BaseProvider Type ...
type BaseProvider struct {
	ID   uuid.UUID `yaml:"id"`
	Name string    `yaml:"name"`
}

//Capability ...
type Capability struct {
	Name           string         `yaml:"name"`
	CapabilityType CapabilityType `yaml:"type"`
	//Attributes []string       `yaml:", attributes"`
}
