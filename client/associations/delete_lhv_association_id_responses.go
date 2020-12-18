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

// DeleteLhvAssociationIDReader is a Reader for the DeleteLhvAssociationID structure.
type DeleteLhvAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLhvAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLhvAssociationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLhvAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteLhvAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteLhvAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLhvAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteLhvAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteLhvAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteLhvAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteLhvAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLhvAssociationIDNoContent creates a DeleteLhvAssociationIDNoContent with default headers values
func NewDeleteLhvAssociationIDNoContent() *DeleteLhvAssociationIDNoContent {
	return &DeleteLhvAssociationIDNoContent{}
}

/*DeleteLhvAssociationIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteLhvAssociationIDNoContent struct {
}

func (o *DeleteLhvAssociationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdNoContent ", 204)
}

func (o *DeleteLhvAssociationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLhvAssociationIDBadRequest creates a DeleteLhvAssociationIDBadRequest with default headers values
func NewDeleteLhvAssociationIDBadRequest() *DeleteLhvAssociationIDBadRequest {
	return &DeleteLhvAssociationIDBadRequest{}
}

/*DeleteLhvAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLhvAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLhvAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDUnauthorized creates a DeleteLhvAssociationIDUnauthorized with default headers values
func NewDeleteLhvAssociationIDUnauthorized() *DeleteLhvAssociationIDUnauthorized {
	return &DeleteLhvAssociationIDUnauthorized{}
}

/*DeleteLhvAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteLhvAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLhvAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDForbidden creates a DeleteLhvAssociationIDForbidden with default headers values
func NewDeleteLhvAssociationIDForbidden() *DeleteLhvAssociationIDForbidden {
	return &DeleteLhvAssociationIDForbidden{}
}

/*DeleteLhvAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLhvAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLhvAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDNotFound creates a DeleteLhvAssociationIDNotFound with default headers values
func NewDeleteLhvAssociationIDNotFound() *DeleteLhvAssociationIDNotFound {
	return &DeleteLhvAssociationIDNotFound{}
}

/*DeleteLhvAssociationIDNotFound handles this case with default header values.

Record not found
*/
type DeleteLhvAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLhvAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDConflict creates a DeleteLhvAssociationIDConflict with default headers values
func NewDeleteLhvAssociationIDConflict() *DeleteLhvAssociationIDConflict {
	return &DeleteLhvAssociationIDConflict{}
}

/*DeleteLhvAssociationIDConflict handles this case with default header values.

Conflict
*/
type DeleteLhvAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteLhvAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDTooManyRequests creates a DeleteLhvAssociationIDTooManyRequests with default headers values
func NewDeleteLhvAssociationIDTooManyRequests() *DeleteLhvAssociationIDTooManyRequests {
	return &DeleteLhvAssociationIDTooManyRequests{}
}

/*DeleteLhvAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteLhvAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLhvAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDInternalServerError creates a DeleteLhvAssociationIDInternalServerError with default headers values
func NewDeleteLhvAssociationIDInternalServerError() *DeleteLhvAssociationIDInternalServerError {
	return &DeleteLhvAssociationIDInternalServerError{}
}

/*DeleteLhvAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteLhvAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLhvAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLhvAssociationIDServiceUnavailable creates a DeleteLhvAssociationIDServiceUnavailable with default headers values
func NewDeleteLhvAssociationIDServiceUnavailable() *DeleteLhvAssociationIDServiceUnavailable {
	return &DeleteLhvAssociationIDServiceUnavailable{}
}

/*DeleteLhvAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteLhvAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteLhvAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /lhv/{associationId}][%d] deleteLhvAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLhvAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
