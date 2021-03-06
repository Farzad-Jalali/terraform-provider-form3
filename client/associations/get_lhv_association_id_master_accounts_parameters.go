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

// NewGetLhvAssociationIDMasterAccountsParams creates a new GetLhvAssociationIDMasterAccountsParams object
// with the default values initialized.
func NewGetLhvAssociationIDMasterAccountsParams() *GetLhvAssociationIDMasterAccountsParams {
	var ()
	return &GetLhvAssociationIDMasterAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLhvAssociationIDMasterAccountsParamsWithTimeout creates a new GetLhvAssociationIDMasterAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLhvAssociationIDMasterAccountsParamsWithTimeout(timeout time.Duration) *GetLhvAssociationIDMasterAccountsParams {
	var ()
	return &GetLhvAssociationIDMasterAccountsParams{

		timeout: timeout,
	}
}

// NewGetLhvAssociationIDMasterAccountsParamsWithContext creates a new GetLhvAssociationIDMasterAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLhvAssociationIDMasterAccountsParamsWithContext(ctx context.Context) *GetLhvAssociationIDMasterAccountsParams {
	var ()
	return &GetLhvAssociationIDMasterAccountsParams{

		Context: ctx,
	}
}

// NewGetLhvAssociationIDMasterAccountsParamsWithHTTPClient creates a new GetLhvAssociationIDMasterAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLhvAssociationIDMasterAccountsParamsWithHTTPClient(client *http.Client) *GetLhvAssociationIDMasterAccountsParams {
	var ()
	return &GetLhvAssociationIDMasterAccountsParams{
		HTTPClient: client,
	}
}

/*GetLhvAssociationIDMasterAccountsParams contains all the parameters to send to the API endpoint
for the get lhv association ID master accounts operation typically these are written to a http.Request
*/
type GetLhvAssociationIDMasterAccountsParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) WithTimeout(timeout time.Duration) *GetLhvAssociationIDMasterAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) WithContext(ctx context.Context) *GetLhvAssociationIDMasterAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) WithHTTPClient(client *http.Client) *GetLhvAssociationIDMasterAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) WithAssociationID(associationID strfmt.UUID) *GetLhvAssociationIDMasterAccountsParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the get lhv association ID master accounts params
func (o *GetLhvAssociationIDMasterAccountsParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLhvAssociationIDMasterAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
