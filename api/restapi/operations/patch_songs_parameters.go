// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"main/internal/models"
)

// NewPatchSongsParams creates a new PatchSongsParams object
//
// There are no default values defined in the spec.
func NewPatchSongsParams() PatchSongsParams {

	return PatchSongsParams{}
}

// PatchSongsParams contains all the bound params for the patch songs operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchSongs
type PatchSongsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*song
	  Required: true
	  In: body
	*/
	Song *models.Song
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchSongsParams() beforehand.
func (o *PatchSongsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Song
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("song", "body", ""))
			} else {
				res = append(res, errors.NewParseError("song", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Song = &body
			}
		}
	} else {
		res = append(res, errors.Required("song", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
