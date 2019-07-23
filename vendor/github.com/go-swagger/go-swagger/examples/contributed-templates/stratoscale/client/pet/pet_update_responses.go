// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetUpdateReader is a Reader for the PetUpdate structure.
type PetUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPetUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPetUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPetUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPetUpdateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPetUpdateCreated creates a PetUpdateCreated with default headers values
func NewPetUpdateCreated() *PetUpdateCreated {
	return &PetUpdateCreated{}
}

/*PetUpdateCreated handles this case with default header values.

Updated successfully
*/
type PetUpdateCreated struct {
	Payload *models.Pet
}

func (o *PetUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateCreated  %+v", 201, o.Payload)
}

func (o *PetUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetUpdateBadRequest creates a PetUpdateBadRequest with default headers values
func NewPetUpdateBadRequest() *PetUpdateBadRequest {
	return &PetUpdateBadRequest{}
}

/*PetUpdateBadRequest handles this case with default header values.

Invalid ID supplied
*/
type PetUpdateBadRequest struct {
}

func (o *PetUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateBadRequest ", 400)
}

func (o *PetUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetUpdateNotFound creates a PetUpdateNotFound with default headers values
func NewPetUpdateNotFound() *PetUpdateNotFound {
	return &PetUpdateNotFound{}
}

/*PetUpdateNotFound handles this case with default header values.

Pet not found
*/
type PetUpdateNotFound struct {
}

func (o *PetUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateNotFound ", 404)
}

func (o *PetUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetUpdateMethodNotAllowed creates a PetUpdateMethodNotAllowed with default headers values
func NewPetUpdateMethodNotAllowed() *PetUpdateMethodNotAllowed {
	return &PetUpdateMethodNotAllowed{}
}

/*PetUpdateMethodNotAllowed handles this case with default header values.

Validation exception
*/
type PetUpdateMethodNotAllowed struct {
}

func (o *PetUpdateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateMethodNotAllowed ", 405)
}

func (o *PetUpdateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
