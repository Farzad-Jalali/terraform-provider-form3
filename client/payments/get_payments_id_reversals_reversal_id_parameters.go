// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// NewGetPaymentsIDReversalsReversalIDParams creates a new GetPaymentsIDReversalsReversalIDParams object
// with the default values initialized.
func NewGetPaymentsIDReversalsReversalIDParams() *GetPaymentsIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReversalsReversalIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDReversalsReversalIDParamsWithTimeout creates a new GetPaymentsIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDReversalsReversalIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReversalsReversalIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDReversalsReversalIDParamsWithContext creates a new GetPaymentsIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDReversalsReversalIDParamsWithContext(ctx context.Context) *GetPaymentsIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReversalsReversalIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDReversalsReversalIDParamsWithHTTPClient creates a new GetPaymentsIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDReversalsReversalIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReversalsReversalIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDReversalsReversalIDParams contains all the parameters to send to the API endpoint
for the get payments ID reversals reversal ID operation typically these are written to a http.Request
*/
type GetPaymentsIDReversalsReversalIDParams struct {

	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDReversalsReversalIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) WithContext(ctx context.Context) *GetPaymentsIDReversalsReversalIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDReversalsReversalIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) WithID(id strfmt.UUID) *GetPaymentsIDReversalsReversalIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReversalID adds the reversalID to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) WithReversalID(reversalID strfmt.UUID) *GetPaymentsIDReversalsReversalIDParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the get payments ID reversals reversal ID params
func (o *GetPaymentsIDReversalsReversalIDParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDReversalsReversalIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
