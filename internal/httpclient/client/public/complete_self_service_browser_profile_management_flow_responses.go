// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// CompleteSelfServiceBrowserProfileManagementFlowReader is a Reader for the CompleteSelfServiceBrowserProfileManagementFlow structure.
type CompleteSelfServiceBrowserProfileManagementFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceBrowserProfileManagementFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceBrowserProfileManagementFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceBrowserProfileManagementFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteSelfServiceBrowserProfileManagementFlowFound creates a CompleteSelfServiceBrowserProfileManagementFlowFound with default headers values
func NewCompleteSelfServiceBrowserProfileManagementFlowFound() *CompleteSelfServiceBrowserProfileManagementFlowFound {
	return &CompleteSelfServiceBrowserProfileManagementFlowFound{}
}

/*CompleteSelfServiceBrowserProfileManagementFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceBrowserProfileManagementFlowFound struct {
}

func (o *CompleteSelfServiceBrowserProfileManagementFlowFound) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/profile/update][%d] completeSelfServiceBrowserProfileManagementFlowFound ", 302)
}

func (o *CompleteSelfServiceBrowserProfileManagementFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceBrowserProfileManagementFlowInternalServerError creates a CompleteSelfServiceBrowserProfileManagementFlowInternalServerError with default headers values
func NewCompleteSelfServiceBrowserProfileManagementFlowInternalServerError() *CompleteSelfServiceBrowserProfileManagementFlowInternalServerError {
	return &CompleteSelfServiceBrowserProfileManagementFlowInternalServerError{}
}

/*CompleteSelfServiceBrowserProfileManagementFlowInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceBrowserProfileManagementFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceBrowserProfileManagementFlowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/profile/update][%d] completeSelfServiceBrowserProfileManagementFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceBrowserProfileManagementFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceBrowserProfileManagementFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
