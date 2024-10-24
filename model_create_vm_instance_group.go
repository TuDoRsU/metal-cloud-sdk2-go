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

type CreateVmInstanceGroup struct {
	// Number of VM instances in the VM Instance Group.
	InstanceCount float64 `json:"instanceCount" yaml:"instanceCount"`
	// Id of the VM Type.
	TypeId float64 `json:"typeId" yaml:"typeId"`
	// Disk size in GB for each VM Instance in the VM Instance Group.
	DiskSizeGB float64 `json:"diskSizeGB" yaml:"diskSizeGB"`
	// Id of the VM Volume Template.
	VolumeTemplateId float64 `json:"volumeTemplateId" yaml:"volumeTemplateId"`
	// Tags for the VM Instance Group.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}