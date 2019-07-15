// Code generated by go-swagger; DO NOT EDIT.

package organisations

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

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostUnitsParams creates a new PostUnitsParams object
// with the default values initialized.
func NewPostUnitsParams() *PostUnitsParams {
	var ()
	return &PostUnitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUnitsParamsWithTimeout creates a new PostUnitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUnitsParamsWithTimeout(timeout time.Duration) *PostUnitsParams {
	var ()
	return &PostUnitsParams{

		timeout: timeout,
	}
}

// NewPostUnitsParamsWithContext creates a new PostUnitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUnitsParamsWithContext(ctx context.Context) *PostUnitsParams {
	var ()
	return &PostUnitsParams{

		Context: ctx,
	}
}

// NewPostUnitsParamsWithHTTPClient creates a new PostUnitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUnitsParamsWithHTTPClient(client *http.Client) *PostUnitsParams {
	var ()
	return &PostUnitsParams{
		HTTPClient: client,
	}
}

/*PostUnitsParams contains all the parameters to send to the API endpoint
for the post units operation typically these are written to a http.Request
*/
type PostUnitsParams struct {

	/*OrganisationCreationRequest*/
	OrganisationCreationRequest *models.OrganisationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post units params
func (o *PostUnitsParams) WithTimeout(timeout time.Duration) *PostUnitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post units params
func (o *PostUnitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post units params
func (o *PostUnitsParams) WithContext(ctx context.Context) *PostUnitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post units params
func (o *PostUnitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post units params
func (o *PostUnitsParams) WithHTTPClient(client *http.Client) *PostUnitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post units params
func (o *PostUnitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganisationCreationRequest adds the organisationCreationRequest to the post units params
func (o *PostUnitsParams) WithOrganisationCreationRequest(organisationCreationRequest *models.OrganisationCreation) *PostUnitsParams {
	o.SetOrganisationCreationRequest(organisationCreationRequest)
	return o
}

// SetOrganisationCreationRequest adds the organisationCreationRequest to the post units params
func (o *PostUnitsParams) SetOrganisationCreationRequest(organisationCreationRequest *models.OrganisationCreation) {
	o.OrganisationCreationRequest = organisationCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostUnitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganisationCreationRequest != nil {
		if err := r.SetBodyParam(o.OrganisationCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
