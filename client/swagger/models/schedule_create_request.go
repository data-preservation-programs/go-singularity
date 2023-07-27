// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScheduleCreateRequest schedule create request
//
// swagger:model schedule.CreateRequest
type ScheduleCreateRequest struct {

	// Allowed piece CIDs in this schedule
	AllowedPieceCids []string `json:"allowedPieceCids"`

	// Dataset name
	DatasetName string `json:"datasetName,omitempty"`

	// Duration in epoch or in duration format, i.e. 1500000, 2400h
	Duration *string `json:"duration,omitempty"`

	// http headers to be passed with the request (i.e. key=value)
	HTTPHeaders []string `json:"httpHeaders"`

	// Whether the deal should be IPNI
	Ipni *bool `json:"ipni,omitempty"`

	// Whether the deal should be kept unsealed
	KeepUnsealed *bool `json:"keepUnsealed,omitempty"`

	// Max pending deal number
	MaxPendingDealNumber int64 `json:"maxPendingDealNumber,omitempty"`

	// Max pending deal size in human readable format, i.e. 100 TiB
	MaxPendingDealSize string `json:"maxPendingDealSize,omitempty"`

	// Notes
	Notes string `json:"notes,omitempty"`

	// Price in FIL per deal
	PricePerDeal float64 `json:"pricePerDeal,omitempty"`

	// Price in FIL  per GiB
	PricePerGb float64 `json:"pricePerGb,omitempty"`

	// Price in FIL per GiB per epoch
	PricePerGbEpoch float64 `json:"pricePerGbEpoch,omitempty"`

	// Provider
	Provider string `json:"provider,omitempty"`

	// Schedule cron patter
	ScheduleCron string `json:"scheduleCron,omitempty"`

	// Number of deals per scheduled time
	ScheduleDealNumber int64 `json:"scheduleDealNumber,omitempty"`

	// Size of deals per schedule trigger in human readable format, i.e. 100 TiB
	ScheduleDealSize string `json:"scheduleDealSize,omitempty"`

	// Deal start delay in epoch or in duration format, i.e. 1000, 72h
	StartDelay *string `json:"startDelay,omitempty"`

	// Total number of deals
	TotalDealNumber int64 `json:"totalDealNumber,omitempty"`

	// Total size of deals in human readable format, i.e. 100 TiB
	TotalDealSize string `json:"totalDealSize,omitempty"`

	// URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car
	URLTemplate string `json:"urlTemplate,omitempty"`

	// Whether the deal should be verified
	Verified *bool `json:"verified,omitempty"`
}

// Validate validates this schedule create request
func (m *ScheduleCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule create request based on context it is used
func (m *ScheduleCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleCreateRequest) UnmarshalBinary(b []byte) error {
	var res ScheduleCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}