// Code generated by go-swagger; DO NOT EDIT.

package single_key_operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// KeyGetHandlerFunc turns a function with the right signature into a key get handler
type KeyGetHandlerFunc func(KeyGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn KeyGetHandlerFunc) Handle(params KeyGetParams) middleware.Responder {
	return fn(params)
}

// KeyGetHandler interface for that can handle valid key get params
type KeyGetHandler interface {
	Handle(KeyGetParams) middleware.Responder
}

// NewKeyGet creates a new http.Handler for the key get operation
func NewKeyGet(ctx *middleware.Context, handler KeyGetHandler) *KeyGet {
	return &KeyGet{Context: ctx, Handler: handler}
}

/* KeyGet swagger:route GET /key/{strKey} single key operation keyGet

Returns a single key Data by given Key as a string

*/
type KeyGet struct {
	Context *middleware.Context
	Handler KeyGetHandler
}

func (o *KeyGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewKeyGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
