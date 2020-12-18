// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewGetPayportIDParams creates a new GetPayportIDParams object
// with the default values initialized.
func NewGetPayportIDParams() *GetPayportIDParams {
	var ()
	return &GetPayportIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPayportIDParamsWithTimeout creates a new GetPayportIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPayportIDParamsWithTimeout(timeout time.Duration) *GetPayportIDParams {
	var ()
	return &GetPayportIDParams{

		timeout: timeout,
	}
}

// NewGetPayportIDParamsWithContext creates a new GetPayportIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPayportIDParamsWithContext(ctx context.Context) *GetPayportIDParams {
	var ()
	return &GetPayportIDParams{

		Context: ctx,
	}
}

// NewGetPayportIDParamsWithHTTPClient creates a new GetPayportIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPayportIDParamsWithHTTPClient(client *http.Client) *GetPayportIDParams {
	var ()
	return &GetPayportIDParams{
		HTTPClient: client,
	}
}

/*GetPayportIDParams contains all the parameters to send to the API endpoint
for the get payport ID operation typically these are written to a http.Request
*/
type GetPayportIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payport ID params
func (o *GetPayportIDParams) WithTimeout(timeout time.Duration) *GetPayportIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payport ID params
func (o *GetPayportIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payport ID params
func (o *GetPayportIDParams) WithContext(ctx context.Context) *GetPayportIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payport ID params
func (o *GetPayportIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payport ID params
func (o *GetPayportIDParams) WithHTTPClient(client *http.Client) *GetPayportIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payport ID params
func (o *GetPayportIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get payport ID params
func (o *GetPayportIDParams) WithID(id strfmt.UUID) *GetPayportIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payport ID params
func (o *GetPayportIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPayportIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
