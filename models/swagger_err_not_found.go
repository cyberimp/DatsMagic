// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwaggerErrNotFound swagger err not found
//
// swagger:model swagger.ErrNotFound
type SwaggerErrNotFound struct {

	// err code
	// Example: 11
	ErrCode int32 `json:"errCode,omitempty"`

	// error
	// Example: not found
	Error string `json:"error,omitempty"`
}

// Validate validates this swagger err not found
func (m *SwaggerErrNotFound) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swagger err not found based on context it is used
func (m *SwaggerErrNotFound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerErrNotFound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerErrNotFound) UnmarshalBinary(b []byte) error {
	var res SwaggerErrNotFound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
