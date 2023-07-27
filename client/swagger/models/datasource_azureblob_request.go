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

// DatasourceAzureblobRequest datasource azureblob request
//
// swagger:model datasource.AzureblobRequest
type DatasourceAzureblobRequest struct {

	// Access tier of blob: hot, cool or archive.
	AccessTier string `json:"accessTier,omitempty"`

	// Azure Storage Account Name.
	Account string `json:"account,omitempty"`

	// Delete archive tier blobs before overwriting.
	ArchiveTierDelete *string `json:"archiveTierDelete,omitempty"`

	// Upload chunk size.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Password for the certificate file (optional).
	ClientCertificatePassword string `json:"clientCertificatePassword,omitempty"`

	// Path to a PEM or PKCS12 certificate file including the private key.
	ClientCertificatePath string `json:"clientCertificatePath,omitempty"`

	// The ID of the client in use.
	ClientID string `json:"clientId,omitempty"`

	// One of the service principal's client secrets
	ClientSecret string `json:"clientSecret,omitempty"`

	// Send the certificate chain when using certificate auth.
	ClientSendCertificateChain *string `json:"clientSendCertificateChain,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Don't store MD5 checksum with object metadata.
	DisableChecksum *string `json:"disableChecksum,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for the service.
	Endpoint string `json:"endpoint,omitempty"`

	// Read credentials from runtime (environment variables, CLI or MSI).
	EnvAuth *string `json:"envAuth,omitempty"`

	// Storage Account Shared Key.
	Key string `json:"key,omitempty"`

	// Size of blob list.
	ListChunk *string `json:"listChunk,omitempty"`

	// How often internal memory buffer pools will be flushed.
	MemoryPoolFlushTime *string `json:"memoryPoolFlushTime,omitempty"`

	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap *string `json:"memoryPoolUseMmap,omitempty"`

	// Object ID of the user-assigned MSI to use, if any.
	MsiClientID string `json:"msiClientId,omitempty"`

	// Azure resource ID of the user-assigned MSI to use, if any.
	MsiMiResID string `json:"msiMiResId,omitempty"`

	// Object ID of the user-assigned MSI to use, if any.
	MsiObjectID string `json:"msiObjectId,omitempty"`

	// If set, don't attempt to check the container exists or create it.
	NoCheckContainer *string `json:"noCheckContainer,omitempty"`

	// If set, do not do HEAD before GET when getting objects.
	NoHeadObject *string `json:"noHeadObject,omitempty"`

	// The user's password
	Password string `json:"password,omitempty"`

	// Public access level of a container: blob or container.
	PublicAccess string `json:"publicAccess,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// SAS URL for container level access only.
	SasURL string `json:"sasUrl,omitempty"`

	// Path to file containing credentials for use with a service principal.
	ServicePrincipalFile string `json:"servicePrincipalFile,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// ID of the service principal's tenant. Also called its directory ID.
	Tenant string `json:"tenant,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *string `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	UploadCutoff string `json:"uploadCutoff,omitempty"`

	// Uses local storage emulator if provided as 'true'.
	UseEmulator *string `json:"useEmulator,omitempty"`

	// Use a managed service identity to authenticate (only works in Azure).
	UseMsi *string `json:"useMsi,omitempty"`

	// User name (usually an email address)
	Username string `json:"username,omitempty"`
}

// Validate validates this datasource azureblob request
func (m *DatasourceAzureblobRequest) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceAzureblobRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceAzureblobRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceAzureblobRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasource azureblob request based on context it is used
func (m *DatasourceAzureblobRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceAzureblobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceAzureblobRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceAzureblobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
