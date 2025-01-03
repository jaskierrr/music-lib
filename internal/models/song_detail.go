// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SongDetail song detail
//
// swagger:model SongDetail
type SongDetail struct {

	// text
	// Example: Ooh baby, don't you know I suffer?
	Text string `json:"Text,omitempty"`

	// link
	// Example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
	Link string `json:"link,omitempty"`

	// release date
	// Example: 2006-07-16
	ReleaseDate string `json:"releaseDate,omitempty"`
}

// Validate validates this song detail
func (m *SongDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this song detail based on context it is used
func (m *SongDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SongDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SongDetail) UnmarshalBinary(b []byte) error {
	var res SongDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
