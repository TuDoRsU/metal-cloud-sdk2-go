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

type FileShareHostBulkOperation struct {
	// Id of the Instance Array Host that will be modified
	InstanceArrayId float64 `json:"instanceArrayId" yaml:"instanceArrayId"`
	// Operation type for the Instance Array Host
	OperationType string `json:"operationType" yaml:"operationType"`
}
