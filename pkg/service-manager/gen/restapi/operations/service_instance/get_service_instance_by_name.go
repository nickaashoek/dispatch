///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetServiceInstanceByNameHandlerFunc turns a function with the right signature into a get service instance by name handler
type GetServiceInstanceByNameHandlerFunc func(GetServiceInstanceByNameParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServiceInstanceByNameHandlerFunc) Handle(params GetServiceInstanceByNameParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetServiceInstanceByNameHandler interface for that can handle valid get service instance by name params
type GetServiceInstanceByNameHandler interface {
	Handle(GetServiceInstanceByNameParams, interface{}) middleware.Responder
}

// NewGetServiceInstanceByName creates a new http.Handler for the get service instance by name operation
func NewGetServiceInstanceByName(ctx *middleware.Context, handler GetServiceInstanceByNameHandler) *GetServiceInstanceByName {
	return &GetServiceInstanceByName{Context: ctx, Handler: handler}
}

/*GetServiceInstanceByName swagger:route GET /serviceinstance/{serviceInstanceName} serviceInstance getServiceInstanceByName

Find service instance by name

Returns a single service instance

*/
type GetServiceInstanceByName struct {
	Context *middleware.Context
	Handler GetServiceInstanceByNameHandler
}

func (o *GetServiceInstanceByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServiceInstanceByNameParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
