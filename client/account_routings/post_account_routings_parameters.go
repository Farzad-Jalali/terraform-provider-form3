// Code generated by go-swagger; DO NOT EDIT.

package account_routings

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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostAccountRoutingsParams creates a new PostAccountRoutingsParams object
// with the default values initialized.
func NewPostAccountRoutingsParams() *PostAccountRoutingsParams {
	var ()
	return &PostAccountRoutingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountRoutingsParamsWithTimeout creates a new PostAccountRoutingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountRoutingsParamsWithTimeout(timeout time.Duration) *PostAccountRoutingsParams {
	var ()
	return &PostAccountRoutingsParams{

		timeout: timeout,
	}
}

// NewPostAccountRoutingsParamsWithContext creates a new PostAccountRoutingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAccountRoutingsParamsWithContext(ctx context.Context) *PostAccountRoutingsParams {
	var ()
	return &PostAccountRoutingsParams{

		Context: ctx,
	}
}

// NewPostAccountRoutingsParamsWithHTTPClient creates a new PostAccountRoutingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAccountRoutingsParamsWithHTTPClient(client *http.Client) *PostAccountRoutingsParams {
	var ()
	return &PostAccountRoutingsParams{
		HTTPClient: client,
	}
}

/*PostAccountRoutingsParams contains all the parameters to send to the API endpoint
for the post account routings operation typically these are written to a http.Request
*/
type PostAccountRoutingsParams struct {

	/*AccountRoutingCreationRequest*/
	AccountRoutingCreationRequest *models.AccountRoutingCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post account routings params
func (o *PostAccountRoutingsParams) WithTimeout(timeout time.Duration) *PostAccountRoutingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post account routings params
func (o *PostAccountRoutingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post account routings params
func (o *PostAccountRoutingsParams) WithContext(ctx context.Context) *PostAccountRoutingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post account routings params
func (o *PostAccountRoutingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post account routings params
func (o *PostAccountRoutingsParams) WithHTTPClient(client *http.Client) *PostAccountRoutingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post account routings params
func (o *PostAccountRoutingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountRoutingCreationRequest adds the accountRoutingCreationRequest to the post account routings params
func (o *PostAccountRoutingsParams) WithAccountRoutingCreationRequest(accountRoutingCreationRequest *models.AccountRoutingCreation) *PostAccountRoutingsParams {
	o.SetAccountRoutingCreationRequest(accountRoutingCreationRequest)
	return o
}

// SetAccountRoutingCreationRequest adds the accountRoutingCreationRequest to the post account routings params
func (o *PostAccountRoutingsParams) SetAccountRoutingCreationRequest(accountRoutingCreationRequest *models.AccountRoutingCreation) {
	o.AccountRoutingCreationRequest = accountRoutingCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountRoutingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountRoutingCreationRequest != nil {
		if err := r.SetBodyParam(o.AccountRoutingCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
