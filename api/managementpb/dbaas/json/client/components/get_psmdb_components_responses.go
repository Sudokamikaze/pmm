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

// GetPSMDBComponentsReader is a Reader for the GetPSMDBComponents structure.
type GetPSMDBComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPSMDBComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPSMDBComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPSMDBComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPSMDBComponentsOK creates a GetPSMDBComponentsOK with default headers values
func NewGetPSMDBComponentsOK() *GetPSMDBComponentsOK {
	return &GetPSMDBComponentsOK{}
}

/*GetPSMDBComponentsOK handles this case with default header values.

A successful response.
*/
type GetPSMDBComponentsOK struct {
	Payload *GetPSMDBComponentsOKBody
}

func (o *GetPSMDBComponentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] getPsmdbComponentsOk  %+v", 200, o.Payload)
}

func (o *GetPSMDBComponentsOK) GetPayload() *GetPSMDBComponentsOKBody {
	return o.Payload
}

func (o *GetPSMDBComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPSMDBComponentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPSMDBComponentsDefault creates a GetPSMDBComponentsDefault with default headers values
func NewGetPSMDBComponentsDefault(code int) *GetPSMDBComponentsDefault {
	return &GetPSMDBComponentsDefault{
		_statusCode: code,
	}
}

/*GetPSMDBComponentsDefault handles this case with default header values.

An unexpected error response.
*/
type GetPSMDBComponentsDefault struct {
	_statusCode int

	Payload *GetPSMDBComponentsDefaultBody
}

// Code gets the status code for the get PSMDB components default response
func (o *GetPSMDBComponentsDefault) Code() int {
	return o._statusCode
}

func (o *GetPSMDBComponentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] GetPSMDBComponents default  %+v", o._statusCode, o.Payload)
}

func (o *GetPSMDBComponentsDefault) GetPayload() *GetPSMDBComponentsDefaultBody {
	return o.Payload
}

func (o *GetPSMDBComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPSMDBComponentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

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

/*GetPSMDBComponentsBody get PSMDB components body
swagger:model GetPSMDBComponentsBody
*/
type GetPSMDBComponentsBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// Version of DB.
	DBVersion string `json:"db_version,omitempty"`
}

