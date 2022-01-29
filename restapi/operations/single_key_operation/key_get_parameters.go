// Code generated by go-swagger; DO NOT EDIT.

package single_key_operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewKeyGetParams creates a new KeyGetParams object
//
// There are no default values defined in the spec.
func NewKeyGetParams() KeyGetParams {

	return KeyGetParams{}
}

// KeyGetParams contains all the bound params for the key get operation
// typically these are obtained from a http.Request
//
// swagger:parameters keyGet
type KeyGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	StrKey string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewKeyGetParams() beforehand.
func (o *KeyGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rStrKey, rhkStrKey, _ := route.Params.GetOK("strKey")
	if err := o.bindStrKey(rStrKey, rhkStrKey, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindStrKey binds and validates parameter StrKey from path.
func (o *KeyGetParams) bindStrKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.StrKey = raw

	return nil
}
