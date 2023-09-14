// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageGphotosConfig storage gphotos config
//
// swagger:model storage.gphotosConfig
type StorageGphotosConfig struct {

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// OAuth Client Id.
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Also view and download archived media.
	IncludeArchived *bool `json:"includeArchived,omitempty"`

	// Set to make the Google Photos backend read only.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// Set to read the size of media items.
	ReadSize *bool `json:"readSize,omitempty"`

	// Year limits the photos to be downloaded to those which are uploaded after the given year.
	StartYear *int64 `json:"startYear,omitempty"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`
}

// Validate validates this storage gphotos config
func (m *StorageGphotosConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage gphotos config based on context it is used
func (m *StorageGphotosConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageGphotosConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGphotosConfig) UnmarshalBinary(b []byte) error {
	var res StorageGphotosConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
