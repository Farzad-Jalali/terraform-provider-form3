// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetRolesRoleIDReader is a Reader for the GetRolesRoleID structure.
type GetRolesRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRolesRoleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRolesRoleIDOK creates a GetRolesRoleIDOK with default headers values
func NewGetRolesRoleIDOK() *GetRolesRoleIDOK {
	return &GetRolesRoleIDOK{}
}

/*GetRolesRoleIDOK handles this case with default header values.

Role details
*/
type GetRolesRoleIDOK struct {
	Payload *models.RoleDetailsResponse
}

func (o *GetRolesRoleIDOK) Error() string {
	return fmt.Sprintf("[GET /roles/{role_id}][%d] getRolesRoleIdOK  %+v", 200, o.Payload)
}

func (o *GetRolesRoleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
