// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckForOperatorUpdateReader is a Reader for the CheckForOperatorUpdate structure.
type CheckForOperatorUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckForOperatorUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckForOperatorUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckForOperatorUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckForOperatorUpdateOK creates a CheckForOperatorUpdateOK with default headers values
func NewCheckForOperatorUpdateOK() *CheckForOperatorUpdateOK {
	return &CheckForOperatorUpdateOK{}
}

/*CheckForOperatorUpdateOK handles this case with default header values.

A successful response.
*/
type CheckForOperatorUpdateOK struct {
	Payload *CheckForOperatorUpdateOKBody
}

func (o *CheckForOperatorUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/CheckForOperatorUpdate][%d] checkForOperatorUpdateOk  %+v", 200, o.Payload)
}

func (o *CheckForOperatorUpdateOK) GetPayload() *CheckForOperatorUpdateOKBody {
	return o.Payload
}

func (o *CheckForOperatorUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckForOperatorUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckForOperatorUpdateDefault creates a CheckForOperatorUpdateDefault with default headers values
func NewCheckForOperatorUpdateDefault(code int) *CheckForOperatorUpdateDefault {
	return &CheckForOperatorUpdateDefault{
		_statusCode: code,
	}
}

/*CheckForOperatorUpdateDefault handles this case with default header values.

An unexpected error response.
*/
type CheckForOperatorUpdateDefault struct {
	_statusCode int

	Payload *CheckForOperatorUpdateDefaultBody
}

// Code gets the status code for the check for operator update default response
func (o *CheckForOperatorUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CheckForOperatorUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/CheckForOperatorUpdate][%d] CheckForOperatorUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *CheckForOperatorUpdateDefault) GetPayload() *CheckForOperatorUpdateDefaultBody {
	return o.Payload
}

func (o *CheckForOperatorUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckForOperatorUpdateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckForOperatorUpdateDefaultBody check for operator update default body
swagger:model CheckForOperatorUpdateDefaultBody
*/
type CheckForOperatorUpdateDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this check for operator update default body
func (o *CheckForOperatorUpdateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckForOperatorUpdateDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckForOperatorUpdate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckForOperatorUpdateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckForOperatorUpdateDefaultBody) UnmarshalBinary(b []byte) error {
	var res CheckForOperatorUpdateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckForOperatorUpdateOKBody check for operator update OK body
swagger:model CheckForOperatorUpdateOKBody
*/
type CheckForOperatorUpdateOKBody struct {

	// The cluster name is used as a key for this map, value stores information about operators update availability.
	UpdateInformation map[string]UpdateInformationAnon `json:"update_information,omitempty"`
}

// Validate validates this check for operator update OK body
func (o *CheckForOperatorUpdateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUpdateInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckForOperatorUpdateOKBody) validateUpdateInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdateInformation) { // not required
		return nil
	}

	for k := range o.UpdateInformation {

		if swag.IsZero(o.UpdateInformation[k]) { // not required
			continue
		}
		if val, ok := o.UpdateInformation[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckForOperatorUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckForOperatorUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res CheckForOperatorUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateInformationAnon OperatorsUpdateInformation is just a container for update information for PXC and PSMDB operators.
swagger:model UpdateInformationAnon
*/
type UpdateInformationAnon struct {

	// pxc_operator_version is not empty if there is an update available, otherwise empty.
	PxcOperatorVersion string `json:"pxc_operator_version,omitempty"`

	// psmdb_operator_version is not empty if there is an update available, otherwise empty.
	PSMDBOperatorVersion string `json:"psmdb_operator_version,omitempty"`
}

// Validate validates this update information anon
func (o *UpdateInformationAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateInformationAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateInformationAnon) UnmarshalBinary(b []byte) error {
	var res UpdateInformationAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
