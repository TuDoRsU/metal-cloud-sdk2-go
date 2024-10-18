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

type CreateVmInstance struct {
	// Id of the VM Instance Group.
	GroupId float64 `json:"groupId" yaml:"groupId"`
	// Id of the VM Type.
	TypeId float64 `json:"typeId" yaml:"typeId"`
	// Tags for the VM Instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
