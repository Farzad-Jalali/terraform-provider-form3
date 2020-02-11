// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostSepareconciliationReader is a Reader for the PostSepareconciliation structure.
type PostSepareconciliationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSepareconciliationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSepareconciliationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSepareconciliationCreated creates a PostSepareconciliationCreated with default headers values
func NewPostSepareconciliationCreated() *PostSepareconciliationCreated {
	return &PostSepareconciliationCreated{}
}

/*PostSepareconciliationCreated handles this case with default header values.

creation response
*/
type PostSepareconciliationCreated struct {
	Payload *models.SepaReconciliationAssociationCreationResponse
}

func (o *PostSepareconciliationCreated) Error() string {
	return fmt.Sprintf("[POST /separeconciliation][%d] postSepareconciliationCreated  %+v", 201, o.Payload)
}

func (o *PostSepareconciliationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaReconciliationAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}