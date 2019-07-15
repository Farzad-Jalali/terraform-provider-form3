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

	models "github.com/form3tech/terraform-provider-form3/models"
)

// NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams creates a new PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams object
// with the default values initialized.
func NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams() *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithTimeout creates a new PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithTimeout(timeout time.Duration) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams{

		timeout: timeout,
	}
}

// NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithContext creates a new PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithContext(ctx context.Context) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams{

		Context: ctx,
	}
}

// NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithHTTPClient creates a new PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParamsWithHTTPClient(client *http.Client) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams{
		HTTPClient: client,
	}
}

/*PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams contains all the parameters to send to the API endpoint
for the patch payments ID returns return ID submissions submission ID operation typically these are written to a http.Request
*/
type PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams struct {

	/*ReturnSubmissionUpdateRequest*/
	ReturnSubmissionUpdateRequest *models.ReturnSubmissionAmendment
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReturnID
	  Return Id

	*/
	ReturnID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithTimeout(timeout time.Duration) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithContext(ctx context.Context) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithHTTPClient(client *http.Client) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnSubmissionUpdateRequest adds the returnSubmissionUpdateRequest to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithReturnSubmissionUpdateRequest(returnSubmissionUpdateRequest *models.ReturnSubmissionAmendment) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetReturnSubmissionUpdateRequest(returnSubmissionUpdateRequest)
	return o
}

// SetReturnSubmissionUpdateRequest adds the returnSubmissionUpdateRequest to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetReturnSubmissionUpdateRequest(returnSubmissionUpdateRequest *models.ReturnSubmissionAmendment) {
	o.ReturnSubmissionUpdateRequest = returnSubmissionUpdateRequest
}

// WithID adds the id to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithID(id strfmt.UUID) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReturnID adds the returnID to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithReturnID(returnID strfmt.UUID) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetReturnID(returnID)
	return o
}

// SetReturnID adds the returnId to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetReturnID(returnID strfmt.UUID) {
	o.ReturnID = returnID
}

// WithSubmissionID adds the submissionID to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WithSubmissionID(submissionID strfmt.UUID) *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the patch payments ID returns return ID submissions submission ID params
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentsIDReturnsReturnIDSubmissionsSubmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnSubmissionUpdateRequest != nil {
		if err := r.SetBodyParam(o.ReturnSubmissionUpdateRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
