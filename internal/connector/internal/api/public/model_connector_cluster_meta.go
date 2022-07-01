/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

import (
	"time"
)

// ConnectorClusterMeta struct for ConnectorClusterMeta
type ConnectorClusterMeta struct {
	Owner      string    `json:"owner,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	Name       string    `json:"name,omitempty"`
}
