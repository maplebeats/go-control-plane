// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v3/tls_spiffe_validator_config.proto

package envoy_extensions_transport_sockets_tls_v3

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on SPIFFECertValidatorConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SPIFFECertValidatorConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetTrustDomains()) < 1 {
		return SPIFFECertValidatorConfigValidationError{
			field:  "TrustDomains",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTrustDomains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SPIFFECertValidatorConfigValidationError{
					field:  fmt.Sprintf("TrustDomains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SPIFFECertValidatorConfigValidationError is the validation error returned by
// SPIFFECertValidatorConfig.Validate if the designated constraints aren't met.
type SPIFFECertValidatorConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SPIFFECertValidatorConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SPIFFECertValidatorConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SPIFFECertValidatorConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SPIFFECertValidatorConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SPIFFECertValidatorConfigValidationError) ErrorName() string {
	return "SPIFFECertValidatorConfigValidationError"
}

// Error satisfies the builtin error interface
func (e SPIFFECertValidatorConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSPIFFECertValidatorConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SPIFFECertValidatorConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SPIFFECertValidatorConfigValidationError{}

// Validate checks the field values on SPIFFECertValidatorConfig_TrustDomain
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *SPIFFECertValidatorConfig_TrustDomain) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return SPIFFECertValidatorConfig_TrustDomainValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetTrustBundle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SPIFFECertValidatorConfig_TrustDomainValidationError{
				field:  "TrustBundle",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SPIFFECertValidatorConfig_TrustDomainValidationError is the validation error
// returned by SPIFFECertValidatorConfig_TrustDomain.Validate if the
// designated constraints aren't met.
type SPIFFECertValidatorConfig_TrustDomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) ErrorName() string {
	return "SPIFFECertValidatorConfig_TrustDomainValidationError"
}

// Error satisfies the builtin error interface
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSPIFFECertValidatorConfig_TrustDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SPIFFECertValidatorConfig_TrustDomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SPIFFECertValidatorConfig_TrustDomainValidationError{}
