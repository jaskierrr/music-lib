// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Song song
//
// swagger:model Song
type Song struct {

	// group
	// Example: Muse
	Group string `json:"group,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// link
	// Example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
	Link string `json:"link,omitempty"`

	// lyrics
	// Example: Ooh baby, don't you know I suffer?
	Lyrics string `json:"lyrics,omitempty"`

	// release date
	// Example: 2006-07-16
	ReleaseDate string `json:"releaseDate,omitempty"`

	// song
	// Example: Supermassive Black Hole
	Song string `json:"song,omitempty"`
}

// Validate validates this song
func (m *Song) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this song based on context it is used
func (m *Song) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Song) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Song) UnmarshalBinary(b []byte) error {
	var res Song
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}