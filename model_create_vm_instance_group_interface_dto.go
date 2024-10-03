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

type CreateVmInstanceGroupInterfaceDto struct {
	// Network ID for the interface. Can be null if the interface is not connected to any network
	NetworkId float64 `json:"networkId,omitempty"`
}
