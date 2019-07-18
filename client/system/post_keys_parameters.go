// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostKeysParams creates a new PostKeysParams object
// with the default values initialized.
func NewPostKeysParams() *PostKeysParams {
	var ()
	return &PostKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKeysParamsWithTimeout creates a new PostKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKeysParamsWithTimeout(timeout time.Duration) *PostKeysParams {
	var ()
	return &PostKeysParams{

		timeout: timeout,
	}
}

// NewPostKeysParamsWithContext creates a new PostKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKeysParamsWithContext(ctx context.Context) *PostKeysParams {
	var ()
	return &PostKeysParams{

		Context: ctx,
	}
}

// NewPostKeysParamsWithHTTPClient creates a new PostKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKeysParamsWithHTTPClient(client *http.Client) *PostKeysParams {
	var ()
	return &PostKeysParams{
		HTTPClient: client,
	}
}

/*PostKeysParams contains all the parameters to send to the API endpoint
for the post keys operation typically these are written to a http.Request
*/
type PostKeysParams struct {

	/*KeyCreationRequest*/
	KeyCreationRequest *models.KeyCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post keys params
func (o *PostKeysParams) WithTimeout(timeout time.Duration) *PostKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post keys params
func (o *PostKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post keys params
func (o *PostKeysParams) WithContext(ctx context.Context) *PostKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post keys params
func (o *PostKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post keys params
func (o *PostKeysParams) WithHTTPClient(client *http.Client) *PostKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post keys params
func (o *PostKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyCreationRequest adds the keyCreationRequest to the post keys params
func (o *PostKeysParams) WithKeyCreationRequest(keyCreationRequest *models.KeyCreation) *PostKeysParams {
	o.SetKeyCreationRequest(keyCreationRequest)
	return o
}

// SetKeyCreationRequest adds the keyCreationRequest to the post keys params
func (o *PostKeysParams) SetKeyCreationRequest(keyCreationRequest *models.KeyCreation) {
	o.KeyCreationRequest = keyCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KeyCreationRequest != nil {
		if err := r.SetBodyParam(o.KeyCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
