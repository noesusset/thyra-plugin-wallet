// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra-plugin-wallet/api/server/models"
)

// RestWalletSignOperationOKCode is the HTTP code returned for type RestWalletSignOperationOK
const RestWalletSignOperationOKCode int = 200

/*
RestWalletSignOperationOK Signature.

swagger:response restWalletSignOperationOK
*/
type RestWalletSignOperationOK struct {

	/*
	  In: Body
	*/
	Payload *models.Signature `json:"body,omitempty"`
}

// NewRestWalletSignOperationOK creates RestWalletSignOperationOK with default headers values
func NewRestWalletSignOperationOK() *RestWalletSignOperationOK {

	return &RestWalletSignOperationOK{}
}

// WithPayload adds the payload to the rest wallet sign operation o k response
func (o *RestWalletSignOperationOK) WithPayload(payload *models.Signature) *RestWalletSignOperationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet sign operation o k response
func (o *RestWalletSignOperationOK) SetPayload(payload *models.Signature) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletSignOperationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestWalletSignOperationBadRequestCode is the HTTP code returned for type RestWalletSignOperationBadRequest
const RestWalletSignOperationBadRequestCode int = 400

/*
RestWalletSignOperationBadRequest Bad request.

swagger:response restWalletSignOperationBadRequest
*/
type RestWalletSignOperationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletSignOperationBadRequest creates RestWalletSignOperationBadRequest with default headers values
func NewRestWalletSignOperationBadRequest() *RestWalletSignOperationBadRequest {

	return &RestWalletSignOperationBadRequest{}
}

// WithPayload adds the payload to the rest wallet sign operation bad request response
func (o *RestWalletSignOperationBadRequest) WithPayload(payload *models.Error) *RestWalletSignOperationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet sign operation bad request response
func (o *RestWalletSignOperationBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletSignOperationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestWalletSignOperationUnprocessableEntityCode is the HTTP code returned for type RestWalletSignOperationUnprocessableEntity
const RestWalletSignOperationUnprocessableEntityCode int = 422

/*
RestWalletSignOperationUnprocessableEntity Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.

swagger:response restWalletSignOperationUnprocessableEntity
*/
type RestWalletSignOperationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletSignOperationUnprocessableEntity creates RestWalletSignOperationUnprocessableEntity with default headers values
func NewRestWalletSignOperationUnprocessableEntity() *RestWalletSignOperationUnprocessableEntity {

	return &RestWalletSignOperationUnprocessableEntity{}
}

// WithPayload adds the payload to the rest wallet sign operation unprocessable entity response
func (o *RestWalletSignOperationUnprocessableEntity) WithPayload(payload *models.Error) *RestWalletSignOperationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet sign operation unprocessable entity response
func (o *RestWalletSignOperationUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletSignOperationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestWalletSignOperationInternalServerErrorCode is the HTTP code returned for type RestWalletSignOperationInternalServerError
const RestWalletSignOperationInternalServerErrorCode int = 500

/*
RestWalletSignOperationInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response restWalletSignOperationInternalServerError
*/
type RestWalletSignOperationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestWalletSignOperationInternalServerError creates RestWalletSignOperationInternalServerError with default headers values
func NewRestWalletSignOperationInternalServerError() *RestWalletSignOperationInternalServerError {

	return &RestWalletSignOperationInternalServerError{}
}

// WithPayload adds the payload to the rest wallet sign operation internal server error response
func (o *RestWalletSignOperationInternalServerError) WithPayload(payload *models.Error) *RestWalletSignOperationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest wallet sign operation internal server error response
func (o *RestWalletSignOperationInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestWalletSignOperationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