// Validate validates this get PSMDB components body
func (o *GetPSMDBComponentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsDefaultBody get PSMDB components default body
swagger:model GetPSMDBComponentsDefaultBody
*/
type GetPSMDBComponentsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get PSMDB components default body
func (o *GetPSMDBComponentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetPSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBody get PSMDB components OK body
swagger:model GetPSMDBComponentsOKBody
*/
type GetPSMDBComponentsOKBody struct {

	// versions
	Versions []*VersionsItems0 `json:"versions"`
}

// Validate validates this get PSMDB components OK body
func (o *GetPSMDBComponentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBody) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPsmdbComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0 Version contains information about operator and components matrix.
swagger:model VersionsItems0
*/
type VersionsItems0 struct {

	// product
	Product string `json:"product,omitempty"`

	// operator
	Operator string `json:"operator,omitempty"`

	// matrix
	Matrix *VersionsItems0Matrix `json:"matrix,omitempty"`
}

// Validate validates this versions items0
func (o *VersionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionsItems0) validateMatrix(formats strfmt.Registry) error {

	if swag.IsZero(o.Matrix) { // not required
		return nil
	}

	if o.Matrix != nil {
		if err := o.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0) UnmarshalBinary(b []byte) error {
	var res VersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0Matrix Matrix contains all available components.
swagger:model VersionsItems0Matrix
*/
type VersionsItems0Matrix struct {

	// mongod
	Mongod map[string]VersionsItems0MatrixMongodAnon `json:"mongod,omitempty"`

	// pxc
	Pxc map[string]VersionsItems0MatrixPxcAnon `json:"pxc,omitempty"`

	// pmm
	PMM map[string]VersionsItems0MatrixPMMAnon `json:"pmm,omitempty"`

	// proxysql
	Proxysql map[string]VersionsItems0MatrixProxysqlAnon `json:"proxysql,omitempty"`

	// haproxy
	Haproxy map[string]VersionsItems0MatrixHaproxyAnon `json:"haproxy,omitempty"`

	// backup
	Backup map[string]VersionsItems0MatrixBackupAnon `json:"backup,omitempty"`

	// operator
	Operator map[string]VersionsItems0MatrixOperatorAnon `json:"operator,omitempty"`

	// log collector
	LogCollector map[string]VersionsItems0MatrixLogCollectorAnon `json:"log_collector,omitempty"`
}

// Validate validates this versions items0 matrix
func (o *VersionsItems0Matrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMM(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogCollector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionsItems0Matrix) validateMongod(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongod) { // not required
		return nil
	}

	for k := range o.Mongod {

		if swag.IsZero(o.Mongod[k]) { // not required
			continue
		}
		if val, ok := o.Mongod[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	for k := range o.Pxc {

		if swag.IsZero(o.Pxc[k]) { // not required
			continue
		}
		if val, ok := o.Pxc[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validatePMM(formats strfmt.Registry) error {

	if swag.IsZero(o.PMM) { // not required
		return nil
	}

	for k := range o.PMM {

		if swag.IsZero(o.PMM[k]) { // not required
			continue
		}
		if val, ok := o.PMM[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	for k := range o.Proxysql {

		if swag.IsZero(o.Proxysql[k]) { // not required
			continue
		}
		if val, ok := o.Proxysql[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validateHaproxy(formats strfmt.Registry) error {

	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	for k := range o.Haproxy {

		if swag.IsZero(o.Haproxy[k]) { // not required
			continue
		}
		if val, ok := o.Haproxy[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validateBackup(formats strfmt.Registry) error {

	if swag.IsZero(o.Backup) { // not required
		return nil
	}

	for k := range o.Backup {

		if swag.IsZero(o.Backup[k]) { // not required
			continue
		}
		if val, ok := o.Backup[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(o.Operator) { // not required
		return nil
	}

	for k := range o.Operator {

		if swag.IsZero(o.Operator[k]) { // not required
			continue
		}
		if val, ok := o.Operator[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *VersionsItems0Matrix) validateLogCollector(formats strfmt.Registry) error {

	if swag.IsZero(o.LogCollector) { // not required
		return nil
	}

	for k := range o.LogCollector {

		if swag.IsZero(o.LogCollector[k]) { // not required
			continue
		}
		if val, ok := o.LogCollector[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0Matrix) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0Matrix) UnmarshalBinary(b []byte) error {
	var res VersionsItems0Matrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixBackupAnon Component contains information about component.
swagger:model VersionsItems0MatrixBackupAnon
*/
type VersionsItems0MatrixBackupAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix backup anon
func (o *VersionsItems0MatrixBackupAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixBackupAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixBackupAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixBackupAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixHaproxyAnon Component contains information about component.
swagger:model VersionsItems0MatrixHaproxyAnon
*/
type VersionsItems0MatrixHaproxyAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix haproxy anon
func (o *VersionsItems0MatrixHaproxyAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixHaproxyAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixHaproxyAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixHaproxyAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixLogCollectorAnon Component contains information about component.
swagger:model VersionsItems0MatrixLogCollectorAnon
*/
type VersionsItems0MatrixLogCollectorAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix log collector anon
func (o *VersionsItems0MatrixLogCollectorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixLogCollectorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixLogCollectorAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixLogCollectorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixMongodAnon Component contains information about component.
swagger:model VersionsItems0MatrixMongodAnon
*/
type VersionsItems0MatrixMongodAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix mongod anon
func (o *VersionsItems0MatrixMongodAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixMongodAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixMongodAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixMongodAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixOperatorAnon Component contains information about component.
swagger:model VersionsItems0MatrixOperatorAnon
*/
type VersionsItems0MatrixOperatorAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix operator anon
func (o *VersionsItems0MatrixOperatorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixOperatorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixOperatorAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixOperatorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixPMMAnon Component contains information about component.
swagger:model VersionsItems0MatrixPMMAnon
*/
type VersionsItems0MatrixPMMAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix PMM anon
func (o *VersionsItems0MatrixPMMAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixPMMAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixPMMAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixPMMAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixProxysqlAnon Component contains information about component.
swagger:model VersionsItems0MatrixProxysqlAnon
*/
type VersionsItems0MatrixProxysqlAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix proxysql anon
func (o *VersionsItems0MatrixProxysqlAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixProxysqlAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixProxysqlAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixProxysqlAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionsItems0MatrixPxcAnon Component contains information about component.
swagger:model VersionsItems0MatrixPxcAnon
*/
type VersionsItems0MatrixPxcAnon struct {

	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this versions items0 matrix pxc anon
func (o *VersionsItems0MatrixPxcAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionsItems0MatrixPxcAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionsItems0MatrixPxcAnon) UnmarshalBinary(b []byte) error {
	var res VersionsItems0MatrixPxcAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
