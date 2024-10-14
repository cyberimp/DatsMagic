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

// ViewGlobalPlayer view global player
//
// swagger:model view.GlobalPlayer
type ViewGlobalPlayer struct {

	// name
	// Example: player1
	Name string `json:"name,omitempty"`

	// points
	// Example: 100
	Points int64 `json:"points,omitempty"`

	// transports
	Transports []*ViewTransport `json:"transports"`
}

// Validate validates this view global player
func (m *ViewGlobalPlayer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewGlobalPlayer) validateTransports(formats strfmt.Registry) error {
	if swag.IsZero(m.Transports) { // not required
		return nil
	}

	for i := 0; i < len(m.Transports); i++ {
		if swag.IsZero(m.Transports[i]) { // not required
			continue
		}

		if m.Transports[i] != nil {
			if err := m.Transports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this view global player based on the context it is used
func (m *ViewGlobalPlayer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewGlobalPlayer) contextValidateTransports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Transports); i++ {

		if m.Transports[i] != nil {

			if swag.IsZero(m.Transports[i]) { // not required
				return nil
			}

			if err := m.Transports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViewGlobalPlayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewGlobalPlayer) UnmarshalBinary(b []byte) error {
	var res ViewGlobalPlayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
