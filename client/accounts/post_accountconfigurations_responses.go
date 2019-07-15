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

// PostAccountconfigurationsReader is a Reader for the PostAccountconfigurations structure.
type PostAccountconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAccountconfigurationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountconfigurationsCreated creates a PostAccountconfigurationsCreated with default headers values
func NewPostAccountconfigurationsCreated() *PostAccountconfigurationsCreated {
	return &PostAccountconfigurationsCreated{}
}

/*PostAccountconfigurationsCreated handles this case with default header values.

AccountConfiguration creation response
*/
type PostAccountconfigurationsCreated struct {
	Payload *models.AccountConfigurationCreationResponse
}

func (o *PostAccountconfigurationsCreated) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsCreated  %+v", 201, o.Payload)
}

func (o *PostAccountconfigurationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountConfigurationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
