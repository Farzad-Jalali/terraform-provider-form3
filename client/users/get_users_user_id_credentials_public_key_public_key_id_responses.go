// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetUsersUserIDCredentialsPublicKeyPublicKeyIDReader is a Reader for the GetUsersUserIDCredentialsPublicKeyPublicKeyID structure.
type GetUsersUserIDCredentialsPublicKeyPublicKeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDOK creates a GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK with default headers values
func NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDOK() *GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK {
	return &GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK{}
}

/*GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK handles this case with default header values.

Public key data
*/
type GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK struct {
	Payload *models.PublicKey
}

func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/public_key/{public_key_id}][%d] getUsersUserIdCredentialsPublicKeyPublicKeyIdOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDCredentialsPublicKeyPublicKeyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}