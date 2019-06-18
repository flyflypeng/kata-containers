// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetMmdsParams creates a new GetMmdsParams object
// with the default values initialized.
func NewGetMmdsParams() *GetMmdsParams {

	return &GetMmdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMmdsParamsWithTimeout creates a new GetMmdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMmdsParamsWithTimeout(timeout time.Duration) *GetMmdsParams {

	return &GetMmdsParams{

		timeout: timeout,
	}
}

// NewGetMmdsParamsWithContext creates a new GetMmdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMmdsParamsWithContext(ctx context.Context) *GetMmdsParams {

	return &GetMmdsParams{

		Context: ctx,
	}
}

// NewGetMmdsParamsWithHTTPClient creates a new GetMmdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMmdsParamsWithHTTPClient(client *http.Client) *GetMmdsParams {

	return &GetMmdsParams{
		HTTPClient: client,
	}
}

/*GetMmdsParams contains all the parameters to send to the API endpoint
for the get mmds operation typically these are written to a http.Request
*/
type GetMmdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mmds params
func (o *GetMmdsParams) WithTimeout(timeout time.Duration) *GetMmdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mmds params
func (o *GetMmdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mmds params
func (o *GetMmdsParams) WithContext(ctx context.Context) *GetMmdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mmds params
func (o *GetMmdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mmds params
func (o *GetMmdsParams) WithHTTPClient(client *http.Client) *GetMmdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mmds params
func (o *GetMmdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMmdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
