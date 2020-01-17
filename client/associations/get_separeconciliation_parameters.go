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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSepareconciliationParams creates a new GetSepareconciliationParams object
// with the default values initialized.
func NewGetSepareconciliationParams() *GetSepareconciliationParams {
	var ()
	return &GetSepareconciliationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepareconciliationParamsWithTimeout creates a new GetSepareconciliationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSepareconciliationParamsWithTimeout(timeout time.Duration) *GetSepareconciliationParams {
	var ()
	return &GetSepareconciliationParams{

		timeout: timeout,
	}
}

// NewGetSepareconciliationParamsWithContext creates a new GetSepareconciliationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSepareconciliationParamsWithContext(ctx context.Context) *GetSepareconciliationParams {
	var ()
	return &GetSepareconciliationParams{

		Context: ctx,
	}
}

// NewGetSepareconciliationParamsWithHTTPClient creates a new GetSepareconciliationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSepareconciliationParamsWithHTTPClient(client *http.Client) *GetSepareconciliationParams {
	var ()
	return &GetSepareconciliationParams{
		HTTPClient: client,
	}
}

/*GetSepareconciliationParams contains all the parameters to send to the API endpoint
for the get separeconciliation operation typically these are written to a http.Request
*/
type GetSepareconciliationParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get separeconciliation params
func (o *GetSepareconciliationParams) WithTimeout(timeout time.Duration) *GetSepareconciliationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get separeconciliation params
func (o *GetSepareconciliationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get separeconciliation params
func (o *GetSepareconciliationParams) WithContext(ctx context.Context) *GetSepareconciliationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get separeconciliation params
func (o *GetSepareconciliationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get separeconciliation params
func (o *GetSepareconciliationParams) WithHTTPClient(client *http.Client) *GetSepareconciliationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get separeconciliation params
func (o *GetSepareconciliationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get separeconciliation params
func (o *GetSepareconciliationParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetSepareconciliationParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get separeconciliation params
func (o *GetSepareconciliationParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepareconciliationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
