// Code generated by go-swagger; DO NOT EDIT.

package object_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetQueryExampleReader is a Reader for the GetQueryExample structure.
type GetQueryExampleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueryExampleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueryExampleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetQueryExampleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueryExampleOK creates a GetQueryExampleOK with default headers values
func NewGetQueryExampleOK() *GetQueryExampleOK {
	return &GetQueryExampleOK{}
}

/*GetQueryExampleOK handles this case with default header values.

A successful response.
*/
type GetQueryExampleOK struct {
	Payload *GetQueryExampleOKBody
}

func (o *GetQueryExampleOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetQueryExample][%d] getQueryExampleOk  %+v", 200, o.Payload)
}

func (o *GetQueryExampleOK) GetPayload() *GetQueryExampleOKBody {
	return o.Payload
}

func (o *GetQueryExampleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetQueryExampleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueryExampleDefault creates a GetQueryExampleDefault with default headers values
func NewGetQueryExampleDefault(code int) *GetQueryExampleDefault {
	return &GetQueryExampleDefault{
		_statusCode: code,
	}
}

/*GetQueryExampleDefault handles this case with default header values.

An error response.
*/
type GetQueryExampleDefault struct {
	_statusCode int

	Payload *GetQueryExampleDefaultBody
}

// Code gets the status code for the get query example default response
func (o *GetQueryExampleDefault) Code() int {
	return o._statusCode
}

func (o *GetQueryExampleDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetQueryExample][%d] GetQueryExample default  %+v", o._statusCode, o.Payload)
}

func (o *GetQueryExampleDefault) GetPayload() *GetQueryExampleDefaultBody {
	return o.Payload
}

