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

// ChangeRDSExporterReader is a Reader for the ChangeRDSExporter structure.
type ChangeRDSExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeRDSExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeRDSExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeRDSExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeRDSExporterOK creates a ChangeRDSExporterOK with default headers values
func NewChangeRDSExporterOK() *ChangeRDSExporterOK {
	return &ChangeRDSExporterOK{}
}

/*ChangeRDSExporterOK handles this case with default header values.

A successful response.
*/
type ChangeRDSExporterOK struct {
	Payload *ChangeRDSExporterOKBody
}

func (o *ChangeRDSExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeRDSExporter][%d] changeRdsExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeRDSExporterOK) GetPayload() *ChangeRDSExporterOKBody {
	return o.Payload
}

func (o *ChangeRDSExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeRDSExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeRDSExporterDefault creates a ChangeRDSExporterDefault with default headers values
func NewChangeRDSExporterDefault(code int) *ChangeRDSExporterDefault {
	return &ChangeRDSExporterDefault{
		_statusCode: code,
	}
}

/*ChangeRDSExporterDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeRDSExporterDefault struct {
	_statusCode int

	Payload *ChangeRDSExporterDefaultBody
}

// Code gets the status code for the change RDS exporter default response
func (o *ChangeRDSExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeRDSExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeRDSExporter][%d] ChangeRDSExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeRDSExporterDefault) GetPayload() *ChangeRDSExporterDefaultBody {
	return o.Payload
}

func (o *ChangeRDSExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeRDSExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeRDSExporterBody change RDS exporter body
swagger:model ChangeRDSExporterBody
*/
type ChangeRDSExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeRDSExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change RDS exporter body
func (o *ChangeRDSExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeRDSExporterBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeRDSExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeRDSExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeRDSExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeRDSExporterDefaultBody change RDS exporter default body
swagger:model ChangeRDSExporterDefaultBody
*/
type ChangeRDSExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change RDS exporter default body
func (o *ChangeRDSExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeRDSExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeRDSExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeRDSExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeRDSExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeRDSExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeRDSExporterOKBody change RDS exporter OK body
swagger:model ChangeRDSExporterOKBody
*/
type ChangeRDSExporterOKBody struct {

	// rds exporter
	RDSExporter *ChangeRDSExporterOKBodyRDSExporter `json:"rds_exporter,omitempty"`
}

// Validate validates this change RDS exporter OK body
func (o *ChangeRDSExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeRDSExporterOKBody) validateRDSExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	if o.RDSExporter != nil {
		if err := o.RDSExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeRdsExporterOk" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeRDSExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeRDSExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeRDSExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeRDSExporterOKBodyRDSExporter RDSExporter runs on Generic or Container Node and exposes RemoteRDS Node metrics.
swagger:model ChangeRDSExporterOKBodyRDSExporter
*/
type ChangeRDSExporterOKBodyRDSExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

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

	// Listen port for scraping metrics (the same for several configurations).
	ListenPort int64 `json:"listen_port,omitempty"`

	// Basic metrics are disabled.
	BasicMetricsDisabled bool `json:"basic_metrics_disabled,omitempty"`

	// Enhanced metrics are disabled.
	EnhancedMetricsDisabled bool `json:"enhanced_metrics_disabled,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`
}

// Validate validates this change RDS exporter OK body RDS exporter
func (o *ChangeRDSExporterOKBodyRDSExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeRdsExporterOkBodyRdsExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeRdsExporterOkBodyRdsExporterTypeStatusPropEnum = append(changeRdsExporterOkBodyRdsExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeRDSExporterOKBodyRDSExporterStatusSTARTING captures enum value "STARTING"
	ChangeRDSExporterOKBodyRDSExporterStatusSTARTING string = "STARTING"

	// ChangeRDSExporterOKBodyRDSExporterStatusRUNNING captures enum value "RUNNING"
	ChangeRDSExporterOKBodyRDSExporterStatusRUNNING string = "RUNNING"

	// ChangeRDSExporterOKBodyRDSExporterStatusWAITING captures enum value "WAITING"
	ChangeRDSExporterOKBodyRDSExporterStatusWAITING string = "WAITING"

	// ChangeRDSExporterOKBodyRDSExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeRDSExporterOKBodyRDSExporterStatusSTOPPING string = "STOPPING"

	// ChangeRDSExporterOKBodyRDSExporterStatusDONE captures enum value "DONE"
	ChangeRDSExporterOKBodyRDSExporterStatusDONE string = "DONE"

	// ChangeRDSExporterOKBodyRDSExporterStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeRDSExporterOKBodyRDSExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeRDSExporterOKBodyRDSExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeRdsExporterOkBodyRdsExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeRDSExporterOKBodyRDSExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeRdsExporterOk"+"."+"rds_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeRDSExporterOKBodyRDSExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeRDSExporterOKBodyRDSExporter) UnmarshalBinary(b []byte) error {
	var res ChangeRDSExporterOKBodyRDSExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeRDSExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeRDSExporterParamsBodyCommon
*/
type ChangeRDSExporterParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change RDS exporter params body common
func (o *ChangeRDSExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeRDSExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeRDSExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeRDSExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
