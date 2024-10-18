/*
 * MetalSoft REST API
 *
 * MetalSoft REST API documentation
 *
 * API version: 2.0
 * Contact: support@metalsoft.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk2

// The address of the account
type AllOfAccountDtoAddress struct {
	// The first line of the account address
	AddressLine1 string `json:"addressLine1,omitempty" yaml:"addressLine1,omitempty"`
	// The second line of the account address
	AddressLine2 string `json:"addressLine2,omitempty" yaml:"addressLine2,omitempty"`
	// The postal code of the account address
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`
	// The state of the account address
	State string `json:"state,omitempty" yaml:"state,omitempty"`
	// The country of the account address
	Country string `json:"country,omitempty" yaml:"country,omitempty"`
}
