// Code generated by go-swagger; DO NOT EDIT.

package pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/signalcd/signalcd/api/v1/models"
)

// UpdatePipelineAgentsOKCode is the HTTP code returned for type UpdatePipelineAgentsOK
const UpdatePipelineAgentsOKCode int = 200

/*UpdatePipelineAgentsOK OK

swagger:response updatePipelineAgentsOK
*/
type UpdatePipelineAgentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pipeline `json:"body,omitempty"`
}

// NewUpdatePipelineAgentsOK creates UpdatePipelineAgentsOK with default headers values
func NewUpdatePipelineAgentsOK() *UpdatePipelineAgentsOK {

	return &UpdatePipelineAgentsOK{}
}

// WithPayload adds the payload to the update pipeline agents o k response
func (o *UpdatePipelineAgentsOK) WithPayload(payload *models.Pipeline) *UpdatePipelineAgentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update pipeline agents o k response
func (o *UpdatePipelineAgentsOK) SetPayload(payload *models.Pipeline) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePipelineAgentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePipelineAgentsBadRequestCode is the HTTP code returned for type UpdatePipelineAgentsBadRequest
const UpdatePipelineAgentsBadRequestCode int = 400

/*UpdatePipelineAgentsBadRequest bad request

swagger:response updatePipelineAgentsBadRequest
*/
type UpdatePipelineAgentsBadRequest struct {
}

// NewUpdatePipelineAgentsBadRequest creates UpdatePipelineAgentsBadRequest with default headers values
func NewUpdatePipelineAgentsBadRequest() *UpdatePipelineAgentsBadRequest {

	return &UpdatePipelineAgentsBadRequest{}
}

// WriteResponse to the client
func (o *UpdatePipelineAgentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdatePipelineAgentsInternalServerErrorCode is the HTTP code returned for type UpdatePipelineAgentsInternalServerError
const UpdatePipelineAgentsInternalServerErrorCode int = 500

/*UpdatePipelineAgentsInternalServerError internal server error

swagger:response updatePipelineAgentsInternalServerError
*/
type UpdatePipelineAgentsInternalServerError struct {
}

// NewUpdatePipelineAgentsInternalServerError creates UpdatePipelineAgentsInternalServerError with default headers values
func NewUpdatePipelineAgentsInternalServerError() *UpdatePipelineAgentsInternalServerError {

	return &UpdatePipelineAgentsInternalServerError{}
}

// WriteResponse to the client
func (o *UpdatePipelineAgentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
