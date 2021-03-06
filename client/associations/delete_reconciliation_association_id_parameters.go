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
	"github.com/go-openapi/swag"
)

// NewDeleteReconciliationAssociationIDParams creates a new DeleteReconciliationAssociationIDParams object
// with the default values initialized.
func NewDeleteReconciliationAssociationIDParams() *DeleteReconciliationAssociationIDParams {
	var ()
	return &DeleteReconciliationAssociationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReconciliationAssociationIDParamsWithTimeout creates a new DeleteReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteReconciliationAssociationIDParamsWithTimeout(timeout time.Duration) *DeleteReconciliationAssociationIDParams {
	var ()
	return &DeleteReconciliationAssociationIDParams{

		timeout: timeout,
	}
}

// NewDeleteReconciliationAssociationIDParamsWithContext creates a new DeleteReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteReconciliationAssociationIDParamsWithContext(ctx context.Context) *DeleteReconciliationAssociationIDParams {
	var ()
	return &DeleteReconciliationAssociationIDParams{

		Context: ctx,
	}
}

// NewDeleteReconciliationAssociationIDParamsWithHTTPClient creates a new DeleteReconciliationAssociationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteReconciliationAssociationIDParamsWithHTTPClient(client *http.Client) *DeleteReconciliationAssociationIDParams {
	var ()
	return &DeleteReconciliationAssociationIDParams{
		HTTPClient: client,
	}
}

/*DeleteReconciliationAssociationIDParams contains all the parameters to send to the API endpoint
for the delete reconciliation association ID operation typically these are written to a http.Request
*/
type DeleteReconciliationAssociationIDParams struct {

	/*AssociationID
	  Association id

	*/
	AssociationID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) WithTimeout(timeout time.Duration) *DeleteReconciliationAssociationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) WithContext(ctx context.Context) *DeleteReconciliationAssociationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) WithHTTPClient(client *http.Client) *DeleteReconciliationAssociationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) WithAssociationID(associationID strfmt.UUID) *DeleteReconciliationAssociationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithVersion adds the version to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) WithVersion(version int64) *DeleteReconciliationAssociationIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete reconciliation association ID params
func (o *DeleteReconciliationAssociationIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReconciliationAssociationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