func (o *GetQueryExampleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetQueryExampleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetQueryExampleBody QueryExampleRequest defines filtering of query examples for specific value of
// dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
swagger:model GetQueryExampleBody
*/
type GetQueryExampleBody struct {

	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `json:"filter_by,omitempty"`

	// one of dimension: queryid | host ...
	GroupBy string `json:"group_by,omitempty"`

	// labels
	Labels []*LabelsItems0 `json:"labels"`

	// limit
	Limit int64 `json:"limit,omitempty"`
}

// Validate validates this get query example body
func (o *GetQueryExampleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetQueryExampleBody) validatePeriodStartFrom(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetQueryExampleBody) validatePeriodStartTo(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetQueryExampleBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetQueryExampleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetQueryExampleBody) UnmarshalBinary(b []byte) error {
	var res GetQueryExampleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetQueryExampleDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model GetQueryExampleDefaultBody
*/
type GetQueryExampleDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get query example default body
func (o *GetQueryExampleDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetQueryExampleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetQueryExampleDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetQueryExampleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetQueryExampleOKBody QueryExampleReply list of query examples.
swagger:model GetQueryExampleOKBody
*/
type GetQueryExampleOKBody struct {

	// query examples
	QueryExamples []*QueryExamplesItems0 `json:"query_examples"`
}

// Validate validates this get query example OK body
func (o *GetQueryExampleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQueryExamples(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetQueryExampleOKBody) validateQueryExamples(formats strfmt.Registry) error {

	if swag.IsZero(o.QueryExamples) { // not required
		return nil
	}

	for i := 0; i < len(o.QueryExamples); i++ {
		if swag.IsZero(o.QueryExamples[i]) { // not required
			continue
		}

		if o.QueryExamples[i] != nil {
			if err := o.QueryExamples[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getQueryExampleOk" + "." + "query_examples" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetQueryExampleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetQueryExampleOKBody) UnmarshalBinary(b []byte) error {
	var res GetQueryExampleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryExamplesItems0 QueryExample shows query examples and their metrics.
swagger:model QueryExamplesItems0
*/
type QueryExamplesItems0 struct {

	// example
	Example string `json:"example,omitempty"`

	// ExampleFormat is format of query example: real or query without values.
	// Enum: [EXAMPLE_FORMAT_INVALID EXAMPLE FINGERPRINT]
	ExampleFormat *string `json:"example_format,omitempty"`

	// ExampleType is a type of query example selected for this query class in given period of time.
	// Enum: [EXAMPLE_TYPE_INVALID RANDOM SLOWEST FASTEST WITH_ERROR]
	ExampleType *string `json:"example_type,omitempty"`

	// is truncated
	IsTruncated int64 `json:"is_truncated,omitempty"`

	// example metrics
	ExampleMetrics string `json:"example_metrics,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// service type
	ServiceType string `json:"service_type,omitempty"`

	// schema
	Schema string `json:"schema,omitempty"`

	// tables
	Tables []string `json:"tables"`
}

// Validate validates this query examples items0
func (o *QueryExamplesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExampleFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExampleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var queryExamplesItems0TypeExampleFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXAMPLE_FORMAT_INVALID","EXAMPLE","FINGERPRINT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryExamplesItems0TypeExampleFormatPropEnum = append(queryExamplesItems0TypeExampleFormatPropEnum, v)
	}
}

const (

	// QueryExamplesItems0ExampleFormatEXAMPLEFORMATINVALID captures enum value "EXAMPLE_FORMAT_INVALID"
	QueryExamplesItems0ExampleFormatEXAMPLEFORMATINVALID string = "EXAMPLE_FORMAT_INVALID"

	// QueryExamplesItems0ExampleFormatEXAMPLE captures enum value "EXAMPLE"
	QueryExamplesItems0ExampleFormatEXAMPLE string = "EXAMPLE"

	// QueryExamplesItems0ExampleFormatFINGERPRINT captures enum value "FINGERPRINT"
	QueryExamplesItems0ExampleFormatFINGERPRINT string = "FINGERPRINT"
)

// prop value enum
func (o *QueryExamplesItems0) validateExampleFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryExamplesItems0TypeExampleFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *QueryExamplesItems0) validateExampleFormat(formats strfmt.Registry) error {

	if swag.IsZero(o.ExampleFormat) { // not required
		return nil
	}

	// value enum
	if err := o.validateExampleFormatEnum("example_format", "body", *o.ExampleFormat); err != nil {
		return err
	}

	return nil
}

var queryExamplesItems0TypeExampleTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXAMPLE_TYPE_INVALID","RANDOM","SLOWEST","FASTEST","WITH_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryExamplesItems0TypeExampleTypePropEnum = append(queryExamplesItems0TypeExampleTypePropEnum, v)
	}
}

const (

	// QueryExamplesItems0ExampleTypeEXAMPLETYPEINVALID captures enum value "EXAMPLE_TYPE_INVALID"
	QueryExamplesItems0ExampleTypeEXAMPLETYPEINVALID string = "EXAMPLE_TYPE_INVALID"

	// QueryExamplesItems0ExampleTypeRANDOM captures enum value "RANDOM"
	QueryExamplesItems0ExampleTypeRANDOM string = "RANDOM"

	// QueryExamplesItems0ExampleTypeSLOWEST captures enum value "SLOWEST"
	QueryExamplesItems0ExampleTypeSLOWEST string = "SLOWEST"

	// QueryExamplesItems0ExampleTypeFASTEST captures enum value "FASTEST"
	QueryExamplesItems0ExampleTypeFASTEST string = "FASTEST"

	// QueryExamplesItems0ExampleTypeWITHERROR captures enum value "WITH_ERROR"
	QueryExamplesItems0ExampleTypeWITHERROR string = "WITH_ERROR"
)

// prop value enum
func (o *QueryExamplesItems0) validateExampleTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryExamplesItems0TypeExampleTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *QueryExamplesItems0) validateExampleType(formats strfmt.Registry) error {

	if swag.IsZero(o.ExampleType) { // not required
		return nil
	}

	// value enum
	if err := o.validateExampleTypeEnum("example_type", "body", *o.ExampleType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryExamplesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryExamplesItems0) UnmarshalBinary(b []byte) error {
	var res QueryExamplesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
