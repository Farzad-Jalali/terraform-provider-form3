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

// GetSepaReconciliationIDReader is a Reader for the GetSepaReconciliationID structure.
type GetSepaReconciliationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepaReconciliationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSepaReconciliationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSepaReconciliationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSepaReconciliationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSepaReconciliationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSepaReconciliationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetSepaReconciliationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSepaReconciliationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSepaReconciliationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSepaReconciliationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepaReconciliationIDOK creates a GetSepaReconciliationIDOK with default headers values
func NewGetSepaReconciliationIDOK() *GetSepaReconciliationIDOK {
	return &GetSepaReconciliationIDOK{}
}

/*GetSepaReconciliationIDOK handles this case with default header values.

Associations details
*/
type GetSepaReconciliationIDOK struct {
	Payload *models.SepaReconciliationAssociationDetailsResponse
}

func (o *GetSepaReconciliationIDOK) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdOK  %+v", 200, o.Payload)
}

func (o *GetSepaReconciliationIDOK) GetPayload() *models.SepaReconciliationAssociationDetailsResponse {
	return o.Payload
}

func (o *GetSepaReconciliationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaReconciliationAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDBadRequest creates a GetSepaReconciliationIDBadRequest with default headers values
func NewGetSepaReconciliationIDBadRequest() *GetSepaReconciliationIDBadRequest {
	return &GetSepaReconciliationIDBadRequest{}
}

/*GetSepaReconciliationIDBadRequest handles this case with default header values.

Bad Request
*/
type GetSepaReconciliationIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSepaReconciliationIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDUnauthorized creates a GetSepaReconciliationIDUnauthorized with default headers values
func NewGetSepaReconciliationIDUnauthorized() *GetSepaReconciliationIDUnauthorized {
	return &GetSepaReconciliationIDUnauthorized{}
}

/*GetSepaReconciliationIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetSepaReconciliationIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSepaReconciliationIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDForbidden creates a GetSepaReconciliationIDForbidden with default headers values
func NewGetSepaReconciliationIDForbidden() *GetSepaReconciliationIDForbidden {
	return &GetSepaReconciliationIDForbidden{}
}

/*GetSepaReconciliationIDForbidden handles this case with default header values.

Forbidden
*/
type GetSepaReconciliationIDForbidden struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdForbidden  %+v", 403, o.Payload)
}

func (o *GetSepaReconciliationIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDNotFound creates a GetSepaReconciliationIDNotFound with default headers values
func NewGetSepaReconciliationIDNotFound() *GetSepaReconciliationIDNotFound {
	return &GetSepaReconciliationIDNotFound{}
}

/*GetSepaReconciliationIDNotFound handles this case with default header values.

Record not found
*/
type GetSepaReconciliationIDNotFound struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSepaReconciliationIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDConflict creates a GetSepaReconciliationIDConflict with default headers values
func NewGetSepaReconciliationIDConflict() *GetSepaReconciliationIDConflict {
	return &GetSepaReconciliationIDConflict{}
}

/*GetSepaReconciliationIDConflict handles this case with default header values.

Conflict
*/
type GetSepaReconciliationIDConflict struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDConflict) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdConflict  %+v", 409, o.Payload)
}

func (o *GetSepaReconciliationIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDTooManyRequests creates a GetSepaReconciliationIDTooManyRequests with default headers values
func NewGetSepaReconciliationIDTooManyRequests() *GetSepaReconciliationIDTooManyRequests {
	return &GetSepaReconciliationIDTooManyRequests{}
}

/*GetSepaReconciliationIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetSepaReconciliationIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSepaReconciliationIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDInternalServerError creates a GetSepaReconciliationIDInternalServerError with default headers values
func NewGetSepaReconciliationIDInternalServerError() *GetSepaReconciliationIDInternalServerError {
	return &GetSepaReconciliationIDInternalServerError{}
}

/*GetSepaReconciliationIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSepaReconciliationIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSepaReconciliationIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaReconciliationIDServiceUnavailable creates a GetSepaReconciliationIDServiceUnavailable with default headers values
func NewGetSepaReconciliationIDServiceUnavailable() *GetSepaReconciliationIDServiceUnavailable {
	return &GetSepaReconciliationIDServiceUnavailable{}
}

/*GetSepaReconciliationIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetSepaReconciliationIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetSepaReconciliationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sepa-reconciliation/{id}][%d] getSepaReconciliationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSepaReconciliationIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaReconciliationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
