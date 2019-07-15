// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// GetLimitsIDReader is a Reader for the GetLimitsID structure.
type GetLimitsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLimitsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLimitsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLimitsIDOK creates a GetLimitsIDOK with default headers values
func NewGetLimitsIDOK() *GetLimitsIDOK {
	return &GetLimitsIDOK{}
}

/*GetLimitsIDOK handles this case with default header values.

Limit details
*/
type GetLimitsIDOK struct {
	Payload *models.LimitDetailsResponse
}

func (o *GetLimitsIDOK) Error() string {
	return fmt.Sprintf("[GET /limits/{id}][%d] getLimitsIdOK  %+v", 200, o.Payload)
}

func (o *GetLimitsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}