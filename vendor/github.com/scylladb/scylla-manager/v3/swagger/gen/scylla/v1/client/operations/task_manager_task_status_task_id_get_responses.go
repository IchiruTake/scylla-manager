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

// TaskManagerTaskStatusTaskIDGetReader is a Reader for the TaskManagerTaskStatusTaskIDGet structure.
type TaskManagerTaskStatusTaskIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaskManagerTaskStatusTaskIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskManagerTaskStatusTaskIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTaskManagerTaskStatusTaskIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaskManagerTaskStatusTaskIDGetOK creates a TaskManagerTaskStatusTaskIDGetOK with default headers values
func NewTaskManagerTaskStatusTaskIDGetOK() *TaskManagerTaskStatusTaskIDGetOK {
	return &TaskManagerTaskStatusTaskIDGetOK{}
}

/*
TaskManagerTaskStatusTaskIDGetOK handles this case with default header values.

Success
*/
type TaskManagerTaskStatusTaskIDGetOK struct {
	Payload *models.TaskStatus
}

func (o *TaskManagerTaskStatusTaskIDGetOK) GetPayload() *models.TaskStatus {
	return o.Payload
}

func (o *TaskManagerTaskStatusTaskIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskManagerTaskStatusTaskIDGetDefault creates a TaskManagerTaskStatusTaskIDGetDefault with default headers values
func NewTaskManagerTaskStatusTaskIDGetDefault(code int) *TaskManagerTaskStatusTaskIDGetDefault {
	return &TaskManagerTaskStatusTaskIDGetDefault{
		_statusCode: code,
	}
}

/*
TaskManagerTaskStatusTaskIDGetDefault handles this case with default header values.

internal server error
*/
type TaskManagerTaskStatusTaskIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the task manager task status task Id get default response
func (o *TaskManagerTaskStatusTaskIDGetDefault) Code() int {
	return o._statusCode
}

func (o *TaskManagerTaskStatusTaskIDGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *TaskManagerTaskStatusTaskIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *TaskManagerTaskStatusTaskIDGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
