// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExpandedService expanded service
// swagger:model ExpandedService
type ExpandedService struct {

	// Creation date of the service
	CreationDate string `json:"creationDate,omitempty"`

	// Currently deployed image
	DeployedImage string `json:"deployedImage,omitempty"`

	// last event types
	LastEventTypes map[string]EventContext `json:"lastEventTypes,omitempty"`

	// Service name
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this expanded service
func (m *ExpandedService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEventTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpandedService) validateLastEventTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEventTypes) { // not required
		return nil
	}

	for k := range m.LastEventTypes {

		if err := validate.Required("lastEventTypes"+"."+k, "body", m.LastEventTypes[k]); err != nil {
			return err
		}
		if val, ok := m.LastEventTypes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpandedService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpandedService) UnmarshalBinary(b []byte) error {
	var res ExpandedService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}