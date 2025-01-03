// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceRestorePostReader is a Reader for the StorageServiceRestorePost structure.
type StorageServiceRestorePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRestorePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRestorePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRestorePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRestorePostOK creates a StorageServiceRestorePostOK with default headers values
func NewStorageServiceRestorePostOK() *StorageServiceRestorePostOK {
	return &StorageServiceRestorePostOK{}
}

/*
StorageServiceRestorePostOK handles this case with default header values.

Task ID that can be used with Task Manager service
*/
type StorageServiceRestorePostOK struct {
	Payload string
}

func (o *StorageServiceRestorePostOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceRestorePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceRestorePostDefault creates a StorageServiceRestorePostDefault with default headers values
func NewStorageServiceRestorePostDefault(code int) *StorageServiceRestorePostDefault {
	return &StorageServiceRestorePostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceRestorePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceRestorePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service restore post default response
func (o *StorageServiceRestorePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRestorePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRestorePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRestorePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
