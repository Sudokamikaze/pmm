// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddProxySQLExporterReader is a Reader for the AddProxySQLExporter structure.
type AddProxySQLExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProxySQLExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProxySQLExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddProxySQLExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProxySQLExporterOK creates a AddProxySQLExporterOK with default headers values
func NewAddProxySQLExporterOK() *AddProxySQLExporterOK {
	return &AddProxySQLExporterOK{}
}

/*AddProxySQLExporterOK handles this case with default header values.

A successful response.
*/
type AddProxySQLExporterOK struct {
	Payload *AddProxySQLExporterOKBody
}

func (o *AddProxySQLExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddProxySQLExporter][%d] addProxySqlExporterOk  %+v", 200, o.Payload)
}

func (o *AddProxySQLExporterOK) GetPayload() *AddProxySQLExporterOKBody {
	return o.Payload
}

func (o *AddProxySQLExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProxySQLExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProxySQLExporterDefault creates a AddProxySQLExporterDefault with default headers values
func NewAddProxySQLExporterDefault(code int) *AddProxySQLExporterDefault {
	return &AddProxySQLExporterDefault{
		_statusCode: code,
	}
}

/*AddProxySQLExporterDefault handles this case with default header values.

An error response.
*/
type AddProxySQLExporterDefault struct {
	_statusCode int

	Payload *AddProxySQLExporterDefaultBody
}

// Code gets the status code for the add proxy SQL exporter default response
func (o *AddProxySQLExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddProxySQLExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddProxySQLExporter][%d] AddProxySQLExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddProxySQLExporterDefault) GetPayload() *AddProxySQLExporterDefaultBody {
	return o.Payload
}

func (o *AddProxySQLExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProxySQLExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddProxySQLExporterBody add proxy SQL exporter body
swagger:model AddProxySQLExporterBody
*/
type AddProxySQLExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// ProxySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`
}

// Validate validates this add proxy SQL exporter body
func (o *AddProxySQLExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLExporterDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddProxySQLExporterDefaultBody
*/
type AddProxySQLExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add proxy SQL exporter default body
func (o *AddProxySQLExporterDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLExporterOKBody add proxy SQL exporter OK body
swagger:model AddProxySQLExporterOKBody
*/
type AddProxySQLExporterOKBody struct {

	// proxysql exporter
	ProxysqlExporter *AddProxySQLExporterOKBodyProxysqlExporter `json:"proxysql_exporter,omitempty"`
}

// Validate validates this add proxy SQL exporter OK body
func (o *AddProxySQLExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysqlExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLExporterOKBody) validateProxysqlExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ProxysqlExporter) { // not required
		return nil
	}

	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlExporterOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLExporterOKBodyProxysqlExporter ProxySQLExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model AddProxySQLExporterOKBodyProxysqlExporter
*/
type AddProxySQLExporterOKBodyProxysqlExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this add proxy SQL exporter OK body proxysql exporter
func (o *AddProxySQLExporterOKBodyProxysqlExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum = append(addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusSTARTING captures enum value "STARTING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusSTARTING string = "STARTING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusRUNNING captures enum value "RUNNING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusRUNNING string = "RUNNING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusWAITING captures enum value "WAITING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusWAITING string = "WAITING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING captures enum value "STOPPING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING string = "STOPPING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusDONE captures enum value "DONE"
	AddProxySQLExporterOKBodyProxysqlExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addProxySqlExporterOk"+"."+"proxysql_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterOKBodyProxysqlExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterOKBodyProxysqlExporter) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterOKBodyProxysqlExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
