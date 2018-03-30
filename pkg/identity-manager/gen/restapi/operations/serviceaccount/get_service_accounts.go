///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetServiceAccountsHandlerFunc turns a function with the right signature into a get service accounts handler
type GetServiceAccountsHandlerFunc func(GetServiceAccountsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServiceAccountsHandlerFunc) Handle(params GetServiceAccountsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetServiceAccountsHandler interface for that can handle valid get service accounts params
type GetServiceAccountsHandler interface {
	Handle(GetServiceAccountsParams, interface{}) middleware.Responder
}

// NewGetServiceAccounts creates a new http.Handler for the get service accounts operation
func NewGetServiceAccounts(ctx *middleware.Context, handler GetServiceAccountsHandler) *GetServiceAccounts {
	return &GetServiceAccounts{Context: ctx, Handler: handler}
}

/*GetServiceAccounts swagger:route GET /v1/iam/serviceaccount serviceaccount getServiceAccounts

List all existing service accounts

*/
type GetServiceAccounts struct {
	Context *middleware.Context
	Handler GetServiceAccountsHandler
}

func (o *GetServiceAccounts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServiceAccountsParams()

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
