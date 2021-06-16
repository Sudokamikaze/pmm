// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/components.proto

package dbaasv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Component) Validate() error {
	return nil
}
func (this *Matrix) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *OperatorVersion) Validate() error {
	if this.Matrix != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Matrix); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Matrix", err)
		}
	}
	return nil
}
func (this *GetPSMDBComponentsRequest) Validate() error {
	return nil
}
func (this *GetPSMDBComponentsResponse) Validate() error {
	for _, item := range this.Versions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Versions", err)
			}
		}
	}
	return nil
}
func (this *GetPXCComponentsRequest) Validate() error {
	return nil
}
func (this *GetPXCComponentsResponse) Validate() error {
	for _, item := range this.Versions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Versions", err)
			}
		}
	}
	return nil
}
func (this *ChangeComponent) Validate() error {
	for _, item := range this.Versions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Versions", err)
			}
		}
	}
	return nil
}
func (this *ChangeComponent_ComponentVersion) Validate() error {
	if this.Version == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must not be an empty string`, this.Version))
	}
	return nil
}
func (this *ChangePSMDBComponentsRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Mongod != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mongod); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mongod", err)
		}
	}
	return nil
}
func (this *ChangePSMDBComponentsResponse) Validate() error {
	return nil
}
func (this *ChangePXCComponentsRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	if this.Haproxy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Haproxy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
		}
	}
	return nil
}
func (this *ChangePXCComponentsResponse) Validate() error {
	return nil
}
func (this *CheckForOperatorUpdateRequest) Validate() error {
	return nil
}
func (this *OperatorsUpdateInformation) Validate() error {
	return nil
}
func (this *AvailableComponentsVersions) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CheckForOperatorUpdateResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
