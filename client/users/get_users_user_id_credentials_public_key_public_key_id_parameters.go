// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParams creates a new GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized.
func NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParams() *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithTimeout creates a new GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithTimeout(timeout time.Duration) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		timeout: timeout,
	}
}

// NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithContext creates a new GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithContext(ctx context.Context) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		Context: ctx,
	}
}

// NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithHTTPClient creates a new GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithHTTPClient(client *http.Client) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams{
		HTTPClient: client,
	}
}

/*GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams contains all the parameters to send to the API endpoint
for the get users user ID credentials public key public key ID operation typically these are written to a http.Request
*/
type GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams struct {

	/*PublicKeyID
	  public_key_id

	*/
	PublicKeyID strfmt.UUID
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithTimeout(timeout time.Duration) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithContext(ctx context.Context) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithHTTPClient(client *http.Client) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublicKeyID adds the publicKeyID to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithPublicKeyID(publicKeyID strfmt.UUID) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetPublicKeyID(publicKeyID)
	return o
}

// SetPublicKeyID adds the publicKeyId to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetPublicKeyID(publicKeyID strfmt.UUID) {
	o.PublicKeyID = publicKeyID
}

// WithUserID adds the userID to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithUserID(userID strfmt.UUID) *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID credentials public key public key ID params
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param public_key_id
	if err := r.SetPathParam("public_key_id", o.PublicKeyID.String()); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
