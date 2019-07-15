// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/terraform-provider-form3/models"
)

// GetSubscriptionsIDReader is a Reader for the GetSubscriptionsID structure.
type GetSubscriptionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubscriptionsIDOK creates a GetSubscriptionsIDOK with default headers values
func NewGetSubscriptionsIDOK() *GetSubscriptionsIDOK {
	return &GetSubscriptionsIDOK{}
}

/*GetSubscriptionsIDOK handles this case with default header values.

Subscription details
*/
type GetSubscriptionsIDOK struct {
	Payload *models.SubscriptionDetailsResponse
}

func (o *GetSubscriptionsIDOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}][%d] getSubscriptionsIdOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
