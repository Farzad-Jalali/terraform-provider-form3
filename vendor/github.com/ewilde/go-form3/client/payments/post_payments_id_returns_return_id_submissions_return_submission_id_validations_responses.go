// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader is a Reader for the PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidations structure.
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated handles this case with default header values.

Return submission validation creation response
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated struct {
	Payload *models.ReturnSubmissionValidationCreationResponse
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnSubmissionValidationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest handles this case with default header values.

Return submission validation creation error
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
