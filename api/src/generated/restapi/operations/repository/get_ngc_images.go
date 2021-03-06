// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rescale-labs/scaleshift/api/src/auth"
)

// GetNgcImagesHandlerFunc turns a function with the right signature into a get ngc images handler
type GetNgcImagesHandlerFunc func(GetNgcImagesParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNgcImagesHandlerFunc) Handle(params GetNgcImagesParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNgcImagesHandler interface for that can handle valid get ngc images params
type GetNgcImagesHandler interface {
	Handle(GetNgcImagesParams, *auth.Principal) middleware.Responder
}

// NewGetNgcImages creates a new http.Handler for the get ngc images operation
func NewGetNgcImages(ctx *middleware.Context, handler GetNgcImagesHandler) *GetNgcImages {
	return &GetNgcImages{Context: ctx, Handler: handler}
}

/*GetNgcImages swagger:route GET /nvidia/repositories/{namespace}/images/{id} repository getNgcImages

returns NGC images


*/
type GetNgcImages struct {
	Context *middleware.Context
	Handler GetNgcImagesHandler
}

func (o *GetNgcImages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNgcImagesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *auth.Principal
	if uprinc != nil {
		principal = uprinc.(*auth.Principal) // this is really a auth.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
