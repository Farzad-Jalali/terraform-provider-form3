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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// NewPostVocalinkreportParams creates a new PostVocalinkreportParams object
// with the default values initialized.
func NewPostVocalinkreportParams() *PostVocalinkreportParams {
	var ()
	return &PostVocalinkreportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVocalinkreportParamsWithTimeout creates a new PostVocalinkreportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVocalinkreportParamsWithTimeout(timeout time.Duration) *PostVocalinkreportParams {
	var ()
	return &PostVocalinkreportParams{

		timeout: timeout,
	}
}

// NewPostVocalinkreportParamsWithContext creates a new PostVocalinkreportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVocalinkreportParamsWithContext(ctx context.Context) *PostVocalinkreportParams {
	var ()
	return &PostVocalinkreportParams{

		Context: ctx,
	}
}

// NewPostVocalinkreportParamsWithHTTPClient creates a new PostVocalinkreportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVocalinkreportParamsWithHTTPClient(client *http.Client) *PostVocalinkreportParams {
	var ()
	return &PostVocalinkreportParams{
		HTTPClient: client,
	}
}

/*PostVocalinkreportParams contains all the parameters to send to the API endpoint
for the post vocalinkreport operation typically these are written to a http.Request
*/
type PostVocalinkreportParams struct {

	/*CreationRequest*/
	CreationRequest *models.VocalinkReportAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vocalinkreport params
func (o *PostVocalinkreportParams) WithTimeout(timeout time.Duration) *PostVocalinkreportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vocalinkreport params
func (o *PostVocalinkreportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vocalinkreport params
func (o *PostVocalinkreportParams) WithContext(ctx context.Context) *PostVocalinkreportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vocalinkreport params
func (o *PostVocalinkreportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vocalinkreport params
func (o *PostVocalinkreportParams) WithHTTPClient(client *http.Client) *PostVocalinkreportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vocalinkreport params
func (o *PostVocalinkreportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post vocalinkreport params
func (o *PostVocalinkreportParams) WithCreationRequest(creationRequest *models.VocalinkReportAssociationCreation) *PostVocalinkreportParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post vocalinkreport params
func (o *PostVocalinkreportParams) SetCreationRequest(creationRequest *models.VocalinkReportAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostVocalinkreportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreationRequest != nil {
		if err := r.SetBodyParam(o.CreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}