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

type CreateAccountDto struct {
	// The ID of the parent account
	ParentAccountId float64 `json:"parentAccountId,omitempty" yaml:"parentAccountId,omitempty"`
	// The name of the account
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// The code of the account
	Code string `json:"code,omitempty" yaml:"code,omitempty"`
	// The fiscal number of the account
	FiscalNumber string `json:"fiscalNumber,omitempty" yaml:"fiscalNumber,omitempty"`
	// The address of the account
	Address *AllOfCreateAccountDtoAddress `json:"address,omitempty" yaml:"address,omitempty"`
}
