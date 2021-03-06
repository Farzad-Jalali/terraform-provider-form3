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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPatchGocardlessIDParams creates a new PatchGocardlessIDParams object
// with the default values initialized.
func NewPatchGocardlessIDParams() *PatchGocardlessIDParams {
	var ()
	return &PatchGocardlessIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGocardlessIDParamsWithTimeout creates a new PatchGocardlessIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchGocardlessIDParamsWithTimeout(timeout time.Duration) *PatchGocardlessIDParams {
	var ()
	return &PatchGocardlessIDParams{

		timeout: timeout,
	}
}

// NewPatchGocardlessIDParamsWithContext creates a new PatchGocardlessIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchGocardlessIDParamsWithContext(ctx context.Context) *PatchGocardlessIDParams {
	var ()
	return &PatchGocardlessIDParams{

		Context: ctx,
	}
}

// NewPatchGocardlessIDParamsWithHTTPClient creates a new PatchGocardlessIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchGocardlessIDParamsWithHTTPClient(client *http.Client) *PatchGocardlessIDParams {
	var ()
	return &PatchGocardlessIDParams{
		HTTPClient: client,
	}
}

/*PatchGocardlessIDParams contains all the parameters to send to the API endpoint
for the patch gocardless ID operation typically these are written to a http.Request
*/
type PatchGocardlessIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID
	/*PatchBody*/
	PatchBody *models.GocardlessAssociationAmendment
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithTimeout(timeout time.Duration) *PatchGocardlessIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithContext(ctx context.Context) *PatchGocardlessIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithHTTPClient(client *http.Client) *PatchGocardlessIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithID(id strfmt.UUID) *PatchGocardlessIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPatchBody adds the patchBody to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithPatchBody(patchBody *models.GocardlessAssociationAmendment) *PatchGocardlessIDParams {
	o.SetPatchBody(patchBody)
	return o
}

// SetPatchBody adds the patchBody to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetPatchBody(patchBody *models.GocardlessAssociationAmendment) {
	o.PatchBody = patchBody
}

// WithVersion adds the version to the patch gocardless ID params
func (o *PatchGocardlessIDParams) WithVersion(version int64) *PatchGocardlessIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch gocardless ID params
func (o *PatchGocardlessIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGocardlessIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.PatchBody != nil {
		if err := r.SetBodyParam(o.PatchBody); err != nil {
			return err
		}
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
