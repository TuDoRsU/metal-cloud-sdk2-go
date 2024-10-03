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

type VmInstanceGroupInterfaceDto struct {
	// Interface ID
	Id float64 `json:"id"`
	// Network ID
	NetworkId float64 `json:"networkId"`
	// Interface index
	InterfaceIndex float64 `json:"interfaceIndex"`
	// Service status of the VM Instance Group Interface.
	ServiceStatus string `json:"serviceStatus"`
	// VM Instance Group ID
	GroupId float64 `json:"groupId"`
	// Infrastructure ID
	InfrastructureId float64 `json:"infrastructureId"`
	// Operation ID
	ChangeId float64 `json:"changeId"`
	// Interface label
	Label float64 `json:"label"`
	// Timestamp of the VM Instance Group Interface creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Timestamp of the VM Instance Group Interface update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Operation object for the VM Instance Group.
	Operation *interface{} `json:"operation"`
}
