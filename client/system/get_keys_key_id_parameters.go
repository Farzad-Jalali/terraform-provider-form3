// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetKeysKeyIDParams creates a new GetKeysKeyIDParams object
// with the default values initialized.
func NewGetKeysKeyIDParams() *GetKeysKeyIDParams {
	var ()
	return &GetKeysKeyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysKeyIDParamsWithTimeout creates a new GetKeysKeyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysKeyIDParamsWithTimeout(timeout time.Duration) *GetKeysKeyIDParams {
	var ()
	return &GetKeysKeyIDParams{

		timeout: timeout,
	}
}

// NewGetKeysKeyIDParamsWithContext creates a new GetKeysKeyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysKeyIDParamsWithContext(ctx context.Context) *GetKeysKeyIDParams {
	var ()
	return &GetKeysKeyIDParams{

		Context: ctx,
	}
}

// NewGetKeysKeyIDParamsWithHTTPClient creates a new GetKeysKeyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysKeyIDParamsWithHTTPClient(client *http.Client) *GetKeysKeyIDParams {
	var ()
	return &GetKeysKeyIDParams{
		HTTPClient: client,
	}
}

/*GetKeysKeyIDParams contains all the parameters to send to the API endpoint
for the get keys key ID operation typically these are written to a http.Request
*/
type GetKeysKeyIDParams struct {

	/*KeyID
	  Key Id

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys key ID params
func (o *GetKeysKeyIDParams) WithTimeout(timeout time.Duration) *GetKeysKeyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys key ID params
func (o *GetKeysKeyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys key ID params
func (o *GetKeysKeyIDParams) WithContext(ctx context.Context) *GetKeysKeyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys key ID params
func (o *GetKeysKeyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys key ID params
func (o *GetKeysKeyIDParams) WithHTTPClient(client *http.Client) *GetKeysKeyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys key ID params
func (o *GetKeysKeyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the get keys key ID params
func (o *GetKeysKeyIDParams) WithKeyID(keyID strfmt.UUID) *GetKeysKeyIDParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the get keys key ID params
func (o *GetKeysKeyIDParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysKeyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key_id
	if err := r.SetPathParam("key_id", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
