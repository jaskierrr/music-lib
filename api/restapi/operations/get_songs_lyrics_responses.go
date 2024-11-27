// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"main/internal/models"
)

// GetSongsLyricsOKCode is the HTTP code returned for type GetSongsLyricsOK
const GetSongsLyricsOKCode int = 200

/*
GetSongsLyricsOK song lyrics

swagger:response getSongsLyricsOK
*/
type GetSongsLyricsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetSongsLyricsOK creates GetSongsLyricsOK with default headers values
func NewGetSongsLyricsOK() *GetSongsLyricsOK {

	return &GetSongsLyricsOK{}
}

// WithPayload adds the payload to the get songs lyrics o k response
func (o *GetSongsLyricsOK) WithPayload(payload string) *GetSongsLyricsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get songs lyrics o k response
func (o *GetSongsLyricsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSongsLyricsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSongsLyricsBadRequestCode is the HTTP code returned for type GetSongsLyricsBadRequest
const GetSongsLyricsBadRequestCode int = 400

/*
GetSongsLyricsBadRequest bad request

swagger:response getSongsLyricsBadRequest
*/
type GetSongsLyricsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetSongsLyricsBadRequest creates GetSongsLyricsBadRequest with default headers values
func NewGetSongsLyricsBadRequest() *GetSongsLyricsBadRequest {

	return &GetSongsLyricsBadRequest{}
}

// WithPayload adds the payload to the get songs lyrics bad request response
func (o *GetSongsLyricsBadRequest) WithPayload(payload *models.ErrorResponse) *GetSongsLyricsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get songs lyrics bad request response
func (o *GetSongsLyricsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSongsLyricsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
