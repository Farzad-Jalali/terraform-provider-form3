// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PatchUsersUserIDReader is a Reader for the PatchUsersUserID structure.
type PatchUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUsersUserIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchUsersUserIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUsersUserIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUsersUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUsersUserIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchUsersUserIDOK creates a PatchUsersUserIDOK with default headers values
func NewPatchUsersUserIDOK() *PatchUsersUserIDOK {
	return &PatchUsersUserIDOK{}
}

/*PatchUsersUserIDOK handles this case with default header values.

User details
*/
type PatchUsersUserIDOK struct {
	Payload *models.UserDetailsResponse
}

func (o *PatchUsersUserIDOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *PatchUsersUserIDOK) GetPayload() *models.UserDetailsResponse {
	return o.Payload
}

func (o *PatchUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDBadRequest creates a PatchUsersUserIDBadRequest with default headers values
func NewPatchUsersUserIDBadRequest() *PatchUsersUserIDBadRequest {
	return &PatchUsersUserIDBadRequest{}
}

/*PatchUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchUsersUserIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUsersUserIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDUnauthorized creates a PatchUsersUserIDUnauthorized with default headers values
func NewPatchUsersUserIDUnauthorized() *PatchUsersUserIDUnauthorized {
	return &PatchUsersUserIDUnauthorized{}
}

/*PatchUsersUserIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchUsersUserIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUsersUserIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDForbidden creates a PatchUsersUserIDForbidden with default headers values
func NewPatchUsersUserIDForbidden() *PatchUsersUserIDForbidden {
	return &PatchUsersUserIDForbidden{}
}

/*PatchUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type PatchUsersUserIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchUsersUserIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDNotFound creates a PatchUsersUserIDNotFound with default headers values
func NewPatchUsersUserIDNotFound() *PatchUsersUserIDNotFound {
	return &PatchUsersUserIDNotFound{}
}

/*PatchUsersUserIDNotFound handles this case with default header values.

Record not found
*/
type PatchUsersUserIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchUsersUserIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDConflict creates a PatchUsersUserIDConflict with default headers values
func NewPatchUsersUserIDConflict() *PatchUsersUserIDConflict {
	return &PatchUsersUserIDConflict{}
}

/*PatchUsersUserIDConflict handles this case with default header values.

Conflict
*/
type PatchUsersUserIDConflict struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdConflict  %+v", 409, o.Payload)
}

func (o *PatchUsersUserIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDTooManyRequests creates a PatchUsersUserIDTooManyRequests with default headers values
func NewPatchUsersUserIDTooManyRequests() *PatchUsersUserIDTooManyRequests {
	return &PatchUsersUserIDTooManyRequests{}
}

/*PatchUsersUserIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchUsersUserIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUsersUserIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDInternalServerError creates a PatchUsersUserIDInternalServerError with default headers values
func NewPatchUsersUserIDInternalServerError() *PatchUsersUserIDInternalServerError {
	return &PatchUsersUserIDInternalServerError{}
}

/*PatchUsersUserIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchUsersUserIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUsersUserIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersUserIDServiceUnavailable creates a PatchUsersUserIDServiceUnavailable with default headers values
func NewPatchUsersUserIDServiceUnavailable() *PatchUsersUserIDServiceUnavailable {
	return &PatchUsersUserIDServiceUnavailable{}
}

/*PatchUsersUserIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchUsersUserIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchUsersUserIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUsersUserIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchUsersUserIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
