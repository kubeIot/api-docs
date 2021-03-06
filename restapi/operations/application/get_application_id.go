package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetApplicationIDHandlerFunc turns a function with the right signature into a get application ID handler
type GetApplicationIDHandlerFunc func(GetApplicationIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationIDHandlerFunc) Handle(params GetApplicationIDParams) middleware.Responder {
	return fn(params)
}

// GetApplicationIDHandler interface for that can handle valid get application ID params
type GetApplicationIDHandler interface {
	Handle(GetApplicationIDParams) middleware.Responder
}

// NewGetApplicationID creates a new http.Handler for the get application ID operation
func NewGetApplicationID(ctx *middleware.Context, handler GetApplicationIDHandler) *GetApplicationID {
	return &GetApplicationID{Context: ctx, Handler: handler}
}

/*GetApplicationID swagger:route GET /application/{id} application getApplicationId

Get application by ID

*/
type GetApplicationID struct {
	Context *middleware.Context
	Handler GetApplicationIDHandler
}

func (o *GetApplicationID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetApplicationIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		// bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
