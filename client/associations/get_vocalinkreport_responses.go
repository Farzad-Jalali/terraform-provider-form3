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

// GetVocalinkreportReader is a Reader for the GetVocalinkreport structure.
type GetVocalinkreportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVocalinkreportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVocalinkreportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVocalinkreportOK creates a GetVocalinkreportOK with default headers values
func NewGetVocalinkreportOK() *GetVocalinkreportOK {
	return &GetVocalinkreportOK{}
}

/*GetVocalinkreportOK handles this case with default header values.

List of associations
*/
type GetVocalinkreportOK struct {
	Payload *models.VocalinkReportAssociationDetailsListResponse
}

func (o *GetVocalinkreportOK) Error() string {
	return fmt.Sprintf("[GET /vocalinkreport][%d] getVocalinkreportOK  %+v", 200, o.Payload)
}

func (o *GetVocalinkreportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VocalinkReportAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
