// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AgentsHandlerFunc turns a function with the right signature into a agents handler
type AgentsHandlerFunc func(AgentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AgentsHandlerFunc) Handle(params AgentsParams) middleware.Responder {
	return fn(params)
}

// AgentsHandler interface for that can handle valid agents params
type AgentsHandler interface {
	Handle(AgentsParams) middleware.Responder
}

// NewAgents creates a new http.Handler for the agents operation
func NewAgents(ctx *middleware.Context, handler AgentsHandler) *Agents {
	return &Agents{Context: ctx, Handler: handler}
}

/*Agents swagger:route GET /agents agents agents

Returns a list of agents

*/
type Agents struct {
	Context *middleware.Context
	Handler AgentsHandler
}

func (o *Agents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAgentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}