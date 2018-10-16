// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	baseclient "github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/baseclient"
	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/models"
)

// GetMtaOperationsReader is a Reader for the GetMtaOperations structure.
type GetMtaOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMtaOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMtaOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, baseclient.BuildErrorResponse(response, consumer, o.formats)
	}
}

// NewGetMtaOperationsOK creates a GetMtaOperationsOK with default headers values
func NewGetMtaOperationsOK() *GetMtaOperationsOK {
	return &GetMtaOperationsOK{}
}

/*GetMtaOperationsOK handles this case with default header values.

OK
*/
type GetMtaOperationsOK struct {
	Payload models.GetMtaOperationsOKBody
}

func (o *GetMtaOperationsOK) Error() string {
	return fmt.Sprintf("[GET /operations][%d] getMtaOperationsOK  %+v", 200, o.Payload)
}

func (o *GetMtaOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
