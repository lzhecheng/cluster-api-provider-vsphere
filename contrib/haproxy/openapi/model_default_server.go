/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DefaultServer struct for DefaultServer
type DefaultServer struct {
	Fall  *int32 `json:"fall,omitempty"`
	Inter *int32 `json:"inter,omitempty"`
	Port  *int32 `json:"port,omitempty"`
	Rise  *int32 `json:"rise,omitempty"`
}
