// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// GetPayportIDReader is a Reader for the GetPayportID structure.
type GetPayportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPayportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPayportIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPayportIDOK creates a GetPayportIDOK with default headers values
func NewGetPayportIDOK() *GetPayportIDOK {
	return &GetPayportIDOK{}
}

/*GetPayportIDOK handles this case with default header values.

Associations details
*/
type GetPayportIDOK struct {
	Payload *models.PayportAssociationDetailsResponse
}

func (o *GetPayportIDOK) Error() string {
	return fmt.Sprintf("[GET /payport/{id}][%d] getPayportIdOK  %+v", 200, o.Payload)
}

func (o *GetPayportIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
