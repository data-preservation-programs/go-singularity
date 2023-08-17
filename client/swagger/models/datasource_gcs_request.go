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

// DatasourceGcsRequest datasource gcs request
//
// swagger:model datasource.GcsRequest
type DatasourceGcsRequest struct {

	// Access public buckets and objects without credentials.
	Anonymous *string `json:"anonymous,omitempty"`

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Access Control List for new buckets.
	BucketACL string `json:"bucketAcl,omitempty"`

	// Access checks should use bucket-level IAM policies.
	BucketPolicyOnly *string `json:"bucketPolicyOnly,omitempty"`

	// OAuth Client Id.
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// If set this will decompress gzip encoded objects.
	Decompress *string `json:"decompress,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for the service.
	Endpoint string `json:"endpoint,omitempty"`

	// Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
	EnvAuth *string `json:"envAuth,omitempty"`

	// Location for the newly created buckets.
	Location string `json:"location,omitempty"`

	// If set, don't attempt to check the bucket exists or create it.
	NoCheckBucket *string `json:"noCheckBucket,omitempty"`

	// Access Control List for new objects.
	ObjectACL string `json:"objectAcl,omitempty"`

	// Project number.
	ProjectNumber string `json:"projectNumber,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// Service Account Credentials JSON blob.
	ServiceAccountCredentials string `json:"serviceAccountCredentials,omitempty"`

	// Service Account Credentials JSON file path.
	ServiceAccountFile string `json:"serviceAccountFile,omitempty"`

	// The path of the source to scan files
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// The storage class to use when storing objects in Google Cloud Storage.
	StorageClass string `json:"storageClass,omitempty"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`
}

// Validate validates this datasource gcs request
func (m *DatasourceGcsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanningState(formats); err != nil {
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

func (m *DatasourceGcsRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceGcsRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceGcsRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceGcsRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource gcs request based on the context it is used
func (m *DatasourceGcsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceGcsRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceGcsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceGcsRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceGcsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
