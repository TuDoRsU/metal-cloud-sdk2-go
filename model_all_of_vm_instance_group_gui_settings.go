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

// GUI settings for the VM Instance Group. This is a JSON object.
type AllOfVmInstanceGroupGuiSettings struct {
	// Row index of the VM Instance Group.
	RowIndex float64 `json:"rowIndex" yaml:"rowIndex"`
	// Column index of the VM Instance Group.
	ColumnIndex float64 `json:"columnIndex" yaml:"columnIndex"`
	// Whether to show the VM Instance Group children in the GUI.
	ShowWidgetChildren bool `json:"showWidgetChildren" yaml:"showWidgetChildren"`
	// Random instance ID.
	RandomInstanceID string `json:"randomInstanceID" yaml:"randomInstanceID"`
	// User agent.
	UserAgent string `json:"userAgent" yaml:"userAgent"`
}
