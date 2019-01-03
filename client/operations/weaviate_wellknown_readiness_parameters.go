/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateWellknownReadinessParams creates a new WeaviateWellknownReadinessParams object
// with the default values initialized.
func NewWeaviateWellknownReadinessParams() *WeaviateWellknownReadinessParams {

	return &WeaviateWellknownReadinessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateWellknownReadinessParamsWithTimeout creates a new WeaviateWellknownReadinessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateWellknownReadinessParamsWithTimeout(timeout time.Duration) *WeaviateWellknownReadinessParams {

	return &WeaviateWellknownReadinessParams{

		timeout: timeout,
	}
}

// NewWeaviateWellknownReadinessParamsWithContext creates a new WeaviateWellknownReadinessParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateWellknownReadinessParamsWithContext(ctx context.Context) *WeaviateWellknownReadinessParams {

	return &WeaviateWellknownReadinessParams{

		Context: ctx,
	}
}

// NewWeaviateWellknownReadinessParamsWithHTTPClient creates a new WeaviateWellknownReadinessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateWellknownReadinessParamsWithHTTPClient(client *http.Client) *WeaviateWellknownReadinessParams {

	return &WeaviateWellknownReadinessParams{
		HTTPClient: client,
	}
}

/*WeaviateWellknownReadinessParams contains all the parameters to send to the API endpoint
for the weaviate wellknown readiness operation typically these are written to a http.Request
*/
type WeaviateWellknownReadinessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) WithTimeout(timeout time.Duration) *WeaviateWellknownReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) WithContext(ctx context.Context) *WeaviateWellknownReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) WithHTTPClient(client *http.Client) *WeaviateWellknownReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate wellknown readiness params
func (o *WeaviateWellknownReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateWellknownReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
