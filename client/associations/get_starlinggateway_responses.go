// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetStarlinggatewayReader is a Reader for the GetStarlinggateway structure.
type GetStarlinggatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStarlinggatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStarlinggatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarlinggatewayOK creates a GetStarlinggatewayOK with default headers values
func NewGetStarlinggatewayOK() *GetStarlinggatewayOK {
	return &GetStarlinggatewayOK{}
}

/*GetStarlinggatewayOK handles this case with default header values.

List of associations
*/
type GetStarlinggatewayOK struct {
	Payload *models.AssociationDetailsListResponse
}

func (o *GetStarlinggatewayOK) Error() string {
	return fmt.Sprintf("[GET /starlinggateway][%d] getStarlinggatewayOK  %+v", 200, o.Payload)
}

func (o *GetStarlinggatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}