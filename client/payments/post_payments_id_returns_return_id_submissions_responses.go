// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostPaymentsIDReturnsReturnIDSubmissionsReader is a Reader for the PostPaymentsIDReturnsReturnIDSubmissions structure.
type PostPaymentsIDReturnsReturnIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReturnIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsCreated creates a PostPaymentsIDReturnsReturnIDSubmissionsCreated with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsCreated() *PostPaymentsIDReturnsReturnIDSubmissionsCreated {
	return &PostPaymentsIDReturnsReturnIDSubmissionsCreated{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsCreated handles this case with default header values.

Return submission creation response
*/
type PostPaymentsIDReturnsReturnIDSubmissionsCreated struct {
	Payload *models.ReturnSubmissionCreationResponse
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsCreated) GetPayload() *models.ReturnSubmissionCreationResponse {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnSubmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsBadRequest creates a PostPaymentsIDReturnsReturnIDSubmissionsBadRequest with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsBadRequest() *PostPaymentsIDReturnsReturnIDSubmissionsBadRequest {
	return &PostPaymentsIDReturnsReturnIDSubmissionsBadRequest{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsIDReturnsReturnIDSubmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsUnauthorized creates a PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsUnauthorized() *PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized {
	return &PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsForbidden creates a PostPaymentsIDReturnsReturnIDSubmissionsForbidden with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsForbidden() *PostPaymentsIDReturnsReturnIDSubmissionsForbidden {
	return &PostPaymentsIDReturnsReturnIDSubmissionsForbidden{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsIDReturnsReturnIDSubmissionsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsNotFound creates a PostPaymentsIDReturnsReturnIDSubmissionsNotFound with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsNotFound() *PostPaymentsIDReturnsReturnIDSubmissionsNotFound {
	return &PostPaymentsIDReturnsReturnIDSubmissionsNotFound{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsIDReturnsReturnIDSubmissionsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsConflict creates a PostPaymentsIDReturnsReturnIDSubmissionsConflict with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsConflict() *PostPaymentsIDReturnsReturnIDSubmissionsConflict {
	return &PostPaymentsIDReturnsReturnIDSubmissionsConflict{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsIDReturnsReturnIDSubmissionsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsConflict) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests creates a PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests() *PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests {
	return &PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsInternalServerError creates a PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsInternalServerError() *PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError {
	return &PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable creates a PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable() *PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable {
	return &PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions][%d] postPaymentsIdReturnsReturnIdSubmissionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
