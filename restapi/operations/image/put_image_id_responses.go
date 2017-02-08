package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*PutImageIDNoContent Image was updated

swagger:response putImageIdNoContent
*/
type PutImageIDNoContent struct {
}

// NewPutImageIDNoContent creates PutImageIDNoContent with default headers values
func NewPutImageIDNoContent() *PutImageIDNoContent {
	return &PutImageIDNoContent{}
}

// WriteResponse to the client
func (o *PutImageIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}