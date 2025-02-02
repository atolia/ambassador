// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

package endpoint

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"

	core "github.com/datawire/ambassador/go/apis/envoy/api/v2/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}

	_ = core.HealthStatus(0)
)

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Endpoint) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAddress()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return EndpointValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetHealthCheckConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return EndpointValidationError{
					field:  "HealthCheckConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// EndpointValidationError is the validation error returned by
// Endpoint.Validate if the designated constraints aren't met.
type EndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointValidationError) ErrorName() string { return "EndpointValidationError" }

// Error satisfies the builtin error interface
func (e EndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointValidationError{}

// Validate checks the field values on LbEndpoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LbEndpoint) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for HealthStatus

	{
		tmp := m.GetMetadata()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return LbEndpointValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetLoadBalancingWeight(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return LbEndpointValidationError{
				field:  "LoadBalancingWeight",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	switch m.HostIdentifier.(type) {

	case *LbEndpoint_Endpoint:

		{
			tmp := m.GetEndpoint()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return LbEndpointValidationError{
						field:  "Endpoint",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *LbEndpoint_EndpointName:
		// no validation rules for EndpointName

	}

	return nil
}

// LbEndpointValidationError is the validation error returned by
// LbEndpoint.Validate if the designated constraints aren't met.
type LbEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LbEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LbEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LbEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LbEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LbEndpointValidationError) ErrorName() string { return "LbEndpointValidationError" }

// Error satisfies the builtin error interface
func (e LbEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLbEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LbEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LbEndpointValidationError{}

// Validate checks the field values on LocalityLbEndpoints with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LocalityLbEndpoints) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetLocality()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return LocalityLbEndpointsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetLbEndpoints() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return LocalityLbEndpointsValidationError{
						field:  fmt.Sprintf("LbEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	if wrapper := m.GetLoadBalancingWeight(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return LocalityLbEndpointsValidationError{
				field:  "LoadBalancingWeight",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if m.GetPriority() > 128 {
		return LocalityLbEndpointsValidationError{
			field:  "Priority",
			reason: "value must be less than or equal to 128",
		}
	}

	{
		tmp := m.GetProximity()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return LocalityLbEndpointsValidationError{
					field:  "Proximity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// LocalityLbEndpointsValidationError is the validation error returned by
// LocalityLbEndpoints.Validate if the designated constraints aren't met.
type LocalityLbEndpointsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbEndpointsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbEndpointsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbEndpointsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbEndpointsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbEndpointsValidationError) ErrorName() string {
	return "LocalityLbEndpointsValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityLbEndpointsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbEndpoints.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbEndpointsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbEndpointsValidationError{}

// Validate checks the field values on Endpoint_HealthCheckConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Endpoint_HealthCheckConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPortValue() > 65535 {
		return Endpoint_HealthCheckConfigValidationError{
			field:  "PortValue",
			reason: "value must be less than or equal to 65535",
		}
	}

	return nil
}

// Endpoint_HealthCheckConfigValidationError is the validation error returned
// by Endpoint_HealthCheckConfig.Validate if the designated constraints aren't met.
type Endpoint_HealthCheckConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Endpoint_HealthCheckConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Endpoint_HealthCheckConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Endpoint_HealthCheckConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Endpoint_HealthCheckConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Endpoint_HealthCheckConfigValidationError) ErrorName() string {
	return "Endpoint_HealthCheckConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Endpoint_HealthCheckConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint_HealthCheckConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Endpoint_HealthCheckConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Endpoint_HealthCheckConfigValidationError{}
