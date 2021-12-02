// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pg/v1/my.proto

package generatedv1

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

// Validate checks the field values on ToUpperRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ToUpperRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetS()) != 5 {
		return ToUpperRequestValidationError{
			field:  "S",
			reason: "value length must be 5 runes",
		}

	}

	// no validation rules for F

	return nil
}

// ToUpperRequestValidationError is the validation error returned by
// ToUpperRequest.Validate if the designated constraints aren't met.
type ToUpperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ToUpperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ToUpperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ToUpperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ToUpperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ToUpperRequestValidationError) ErrorName() string { return "ToUpperRequestValidationError" }

// Error satisfies the builtin error interface
func (e ToUpperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToUpperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ToUpperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ToUpperRequestValidationError{}

// Validate checks the field values on ToUpperResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ToUpperResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for S

	return nil
}

// ToUpperResponseValidationError is the validation error returned by
// ToUpperResponse.Validate if the designated constraints aren't met.
type ToUpperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ToUpperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ToUpperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ToUpperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ToUpperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ToUpperResponseValidationError) ErrorName() string { return "ToUpperResponseValidationError" }

// Error satisfies the builtin error interface
func (e ToUpperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToUpperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ToUpperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ToUpperResponseValidationError{}

// Validate checks the field values on NoopMsg with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NoopMsg) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetF()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoopMsgValidationError{
				field:  "F",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NoopMsgValidationError is the validation error returned by NoopMsg.Validate
// if the designated constraints aren't met.
type NoopMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoopMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoopMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoopMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoopMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoopMsgValidationError) ErrorName() string { return "NoopMsgValidationError" }

// Error satisfies the builtin error interface
func (e NoopMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoopMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoopMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoopMsgValidationError{}
