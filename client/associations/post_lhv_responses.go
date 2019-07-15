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

// PostLhvReader is a Reader for the PostLhv structure.
type PostLhvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLhvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostLhvCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLhvCreated creates a PostLhvCreated with default headers values
func NewPostLhvCreated() *PostLhvCreated {
	return &PostLhvCreated{}
}

/*PostLhvCreated handles this case with default header values.

creation response
*/
type PostLhvCreated struct {
	Payload *models.LhvAssociationCreationResponse
}

func (o *PostLhvCreated) Error() string {
	return fmt.Sprintf("[POST /lhv][%d] postLhvCreated  %+v", 201, o.Payload)
}

func (o *PostLhvCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
