// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetPaymentsIDReversalsReversalIDReader is a Reader for the GetPaymentsIDReversalsReversalID structure.
type GetPaymentsIDReversalsReversalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDReversalsReversalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDReversalsReversalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDReversalsReversalIDOK creates a GetPaymentsIDReversalsReversalIDOK with default headers values
func NewGetPaymentsIDReversalsReversalIDOK() *GetPaymentsIDReversalsReversalIDOK {
	return &GetPaymentsIDReversalsReversalIDOK{}
}

/*GetPaymentsIDReversalsReversalIDOK handles this case with default header values.

Reversal details
*/
type GetPaymentsIDReversalsReversalIDOK struct {
	Payload *models.ReversalDetailsResponse
}

func (o *GetPaymentsIDReversalsReversalIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDReversalsReversalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
