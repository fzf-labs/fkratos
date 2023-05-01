// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_sys/v1/sys_job.proto

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

// Validate checks the field values on SysJobInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobInfoMultiError, or
// nil if none found.
func (m *SysJobInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Code

	// no validation rules for Remark

	// no validation rules for Status

	// no validation rules for Sort

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return SysJobInfoMultiError(errors)
	}

	return nil
}

// SysJobInfoMultiError is an error wrapping multiple validation errors
// returned by SysJobInfo.ValidateAll() if the designated constraints aren't met.
type SysJobInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobInfoMultiError) AllErrors() []error { return m }

// SysJobInfoValidationError is the validation error returned by
// SysJobInfo.Validate if the designated constraints aren't met.
type SysJobInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobInfoValidationError) ErrorName() string { return "SysJobInfoValidationError" }

// Error satisfies the builtin error interface
func (e SysJobInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobInfoValidationError{}

// Validate checks the field values on SysJobListReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobListReqMultiError, or
// nil if none found.
func (m *SysJobListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for QuickSearch

	for idx, item := range m.GetSearch() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SysJobListReqValidationError{
						field:  fmt.Sprintf("Search[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysJobListReqValidationError{
						field:  fmt.Sprintf("Search[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysJobListReqValidationError{
					field:  fmt.Sprintf("Search[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Order

	if len(errors) > 0 {
		return SysJobListReqMultiError(errors)
	}

	return nil
}

// SysJobListReqMultiError is an error wrapping multiple validation errors
// returned by SysJobListReq.ValidateAll() if the designated constraints
// aren't met.
type SysJobListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobListReqMultiError) AllErrors() []error { return m }

// SysJobListReqValidationError is the validation error returned by
// SysJobListReq.Validate if the designated constraints aren't met.
type SysJobListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobListReqValidationError) ErrorName() string { return "SysJobListReqValidationError" }

// Error satisfies the builtin error interface
func (e SysJobListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobListReqValidationError{}

// Validate checks the field values on SysJobListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysJobListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysJobListReplyMultiError, or nil if none found.
func (m *SysJobListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobListReply) validate(all bool) error {
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
					errors = append(errors, SysJobListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysJobListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysJobListReplyValidationError{
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
				errors = append(errors, SysJobListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysJobListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysJobListReplyValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysJobListReplyMultiError(errors)
	}

	return nil
}

// SysJobListReplyMultiError is an error wrapping multiple validation errors
// returned by SysJobListReply.ValidateAll() if the designated constraints
// aren't met.
type SysJobListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobListReplyMultiError) AllErrors() []error { return m }

// SysJobListReplyValidationError is the validation error returned by
// SysJobListReply.Validate if the designated constraints aren't met.
type SysJobListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobListReplyValidationError) ErrorName() string { return "SysJobListReplyValidationError" }

// Error satisfies the builtin error interface
func (e SysJobListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobListReplyValidationError{}

// Validate checks the field values on SysJobInfoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobInfoReqMultiError, or
// nil if none found.
func (m *SysJobInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysJobInfoReqMultiError(errors)
	}

	return nil
}

// SysJobInfoReqMultiError is an error wrapping multiple validation errors
// returned by SysJobInfoReq.ValidateAll() if the designated constraints
// aren't met.
type SysJobInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobInfoReqMultiError) AllErrors() []error { return m }

// SysJobInfoReqValidationError is the validation error returned by
// SysJobInfoReq.Validate if the designated constraints aren't met.
type SysJobInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobInfoReqValidationError) ErrorName() string { return "SysJobInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e SysJobInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobInfoReqValidationError{}

// Validate checks the field values on SysJobInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysJobInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobInfoReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysJobInfoReplyMultiError, or nil if none found.
func (m *SysJobInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysJobInfoReplyValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysJobInfoReplyValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysJobInfoReplyValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysJobInfoReplyMultiError(errors)
	}

	return nil
}

// SysJobInfoReplyMultiError is an error wrapping multiple validation errors
// returned by SysJobInfoReply.ValidateAll() if the designated constraints
// aren't met.
type SysJobInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobInfoReplyMultiError) AllErrors() []error { return m }

// SysJobInfoReplyValidationError is the validation error returned by
// SysJobInfoReply.Validate if the designated constraints aren't met.
type SysJobInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobInfoReplyValidationError) ErrorName() string { return "SysJobInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e SysJobInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobInfoReplyValidationError{}

// Validate checks the field values on SysJobStoreReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobStoreReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobStoreReqMultiError,
// or nil if none found.
func (m *SysJobStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Code

	// no validation rules for Remark

	// no validation rules for Status

	// no validation rules for Sort

	if len(errors) > 0 {
		return SysJobStoreReqMultiError(errors)
	}

	return nil
}

// SysJobStoreReqMultiError is an error wrapping multiple validation errors
// returned by SysJobStoreReq.ValidateAll() if the designated constraints
// aren't met.
type SysJobStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobStoreReqMultiError) AllErrors() []error { return m }

