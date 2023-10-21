// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_common/v1/sensitive_word.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on SensitiveWordInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordInfoMultiError, or nil if none found.
func (m *SensitiveWordInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Word

	// no validation rules for Desc

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return SensitiveWordInfoMultiError(errors)
	}

	return nil
}

// SensitiveWordInfoMultiError is an error wrapping multiple validation errors
// returned by SensitiveWordInfo.ValidateAll() if the designated constraints
// aren't met.
type SensitiveWordInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordInfoMultiError) AllErrors() []error { return m }

// SensitiveWordInfoValidationError is the validation error returned by
// SensitiveWordInfo.Validate if the designated constraints aren't met.
type SensitiveWordInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordInfoValidationError) ErrorName() string {
	return "SensitiveWordInfoValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordInfoValidationError{}

// Validate checks the field values on SensitiveWordListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordListReqMultiError, or nil if none found.
func (m *SensitiveWordListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SensitiveWordListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SensitiveWordListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SensitiveWordListReqValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SensitiveWordListReqMultiError(errors)
	}

	return nil
}

// SensitiveWordListReqMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordListReq.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordListReqMultiError) AllErrors() []error { return m }

// SensitiveWordListReqValidationError is the validation error returned by
// SensitiveWordListReq.Validate if the designated constraints aren't met.
type SensitiveWordListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordListReqValidationError) ErrorName() string {
	return "SensitiveWordListReqValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordListReqValidationError{}

// Validate checks the field values on SensitiveWordListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordListReplyMultiError, or nil if none found.
func (m *SensitiveWordListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SensitiveWordListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SensitiveWordListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SensitiveWordListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SensitiveWordListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SensitiveWordListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SensitiveWordListReplyValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SensitiveWordListReplyMultiError(errors)
	}

	return nil
}

// SensitiveWordListReplyMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordListReply.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordListReplyMultiError) AllErrors() []error { return m }

// SensitiveWordListReplyValidationError is the validation error returned by
// SensitiveWordListReply.Validate if the designated constraints aren't met.
type SensitiveWordListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordListReplyValidationError) ErrorName() string {
	return "SensitiveWordListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordListReplyValidationError{}

// Validate checks the field values on SensitiveWordStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordStoreReqMultiError, or nil if none found.
func (m *SensitiveWordStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Word

	// no validation rules for Desc

	if len(errors) > 0 {
		return SensitiveWordStoreReqMultiError(errors)
	}

	return nil
}

// SensitiveWordStoreReqMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordStoreReq.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordStoreReqMultiError) AllErrors() []error { return m }

// SensitiveWordStoreReqValidationError is the validation error returned by
// SensitiveWordStoreReq.Validate if the designated constraints aren't met.
type SensitiveWordStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordStoreReqValidationError) ErrorName() string {
	return "SensitiveWordStoreReqValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordStoreReqValidationError{}

// Validate checks the field values on SensitiveWordStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordStoreReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordStoreReplyMultiError, or nil if none found.
func (m *SensitiveWordStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SensitiveWordStoreReplyMultiError(errors)
	}

	return nil
}

// SensitiveWordStoreReplyMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordStoreReply.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordStoreReplyMultiError) AllErrors() []error { return m }

// SensitiveWordStoreReplyValidationError is the validation error returned by
// SensitiveWordStoreReply.Validate if the designated constraints aren't met.
type SensitiveWordStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordStoreReplyValidationError) ErrorName() string {
	return "SensitiveWordStoreReplyValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordStoreReplyValidationError{}

// Validate checks the field values on SensitiveWordDelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordDelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordDelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordDelReqMultiError, or nil if none found.
func (m *SensitiveWordDelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordDelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SensitiveWordDelReqMultiError(errors)
	}

	return nil
}

// SensitiveWordDelReqMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordDelReq.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordDelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordDelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordDelReqMultiError) AllErrors() []error { return m }

// SensitiveWordDelReqValidationError is the validation error returned by
// SensitiveWordDelReq.Validate if the designated constraints aren't met.
type SensitiveWordDelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordDelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordDelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordDelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordDelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordDelReqValidationError) ErrorName() string {
	return "SensitiveWordDelReqValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordDelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordDelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordDelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordDelReqValidationError{}

// Validate checks the field values on SensitiveWordDelReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordDelReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordDelReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordDelReplyMultiError, or nil if none found.
func (m *SensitiveWordDelReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordDelReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SensitiveWordDelReplyMultiError(errors)
	}

	return nil
}

// SensitiveWordDelReplyMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordDelReply.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordDelReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordDelReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordDelReplyMultiError) AllErrors() []error { return m }

// SensitiveWordDelReplyValidationError is the validation error returned by
// SensitiveWordDelReply.Validate if the designated constraints aren't met.
type SensitiveWordDelReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordDelReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordDelReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordDelReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordDelReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordDelReplyValidationError) ErrorName() string {
	return "SensitiveWordDelReplyValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordDelReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordDelReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordDelReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordDelReplyValidationError{}

// Validate checks the field values on SensitiveWordCheckReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordCheckReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordCheckReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordCheckReqMultiError, or nil if none found.
func (m *SensitiveWordCheckReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordCheckReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Word

	if len(errors) > 0 {
		return SensitiveWordCheckReqMultiError(errors)
	}

	return nil
}

// SensitiveWordCheckReqMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordCheckReq.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordCheckReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordCheckReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordCheckReqMultiError) AllErrors() []error { return m }

// SensitiveWordCheckReqValidationError is the validation error returned by
// SensitiveWordCheckReq.Validate if the designated constraints aren't met.
type SensitiveWordCheckReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordCheckReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordCheckReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordCheckReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordCheckReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordCheckReqValidationError) ErrorName() string {
	return "SensitiveWordCheckReqValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordCheckReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordCheckReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordCheckReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordCheckReqValidationError{}

// Validate checks the field values on SensitiveWordCheckResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SensitiveWordCheckResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SensitiveWordCheckResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SensitiveWordCheckRespMultiError, or nil if none found.
func (m *SensitiveWordCheckResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SensitiveWordCheckResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	// no validation rules for Replace

	// no validation rules for Filter

	if len(errors) > 0 {
		return SensitiveWordCheckRespMultiError(errors)
	}

	return nil
}

// SensitiveWordCheckRespMultiError is an error wrapping multiple validation
// errors returned by SensitiveWordCheckResp.ValidateAll() if the designated
// constraints aren't met.
type SensitiveWordCheckRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SensitiveWordCheckRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SensitiveWordCheckRespMultiError) AllErrors() []error { return m }

// SensitiveWordCheckRespValidationError is the validation error returned by
// SensitiveWordCheckResp.Validate if the designated constraints aren't met.
type SensitiveWordCheckRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SensitiveWordCheckRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SensitiveWordCheckRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SensitiveWordCheckRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SensitiveWordCheckRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SensitiveWordCheckRespValidationError) ErrorName() string {
	return "SensitiveWordCheckRespValidationError"
}

// Error satisfies the builtin error interface
func (e SensitiveWordCheckRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSensitiveWordCheckResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SensitiveWordCheckRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SensitiveWordCheckRespValidationError{}
