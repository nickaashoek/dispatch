///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetBaseImageByNameOKCode is the HTTP code returned for type GetBaseImageByNameOK
const GetBaseImageByNameOKCode int = 200

/*GetBaseImageByNameOK successful operation

swagger:response getBaseImageByNameOK
*/
type GetBaseImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *v1.BaseImage `json:"body,omitempty"`
}

// NewGetBaseImageByNameOK creates GetBaseImageByNameOK with default headers values
func NewGetBaseImageByNameOK() *GetBaseImageByNameOK {

	return &GetBaseImageByNameOK{}
}

// WithPayload adds the payload to the get base image by name o k response
func (o *GetBaseImageByNameOK) WithPayload(payload *v1.BaseImage) *GetBaseImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name o k response
func (o *GetBaseImageByNameOK) SetPayload(payload *v1.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBaseImageByNameBadRequestCode is the HTTP code returned for type GetBaseImageByNameBadRequest
const GetBaseImageByNameBadRequestCode int = 400

/*GetBaseImageByNameBadRequest Invalid ID supplied

swagger:response getBaseImageByNameBadRequest
*/
type GetBaseImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetBaseImageByNameBadRequest creates GetBaseImageByNameBadRequest with default headers values
func NewGetBaseImageByNameBadRequest() *GetBaseImageByNameBadRequest {

	return &GetBaseImageByNameBadRequest{}
}

// WithPayload adds the payload to the get base image by name bad request response
func (o *GetBaseImageByNameBadRequest) WithPayload(payload *v1.Error) *GetBaseImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name bad request response
func (o *GetBaseImageByNameBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBaseImageByNameUnauthorizedCode is the HTTP code returned for type GetBaseImageByNameUnauthorized
const GetBaseImageByNameUnauthorizedCode int = 401

/*GetBaseImageByNameUnauthorized Unauthorized Request

swagger:response getBaseImageByNameUnauthorized
*/
type GetBaseImageByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetBaseImageByNameUnauthorized creates GetBaseImageByNameUnauthorized with default headers values
func NewGetBaseImageByNameUnauthorized() *GetBaseImageByNameUnauthorized {

	return &GetBaseImageByNameUnauthorized{}
}

// WithPayload adds the payload to the get base image by name unauthorized response
func (o *GetBaseImageByNameUnauthorized) WithPayload(payload *v1.Error) *GetBaseImageByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name unauthorized response
func (o *GetBaseImageByNameUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBaseImageByNameForbiddenCode is the HTTP code returned for type GetBaseImageByNameForbidden
const GetBaseImageByNameForbiddenCode int = 403

/*GetBaseImageByNameForbidden access to this resource is forbidden

swagger:response getBaseImageByNameForbidden
*/
type GetBaseImageByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetBaseImageByNameForbidden creates GetBaseImageByNameForbidden with default headers values
func NewGetBaseImageByNameForbidden() *GetBaseImageByNameForbidden {

	return &GetBaseImageByNameForbidden{}
}

// WithPayload adds the payload to the get base image by name forbidden response
func (o *GetBaseImageByNameForbidden) WithPayload(payload *v1.Error) *GetBaseImageByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name forbidden response
func (o *GetBaseImageByNameForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBaseImageByNameNotFoundCode is the HTTP code returned for type GetBaseImageByNameNotFound
const GetBaseImageByNameNotFoundCode int = 404

/*GetBaseImageByNameNotFound Base image not found

swagger:response getBaseImageByNameNotFound
*/
type GetBaseImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetBaseImageByNameNotFound creates GetBaseImageByNameNotFound with default headers values
func NewGetBaseImageByNameNotFound() *GetBaseImageByNameNotFound {

	return &GetBaseImageByNameNotFound{}
}

// WithPayload adds the payload to the get base image by name not found response
func (o *GetBaseImageByNameNotFound) WithPayload(payload *v1.Error) *GetBaseImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name not found response
func (o *GetBaseImageByNameNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBaseImageByNameDefault Generic error response

swagger:response getBaseImageByNameDefault
*/
type GetBaseImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetBaseImageByNameDefault creates GetBaseImageByNameDefault with default headers values
func NewGetBaseImageByNameDefault(code int) *GetBaseImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBaseImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get base image by name default response
func (o *GetBaseImageByNameDefault) WithStatusCode(code int) *GetBaseImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get base image by name default response
func (o *GetBaseImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get base image by name default response
func (o *GetBaseImageByNameDefault) WithPayload(payload *v1.Error) *GetBaseImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base image by name default response
func (o *GetBaseImageByNameDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
