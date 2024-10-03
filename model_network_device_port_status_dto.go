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

type NetworkDevicePortStatusDto struct {
	// The ports of the network device that will have their status changed
	Ports []string `json:"ports"`
	// The new status of the ports
	Status bool `json:"status"`
}