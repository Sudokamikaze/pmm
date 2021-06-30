// Code generated by go-swagger; DO NOT EDIT.

package agent_local

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

// StatusReader is a Reader for the Status structure.
type StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusOK creates a StatusOK with default headers values
func NewStatusOK() *StatusOK {
	return &StatusOK{}
}

/*StatusOK handles this case with default header values.

A successful response.
*/
type StatusOK struct {
	Payload *StatusOKBody
}

func (o *StatusOK) Error() string {
	return fmt.Sprintf("[POST /local/Status][%d] statusOk  %+v", 200, o.Payload)
}

func (o *StatusOK) GetPayload() *StatusOKBody {
	return o.Payload
}

func (o *StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusDefault creates a StatusDefault with default headers values
func NewStatusDefault(code int) *StatusDefault {
	return &StatusDefault{
		_statusCode: code,
	}
}

/*StatusDefault handles this case with default header values.

An unexpected error response.
*/
type StatusDefault struct {
	_statusCode int

	Payload *StatusDefaultBody
}

// Code gets the status code for the status default response
func (o *StatusDefault) Code() int {
	return o._statusCode
}

func (o *StatusDefault) Error() string {
	return fmt.Sprintf("[POST /local/Status][%d] Status default  %+v", o._statusCode, o.Payload)
}

func (o *StatusDefault) GetPayload() *StatusDefaultBody {
	return o.Payload
}

func (o *StatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AgentsInfoItems0 AgentInfo contains information about Agent managed by pmm-agent.
swagger:model AgentsInfoItems0
*/
type AgentsInfoItems0 struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// AgentType describes supported Agent types.
	// Enum: [AGENT_TYPE_INVALID PMM_AGENT VM_AGENT NODE_EXPORTER MYSQLD_EXPORTER MONGODB_EXPORTER POSTGRES_EXPORTER PROXYSQL_EXPORTER QAN_MYSQL_PERFSCHEMA_AGENT QAN_MYSQL_SLOWLOG_AGENT QAN_MONGODB_PROFILER_AGENT QAN_POSTGRESQL_PGSTATEMENTS_AGENT QAN_POSTGRESQL_PGSTATMONITOR_AGENT RDS_EXPORTER EXTERNAL_EXPORTER AZURE_DATABASE_EXPORTER]
	AgentType *string `json:"agent_type,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// The current listen port of this Agent (exporter or vmagent).
	// Zero for other Agent types, or if unknown or not yet supported.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this agents info items0
func (o *AgentsInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var agentsInfoItems0TypeAgentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_TYPE_INVALID","PMM_AGENT","VM_AGENT","NODE_EXPORTER","MYSQLD_EXPORTER","MONGODB_EXPORTER","POSTGRES_EXPORTER","PROXYSQL_EXPORTER","QAN_MYSQL_PERFSCHEMA_AGENT","QAN_MYSQL_SLOWLOG_AGENT","QAN_MONGODB_PROFILER_AGENT","QAN_POSTGRESQL_PGSTATEMENTS_AGENT","QAN_POSTGRESQL_PGSTATMONITOR_AGENT","RDS_EXPORTER","EXTERNAL_EXPORTER","AZURE_DATABASE_EXPORTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentsInfoItems0TypeAgentTypePropEnum = append(agentsInfoItems0TypeAgentTypePropEnum, v)
	}
}

const (

	// AgentsInfoItems0AgentTypeAGENTTYPEINVALID captures enum value "AGENT_TYPE_INVALID"
	AgentsInfoItems0AgentTypeAGENTTYPEINVALID string = "AGENT_TYPE_INVALID"

	// AgentsInfoItems0AgentTypePMMAGENT captures enum value "PMM_AGENT"
	AgentsInfoItems0AgentTypePMMAGENT string = "PMM_AGENT"

	// AgentsInfoItems0AgentTypeVMAGENT captures enum value "VM_AGENT"
	AgentsInfoItems0AgentTypeVMAGENT string = "VM_AGENT"

	// AgentsInfoItems0AgentTypeNODEEXPORTER captures enum value "NODE_EXPORTER"
	AgentsInfoItems0AgentTypeNODEEXPORTER string = "NODE_EXPORTER"

	// AgentsInfoItems0AgentTypeMYSQLDEXPORTER captures enum value "MYSQLD_EXPORTER"
	AgentsInfoItems0AgentTypeMYSQLDEXPORTER string = "MYSQLD_EXPORTER"

	// AgentsInfoItems0AgentTypeMONGODBEXPORTER captures enum value "MONGODB_EXPORTER"
	AgentsInfoItems0AgentTypeMONGODBEXPORTER string = "MONGODB_EXPORTER"

	// AgentsInfoItems0AgentTypePOSTGRESEXPORTER captures enum value "POSTGRES_EXPORTER"
	AgentsInfoItems0AgentTypePOSTGRESEXPORTER string = "POSTGRES_EXPORTER"

	// AgentsInfoItems0AgentTypePROXYSQLEXPORTER captures enum value "PROXYSQL_EXPORTER"
	AgentsInfoItems0AgentTypePROXYSQLEXPORTER string = "PROXYSQL_EXPORTER"

	// AgentsInfoItems0AgentTypeQANMYSQLPERFSCHEMAAGENT captures enum value "QAN_MYSQL_PERFSCHEMA_AGENT"
	AgentsInfoItems0AgentTypeQANMYSQLPERFSCHEMAAGENT string = "QAN_MYSQL_PERFSCHEMA_AGENT"

	// AgentsInfoItems0AgentTypeQANMYSQLSLOWLOGAGENT captures enum value "QAN_MYSQL_SLOWLOG_AGENT"
	AgentsInfoItems0AgentTypeQANMYSQLSLOWLOGAGENT string = "QAN_MYSQL_SLOWLOG_AGENT"

	// AgentsInfoItems0AgentTypeQANMONGODBPROFILERAGENT captures enum value "QAN_MONGODB_PROFILER_AGENT"
	AgentsInfoItems0AgentTypeQANMONGODBPROFILERAGENT string = "QAN_MONGODB_PROFILER_AGENT"

	// AgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATEMENTSAGENT captures enum value "QAN_POSTGRESQL_PGSTATEMENTS_AGENT"
	AgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATEMENTSAGENT string = "QAN_POSTGRESQL_PGSTATEMENTS_AGENT"

	// AgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATMONITORAGENT captures enum value "QAN_POSTGRESQL_PGSTATMONITOR_AGENT"
	AgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATMONITORAGENT string = "QAN_POSTGRESQL_PGSTATMONITOR_AGENT"

	// AgentsInfoItems0AgentTypeRDSEXPORTER captures enum value "RDS_EXPORTER"
	AgentsInfoItems0AgentTypeRDSEXPORTER string = "RDS_EXPORTER"

	// AgentsInfoItems0AgentTypeEXTERNALEXPORTER captures enum value "EXTERNAL_EXPORTER"
	AgentsInfoItems0AgentTypeEXTERNALEXPORTER string = "EXTERNAL_EXPORTER"

	// AgentsInfoItems0AgentTypeAZUREDATABASEEXPORTER captures enum value "AZURE_DATABASE_EXPORTER"
	AgentsInfoItems0AgentTypeAZUREDATABASEEXPORTER string = "AZURE_DATABASE_EXPORTER"
)

// prop value enum
func (o *AgentsInfoItems0) validateAgentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agentsInfoItems0TypeAgentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AgentsInfoItems0) validateAgentType(formats strfmt.Registry) error {

	if swag.IsZero(o.AgentType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAgentTypeEnum("agent_type", "body", *o.AgentType); err != nil {
		return err
	}

	return nil
}

var agentsInfoItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentsInfoItems0TypeStatusPropEnum = append(agentsInfoItems0TypeStatusPropEnum, v)
	}
}

const (

	// AgentsInfoItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AgentsInfoItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AgentsInfoItems0StatusSTARTING captures enum value "STARTING"
	AgentsInfoItems0StatusSTARTING string = "STARTING"

	// AgentsInfoItems0StatusRUNNING captures enum value "RUNNING"
	AgentsInfoItems0StatusRUNNING string = "RUNNING"

	// AgentsInfoItems0StatusWAITING captures enum value "WAITING"
	AgentsInfoItems0StatusWAITING string = "WAITING"

	// AgentsInfoItems0StatusSTOPPING captures enum value "STOPPING"
	AgentsInfoItems0StatusSTOPPING string = "STOPPING"

	// AgentsInfoItems0StatusDONE captures enum value "DONE"
	AgentsInfoItems0StatusDONE string = "DONE"

	// AgentsInfoItems0StatusUNKNOWN captures enum value "UNKNOWN"
	AgentsInfoItems0StatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AgentsInfoItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agentsInfoItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AgentsInfoItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AgentsInfoItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AgentsInfoItems0) UnmarshalBinary(b []byte) error {
	var res AgentsInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusBody status body
swagger:model StatusBody
*/
type StatusBody struct {

	// Returns network info (latency and clock_drift) if true.
	GetNetworkInfo bool `json:"get_network_info,omitempty"`
}

// Validate validates this status body
func (o *StatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusBody) UnmarshalBinary(b []byte) error {
	var res StatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusDefaultBody status default body
swagger:model StatusDefaultBody
*/
type StatusDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this status default body
func (o *StatusDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("Status default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res StatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBody status OK body
swagger:model StatusOKBody
*/
type StatusOKBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// runs on node id
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// agents info
	AgentsInfo []*AgentsInfoItems0 `json:"agents_info"`

	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `json:"config_filepath,omitempty"`

	// PMM Agent version.
	AgentVersion string `json:"agent_version,omitempty"`

	// server info
	ServerInfo *StatusOKBodyServerInfo `json:"server_info,omitempty"`
}

// Validate validates this status OK body
func (o *StatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusOKBody) validateAgentsInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.AgentsInfo) { // not required
		return nil
	}

	for i := 0; i < len(o.AgentsInfo); i++ {
		if swag.IsZero(o.AgentsInfo[i]) { // not required
			continue
		}

		if o.AgentsInfo[i] != nil {
			if err := o.AgentsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statusOk" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *StatusOKBody) validateServerInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.ServerInfo) { // not required
		return nil
	}

	if o.ServerInfo != nil {
		if err := o.ServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOk" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBody) UnmarshalBinary(b []byte) error {
	var res StatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBodyServerInfo ServerInfo contains information about the PMM Server.
swagger:model StatusOKBodyServerInfo
*/
type StatusOKBodyServerInfo struct {

	// PMM Server URL in a form https://HOST:PORT/.
	URL string `json:"url,omitempty"`

	// PMM Server's TLS certificate validation should be skipped if true.
	InsecureTLS bool `json:"insecure_tls,omitempty"`

	// True if pmm-agent is currently connected to the server.
	Connected bool `json:"connected,omitempty"`

	// PMM Server version (if agent is connected).
	Version string `json:"version,omitempty"`

	// Ping time from pmm-agent to pmm-managed (if agent is connected).
	Latency string `json:"latency,omitempty"`

	// Clock drift from PMM Server (if agent is connected).
	ClockDrift string `json:"clock_drift,omitempty"`
}

// Validate validates this status OK body server info
func (o *StatusOKBodyServerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyServerInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyServerInfo) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
