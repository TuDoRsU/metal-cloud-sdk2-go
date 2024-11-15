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

type ServerRegistrationDto struct {
	// The BMC hostname of the server.
	BmcHostname string `json:"bmcHostname" yaml:"bmcHostname"`
	// The site id where the server is located.
	SiteId float64 `json:"siteId" yaml:"siteId"`
	// The BMC username to use.
	BmcUser string `json:"bmcUser,omitempty" yaml:"bmcUser,omitempty"`
	// The BMC password to use.
	BmcPassword string `json:"bmcPassword,omitempty" yaml:"bmcPassword,omitempty"`
	// The vendor of the server.
	Vendor string `json:"vendor,omitempty" yaml:"vendor,omitempty"`
	// The model of the server.
	Model string `json:"model,omitempty" yaml:"model,omitempty"`
	// The MAC address of the server.
	MacAddress string `json:"macAddress,omitempty" yaml:"macAddress,omitempty"`
	// The UUID of the server.
	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	// The Serial Number of the server.
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`
}
