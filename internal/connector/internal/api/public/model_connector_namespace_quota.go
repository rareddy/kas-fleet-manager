/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ConnectorNamespaceQuota struct for ConnectorNamespaceQuota
type ConnectorNamespaceQuota struct {
	Connectors int32 `json:"connectors,omitempty"`
	// Memory quota for limits or requests
	MemoryRequests string `json:"memory_requests,omitempty"`
	// Memory quota for limits or requests
	MemoryLimits string `json:"memory_limits,omitempty"`
	// CPU quota for limits or requests
	CpuRequests string `json:"cpu_requests,omitempty"`
	// CPU quota for limits or requests
	CpuLimits string `json:"cpu_limits,omitempty"`
}
