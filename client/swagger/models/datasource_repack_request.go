// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasourceRepackRequest datasource repack request
//
// swagger:model datasource.RepackRequest
type DatasourceRepackRequest struct {

	// pack job Id
	PackJobID int64 `json:"packJobId,omitempty"`
}

// Validate validates this datasource repack request
func (m *DatasourceRepackRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datasource repack request based on context it is used
func (m *DatasourceRepackRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceRepackRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceRepackRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceRepackRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}