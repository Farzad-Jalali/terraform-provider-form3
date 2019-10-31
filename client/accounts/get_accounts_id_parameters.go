// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetAccountsIDParams creates a new GetAccountsIDParams object
// with the default values initialized.
func NewGetAccountsIDParams() *GetAccountsIDParams {
	var ()
	return &GetAccountsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsIDParamsWithTimeout creates a new GetAccountsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsIDParamsWithTimeout(timeout time.Duration) *GetAccountsIDParams {
	var ()
	return &GetAccountsIDParams{

		timeout: timeout,
	}
}

// NewGetAccountsIDParamsWithContext creates a new GetAccountsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsIDParamsWithContext(ctx context.Context) *GetAccountsIDParams {
	var ()
	return &GetAccountsIDParams{

		Context: ctx,
	}
}

// NewGetAccountsIDParamsWithHTTPClient creates a new GetAccountsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsIDParamsWithHTTPClient(client *http.Client) *GetAccountsIDParams {
	var ()
	return &GetAccountsIDParams{
		HTTPClient: client,
	}
}

/*GetAccountsIDParams contains all the parameters to send to the API endpoint
for the get accounts ID operation typically these are written to a http.Request
*/
type GetAccountsIDParams struct {

	/*ID
	  Account Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accounts ID params
func (o *GetAccountsIDParams) WithTimeout(timeout time.Duration) *GetAccountsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts ID params
func (o *GetAccountsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts ID params
func (o *GetAccountsIDParams) WithContext(ctx context.Context) *GetAccountsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts ID params
func (o *GetAccountsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts ID params
func (o *GetAccountsIDParams) WithHTTPClient(client *http.Client) *GetAccountsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts ID params
func (o *GetAccountsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get accounts ID params
func (o *GetAccountsIDParams) WithID(id strfmt.UUID) *GetAccountsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get accounts ID params
func (o *GetAccountsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
