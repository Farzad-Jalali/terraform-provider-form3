// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// PostAccountsReader is a Reader for the PostAccounts structure.
type PostAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountsCreated creates a PostAccountsCreated with default headers values
func NewPostAccountsCreated() *PostAccountsCreated {
	return &PostAccountsCreated{}
}

/*PostAccountsCreated handles this case with default header values.

Account creation response
*/
type PostAccountsCreated struct {
	Payload *models.AccountCreationResponse
}

func (o *PostAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] postAccountsCreated  %+v", 201, o.Payload)
}

func (o *PostAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
