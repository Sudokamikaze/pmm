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

// AddAzureDatabaseExporterReader is a Reader for the AddAzureDatabaseExporter structure.
type AddAzureDatabaseExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAzureDatabaseExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAzureDatabaseExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddAzureDatabaseExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAzureDatabaseExporterOK creates a AddAzureDatabaseExporterOK with default headers values
func NewAddAzureDatabaseExporterOK() *AddAzureDatabaseExporterOK {
	return &AddAzureDatabaseExporterOK{}
}

/*AddAzureDatabaseExporterOK handles this case with default header values.

A successful response.
*/
type AddAzureDatabaseExporterOK struct {
	Payload *AddAzureDatabaseExporterOKBody
}

func (o *AddAzureDatabaseExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddAzureDatabaseExporter][%d] addAzureDatabaseExporterOk  %+v", 200, o.Payload)
}

func (o *AddAzureDatabaseExporterOK) GetPayload() *AddAzureDatabaseExporterOKBody {
	return o.Payload
}

func (o *AddAzureDatabaseExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddAzureDatabaseExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAzureDatabaseExporterDefault creates a AddAzureDatabaseExporterDefault with default headers values
func NewAddAzureDatabaseExporterDefault(code int) *AddAzureDatabaseExporterDefault {
	return &AddAzureDatabaseExporterDefault{
		_statusCode: code,
	}
}

/*AddAzureDatabaseExporterDefault handles this case with default header values.

An unexpected error response.
*/
type AddAzureDatabaseExporterDefault struct {
	_statusCode int

	Payload *AddAzureDatabaseExporterDefaultBody
}

// Code gets the status code for the add azure database exporter default response
func (o *AddAzureDatabaseExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddAzureDatabaseExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddAzureDatabaseExporter][%d] AddAzureDatabaseExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddAzureDatabaseExporterDefault) GetPayload() *AddAzureDatabaseExporterDefaultBody {
	return o.Payload
}

func (o *AddAzureDatabaseExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddAzureDatabaseExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddAzureDatabaseExporterBody add azure database exporter body
swagger:model AddAzureDatabaseExporterBody
*/
type AddAzureDatabaseExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// Azure client ID
	AzureClientID string `json:"azure_client_id,omitempty"`

	// Azure client secret
	AzureClientSecret string `json:"azure_client_secret,omitempty"`

	// Azure tanant ID
	AzureTenantID string `json:"azure_tenant_id,omitempty"`

	// Azure subscription ID
	AzureSubscriptionID string `json:"azure_subscription_id,omitempty"`

	// Azure resource group.
	AzureResourceGroup string `json:"azure_resource_group,omitempty"`

	// Azure resource type (mysql, maria, postgres)
	AzureDatabaseResourceType string `json:"azure_database_resource_type,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`
}

// Validate validates this add azure database exporter body
func (o *AddAzureDatabaseExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseExporterBody) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddAzureDatabaseExporterDefaultBody add azure database exporter default body
swagger:model AddAzureDatabaseExporterDefaultBody
*/
type AddAzureDatabaseExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add azure database exporter default body
func (o *AddAzureDatabaseExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAzureDatabaseExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddAzureDatabaseExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddAzureDatabaseExporterOKBody add azure database exporter OK body
swagger:model AddAzureDatabaseExporterOKBody
*/
type AddAzureDatabaseExporterOKBody struct {

	// azure database exporter
	AzureDatabaseExporter *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter `json:"azure_database_exporter,omitempty"`
}

// Validate validates this add azure database exporter OK body
func (o *AddAzureDatabaseExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureDatabaseExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddAzureDatabaseExporterOKBody) validateAzureDatabaseExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.AzureDatabaseExporter) { // not required
		return nil
	}

	if o.AzureDatabaseExporter != nil {
		if err := o.AzureDatabaseExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addAzureDatabaseExporterOk" + "." + "azure_database_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddAzureDatabaseExporterOKBodyAzureDatabaseExporter AzureDatabaseExporter runs on Generic or Container Node and exposes RemoteAzure Node metrics.
swagger:model AddAzureDatabaseExporterOKBodyAzureDatabaseExporter
*/
type AddAzureDatabaseExporterOKBodyAzureDatabaseExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// Azure database subscription ID.
	AzureDatabaseSubscriptionID string `json:"azure_database_subscription_id,omitempty"`

	// Azure database resource type (mysql, maria, postgres)
	AzureDatabaseResourceType string `json:"azure_database_resource_type,omitempty"`

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

	// Listen port for scraping metrics (the same for several configurations).
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if the exporter operates in push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`
}

// Validate validates this add azure database exporter OK body azure database exporter
func (o *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum = append(addAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTARTING captures enum value "STARTING"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTARTING string = "STARTING"

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusRUNNING captures enum value "RUNNING"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusRUNNING string = "RUNNING"

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusWAITING captures enum value "WAITING"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusWAITING string = "WAITING"

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTOPPING captures enum value "STOPPING"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTOPPING string = "STOPPING"

	// AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusDONE captures enum value "DONE"
	AddAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addAzureDatabaseExporterOk"+"."+"azure_database_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddAzureDatabaseExporterOKBodyAzureDatabaseExporter) UnmarshalBinary(b []byte) error {
	var res AddAzureDatabaseExporterOKBodyAzureDatabaseExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
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
