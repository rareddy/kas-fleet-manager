/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ConnectorDesiredState the model 'ConnectorDesiredState'
type ConnectorDesiredState string

// List of ConnectorDesiredState
const (
	CONNECTORDESIREDSTATE_UNASSIGNED ConnectorDesiredState = "unassigned"
	CONNECTORDESIREDSTATE_READY      ConnectorDesiredState = "ready"
	CONNECTORDESIREDSTATE_STOPPED    ConnectorDesiredState = "stopped"
	CONNECTORDESIREDSTATE_DELETED    ConnectorDesiredState = "deleted"
)
