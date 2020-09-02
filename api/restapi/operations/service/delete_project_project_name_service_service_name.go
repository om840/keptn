// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/keptn/keptn/api/models"
)

// DeleteProjectProjectNameServiceServiceNameHandlerFunc turns a function with the right signature into a delete project project name service service name handler
type DeleteProjectProjectNameServiceServiceNameHandlerFunc func(DeleteProjectProjectNameServiceServiceNameParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProjectProjectNameServiceServiceNameHandlerFunc) Handle(params DeleteProjectProjectNameServiceServiceNameParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteProjectProjectNameServiceServiceNameHandler interface for that can handle valid delete project project name service service name params
type DeleteProjectProjectNameServiceServiceNameHandler interface {
	Handle(DeleteProjectProjectNameServiceServiceNameParams, *models.Principal) middleware.Responder
}

// NewDeleteProjectProjectNameServiceServiceName creates a new http.Handler for the delete project project name service service name operation
func NewDeleteProjectProjectNameServiceServiceName(ctx *middleware.Context, handler DeleteProjectProjectNameServiceServiceNameHandler) *DeleteProjectProjectNameServiceServiceName {
	return &DeleteProjectProjectNameServiceServiceName{Context: ctx, Handler: handler}
}

/*DeleteProjectProjectNameServiceServiceName swagger:route DELETE /project/{projectName}/service/{serviceName} Service deleteProjectProjectNameServiceServiceName

Deletes a service

*/
type DeleteProjectProjectNameServiceServiceName struct {
	Context *middleware.Context
	Handler DeleteProjectProjectNameServiceServiceNameHandler
}

func (o *DeleteProjectProjectNameServiceServiceName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProjectProjectNameServiceServiceNameParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}