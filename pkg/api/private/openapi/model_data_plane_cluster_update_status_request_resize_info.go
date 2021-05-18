/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DataPlaneClusterUpdateStatusRequestResizeInfo struct for DataPlaneClusterUpdateStatusRequestResizeInfo
type DataPlaneClusterUpdateStatusRequestResizeInfo struct {
	NodeDelta *int32                                              `json:"nodeDelta,omitempty"`
	Delta     *DataPlaneClusterUpdateStatusRequestResizeInfoDelta `json:"delta,omitempty"`
}
