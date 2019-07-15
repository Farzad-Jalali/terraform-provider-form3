// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// GetConfirmationOfPayeeReader is a Reader for the GetConfirmationOfPayee structure.
type GetConfirmationOfPayeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfirmationOfPayeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfirmationOfPayeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfirmationOfPayeeOK creates a GetConfirmationOfPayeeOK with default headers values
func NewGetConfirmationOfPayeeOK() *GetConfirmationOfPayeeOK {
	return &GetConfirmationOfPayeeOK{}
}

/*GetConfirmationOfPayeeOK handles this case with default header values.

List of associations
*/
type GetConfirmationOfPayeeOK struct {
	Payload *models.CoPAssociationDetailsListResponse
}

func (o *GetConfirmationOfPayeeOK) Error() string {
	return fmt.Sprintf("[GET /confirmation-of-payee][%d] getConfirmationOfPayeeOK  %+v", 200, o.Payload)
}

func (o *GetConfirmationOfPayeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoPAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
