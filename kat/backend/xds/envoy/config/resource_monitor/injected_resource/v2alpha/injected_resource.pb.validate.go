// Code generated by protoc-gen-validate
// source: envoy/config/resource_monitor/injected_resource/v2alpha/injected_resource.proto
// DO NOT EDIT!!!

package v2alpha

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
)

// Validate checks the field values on InjectedResourceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InjectedResourceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilename()) < 1 {
		return InjectedResourceConfigValidationError{
			Field:  "Filename",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// InjectedResourceConfigValidationError is the validation error returned by
// InjectedResourceConfig.Validate if the designated constraints aren't met.
type InjectedResourceConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e InjectedResourceConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInjectedResourceConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = InjectedResourceConfigValidationError{}