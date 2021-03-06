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

// GetProductsIDReader is a Reader for the GetProductsID structure.
type GetProductsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProductsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProductsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProductsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProductsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProductsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetProductsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProductsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetProductsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductsIDOK creates a GetProductsIDOK with default headers values
func NewGetProductsIDOK() *GetProductsIDOK {
	return &GetProductsIDOK{}
}

/*GetProductsIDOK handles this case with default header values.

Associations details
*/
type GetProductsIDOK struct {
	Payload *models.ProductsAssociationDetailsResponse
}

func (o *GetProductsIDOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdOK  %+v", 200, o.Payload)
}

func (o *GetProductsIDOK) GetPayload() *models.ProductsAssociationDetailsResponse {
	return o.Payload
}

func (o *GetProductsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductsAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDBadRequest creates a GetProductsIDBadRequest with default headers values
func NewGetProductsIDBadRequest() *GetProductsIDBadRequest {
	return &GetProductsIDBadRequest{}
}

/*GetProductsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetProductsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetProductsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetProductsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDUnauthorized creates a GetProductsIDUnauthorized with default headers values
func NewGetProductsIDUnauthorized() *GetProductsIDUnauthorized {
	return &GetProductsIDUnauthorized{}
}

/*GetProductsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetProductsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetProductsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProductsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDForbidden creates a GetProductsIDForbidden with default headers values
func NewGetProductsIDForbidden() *GetProductsIDForbidden {
	return &GetProductsIDForbidden{}
}

/*GetProductsIDForbidden handles this case with default header values.

Forbidden
*/
type GetProductsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetProductsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetProductsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDNotFound creates a GetProductsIDNotFound with default headers values
func NewGetProductsIDNotFound() *GetProductsIDNotFound {
	return &GetProductsIDNotFound{}
}

/*GetProductsIDNotFound handles this case with default header values.

Record not found
*/
type GetProductsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetProductsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetProductsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDConflict creates a GetProductsIDConflict with default headers values
func NewGetProductsIDConflict() *GetProductsIDConflict {
	return &GetProductsIDConflict{}
}

/*GetProductsIDConflict handles this case with default header values.

Conflict
*/
type GetProductsIDConflict struct {
	Payload *models.APIError
}

func (o *GetProductsIDConflict) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdConflict  %+v", 409, o.Payload)
}

func (o *GetProductsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDTooManyRequests creates a GetProductsIDTooManyRequests with default headers values
func NewGetProductsIDTooManyRequests() *GetProductsIDTooManyRequests {
	return &GetProductsIDTooManyRequests{}
}

/*GetProductsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetProductsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetProductsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProductsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDInternalServerError creates a GetProductsIDInternalServerError with default headers values
func NewGetProductsIDInternalServerError() *GetProductsIDInternalServerError {
	return &GetProductsIDInternalServerError{}
}

/*GetProductsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetProductsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetProductsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProductsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsIDServiceUnavailable creates a GetProductsIDServiceUnavailable with default headers values
func NewGetProductsIDServiceUnavailable() *GetProductsIDServiceUnavailable {
	return &GetProductsIDServiceUnavailable{}
}

/*GetProductsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetProductsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetProductsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProductsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetProductsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
