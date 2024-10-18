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

type VmInstanceGroup struct {
	// Id of the VM Instance Group.
	Id float64 `json:"id" yaml:"id"`
	// Id of the Infrastructure.
	InfrastructureId float64 `json:"infrastructureId" yaml:"infrastructureId"`
	// Status of the VM Instance Group.
	ServiceStatus string `json:"serviceStatus" yaml:"serviceStatus"`
	// Id of the VM Instance Group change object.
	ChangeId float64 `json:"changeId" yaml:"changeId"`
	// Name of the VM Instance Group.
	Label string `json:"label" yaml:"label"`
	// Number of VM instances in the VM Instance Group.
	InstanceCount float64 `json:"instanceCount" yaml:"instanceCount"`
	// Timestamp of the VM Instance Group creation.
	CreatedTimestamp string `json:"createdTimestamp" yaml:"createdTimestamp"`
	// Timestamp of the VM Instance Group update.
	UpdatedTimestamp string `json:"updatedTimestamp" yaml:"updatedTimestamp"`
	// Operation object for the VM Instance Group.
	Operation *interface{} `json:"operation" yaml:"operation"`
	// Array of VM instances in the VM Instance Group.
	VmInstance []string `json:"vmInstance" yaml:"vmInstance"`
	// Tags for the VM Instance Group.
	Tags []string `json:"tags" yaml:"tags"`
	// Disk size in GB for each VM Instance in the VM Instance Group.
	DiskSizeGB float64 `json:"diskSizeGB" yaml:"diskSizeGB"`
	// Id of the template used by the VM Instance Group.
	VolumeTemplateId float64 `json:"volumeTemplateId,omitempty" yaml:"volumeTemplateId,omitempty"`
	// Network Id to Network Profile Id for the VM Instance Group. This is a JSON object.
	NetworkIdToNetworkProfileId *interface{} `json:"networkIdToNetworkProfileId,omitempty" yaml:"networkIdToNetworkProfileId,omitempty"`
	// GUI settings for the VM Instance Group. This is a JSON object.
	GuiSettings *AllOfVmInstanceGroupGuiSettings `json:"guiSettings,omitempty" yaml:"guiSettings,omitempty"`
	Links *interface{} `json:"links" yaml:"links"`
}
