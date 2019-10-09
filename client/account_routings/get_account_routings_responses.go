// Code generated by go-swagger; DO NOT EDIT.

package account_routings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetAccountRoutingsReader is a Reader for the GetAccountRoutings structure.
type GetAccountRoutingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountRoutingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountRoutingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountRoutingsOK creates a GetAccountRoutingsOK with default headers values
func NewGetAccountRoutingsOK() *GetAccountRoutingsOK {
	return &GetAccountRoutingsOK{}
}

/*GetAccountRoutingsOK handles this case with default header values.

List of account routing details
*/
type GetAccountRoutingsOK struct {
	Payload *models.AccountRoutingDetailsListResponse
}

func (o *GetAccountRoutingsOK) Error() string {
	return fmt.Sprintf("[GET /account_routings][%d] getAccountRoutingsOK  %+v", 200, o.Payload)
}

func (o *GetAccountRoutingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountRoutingDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
