package presenters

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/api/admin/private"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/api/public"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/db"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"strings"
	"time"
)

var AllNamespaceTenantKinds = map[string]public.ConnectorNamespaceTenantKind{
	string(public.CONNECTORNAMESPACETENANTKIND_USER):         public.CONNECTORNAMESPACETENANTKIND_USER,
	string(public.CONNECTORNAMESPACETENANTKIND_ORGANISATION): public.CONNECTORNAMESPACETENANTKIND_ORGANISATION,
}

func ConvertConnectorNamespaceRequest(namespaceRequest *public.ConnectorNamespaceRequest,
	userID string, organisationID string) (*dbapi.ConnectorNamespace, *errors.ServiceError) {
	result := &dbapi.ConnectorNamespace{
		Model: db.Model{
			ID: api.NewID(),
		},
		Name:      namespaceRequest.Name,
		ClusterId: namespaceRequest.ClusterId,
		Owner:     userID,
	}
	result.Annotations = make([]dbapi.ConnectorNamespaceAnnotation, len(namespaceRequest.Annotations))
	for i, annotation := range namespaceRequest.Annotations {
		result.Annotations[i].NamespaceId = result.ID
		result.Annotations[i].Name = annotation.Name
		result.Annotations[i].Value = annotation.Value
	}
	switch namespaceRequest.Kind {
	case public.CONNECTORNAMESPACETENANTKIND_USER:
		result.TenantUserId = &userID
		result.TenantUser = &dbapi.ConnectorTenantUser{
			Model: db.Model{
				ID: userID,
			},
		}
	case public.CONNECTORNAMESPACETENANTKIND_ORGANISATION:
		if organisationID == "" {
			return nil, errors.BadRequest("missing organization for tenant organisation namespace")
		}
		result.TenantOrganisationId = &organisationID
		result.TenantOrganisation = &dbapi.ConnectorTenantOrganisation{
			Model: db.Model{
				ID: organisationID,
			},
		}
	default:
		return nil, errors.BadRequest("invalid tenant kind: %s", namespaceRequest.Kind)
	}

	return result, nil
}

func ConvertConnectorNamespaceEvalRequest(namespaceRequest *public.ConnectorNamespaceEvalRequest, userID string) *dbapi.ConnectorNamespace {
	result := &dbapi.ConnectorNamespace{
		Model: db.Model{
			ID: api.NewID(),
		},
		Name:  namespaceRequest.Name,
		Owner: userID,
	}
	result.Annotations = make([]dbapi.ConnectorNamespaceAnnotation, len(namespaceRequest.Annotations))
	for i, annotation := range namespaceRequest.Annotations {
		result.Annotations[i].NamespaceId = result.ID
		result.Annotations[i].Name = annotation.Name
		result.Annotations[i].Value = annotation.Value
	}
	result.TenantUserId = &result.Owner
	result.TenantUser = &dbapi.ConnectorTenantUser{
		Model: db.Model{
			ID: userID,
		},
	}

	return result
}

func ConvertConnectorNamespaceWithTenantRequest(namespaceRequest *private.ConnectorNamespaceWithTenantRequest) (*dbapi.ConnectorNamespace, *errors.ServiceError) {
	result := &dbapi.ConnectorNamespace{
		Model: db.Model{
			ID: api.NewID(),
		},
		Name:      namespaceRequest.Name,
		ClusterId: namespaceRequest.ClusterId,
	}
	switch namespaceRequest.Tenant.Kind {
	case private.USER:
		result.TenantUserId = &namespaceRequest.Tenant.Id
		result.TenantUser = &dbapi.ConnectorTenantUser{
			Model: db.Model{
				ID: *result.TenantUserId,
			},
		}
	case private.ORGANISATION:
		result.TenantOrganisationId = &namespaceRequest.Tenant.Id
		result.TenantOrganisation = &dbapi.ConnectorTenantOrganisation{
			Model: db.Model{
				ID: *result.TenantOrganisationId,
			},
		}
	default:
		return nil, errors.BadRequest("invalid kind %s", namespaceRequest.Tenant.Kind)
	}
	result.Annotations = make([]dbapi.ConnectorNamespaceAnnotation, len(namespaceRequest.Annotations))
	for i, annotation := range namespaceRequest.Annotations {
		result.Annotations[i].NamespaceId = result.ID
		result.Annotations[i].Name = annotation.Name
		result.Annotations[i].Value = annotation.Value
	}

	return result, nil
}

func PresentConnectorNamespace(namespace *dbapi.ConnectorNamespace) public.ConnectorNamespace {
	annotations := make([]public.ConnectorNamespaceRequestMetaAnnotations, len(namespace.Annotations))
	for i, anno := range namespace.Annotations {
		annotations[i].Name = anno.Name
		annotations[i].Value = anno.Value
	}

	reference := PresentReference(namespace.ID, namespace)
	result := public.ConnectorNamespace{
		Id:   reference.Id,
		Kind: reference.Kind,
		Href: reference.Href,

		CreatedAt:  namespace.CreatedAt,
		ModifiedAt: namespace.UpdatedAt,
		Owner:      namespace.Owner,
		Version:    namespace.Version,

		Name:        namespace.Name,
		ClusterId:   namespace.ClusterId,
		Tenant:      public.ConnectorNamespaceTenant{},
		Annotations: annotations,
	}
	if namespace.TenantUser != nil {
		result.Tenant.Kind = public.CONNECTORNAMESPACETENANTKIND_USER
		result.Tenant.Id = namespace.TenantUser.ID
	}
	if namespace.TenantOrganisation != nil {
		result.Tenant.Kind = public.CONNECTORNAMESPACETENANTKIND_ORGANISATION
		result.Tenant.Id = namespace.TenantOrganisation.ID
	}
	if namespace.Expiration != nil {
		result.Expiration = getTimestamp(*namespace.Expiration)
	}

	return result
}

func PresentPrivateConnectorNamespace(namespace *dbapi.ConnectorNamespace) private.ConnectorNamespace {
	annotations := make([]private.ConnectorNamespaceRequestMetaAnnotations, len(namespace.Annotations))
	for i, anno := range namespace.Annotations {
		annotations[i].Name = anno.Name
		annotations[i].Value = anno.Value
	}

	reference := PresentReference(namespace.ID, namespace)
	result := private.ConnectorNamespace{
		Id:   reference.Id,
		Kind: reference.Kind,
		Href: reference.Href,

		CreatedAt:  namespace.CreatedAt,
		ModifiedAt: namespace.UpdatedAt,
		Owner:      namespace.Owner,
		Version:    namespace.Version,

		Name:        namespace.Name,
		ClusterId:   namespace.ClusterId,
		Tenant:      private.ConnectorNamespaceTenant{},
		Annotations: annotations,
	}
	if namespace.TenantUser != nil {
		result.Tenant.Kind = private.USER
		result.Tenant.Id = namespace.TenantUser.ID
	}
	if namespace.TenantOrganisationId != nil {
		result.Tenant.Kind = private.ORGANISATION
		result.Tenant.Id = namespace.TenantOrganisation.ID
	}
	if namespace.Expiration != nil {
		result.Expiration = getTimestamp(*namespace.Expiration)
	}

	return result
}

func getTimestamp(expiration time.Time) string {
	bytes, _ := json.Marshal(expiration)
	return strings.Trim(string(bytes), "\"")
}