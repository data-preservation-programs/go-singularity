// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourceSwiftRequest datasource swift request
//
// swagger:model datasource.SwiftRequest
type DatasourceSwiftRequest struct {

	// Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	ApplicationCredentialID string `json:"applicationCredentialId,omitempty"`

	// Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	ApplicationCredentialName string `json:"applicationCredentialName,omitempty"`

	// Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	ApplicationCredentialSecret string `json:"applicationCredentialSecret,omitempty"`

	// Authentication URL for server (OS_AUTH_URL).
	Auth string `json:"auth,omitempty"`

	// Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	AuthToken string `json:"authToken,omitempty"`

	// AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	AuthVersion *string `json:"authVersion,omitempty"`

	// Above this size files will be chunked into a _segments container.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	Domain string `json:"domain,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	EndpointType *string `json:"endpointType,omitempty"`

	// Get swift credentials from environment variables in standard OpenStack form.
	EnvAuth *string `json:"envAuth,omitempty"`

	// API key or password (OS_PASSWORD).
	Key string `json:"key,omitempty"`

	// If true avoid calling abort upload on a failure.
	LeavePartsOnError *string `json:"leavePartsOnError,omitempty"`

	// Don't chunk files during streaming upload.
	NoChunk *string `json:"noChunk,omitempty"`

	// Disable support for static and dynamic large objects
	NoLargeObjects *string `json:"noLargeObjects,omitempty"`

	// Region name - optional (OS_REGION_NAME).
	Region string `json:"region,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// The storage policy to use when creating a new container.
	StoragePolicy string `json:"storagePolicy,omitempty"`

	// Storage URL - optional (OS_STORAGE_URL).
	StorageURL string `json:"storageUrl,omitempty"`

	// Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	Tenant string `json:"tenant,omitempty"`

	// Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	TenantDomain string `json:"tenantDomain,omitempty"`

	// Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	TenantID string `json:"tenantId,omitempty"`

	// User name to log in (OS_USERNAME).
	User string `json:"user,omitempty"`

	// User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	UserID string `json:"userId,omitempty"`
}

// Validate validates this datasource swift request
func (m *DatasourceSwiftRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceSwiftRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSwiftRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSwiftRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasource swift request based on context it is used
func (m *DatasourceSwiftRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceSwiftRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceSwiftRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceSwiftRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
