// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSetPXCComponentsParams creates a new SetPXCComponentsParams object
// with the default values initialized.
func NewSetPXCComponentsParams() *SetPXCComponentsParams {
	var ()
	return &SetPXCComponentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPXCComponentsParamsWithTimeout creates a new SetPXCComponentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPXCComponentsParamsWithTimeout(timeout time.Duration) *SetPXCComponentsParams {
	var ()
	return &SetPXCComponentsParams{

		timeout: timeout,
	}
}

// NewSetPXCComponentsParamsWithContext creates a new SetPXCComponentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPXCComponentsParamsWithContext(ctx context.Context) *SetPXCComponentsParams {
	var ()
	return &SetPXCComponentsParams{

		Context: ctx,
	}
}

// NewSetPXCComponentsParamsWithHTTPClient creates a new SetPXCComponentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPXCComponentsParamsWithHTTPClient(client *http.Client) *SetPXCComponentsParams {
	var ()
	return &SetPXCComponentsParams{
		HTTPClient: client,
	}
}

/*SetPXCComponentsParams contains all the parameters to send to the API endpoint
for the set p x c components operation typically these are written to a http.Request
*/
type SetPXCComponentsParams struct {

	/*Body*/
	Body SetPXCComponentsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set p x c components params
func (o *SetPXCComponentsParams) WithTimeout(timeout time.Duration) *SetPXCComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set p x c components params
func (o *SetPXCComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set p x c components params
func (o *SetPXCComponentsParams) WithContext(ctx context.Context) *SetPXCComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set p x c components params
func (o *SetPXCComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set p x c components params
func (o *SetPXCComponentsParams) WithHTTPClient(client *http.Client) *SetPXCComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set p x c components params
func (o *SetPXCComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set p x c components params
func (o *SetPXCComponentsParams) WithBody(body SetPXCComponentsBody) *SetPXCComponentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set p x c components params
func (o *SetPXCComponentsParams) SetBody(body SetPXCComponentsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetPXCComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
