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

type ServerRegistrationResponseDto struct {
	// The id of the server.
	Id float64 `json:"id" yaml:"id"`
	// The Serial Number of the server.
	SerialNumber string `json:"serialNumber" yaml:"serialNumber"`
	// The UUID of the server.
	Uuid string `json:"uuid" yaml:"uuid"`
	// The job info of the server.
	JobInfo *AllOfServerRegistrationResponseDtoJobInfo `json:"jobInfo,omitempty" yaml:"jobInfo,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty" yaml:"links,omitempty"`
}
