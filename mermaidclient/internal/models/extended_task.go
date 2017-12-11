// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtendedTask extended task
// swagger:model extendedTask
type ExtendedTask struct {

	// cause
	Cause string `json:"cause,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// end time
	EndTime string `json:"end_time,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// metadata
	Metadata string `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// schedule
	Schedule *Schedule `json:"schedule,omitempty"`

	// start time
	StartTime string `json:"start_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this extended task
func (m *ExtendedTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedTask) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {

		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedTask) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtendedTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtendedTask) UnmarshalBinary(b []byte) error {
	var res ExtendedTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
