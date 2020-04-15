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

// PostPaymentsIDReversalsReader is a Reader for the PostPaymentsIDReversals structure.
type PostPaymentsIDReversalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReversalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPaymentsIDReversalsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentsIDReversalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentsIDReversalsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPaymentsIDReversalsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentsIDReversalsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentsIDReversalsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPaymentsIDReversalsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentsIDReversalsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPaymentsIDReversalsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReversalsCreated creates a PostPaymentsIDReversalsCreated with default headers values
func NewPostPaymentsIDReversalsCreated() *PostPaymentsIDReversalsCreated {
	return &PostPaymentsIDReversalsCreated{}
}

/*PostPaymentsIDReversalsCreated handles this case with default header values.

Reversal creation response
*/
type PostPaymentsIDReversalsCreated struct {
	Payload *models.ReversalCreationResponse
}

func (o *PostPaymentsIDReversalsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReversalsCreated) GetPayload() *models.ReversalCreationResponse {
	return o.Payload
}

func (o *PostPaymentsIDReversalsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsBadRequest creates a PostPaymentsIDReversalsBadRequest with default headers values
func NewPostPaymentsIDReversalsBadRequest() *PostPaymentsIDReversalsBadRequest {
	return &PostPaymentsIDReversalsBadRequest{}
}

/*PostPaymentsIDReversalsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsIDReversalsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReversalsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsUnauthorized creates a PostPaymentsIDReversalsUnauthorized with default headers values
func NewPostPaymentsIDReversalsUnauthorized() *PostPaymentsIDReversalsUnauthorized {
	return &PostPaymentsIDReversalsUnauthorized{}
}

/*PostPaymentsIDReversalsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsIDReversalsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsIDReversalsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsForbidden creates a PostPaymentsIDReversalsForbidden with default headers values
func NewPostPaymentsIDReversalsForbidden() *PostPaymentsIDReversalsForbidden {
	return &PostPaymentsIDReversalsForbidden{}
}

/*PostPaymentsIDReversalsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsIDReversalsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsIDReversalsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsNotFound creates a PostPaymentsIDReversalsNotFound with default headers values
func NewPostPaymentsIDReversalsNotFound() *PostPaymentsIDReversalsNotFound {
	return &PostPaymentsIDReversalsNotFound{}
}

/*PostPaymentsIDReversalsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsIDReversalsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsIDReversalsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsConflict creates a PostPaymentsIDReversalsConflict with default headers values
func NewPostPaymentsIDReversalsConflict() *PostPaymentsIDReversalsConflict {
	return &PostPaymentsIDReversalsConflict{}
}

/*PostPaymentsIDReversalsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsIDReversalsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsConflict) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsIDReversalsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsTooManyRequests creates a PostPaymentsIDReversalsTooManyRequests with default headers values
func NewPostPaymentsIDReversalsTooManyRequests() *PostPaymentsIDReversalsTooManyRequests {
	return &PostPaymentsIDReversalsTooManyRequests{}
}

/*PostPaymentsIDReversalsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsIDReversalsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsIDReversalsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsInternalServerError creates a PostPaymentsIDReversalsInternalServerError with default headers values
func NewPostPaymentsIDReversalsInternalServerError() *PostPaymentsIDReversalsInternalServerError {
	return &PostPaymentsIDReversalsInternalServerError{}
}

/*PostPaymentsIDReversalsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsIDReversalsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsIDReversalsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsServiceUnavailable creates a PostPaymentsIDReversalsServiceUnavailable with default headers values
func NewPostPaymentsIDReversalsServiceUnavailable() *PostPaymentsIDReversalsServiceUnavailable {
	return &PostPaymentsIDReversalsServiceUnavailable{}
}

/*PostPaymentsIDReversalsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsIDReversalsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals][%d] postPaymentsIdReversalsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsIDReversalsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReversalsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
