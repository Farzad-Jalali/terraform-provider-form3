// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/models"
)

// GetSepasctReader is a Reader for the GetSepasct structure.
type GetSepasctReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepasctReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSepasctOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepasctOK creates a GetSepasctOK with default headers values
func NewGetSepasctOK() *GetSepasctOK {
	return &GetSepasctOK{}
}

/*GetSepasctOK handles this case with default header values.

List of associations
*/
type GetSepasctOK struct {
	Payload *models.SepaSctAssociationDetailsListResponse
}

func (o *GetSepasctOK) Error() string {
	return fmt.Sprintf("[GET /sepasct][%d] getSepasctOK  %+v", 200, o.Payload)
}

func (o *GetSepasctOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaSctAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
