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
	"github.com/go-openapi/strfmt"
)

// NewGetLhvAssociationIDAgencySynchronisationsParams creates a new GetLhvAssociationIDAgencySynchronisationsParams object
// with the default values initialized.
func NewGetLhvAssociationIDAgencySynchronisationsParams() *GetLhvAssociationIDAgencySynchronisationsParams {
	var ()
	return &GetLhvAssociationIDAgencySynchronisationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLhvAssociationIDAgencySynchronisationsParamsWithTimeout creates a new GetLhvAssociationIDAgencySynchronisationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLhvAssociationIDAgencySynchronisationsParamsWithTimeout(timeout time.Duration) *GetLhvAssociationIDAgencySynchronisationsParams {
	var ()
	return &GetLhvAssociationIDAgencySynchronisationsParams{

		timeout: timeout,
	}
}

// NewGetLhvAssociationIDAgencySynchronisationsParamsWithContext creates a new GetLhvAssociationIDAgencySynchronisationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLhvAssociationIDAgencySynchronisationsParamsWithContext(ctx context.Context) *GetLhvAssociationIDAgencySynchronisationsParams {
	var ()
	return &GetLhvAssociationIDAgencySynchronisationsParams{

		Context: ctx,
	}
}

// NewGetLhvAssociationIDAgencySynchronisationsParamsWithHTTPClient creates a new GetLhvAssociationIDAgencySynchronisationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLhvAssociationIDAgencySynchronisationsParamsWithHTTPClient(client *http.Client) *GetLhvAssociationIDAgencySynchronisationsParams {
	var ()
	return &GetLhvAssociationIDAgencySynchronisationsParams{
		HTTPClient: client,
	}
}

/*GetLhvAssociationIDAgencySynchronisationsParams contains all the parameters to send to the API endpoint
for the get lhv association ID agency synchronisations operation typically these are written to a http.Request
*/
type GetLhvAssociationIDAgencySynchronisationsParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) WithTimeout(timeout time.Duration) *GetLhvAssociationIDAgencySynchronisationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) WithContext(ctx context.Context) *GetLhvAssociationIDAgencySynchronisationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) WithHTTPClient(client *http.Client) *GetLhvAssociationIDAgencySynchronisationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) WithAssociationID(associationID strfmt.UUID) *GetLhvAssociationIDAgencySynchronisationsParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the get lhv association ID agency synchronisations params
func (o *GetLhvAssociationIDAgencySynchronisationsParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLhvAssociationIDAgencySynchronisationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
