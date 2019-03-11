// Code generated by go-swagger; DO NOT EDIT.

package ace

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

// NewDeleteRolesRoleIDAcesAceIDParams creates a new DeleteRolesRoleIDAcesAceIDParams object
// with the default values initialized.
func NewDeleteRolesRoleIDAcesAceIDParams() *DeleteRolesRoleIDAcesAceIDParams {
	var ()
	return &DeleteRolesRoleIDAcesAceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRolesRoleIDAcesAceIDParamsWithTimeout creates a new DeleteRolesRoleIDAcesAceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRolesRoleIDAcesAceIDParamsWithTimeout(timeout time.Duration) *DeleteRolesRoleIDAcesAceIDParams {
	var ()
	return &DeleteRolesRoleIDAcesAceIDParams{

		timeout: timeout,
	}
}

// NewDeleteRolesRoleIDAcesAceIDParamsWithContext creates a new DeleteRolesRoleIDAcesAceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRolesRoleIDAcesAceIDParamsWithContext(ctx context.Context) *DeleteRolesRoleIDAcesAceIDParams {
	var ()
	return &DeleteRolesRoleIDAcesAceIDParams{

		Context: ctx,
	}
}

// NewDeleteRolesRoleIDAcesAceIDParamsWithHTTPClient creates a new DeleteRolesRoleIDAcesAceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRolesRoleIDAcesAceIDParamsWithHTTPClient(client *http.Client) *DeleteRolesRoleIDAcesAceIDParams {
	var ()
	return &DeleteRolesRoleIDAcesAceIDParams{
		HTTPClient: client,
	}
}

/*DeleteRolesRoleIDAcesAceIDParams contains all the parameters to send to the API endpoint
for the delete roles role ID aces ace ID operation typically these are written to a http.Request
*/
type DeleteRolesRoleIDAcesAceIDParams struct {

	/*AceID
	  Ace Id

	*/
	AceID strfmt.UUID
	/*RoleID
	  Role Id

	*/
	RoleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) WithTimeout(timeout time.Duration) *DeleteRolesRoleIDAcesAceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) WithContext(ctx context.Context) *DeleteRolesRoleIDAcesAceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) WithHTTPClient(client *http.Client) *DeleteRolesRoleIDAcesAceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAceID adds the aceID to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) WithAceID(aceID strfmt.UUID) *DeleteRolesRoleIDAcesAceIDParams {
	o.SetAceID(aceID)
	return o
}

// SetAceID adds the aceId to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) SetAceID(aceID strfmt.UUID) {
	o.AceID = aceID
}

// WithRoleID adds the roleID to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) WithRoleID(roleID strfmt.UUID) *DeleteRolesRoleIDAcesAceIDParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the delete roles role ID aces ace ID params
func (o *DeleteRolesRoleIDAcesAceIDParams) SetRoleID(roleID strfmt.UUID) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRolesRoleIDAcesAceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ace_id
	if err := r.SetPathParam("ace_id", o.AceID.String()); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
