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

// Shutdown options
type AllOfInfrastructureDeployOptionsShutdownOptions struct {
	// Attempt soft shutdown
	AttemptSoftShutdown bool `json:"attemptSoftShutdown" yaml:"attemptSoftShutdown"`
	// Soft shutdown timeout in seconds
	SoftShutdownTimeout float64 `json:"softShutdownTimeout" yaml:"softShutdownTimeout"`
	// Force shutdown
	ForceShutdown bool `json:"forceShutdown" yaml:"forceShutdown"`
}
