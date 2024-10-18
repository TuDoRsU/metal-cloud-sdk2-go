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

type UpdateVmType struct {
	// Display name of the VM type
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Label of the VM type
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	// Tags for the VM Type. This is a JSON object.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	// Flag to indicate if the VM Type is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental float64 `json:"isExperimental,omitempty" yaml:"isExperimental,omitempty"`
	// Flag to indicate if the VM Type is for unmanaged VMs only. 1 for true, 0 for false. Default is 0.
	ForUnmanagedVMsOnly float64 `json:"forUnmanagedVMsOnly,omitempty" yaml:"forUnmanagedVMsOnly,omitempty"`
}
