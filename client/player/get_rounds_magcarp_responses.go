// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"DatsMagic/models"
)

// GetRoundsMagcarpReader is a Reader for the GetRoundsMagcarp structure.
type GetRoundsMagcarpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoundsMagcarpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoundsMagcarpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoundsMagcarpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /rounds/magcarp] GetRoundsMagcarp", response, response.Code())
	}
}

// NewGetRoundsMagcarpOK creates a GetRoundsMagcarpOK with default headers values
func NewGetRoundsMagcarpOK() *GetRoundsMagcarpOK {
	return &GetRoundsMagcarpOK{}
}

/*
GetRoundsMagcarpOK describes a response with status code 200, with default header values.

OK
*/
type GetRoundsMagcarpOK struct {
	Payload *models.DtoRoundList
}

// IsSuccess returns true when this get rounds magcarp o k response has a 2xx status code
func (o *GetRoundsMagcarpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rounds magcarp o k response has a 3xx status code
func (o *GetRoundsMagcarpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rounds magcarp o k response has a 4xx status code
func (o *GetRoundsMagcarpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rounds magcarp o k response has a 5xx status code
func (o *GetRoundsMagcarpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rounds magcarp o k response a status code equal to that given
func (o *GetRoundsMagcarpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rounds magcarp o k response
func (o *GetRoundsMagcarpOK) Code() int {
	return 200
}

func (o *GetRoundsMagcarpOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /rounds/magcarp][%d] getRoundsMagcarpOK %s", 200, payload)
}

func (o *GetRoundsMagcarpOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /rounds/magcarp][%d] getRoundsMagcarpOK %s", 200, payload)
}

func (o *GetRoundsMagcarpOK) GetPayload() *models.DtoRoundList {
	return o.Payload
}

func (o *GetRoundsMagcarpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoRoundList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoundsMagcarpBadRequest creates a GetRoundsMagcarpBadRequest with default headers values
func NewGetRoundsMagcarpBadRequest() *GetRoundsMagcarpBadRequest {
	return &GetRoundsMagcarpBadRequest{}
}

/*
GetRoundsMagcarpBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRoundsMagcarpBadRequest struct {
	Payload *models.PuberrPubErr
}

// IsSuccess returns true when this get rounds magcarp bad request response has a 2xx status code
func (o *GetRoundsMagcarpBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rounds magcarp bad request response has a 3xx status code
func (o *GetRoundsMagcarpBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rounds magcarp bad request response has a 4xx status code
func (o *GetRoundsMagcarpBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rounds magcarp bad request response has a 5xx status code
func (o *GetRoundsMagcarpBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get rounds magcarp bad request response a status code equal to that given
func (o *GetRoundsMagcarpBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get rounds magcarp bad request response
func (o *GetRoundsMagcarpBadRequest) Code() int {
	return 400
}

func (o *GetRoundsMagcarpBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /rounds/magcarp][%d] getRoundsMagcarpBadRequest %s", 400, payload)
}

func (o *GetRoundsMagcarpBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /rounds/magcarp][%d] getRoundsMagcarpBadRequest %s", 400, payload)
}

func (o *GetRoundsMagcarpBadRequest) GetPayload() *models.PuberrPubErr {
	return o.Payload
}

func (o *GetRoundsMagcarpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PuberrPubErr)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
