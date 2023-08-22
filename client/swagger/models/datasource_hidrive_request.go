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

// DatasourceHidriveRequest datasource hidrive request
//
// swagger:model datasource.HidriveRequest
type DatasourceHidriveRequest struct {

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Chunksize for chunked uploads.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// OAuth Client Id.
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Do not fetch number of objects in directories unless it is absolutely necessary.
	DisableFetchingMemberCount *string `json:"disableFetchingMemberCount,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for the service.
	Endpoint *string `json:"endpoint,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// The root/parent folder for all paths.
	RootPrefix *string `json:"rootPrefix,omitempty"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// Access permissions that rclone should use when requesting access from HiDrive.
	ScopeAccess *string `json:"scopeAccess,omitempty"`

	// User-level that rclone should use when requesting access from HiDrive.
	ScopeRole *string `json:"scopeRole,omitempty"`

	// The path of the source to scan files
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`

	// Concurrency for chunked uploads.
	UploadConcurrency *string `json:"uploadConcurrency,omitempty"`

	// Cutoff/Threshold for chunked uploads.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`
}

// Validate validates this datasource hidrive request
func (m *DatasourceHidriveRequest) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceHidriveRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceHidriveRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceHidriveRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceHidriveRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource hidrive request based on the context it is used
func (m *DatasourceHidriveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceHidriveRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceHidriveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceHidriveRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceHidriveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
