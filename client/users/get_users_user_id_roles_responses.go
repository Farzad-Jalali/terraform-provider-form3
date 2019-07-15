// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// GetUsersUserIDRolesReader is a Reader for the GetUsersUserIDRoles structure.
type GetUsersUserIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDRolesOK creates a GetUsersUserIDRolesOK with default headers values
func NewGetUsersUserIDRolesOK() *GetUsersUserIDRolesOK {
	return &GetUsersUserIDRolesOK{}
}

/*GetUsersUserIDRolesOK handles this case with default header values.

List of roles for user
*/
type GetUsersUserIDRolesOK struct {
	Payload *models.UserRoleListResponse
}

func (o *GetUsersUserIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
