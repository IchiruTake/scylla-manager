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

// TaskManagerListModulesGetReader is a Reader for the TaskManagerListModulesGet structure.
type TaskManagerListModulesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaskManagerListModulesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskManagerListModulesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTaskManagerListModulesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaskManagerListModulesGetOK creates a TaskManagerListModulesGetOK with default headers values
func NewTaskManagerListModulesGetOK() *TaskManagerListModulesGetOK {
	return &TaskManagerListModulesGetOK{}
}

/*
TaskManagerListModulesGetOK handles this case with default header values.

Success
*/
type TaskManagerListModulesGetOK struct {
	Payload []string
}

func (o *TaskManagerListModulesGetOK) GetPayload() []string {
	return o.Payload
}

func (o *TaskManagerListModulesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskManagerListModulesGetDefault creates a TaskManagerListModulesGetDefault with default headers values
func NewTaskManagerListModulesGetDefault(code int) *TaskManagerListModulesGetDefault {
	return &TaskManagerListModulesGetDefault{
		_statusCode: code,
	}
}

/*
TaskManagerListModulesGetDefault handles this case with default header values.

internal server error
*/
type TaskManagerListModulesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the task manager list modules get default response
func (o *TaskManagerListModulesGetDefault) Code() int {
	return o._statusCode
}

func (o *TaskManagerListModulesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *TaskManagerListModulesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *TaskManagerListModulesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
