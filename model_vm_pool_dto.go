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

type VmPoolDto struct {
	// VM Pool ID
	Id float64 `json:"id"`
	// Id of the site for the VM
	SiteId float64 `json:"siteId"`
	// Datacenter of the VM
	DatacenterName string `json:"datacenterName"`
	// Host of the VM Pool
	ManagementHost string `json:"managementHost"`
	// Name of the VM Pool
	Name string `json:"name"`
	// Description of the VM Pool
	Description string `json:"description"`
	// Type of the VM Pool
	Type_ string `json:"type"`
	// Certificate of the VM Pool
	Certificate string `json:"certificate"`
	// Status of the VM Pool
	Status string `json:"status"`
	// External identifier of the VM Pool
	ExternalIdentifier string `json:"externalIdentifier"`
	// Number of total RAM GB in the VM Pool
	TotalRamGB float64 `json:"totalRamGB"`
	// Number of used RAM GB in the VM Pool
	UsedRamGB float64 `json:"usedRamGB"`
	// Number of free RAM GB in the VM Pool
	FreeRamGB float64 `json:"freeRamGB"`
	// Number of total disk space GB in the VM Pool
	TotalSpaceGB float64 `json:"totalSpaceGB"`
	// Number of used disk space GB in the VM Pool
	UsedSpaceGB float64 `json:"usedSpaceGB"`
	// Number of free disk space GB in the VM Pool
	FreeSpaceGB float64 `json:"freeSpaceGB"`
	// Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0.
	InMaintenance float64 `json:"inMaintenance"`
	// Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental float64 `json:"isExperimental"`
	// Timestamp when the VM Pool was created
	CreatedTimestamp string `json:"createdTimestamp"`
	// Timestamp when the VM Pool was updated
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Tags for the VM Pool.
	Tags []string `json:"tags"`
}
