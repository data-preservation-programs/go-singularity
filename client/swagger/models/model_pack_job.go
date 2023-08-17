// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelPackJob model pack job
//
// swagger:model model.PackJob
type ModelPackJob struct {

	// cars
	Cars []*ModelCar `json:"cars"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// file ranges
	FileRanges []*ModelFileRange `json:"fileRanges"`

	// id
	ID int64 `json:"id,omitempty"`

	// packing state
	PackingState ModelWorkState `json:"packingState,omitempty"`

	// packing worker Id
	PackingWorkerID string `json:"packingWorkerId,omitempty"`

	// source Id
	SourceID int64 `json:"sourceId,omitempty"`
}

// Validate validates this model pack job
func (m *ModelPackJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackingState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelPackJob) validateCars(formats strfmt.Registry) error {
	if swag.IsZero(m.Cars) { // not required
		return nil
	}

	for i := 0; i < len(m.Cars); i++ {
		if swag.IsZero(m.Cars[i]) { // not required
			continue
		}

		if m.Cars[i] != nil {
			if err := m.Cars[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelPackJob) validateFileRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.FileRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.FileRanges); i++ {
		if swag.IsZero(m.FileRanges[i]) { // not required
			continue
		}

		if m.FileRanges[i] != nil {
			if err := m.FileRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelPackJob) validatePackingState(formats strfmt.Registry) error {
	if swag.IsZero(m.PackingState) { // not required
		return nil
	}

	if err := m.PackingState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packingState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packingState")
		}
		return err
	}

	return nil
}

// ContextValidate validate this model pack job based on the context it is used
func (m *ModelPackJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackingState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelPackJob) contextValidateCars(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cars); i++ {

		if m.Cars[i] != nil {

			if swag.IsZero(m.Cars[i]) { // not required
				return nil
			}

			if err := m.Cars[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelPackJob) contextValidateFileRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileRanges); i++ {

		if m.FileRanges[i] != nil {

			if swag.IsZero(m.FileRanges[i]) { // not required
				return nil
			}

			if err := m.FileRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelPackJob) contextValidatePackingState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PackingState) { // not required
		return nil
	}

	if err := m.PackingState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packingState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packingState")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelPackJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelPackJob) UnmarshalBinary(b []byte) error {
	var res ModelPackJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
