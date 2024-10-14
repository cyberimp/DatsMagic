// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwaggerErrNotAuthorized swagger err not authorized
//
// swagger:model swagger.ErrNotAuthorized
type SwaggerErrNotAuthorized struct {

	// err code
	// Example: 14
	ErrCode int32 `json:"errCode,omitempty"`

	// error
	// Example: not authorized
	Error string `json:"error,omitempty"`
}

// Validate validates this swagger err not authorized
func (m *SwaggerErrNotAuthorized) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swagger err not authorized based on context it is used
func (m *SwaggerErrNotAuthorized) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerErrNotAuthorized) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerErrNotAuthorized) UnmarshalBinary(b []byte) error {
	var res SwaggerErrNotAuthorized
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
