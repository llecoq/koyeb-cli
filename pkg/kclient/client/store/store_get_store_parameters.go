// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStoreGetStoreParams creates a new StoreGetStoreParams object
// with the default values initialized.
func NewStoreGetStoreParams() *StoreGetStoreParams {
	var ()
	return &StoreGetStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStoreGetStoreParamsWithTimeout creates a new StoreGetStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStoreGetStoreParamsWithTimeout(timeout time.Duration) *StoreGetStoreParams {
	var ()
	return &StoreGetStoreParams{

		timeout: timeout,
	}
}

// NewStoreGetStoreParamsWithContext creates a new StoreGetStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewStoreGetStoreParamsWithContext(ctx context.Context) *StoreGetStoreParams {
	var ()
	return &StoreGetStoreParams{

		Context: ctx,
	}
}

// NewStoreGetStoreParamsWithHTTPClient creates a new StoreGetStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStoreGetStoreParamsWithHTTPClient(client *http.Client) *StoreGetStoreParams {
	var ()
	return &StoreGetStoreParams{
		HTTPClient: client,
	}
}

/*StoreGetStoreParams contains all the parameters to send to the API endpoint
for the store get store operation typically these are written to a http.Request
*/
type StoreGetStoreParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the store get store params
func (o *StoreGetStoreParams) WithTimeout(timeout time.Duration) *StoreGetStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the store get store params
func (o *StoreGetStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the store get store params
func (o *StoreGetStoreParams) WithContext(ctx context.Context) *StoreGetStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the store get store params
func (o *StoreGetStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the store get store params
func (o *StoreGetStoreParams) WithHTTPClient(client *http.Client) *StoreGetStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the store get store params
func (o *StoreGetStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the store get store params
func (o *StoreGetStoreParams) WithID(id string) *StoreGetStoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the store get store params
func (o *StoreGetStoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StoreGetStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
