// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostPaymentsIDAdmissionsReader is a Reader for the PostPaymentsIDAdmissions structure.
type PostPaymentsIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDAdmissionsCreated creates a PostPaymentsIDAdmissionsCreated with default headers values
func NewPostPaymentsIDAdmissionsCreated() *PostPaymentsIDAdmissionsCreated {
	return &PostPaymentsIDAdmissionsCreated{}
}

/*PostPaymentsIDAdmissionsCreated handles this case with default header values.

Admission creation response
*/
type PostPaymentsIDAdmissionsCreated struct {
	Payload *models.PaymentAdmissionCreationResponse
}

func (o *PostPaymentsIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/admissions][%d] postPaymentsIdAdmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentAdmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDAdmissionsBadRequest creates a PostPaymentsIDAdmissionsBadRequest with default headers values
func NewPostPaymentsIDAdmissionsBadRequest() *PostPaymentsIDAdmissionsBadRequest {
	return &PostPaymentsIDAdmissionsBadRequest{}
}

/*PostPaymentsIDAdmissionsBadRequest handles this case with default header values.

Admission creation error
*/
type PostPaymentsIDAdmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/admissions][%d] postPaymentsIdAdmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
