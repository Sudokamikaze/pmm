// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// AddQANMySQLPerfSchemaAgentReader is a Reader for the AddQANMySQLPerfSchemaAgent structure.
type AddQANMySQLPerfSchemaAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddQANMySQLPerfSchemaAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddQANMySQLPerfSchemaAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddQANMySQLPerfSchemaAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddQANMySQLPerfSchemaAgentOK creates a AddQANMySQLPerfSchemaAgentOK with default headers values
func NewAddQANMySQLPerfSchemaAgentOK() *AddQANMySQLPerfSchemaAgentOK {
	return &AddQANMySQLPerfSchemaAgentOK{}
}

/*AddQANMySQLPerfSchemaAgentOK handles this case with default header values.

A successful response.
*/
type AddQANMySQLPerfSchemaAgentOK struct {
	Payload *AddQANMySQLPerfSchemaAgentOKBody
}

func (o *AddQANMySQLPerfSchemaAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMySQLPerfSchemaAgent][%d] addQanMySqlPerfSchemaAgentOk  %+v", 200, o.Payload)
}

func (o *AddQANMySQLPerfSchemaAgentOK) GetPayload() *AddQANMySQLPerfSchemaAgentOKBody {
	return o.Payload
}

func (o *AddQANMySQLPerfSchemaAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMySQLPerfSchemaAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddQANMySQLPerfSchemaAgentDefault creates a AddQANMySQLPerfSchemaAgentDefault with default headers values
func NewAddQANMySQLPerfSchemaAgentDefault(code int) *AddQANMySQLPerfSchemaAgentDefault {
	return &AddQANMySQLPerfSchemaAgentDefault{
		_statusCode: code,
	}
}

/*AddQANMySQLPerfSchemaAgentDefault handles this case with default header values.

An unexpected error response.
*/
type AddQANMySQLPerfSchemaAgentDefault struct {
	_statusCode int

	Payload *AddQANMySQLPerfSchemaAgentDefaultBody
}

// Code gets the status code for the add QAN my SQL perf schema agent default response
func (o *AddQANMySQLPerfSchemaAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddQANMySQLPerfSchemaAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMySQLPerfSchemaAgent][%d] AddQANMySQLPerfSchemaAgent default  %+v", o._statusCode, o.Payload)
}

func (o *AddQANMySQLPerfSchemaAgentDefault) GetPayload() *AddQANMySQLPerfSchemaAgentDefaultBody {
	return o.Payload
}

func (o *AddQANMySQLPerfSchemaAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMySQLPerfSchemaAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddQANMySQLPerfSchemaAgentBody add QAN my SQL perf schema agent body
swagger:model AddQANMySQLPerfSchemaAgentBody
*/
type AddQANMySQLPerfSchemaAgentBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`

	// MySQL password for getting performance data.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`
}

// Validate validates this add QAN my SQL perf schema agent body
func (o *AddQANMySQLPerfSchemaAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLPerfSchemaAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLPerfSchemaAgentDefaultBody add QAN my SQL perf schema agent default body
swagger:model AddQANMySQLPerfSchemaAgentDefaultBody
*/
type AddQANMySQLPerfSchemaAgentDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add QAN my SQL perf schema agent default body
func (o *AddQANMySQLPerfSchemaAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMySQLPerfSchemaAgentDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddQANMySQLPerfSchemaAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLPerfSchemaAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLPerfSchemaAgentOKBody add QAN my SQL perf schema agent OK body
swagger:model AddQANMySQLPerfSchemaAgentOKBody
*/
type AddQANMySQLPerfSchemaAgentOKBody struct {

	// qan mysql perfschema agent
	QANMysqlPerfschemaAgent *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent `json:"qan_mysql_perfschema_agent,omitempty"`
}

// Validate validates this add QAN my SQL perf schema agent OK body
func (o *AddQANMySQLPerfSchemaAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMysqlPerfschemaAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMySQLPerfSchemaAgentOKBody) validateQANMysqlPerfschemaAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschemaAgent) { // not required
		return nil
	}

	if o.QANMysqlPerfschemaAgent != nil {
		if err := o.QANMysqlPerfschemaAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanMySqlPerfSchemaAgentOk" + "." + "qan_mysql_perfschema_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLPerfSchemaAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent
*/
type AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

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
}

// Validate validates this add QAN my SQL perf schema agent OK body QAN mysql perfschema agent
func (o *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum = append(addQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTARTING captures enum value "STARTING"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTARTING string = "STARTING"

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusRUNNING captures enum value "RUNNING"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusRUNNING string = "RUNNING"

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusWAITING captures enum value "WAITING"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusWAITING string = "WAITING"

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTOPPING captures enum value "STOPPING"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTOPPING string = "STOPPING"

	// AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusDONE captures enum value "DONE"
	AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addQanMySqlPerfSchemaAgentOk"+"."+"qan_mysql_perfschema_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
