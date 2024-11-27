// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSongsHandlerFunc turns a function with the right signature into a delete songs handler
type DeleteSongsHandlerFunc func(DeleteSongsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSongsHandlerFunc) Handle(params DeleteSongsParams) middleware.Responder {
	return fn(params)
}

// DeleteSongsHandler interface for that can handle valid delete songs params
type DeleteSongsHandler interface {
	Handle(DeleteSongsParams) middleware.Responder
}

// NewDeleteSongs creates a new http.Handler for the delete songs operation
func NewDeleteSongs(ctx *middleware.Context, handler DeleteSongsHandler) *DeleteSongs {
	return &DeleteSongs{Context: ctx, Handler: handler}
}

/*
	DeleteSongs swagger:route DELETE /songs deleteSongs

Delete song
*/
type DeleteSongs struct {
	Context *middleware.Context
	Handler DeleteSongsHandler
}

func (o *DeleteSongs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSongsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
