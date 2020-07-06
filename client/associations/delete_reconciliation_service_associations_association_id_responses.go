// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeleteReconciliationServiceAssociationsAssociationIDReader is a Reader for the DeleteReconciliationServiceAssociationsAssociationID structure.
type DeleteReconciliationServiceAssociationsAssociationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReconciliationServiceAssociationsAssociationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteReconciliationServiceAssociationsAssociationIDNoContent creates a DeleteReconciliationServiceAssociationsAssociationIDNoContent with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDNoContent() *DeleteReconciliationServiceAssociationsAssociationIDNoContent {
	return &DeleteReconciliationServiceAssociationsAssociationIDNoContent{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteReconciliationServiceAssociationsAssociationIDNoContent struct {
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdNoContent ", 204)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDBadRequest creates a DeleteReconciliationServiceAssociationsAssociationIDBadRequest with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDBadRequest() *DeleteReconciliationServiceAssociationsAssociationIDBadRequest {
	return &DeleteReconciliationServiceAssociationsAssociationIDBadRequest{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteReconciliationServiceAssociationsAssociationIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDUnauthorized creates a DeleteReconciliationServiceAssociationsAssociationIDUnauthorized with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDUnauthorized() *DeleteReconciliationServiceAssociationsAssociationIDUnauthorized {
	return &DeleteReconciliationServiceAssociationsAssociationIDUnauthorized{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteReconciliationServiceAssociationsAssociationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDForbidden creates a DeleteReconciliationServiceAssociationsAssociationIDForbidden with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDForbidden() *DeleteReconciliationServiceAssociationsAssociationIDForbidden {
	return &DeleteReconciliationServiceAssociationsAssociationIDForbidden{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteReconciliationServiceAssociationsAssociationIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDNotFound creates a DeleteReconciliationServiceAssociationsAssociationIDNotFound with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDNotFound() *DeleteReconciliationServiceAssociationsAssociationIDNotFound {
	return &DeleteReconciliationServiceAssociationsAssociationIDNotFound{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDNotFound handles this case with default header values.

Record not found
*/
type DeleteReconciliationServiceAssociationsAssociationIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDConflict creates a DeleteReconciliationServiceAssociationsAssociationIDConflict with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDConflict() *DeleteReconciliationServiceAssociationsAssociationIDConflict {
	return &DeleteReconciliationServiceAssociationsAssociationIDConflict{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDConflict handles this case with default header values.

Conflict
*/
type DeleteReconciliationServiceAssociationsAssociationIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDTooManyRequests creates a DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDTooManyRequests() *DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests {
	return &DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDInternalServerError creates a DeleteReconciliationServiceAssociationsAssociationIDInternalServerError with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDInternalServerError() *DeleteReconciliationServiceAssociationsAssociationIDInternalServerError {
	return &DeleteReconciliationServiceAssociationsAssociationIDInternalServerError{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteReconciliationServiceAssociationsAssociationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable creates a DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable with default headers values
func NewDeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable() *DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable {
	return &DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable{}
}

/*DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /reconciliation_service/associations/{associationId}][%d] deleteReconciliationServiceAssociationsAssociationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteReconciliationServiceAssociationsAssociationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
