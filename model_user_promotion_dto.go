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

type UserPromotionDto struct {
	// The promotion for the user
	Promotion string `json:"promotion,omitempty" yaml:"promotion,omitempty"`
	// Whether to remove the promotion
	RemovePromotion bool `json:"removePromotion,omitempty" yaml:"removePromotion,omitempty"`
}
