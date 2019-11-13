// Code generated by go-swagger; DO NOT EDIT.

package my_sql

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

// AddMySQLReader is a Reader for the AddMySQL structure.
type AddMySQLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMySQLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMySQLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMySQLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMySQLOK creates a AddMySQLOK with default headers values
func NewAddMySQLOK() *AddMySQLOK {
	return &AddMySQLOK{}
}

/*AddMySQLOK handles this case with default header values.

A successful response.
*/
type AddMySQLOK struct {
	Payload *AddMySQLOKBody
}

func (o *AddMySQLOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/MySQL/Add][%d] addMySqlOk  %+v", 200, o.Payload)
}

func (o *AddMySQLOK) GetPayload() *AddMySQLOKBody {
	return o.Payload
}

func (o *AddMySQLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMySQLDefault creates a AddMySQLDefault with default headers values
func NewAddMySQLDefault(code int) *AddMySQLDefault {
	return &AddMySQLDefault{
		_statusCode: code,
	}
}

/*AddMySQLDefault handles this case with default header values.

An error response.
*/
type AddMySQLDefault struct {
	_statusCode int

	Payload *AddMySQLDefaultBody
}

// Code gets the status code for the add my SQL default response
func (o *AddMySQLDefault) Code() int {
	return o._statusCode
}

func (o *AddMySQLDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/MySQL/Add][%d] AddMySQL default  %+v", o._statusCode, o.Payload)
}

func (o *AddMySQLDefault) GetPayload() *AddMySQLDefaultBody {
	return o.Payload
}

func (o *AddMySQLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMySQLBody add my SQL body
swagger:model AddMySQLBody
*/
type AddMySQLBody struct {

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node and Service access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Service Access port. Required.
	Port int64 `json:"port,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// If true, adds qan-mysql-perfschema-agent for provided service.
	QANMysqlPerfschema bool `json:"qan_mysql_perfschema,omitempty"`

	// If true, adds qan-mysql-slowlog-agent for provided service.
	QANMysqlSlowlog bool `json:"qan_mysql_slowlog,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// If qan-mysql-slowlog-agent is added, slowlog file is rotated at this size if > 0.
	// If zero, default value 1GB is used. Use negative value to disable rotation.
	MaxSlowlogFileSize string `json:"max_slowlog_file_size,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// add node
	AddNode *AddMySQLParamsBodyAddNode `json:"add_node,omitempty"`
}

// Validate validates this add my SQL body
func (o *AddMySQLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLBody) validateAddNode(formats strfmt.Registry) error {

	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddMySQLDefaultBody
*/
type AddMySQLDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add my SQL default body
func (o *AddMySQLDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLOKBody add my SQL OK body
swagger:model AddMySQLOKBody
*/
type AddMySQLOKBody struct {

	// mysqld exporter
	MysqldExporter *AddMySQLOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`

	// qan mysql perfschema
	QANMysqlPerfschema *AddMySQLOKBodyQANMysqlPerfschema `json:"qan_mysql_perfschema,omitempty"`

	// qan mysql slowlog
	QANMysqlSlowlog *AddMySQLOKBodyQANMysqlSlowlog `json:"qan_mysql_slowlog,omitempty"`

	// service
	Service *AddMySQLOKBodyService `json:"service,omitempty"`
}

// Validate validates this add my SQL OK body
func (o *AddMySQLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlPerfschema(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlSlowlog(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqlOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddMySQLOKBody) validateQANMysqlPerfschema(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschema) { // not required
		return nil
	}

	if o.QANMysqlPerfschema != nil {
		if err := o.QANMysqlPerfschema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqlOk" + "." + "qan_mysql_perfschema")
			}
			return err
		}
	}

	return nil
}

func (o *AddMySQLOKBody) validateQANMysqlSlowlog(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlSlowlog) { // not required
		return nil
	}

	if o.QANMysqlSlowlog != nil {
		if err := o.QANMysqlSlowlog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqlOk" + "." + "qan_mysql_slowlog")
			}
			return err
		}
	}

	return nil
}

