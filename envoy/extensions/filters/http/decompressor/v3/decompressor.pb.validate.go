// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/decompressor/v3/decompressor.proto

package envoy_extensions_filters_http_decompressor_v3

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

// Validate checks the field values on Decompressor with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Decompressor) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDecompressorLibrary() == nil {
		return DecompressorValidationError{
			field:  "DecompressorLibrary",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDecompressorLibrary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecompressorValidationError{
				field:  "DecompressorLibrary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequestDirectionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecompressorValidationError{
				field:  "RequestDirectionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponseDirectionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecompressorValidationError{
				field:  "ResponseDirectionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DecompressorValidationError is the validation error returned by
// Decompressor.Validate if the designated constraints aren't met.
type DecompressorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecompressorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecompressorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecompressorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecompressorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecompressorValidationError) ErrorName() string { return "DecompressorValidationError" }

// Error satisfies the builtin error interface
func (e DecompressorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecompressor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecompressorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecompressorValidationError{}

// Validate checks the field values on Decompressor_CommonDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *Decompressor_CommonDirectionConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Decompressor_CommonDirectionConfigValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Decompressor_CommonDirectionConfigValidationError is the validation error
// returned by Decompressor_CommonDirectionConfig.Validate if the designated
// constraints aren't met.
type Decompressor_CommonDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Decompressor_CommonDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Decompressor_CommonDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Decompressor_CommonDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Decompressor_CommonDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Decompressor_CommonDirectionConfigValidationError) ErrorName() string {
	return "Decompressor_CommonDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Decompressor_CommonDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecompressor_CommonDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Decompressor_CommonDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Decompressor_CommonDirectionConfigValidationError{}

// Validate checks the field values on Decompressor_RequestDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *Decompressor_RequestDirectionConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Decompressor_RequestDirectionConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAdvertiseAcceptEncoding()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Decompressor_RequestDirectionConfigValidationError{
				field:  "AdvertiseAcceptEncoding",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Decompressor_RequestDirectionConfigValidationError is the validation error
// returned by Decompressor_RequestDirectionConfig.Validate if the designated
// constraints aren't met.
type Decompressor_RequestDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Decompressor_RequestDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Decompressor_RequestDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Decompressor_RequestDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Decompressor_RequestDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Decompressor_RequestDirectionConfigValidationError) ErrorName() string {
	return "Decompressor_RequestDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Decompressor_RequestDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecompressor_RequestDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Decompressor_RequestDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Decompressor_RequestDirectionConfigValidationError{}

// Validate checks the field values on Decompressor_ResponseDirectionConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *Decompressor_ResponseDirectionConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Decompressor_ResponseDirectionConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Decompressor_ResponseDirectionConfigValidationError is the validation error
// returned by Decompressor_ResponseDirectionConfig.Validate if the designated
// constraints aren't met.
type Decompressor_ResponseDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Decompressor_ResponseDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Decompressor_ResponseDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Decompressor_ResponseDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Decompressor_ResponseDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Decompressor_ResponseDirectionConfigValidationError) ErrorName() string {
	return "Decompressor_ResponseDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Decompressor_ResponseDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecompressor_ResponseDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Decompressor_ResponseDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Decompressor_ResponseDirectionConfigValidationError{}
