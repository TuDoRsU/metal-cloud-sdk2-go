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

// Resource Pool statistics for users, servers, and subnet pools
type AllOfResourcePoolWithStatsDtoStatistics struct {
	Users float64 `json:"users" yaml:"users"`
	Servers float64 `json:"servers" yaml:"servers"`
	SubnetPools float64 `json:"subnetPools" yaml:"subnetPools"`
}