func (o *AddMySQLOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLOKBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model AddMySQLOKBodyMysqldExporter
*/
type AddMySQLOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
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

// Validate validates this add my SQL OK body mysqld exporter
func (o *AddMySQLOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySqlOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySqlOkBodyMysqldExporterTypeStatusPropEnum = append(addMySqlOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMySQLOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMySQLOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMySQLOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	AddMySQLOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// AddMySQLOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	AddMySQLOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// AddMySQLOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	AddMySQLOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// AddMySQLOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	AddMySQLOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// AddMySQLOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	AddMySQLOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMySQLOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMySqlOkBodyMysqldExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMySQLOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMySqlOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res AddMySQLOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLOKBodyQANMysqlPerfschema QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddMySQLOKBodyQANMysqlPerfschema
*/
type AddMySQLOKBodyQANMysqlPerfschema struct {

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

// Validate validates this add my SQL OK body QAN mysql perfschema
func (o *AddMySQLOKBodyQANMysqlPerfschema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySqlOkBodyQanMysqlPerfschemaTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySqlOkBodyQanMysqlPerfschemaTypeStatusPropEnum = append(addMySqlOkBodyQanMysqlPerfschemaTypeStatusPropEnum, v)
	}
}

const (

	// AddMySQLOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMySQLOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMySQLOKBodyQANMysqlPerfschemaStatusSTARTING captures enum value "STARTING"
	AddMySQLOKBodyQANMysqlPerfschemaStatusSTARTING string = "STARTING"

	// AddMySQLOKBodyQANMysqlPerfschemaStatusRUNNING captures enum value "RUNNING"
	AddMySQLOKBodyQANMysqlPerfschemaStatusRUNNING string = "RUNNING"

	// AddMySQLOKBodyQANMysqlPerfschemaStatusWAITING captures enum value "WAITING"
	AddMySQLOKBodyQANMysqlPerfschemaStatusWAITING string = "WAITING"

	// AddMySQLOKBodyQANMysqlPerfschemaStatusSTOPPING captures enum value "STOPPING"
	AddMySQLOKBodyQANMysqlPerfschemaStatusSTOPPING string = "STOPPING"

	// AddMySQLOKBodyQANMysqlPerfschemaStatusDONE captures enum value "DONE"
	AddMySQLOKBodyQANMysqlPerfschemaStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMySQLOKBodyQANMysqlPerfschema) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMySqlOkBodyQanMysqlPerfschemaTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMySQLOKBodyQANMysqlPerfschema) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMySqlOk"+"."+"qan_mysql_perfschema"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLOKBodyQANMysqlPerfschema) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLOKBodyQANMysqlPerfschema) UnmarshalBinary(b []byte) error {
	var res AddMySQLOKBodyQANMysqlPerfschema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLOKBodyQANMysqlSlowlog QANMySQLSlowlogAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddMySQLOKBodyQANMysqlSlowlog
*/
type AddMySQLOKBodyQANMysqlSlowlog struct {

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

	// query examples disabled
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Slowlog file is rotated at this size if > 0.
	MaxSlowlogFileSize string `json:"max_slowlog_file_size,omitempty"`

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

// Validate validates this add my SQL OK body QAN mysql slowlog
func (o *AddMySQLOKBodyQANMysqlSlowlog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySqlOkBodyQanMysqlSlowlogTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySqlOkBodyQanMysqlSlowlogTypeStatusPropEnum = append(addMySqlOkBodyQanMysqlSlowlogTypeStatusPropEnum, v)
	}
}

const (

	// AddMySQLOKBodyQANMysqlSlowlogStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMySQLOKBodyQANMysqlSlowlogStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMySQLOKBodyQANMysqlSlowlogStatusSTARTING captures enum value "STARTING"
	AddMySQLOKBodyQANMysqlSlowlogStatusSTARTING string = "STARTING"

	// AddMySQLOKBodyQANMysqlSlowlogStatusRUNNING captures enum value "RUNNING"
	AddMySQLOKBodyQANMysqlSlowlogStatusRUNNING string = "RUNNING"

	// AddMySQLOKBodyQANMysqlSlowlogStatusWAITING captures enum value "WAITING"
	AddMySQLOKBodyQANMysqlSlowlogStatusWAITING string = "WAITING"

	// AddMySQLOKBodyQANMysqlSlowlogStatusSTOPPING captures enum value "STOPPING"
	AddMySQLOKBodyQANMysqlSlowlogStatusSTOPPING string = "STOPPING"

	// AddMySQLOKBodyQANMysqlSlowlogStatusDONE captures enum value "DONE"
	AddMySQLOKBodyQANMysqlSlowlogStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMySQLOKBodyQANMysqlSlowlog) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMySqlOkBodyQanMysqlSlowlogTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMySQLOKBodyQANMysqlSlowlog) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMySqlOk"+"."+"qan_mysql_slowlog"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLOKBodyQANMysqlSlowlog) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLOKBodyQANMysqlSlowlog) UnmarshalBinary(b []byte) error {
	var res AddMySQLOKBodyQANMysqlSlowlog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLOKBodyService MySQLService represents a generic MySQL instance.
swagger:model AddMySQLOKBodyService
*/
type AddMySQLOKBodyService struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add my SQL OK body service
func (o *AddMySQLOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddMySQLOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddMySQLParamsBodyAddNode
*/
type AddMySQLParamsBodyAddNode struct {

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add my SQL params body add node
func (o *AddMySQLParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySqlParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySqlParamsBodyAddNodeTypeNodeTypePropEnum = append(addMySqlParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddMySQLParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddMySQLParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddMySQLParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddMySQLParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddMySQLParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddMySQLParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddMySQLParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddMySQLParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"
)

// prop value enum
func (o *AddMySQLParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMySqlParamsBodyAddNodeTypeNodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMySQLParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"add_node"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddMySQLParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