// SysJobStoreReqValidationError is the validation error returned by
// SysJobStoreReq.Validate if the designated constraints aren't met.
type SysJobStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobStoreReqValidationError) ErrorName() string { return "SysJobStoreReqValidationError" }

// Error satisfies the builtin error interface
func (e SysJobStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobStoreReqValidationError{}

// Validate checks the field values on SysJobStoreReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysJobStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysJobStoreReplyMultiError, or nil if none found.
func (m *SysJobStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysJobStoreReplyMultiError(errors)
	}

	return nil
}

// SysJobStoreReplyMultiError is an error wrapping multiple validation errors
// returned by SysJobStoreReply.ValidateAll() if the designated constraints
// aren't met.
type SysJobStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobStoreReplyMultiError) AllErrors() []error { return m }

// SysJobStoreReplyValidationError is the validation error returned by
// SysJobStoreReply.Validate if the designated constraints aren't met.
type SysJobStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobStoreReplyValidationError) ErrorName() string { return "SysJobStoreReplyValidationError" }

// Error satisfies the builtin error interface
func (e SysJobStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobStoreReplyValidationError{}

// Validate checks the field values on SysJobDelReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobDelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobDelReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobDelReqMultiError, or
// nil if none found.
func (m *SysJobDelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobDelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysJobDelReqMultiError(errors)
	}

	return nil
}

// SysJobDelReqMultiError is an error wrapping multiple validation errors
// returned by SysJobDelReq.ValidateAll() if the designated constraints aren't met.
type SysJobDelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobDelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobDelReqMultiError) AllErrors() []error { return m }

// SysJobDelReqValidationError is the validation error returned by
// SysJobDelReq.Validate if the designated constraints aren't met.
type SysJobDelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobDelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobDelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobDelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobDelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobDelReqValidationError) ErrorName() string { return "SysJobDelReqValidationError" }

// Error satisfies the builtin error interface
func (e SysJobDelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobDelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobDelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobDelReqValidationError{}

// Validate checks the field values on SysJobDelReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysJobDelReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysJobDelReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysJobDelReplyMultiError,
// or nil if none found.
func (m *SysJobDelReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SysJobDelReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysJobDelReplyMultiError(errors)
	}

	return nil
}

// SysJobDelReplyMultiError is an error wrapping multiple validation errors
// returned by SysJobDelReply.ValidateAll() if the designated constraints
// aren't met.
type SysJobDelReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysJobDelReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysJobDelReplyMultiError) AllErrors() []error { return m }

// SysJobDelReplyValidationError is the validation error returned by
// SysJobDelReply.Validate if the designated constraints aren't met.
type SysJobDelReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysJobDelReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysJobDelReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysJobDelReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysJobDelReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysJobDelReplyValidationError) ErrorName() string { return "SysJobDelReplyValidationError" }

// Error satisfies the builtin error interface
func (e SysJobDelReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysJobDelReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysJobDelReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysJobDelReplyValidationError{}
