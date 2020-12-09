// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigClusterNameReader is a Reader for the FindConfigClusterName structure.
type FindConfigClusterNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigClusterNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigClusterNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigClusterNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigClusterNameOK creates a FindConfigClusterNameOK with default headers values
func NewFindConfigClusterNameOK() *FindConfigClusterNameOK {
	return &FindConfigClusterNameOK{}
}

/*FindConfigClusterNameOK handles this case with default header values.

Config value
*/
type FindConfigClusterNameOK struct {
	Payload string
}

func (o *FindConfigClusterNameOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigClusterNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigClusterNameDefault creates a FindConfigClusterNameDefault with default headers values
func NewFindConfigClusterNameDefault(code int) *FindConfigClusterNameDefault {
	return &FindConfigClusterNameDefault{
		_statusCode: code,
	}
}

/*FindConfigClusterNameDefault handles this case with default header values.

unexpected error
*/
type FindConfigClusterNameDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config cluster name default response
func (o *FindConfigClusterNameDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigClusterNameDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigClusterNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigClusterNameDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}