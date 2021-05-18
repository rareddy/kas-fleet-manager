/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectorClusterStatusOperators struct for ConnectorClusterStatusOperators
type ConnectorClusterStatusOperators struct {
	// the id of the operator
	Id string `json:"id,omitempty"`
	// the version of the operator
	Version string `json:"version,omitempty"`
	// the namespace to which the operator has been installed
	Namespace string `json:"namespace,omitempty"`
	// the status of the operator
	Status string `json:"status,omitempty"`
}
