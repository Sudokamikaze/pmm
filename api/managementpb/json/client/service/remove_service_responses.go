// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoveServiceReader is a Reader for the RemoveService structure.
type RemoveServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveServiceOK creates a RemoveServiceOK with default headers values
func NewRemoveServiceOK() *RemoveServiceOK {
	return &RemoveServiceOK{}
}

/*RemoveServiceOK handles this case with default header values.

A successful response.
*/
type RemoveServiceOK struct {
	Payload interface{}
}

func (o *RemoveServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/Remove][%d] removeServiceOk  %+v", 200, o.Payload)
}

func (o *RemoveServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveServiceDefault creates a RemoveServiceDefault with default headers values
func NewRemoveServiceDefault(code int) *RemoveServiceDefault {
	return &RemoveServiceDefault{
		_statusCode: code,
	}
}

/*RemoveServiceDefault handles this case with default header values.

An unexpected error response.
*/
type RemoveServiceDefault struct {
	_statusCode int

	Payload *RemoveServiceDefaultBody
}

// Code gets the status code for the remove service default response
func (o *RemoveServiceDefault) Code() int {
	return o._statusCode
}

func (o *RemoveServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/Remove][%d] RemoveService default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveServiceDefault) GetPayload() *RemoveServiceDefaultBody {
	return o.Payload
}

func (o *RemoveServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveServiceBody remove service body
swagger:model RemoveServiceBody
*/
type RemoveServiceBody struct {

	// ServiceType describes supported Service types.
	// Enum: [SERVICE_TYPE_INVALID MYSQL_SERVICE MONGODB_SERVICE POSTGRESQL_SERVICE PROXYSQL_SERVICE HAPROXY_SERVICE EXTERNAL_SERVICE]
	ServiceType *string `json:"service_type,omitempty"`

	// Service ID or Service Name is required.
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this remove service body
func (o *RemoveServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var removeServiceBodyTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_TYPE_INVALID","MYSQL_SERVICE","MONGODB_SERVICE","POSTGRESQL_SERVICE","PROXYSQL_SERVICE","HAPROXY_SERVICE","EXTERNAL_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		removeServiceBodyTypeServiceTypePropEnum = append(removeServiceBodyTypeServiceTypePropEnum, v)
	}
}

const (

	// RemoveServiceBodyServiceTypeSERVICETYPEINVALID captures enum value "SERVICE_TYPE_INVALID"
	RemoveServiceBodyServiceTypeSERVICETYPEINVALID string = "SERVICE_TYPE_INVALID"

	// RemoveServiceBodyServiceTypeMYSQLSERVICE captures enum value "MYSQL_SERVICE"
	RemoveServiceBodyServiceTypeMYSQLSERVICE string = "MYSQL_SERVICE"

	// RemoveServiceBodyServiceTypeMONGODBSERVICE captures enum value "MONGODB_SERVICE"
	RemoveServiceBodyServiceTypeMONGODBSERVICE string = "MONGODB_SERVICE"

	// RemoveServiceBodyServiceTypePOSTGRESQLSERVICE captures enum value "POSTGRESQL_SERVICE"
	RemoveServiceBodyServiceTypePOSTGRESQLSERVICE string = "POSTGRESQL_SERVICE"

	// RemoveServiceBodyServiceTypePROXYSQLSERVICE captures enum value "PROXYSQL_SERVICE"
	RemoveServiceBodyServiceTypePROXYSQLSERVICE string = "PROXYSQL_SERVICE"

	// RemoveServiceBodyServiceTypeHAPROXYSERVICE captures enum value "HAPROXY_SERVICE"
	RemoveServiceBodyServiceTypeHAPROXYSERVICE string = "HAPROXY_SERVICE"

	// RemoveServiceBodyServiceTypeEXTERNALSERVICE captures enum value "EXTERNAL_SERVICE"
	RemoveServiceBodyServiceTypeEXTERNALSERVICE string = "EXTERNAL_SERVICE"
)

// prop value enum
func (o *RemoveServiceBody) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, removeServiceBodyTypeServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RemoveServiceBody) validateServiceType(formats strfmt.Registry) error {

	if swag.IsZero(o.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := o.validateServiceTypeEnum("body"+"."+"service_type", "body", *o.ServiceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveServiceDefaultBody remove service default body
swagger:model RemoveServiceDefaultBody
*/
type RemoveServiceDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this remove service default body
func (o *RemoveServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveServiceDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("RemoveService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
