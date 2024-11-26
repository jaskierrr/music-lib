// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSongsHandlerFunc turns a function with the right signature into a get songs handler
type GetSongsHandlerFunc func(GetSongsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSongsHandlerFunc) Handle(params GetSongsParams) middleware.Responder {
	return fn(params)
}

// GetSongsHandler interface for that can handle valid get songs params
type GetSongsHandler interface {
	Handle(GetSongsParams) middleware.Responder
}

// NewGetSongs creates a new http.Handler for the get songs operation
func NewGetSongs(ctx *middleware.Context, handler GetSongsHandler) *GetSongs {
	return &GetSongs{Context: ctx, Handler: handler}
}

/*
	GetSongs swagger:route GET /songs getSongs

Return songs
*/
type GetSongs struct {
	Context *middleware.Context
	Handler GetSongsHandler
}

func (o *GetSongs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSongsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}