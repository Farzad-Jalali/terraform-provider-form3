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

// GetStarlingIDReader is a Reader for the GetStarlingID structure.
type GetStarlingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStarlingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStarlingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStarlingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStarlingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStarlingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStarlingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetStarlingIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStarlingIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStarlingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetStarlingIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarlingIDOK creates a GetStarlingIDOK with default headers values
func NewGetStarlingIDOK() *GetStarlingIDOK {
	return &GetStarlingIDOK{}
}

/*GetStarlingIDOK handles this case with default header values.

Associations details
*/
type GetStarlingIDOK struct {
	Payload *models.AssociationDetailsResponse
}

func (o *GetStarlingIDOK) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdOK  %+v", 200, o.Payload)
}

func (o *GetStarlingIDOK) GetPayload() *models.AssociationDetailsResponse {
	return o.Payload
}

func (o *GetStarlingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDBadRequest creates a GetStarlingIDBadRequest with default headers values
func NewGetStarlingIDBadRequest() *GetStarlingIDBadRequest {
	return &GetStarlingIDBadRequest{}
}

/*GetStarlingIDBadRequest handles this case with default header values.

Bad Request
*/
type GetStarlingIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetStarlingIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetStarlingIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDUnauthorized creates a GetStarlingIDUnauthorized with default headers values
func NewGetStarlingIDUnauthorized() *GetStarlingIDUnauthorized {
	return &GetStarlingIDUnauthorized{}
}

/*GetStarlingIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetStarlingIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetStarlingIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStarlingIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDForbidden creates a GetStarlingIDForbidden with default headers values
func NewGetStarlingIDForbidden() *GetStarlingIDForbidden {
	return &GetStarlingIDForbidden{}
}

/*GetStarlingIDForbidden handles this case with default header values.

Forbidden
*/
type GetStarlingIDForbidden struct {
	Payload *models.APIError
}

func (o *GetStarlingIDForbidden) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdForbidden  %+v", 403, o.Payload)
}

func (o *GetStarlingIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDNotFound creates a GetStarlingIDNotFound with default headers values
func NewGetStarlingIDNotFound() *GetStarlingIDNotFound {
	return &GetStarlingIDNotFound{}
}

/*GetStarlingIDNotFound handles this case with default header values.

Record not found
*/
type GetStarlingIDNotFound struct {
	Payload *models.APIError
}

func (o *GetStarlingIDNotFound) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdNotFound  %+v", 404, o.Payload)
}

func (o *GetStarlingIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDConflict creates a GetStarlingIDConflict with default headers values
func NewGetStarlingIDConflict() *GetStarlingIDConflict {
	return &GetStarlingIDConflict{}
}

/*GetStarlingIDConflict handles this case with default header values.

Conflict
*/
type GetStarlingIDConflict struct {
	Payload *models.APIError
}

func (o *GetStarlingIDConflict) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdConflict  %+v", 409, o.Payload)
}

func (o *GetStarlingIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDTooManyRequests creates a GetStarlingIDTooManyRequests with default headers values
func NewGetStarlingIDTooManyRequests() *GetStarlingIDTooManyRequests {
	return &GetStarlingIDTooManyRequests{}
}

/*GetStarlingIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetStarlingIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetStarlingIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStarlingIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDInternalServerError creates a GetStarlingIDInternalServerError with default headers values
func NewGetStarlingIDInternalServerError() *GetStarlingIDInternalServerError {
	return &GetStarlingIDInternalServerError{}
}

/*GetStarlingIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStarlingIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetStarlingIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStarlingIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStarlingIDServiceUnavailable creates a GetStarlingIDServiceUnavailable with default headers values
func NewGetStarlingIDServiceUnavailable() *GetStarlingIDServiceUnavailable {
	return &GetStarlingIDServiceUnavailable{}
}

/*GetStarlingIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetStarlingIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetStarlingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /starling/{id}][%d] getStarlingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStarlingIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetStarlingIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
