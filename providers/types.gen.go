// Package providers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package providers

// Defines values for AuthType.
const (
	Oauth2 AuthType = "oauth2"
)

// AuthType defines model for AuthType.
type AuthType string

// CatalogType defines model for CatalogType.
type CatalogType map[string]ProviderInfo

// OauthOpts defines model for OauthOpts.
type OauthOpts struct {
	AuthURL  string `json:"authURL"`
	TokenURL string `json:"tokenURL"`
}

// Provider defines model for Provider.
type Provider = string

// ProviderInfo defines model for ProviderInfo.
type ProviderInfo struct {
	AuthType     AuthType     `json:"authType"`
	BaseURL      string       `json:"baseURL"`
	OauthOpts    OauthOpts    `json:"oauthOpts"`
	ProviderOpts ProviderOpts `json:"providerOpts"`
	Support      Support      `json:"support"`
}

// ProviderOpts defines model for ProviderOpts.
type ProviderOpts map[string]string

// Support defines model for Support.
type Support struct {
	BulkWrite bool `json:"bulkWrite"`
	Proxy     bool `json:"proxy"`
	Read      bool `json:"read"`
	Subscribe bool `json:"subscribe"`
	Write     bool `json:"write"`
}
