// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"log"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateKey create key
//
// swagger:model CreateKey
type CreateKey struct {

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// expire
	// Example: 2019-05-17
	// Pattern: /(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})/
	Expire string `json:"Expire,omitempty"`
}

// Validate validates this create key
func (m *CreateKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpire(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateKey) validateExpire(formats strfmt.Registry) error {
	if swag.IsZero(m.Expire) { // not required
		return nil
	}

	log.Println(m.Expire)

	if err := validate.Pattern("Expire", "body", m.Expire, `(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})`); err != nil {
		log.Println(err.Error())

		return err
	}

	return nil
}

// ContextValidate validates this create key based on context it is used
func (m *CreateKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKey) UnmarshalBinary(b []byte) error {
	var res CreateKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
