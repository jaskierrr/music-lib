// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSongsLyricsHandlerFunc turns a function with the right signature into a get songs lyrics handler
type GetSongsLyricsHandlerFunc func(GetSongsLyricsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSongsLyricsHandlerFunc) Handle(params GetSongsLyricsParams) middleware.Responder {
	return fn(params)
}

// GetSongsLyricsHandler interface for that can handle valid get songs lyrics params
type GetSongsLyricsHandler interface {
	Handle(GetSongsLyricsParams) middleware.Responder
}

// NewGetSongsLyrics creates a new http.Handler for the get songs lyrics operation
func NewGetSongsLyrics(ctx *middleware.Context, handler GetSongsLyricsHandler) *GetSongsLyrics {
	return &GetSongsLyrics{Context: ctx, Handler: handler}
}

/*
	GetSongsLyrics swagger:route GET /songs/lyrics getSongsLyrics

Return song lyrics
*/
type GetSongsLyrics struct {
	Context *middleware.Context
	Handler GetSongsLyricsHandler
}

func (o *GetSongsLyrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSongsLyricsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
