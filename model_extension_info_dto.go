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

type ExtensionInfoDto struct {
	// The extension ID
	Id float64 `json:"id,omitempty"`
	// The extension unique slug
	Slug string `json:"slug,omitempty"`
	// The extension name
	Name string `json:"name,omitempty"`
	// The extension description
	Description string `json:"description,omitempty"`
	// Extension status
	Status string `json:"status,omitempty"`
}