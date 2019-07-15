// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// GetPositionsReader is a Reader for the GetPositions structure.
type GetPositionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPositionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPositionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPositionsOK creates a GetPositionsOK with default headers values
func NewGetPositionsOK() *GetPositionsOK {
	return &GetPositionsOK{}
}

/*GetPositionsOK handles this case with default header values.

List of position details
*/
type GetPositionsOK struct {
	Payload *models.PositionDetailsListResponse
}

func (o *GetPositionsOK) Error() string {
	return fmt.Sprintf("[GET /positions][%d] getPositionsOK  %+v", 200, o.Payload)
}

func (o *GetPositionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PositionDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
