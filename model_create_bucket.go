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

type CreateBucket struct {
	// Label of the Bucket.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	// Disk size in GB for Bucket
	SizeGB float64 `json:"sizeGB" yaml:"sizeGB"`
}