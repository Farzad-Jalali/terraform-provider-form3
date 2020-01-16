// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSepaReconciliationIDReader is a Reader for the DeleteSepaReconciliationID structure.
type DeleteSepaReconciliationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepaReconciliationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSepaReconciliationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSepaReconciliationIDNoContent creates a DeleteSepaReconciliationIDNoContent with default headers values
func NewDeleteSepaReconciliationIDNoContent() *DeleteSepaReconciliationIDNoContent {
	return &DeleteSepaReconciliationIDNoContent{}
}

/*DeleteSepaReconciliationIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteSepaReconciliationIDNoContent struct {
}

func (o *DeleteSepaReconciliationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sepa-reconciliation/{id}][%d] deleteSepaReconciliationIdNoContent ", 204)
}

func (o *DeleteSepaReconciliationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
