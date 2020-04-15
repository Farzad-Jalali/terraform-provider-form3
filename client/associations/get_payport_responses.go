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

// GetPayportReader is a Reader for the GetPayport structure.
type GetPayportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPayportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPayportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPayportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPayportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPayportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPayportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPayportConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPayportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPayportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPayportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPayportOK creates a GetPayportOK with default headers values
func NewGetPayportOK() *GetPayportOK {
	return &GetPayportOK{}
}

/*GetPayportOK handles this case with default header values.

List of associations
*/
type GetPayportOK struct {
	Payload *models.PayportAssociationDetailsListResponse
}

func (o *GetPayportOK) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportOK  %+v", 200, o.Payload)
}

func (o *GetPayportOK) GetPayload() *models.PayportAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetPayportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportBadRequest creates a GetPayportBadRequest with default headers values
func NewGetPayportBadRequest() *GetPayportBadRequest {
	return &GetPayportBadRequest{}
}

/*GetPayportBadRequest handles this case with default header values.

Bad Request
*/
type GetPayportBadRequest struct {
	Payload *models.APIError
}

func (o *GetPayportBadRequest) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportBadRequest  %+v", 400, o.Payload)
}

func (o *GetPayportBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportUnauthorized creates a GetPayportUnauthorized with default headers values
func NewGetPayportUnauthorized() *GetPayportUnauthorized {
	return &GetPayportUnauthorized{}
}

/*GetPayportUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetPayportUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPayportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPayportUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportForbidden creates a GetPayportForbidden with default headers values
func NewGetPayportForbidden() *GetPayportForbidden {
	return &GetPayportForbidden{}
}

/*GetPayportForbidden handles this case with default header values.

Forbidden
*/
type GetPayportForbidden struct {
	Payload *models.APIError
}

func (o *GetPayportForbidden) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportForbidden  %+v", 403, o.Payload)
}

func (o *GetPayportForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportNotFound creates a GetPayportNotFound with default headers values
func NewGetPayportNotFound() *GetPayportNotFound {
	return &GetPayportNotFound{}
}

/*GetPayportNotFound handles this case with default header values.

Record not found
*/
type GetPayportNotFound struct {
	Payload *models.APIError
}

func (o *GetPayportNotFound) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportNotFound  %+v", 404, o.Payload)
}

func (o *GetPayportNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportConflict creates a GetPayportConflict with default headers values
func NewGetPayportConflict() *GetPayportConflict {
	return &GetPayportConflict{}
}

/*GetPayportConflict handles this case with default header values.

Conflict
*/
type GetPayportConflict struct {
	Payload *models.APIError
}

func (o *GetPayportConflict) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportConflict  %+v", 409, o.Payload)
}

func (o *GetPayportConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportTooManyRequests creates a GetPayportTooManyRequests with default headers values
func NewGetPayportTooManyRequests() *GetPayportTooManyRequests {
	return &GetPayportTooManyRequests{}
}

/*GetPayportTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPayportTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPayportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPayportTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportInternalServerError creates a GetPayportInternalServerError with default headers values
func NewGetPayportInternalServerError() *GetPayportInternalServerError {
	return &GetPayportInternalServerError{}
}

/*GetPayportInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPayportInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPayportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPayportInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPayportServiceUnavailable creates a GetPayportServiceUnavailable with default headers values
func NewGetPayportServiceUnavailable() *GetPayportServiceUnavailable {
	return &GetPayportServiceUnavailable{}
}

/*GetPayportServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPayportServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPayportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPayportServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPayportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
