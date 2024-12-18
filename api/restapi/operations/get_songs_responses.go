// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"main/internal/models"
)

// GetSongsOKCode is the HTTP code returned for type GetSongsOK
const GetSongsOKCode int = 200

/*
GetSongsOK song list

swagger:response getSongsOK
*/
type GetSongsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Song `json:"body,omitempty"`
}

// NewGetSongsOK creates GetSongsOK with default headers values
func NewGetSongsOK() *GetSongsOK {

	return &GetSongsOK{}
}

// WithPayload adds the payload to the get songs o k response
func (o *GetSongsOK) WithPayload(payload []*models.Song) *GetSongsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get songs o k response
func (o *GetSongsOK) SetPayload(payload []*models.Song) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSongsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Song, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSongsBadRequestCode is the HTTP code returned for type GetSongsBadRequest
const GetSongsBadRequestCode int = 400

/*
GetSongsBadRequest bad request

swagger:response getSongsBadRequest
*/
type GetSongsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetSongsBadRequest creates GetSongsBadRequest with default headers values
func NewGetSongsBadRequest() *GetSongsBadRequest {

	return &GetSongsBadRequest{}
}

// WithPayload adds the payload to the get songs bad request response
func (o *GetSongsBadRequest) WithPayload(payload *models.ErrorResponse) *GetSongsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get songs bad request response
func (o *GetSongsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSongsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
