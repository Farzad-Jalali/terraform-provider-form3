// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// NewGetPaymentsIDAdmissionsAdmissionIDParams creates a new GetPaymentsIDAdmissionsAdmissionIDParams object
// with the default values initialized.
func NewGetPaymentsIDAdmissionsAdmissionIDParams() *GetPaymentsIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDAdmissionsAdmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDAdmissionsAdmissionIDParamsWithTimeout creates a new GetPaymentsIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDAdmissionsAdmissionIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDAdmissionsAdmissionIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDAdmissionsAdmissionIDParamsWithContext creates a new GetPaymentsIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDAdmissionsAdmissionIDParamsWithContext(ctx context.Context) *GetPaymentsIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDAdmissionsAdmissionIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDAdmissionsAdmissionIDParamsWithHTTPClient creates a new GetPaymentsIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDAdmissionsAdmissionIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDAdmissionsAdmissionIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDAdmissionsAdmissionIDParams contains all the parameters to send to the API endpoint
for the get payments ID admissions admission ID operation typically these are written to a http.Request
*/
type GetPaymentsIDAdmissionsAdmissionIDParams struct {

	/*AdmissionID
	  Admission Id

	*/
	AdmissionID strfmt.UUID
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDAdmissionsAdmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WithContext(ctx context.Context) *GetPaymentsIDAdmissionsAdmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDAdmissionsAdmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdmissionID adds the admissionID to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WithAdmissionID(admissionID strfmt.UUID) *GetPaymentsIDAdmissionsAdmissionIDParams {
	o.SetAdmissionID(admissionID)
	return o
}

// SetAdmissionID adds the admissionId to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) SetAdmissionID(admissionID strfmt.UUID) {
	o.AdmissionID = admissionID
}

// WithID adds the id to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WithID(id strfmt.UUID) *GetPaymentsIDAdmissionsAdmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID admissions admission ID params
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDAdmissionsAdmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}