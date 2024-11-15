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

type UpdateVmInstanceGroup struct {
	// Label for the VM Instance Group.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	// Tags for the VM Instance Group.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	GuiSettings *VmInstanceGroupGuiSettingsDto `json:"guiSettings,omitempty" yaml:"guiSettings,omitempty"`
	// Interfaces for the VM Instance Group
	VmInstanceGroupInterfaces []UpdateVmInstanceGroupInterface `json:"vmInstanceGroupInterfaces,omitempty" yaml:"vmInstanceGroupInterfaces,omitempty"`
	// Custom variables for the VM Instance.
	CustomVariables *interface{} `json:"customVariables,omitempty" yaml:"customVariables,omitempty"`
}
