// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_sys/v1/sys_log.proto

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

// Validate checks the field values on SysLog with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLog with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SysLogMultiError, or nil if none found.
func (m *SysLog) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AdminID

	// no validation rules for Username

	// no validation rules for Ip

	// no validation rules for Uri

	// no validation rules for UriDesc

	// no validation rules for Useragent

	// no validation rules for HttpCode

	// no validation rules for Req

	// no validation rules for Resp

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return SysLogMultiError(errors)
	}

	return nil
}

// SysLogMultiError is an error wrapping multiple validation errors returned by
// SysLog.ValidateAll() if the designated constraints aren't met.
type SysLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLogMultiError) AllErrors() []error { return m }

// SysLogValidationError is the validation error returned by SysLog.Validate if
// the designated constraints aren't met.
type SysLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLogValidationError) ErrorName() string { return "SysLogValidationError" }

// Error satisfies the builtin error interface
func (e SysLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLogValidationError{}

// Validate checks the field values on SysLogListReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysLogListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLogListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysLogListReqMultiError, or
// nil if none found.
func (m *SysLogListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLogListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return SysLogListReqMultiError(errors)
	}

	return nil
}

// SysLogListReqMultiError is an error wrapping multiple validation errors
// returned by SysLogListReq.ValidateAll() if the designated constraints
// aren't met.
type SysLogListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLogListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLogListReqMultiError) AllErrors() []error { return m }

// SysLogListReqValidationError is the validation error returned by
// SysLogListReq.Validate if the designated constraints aren't met.
type SysLogListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLogListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLogListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLogListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLogListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLogListReqValidationError) ErrorName() string { return "SysLogListReqValidationError" }

// Error satisfies the builtin error interface
func (e SysLogListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLogListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLogListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLogListReqValidationError{}

// Validate checks the field values on SysLogListResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysLogListResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLogListResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysLogListRespMultiError,
// or nil if none found.
func (m *SysLogListResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLogListResp) validate(all bool) error {
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
					errors = append(errors, SysLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysLogListRespValidationError{
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
				errors = append(errors, SysLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysLogListRespValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysLogListRespMultiError(errors)
	}

	return nil
}

// SysLogListRespMultiError is an error wrapping multiple validation errors
// returned by SysLogListResp.ValidateAll() if the designated constraints
// aren't met.
type SysLogListRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLogListRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLogListRespMultiError) AllErrors() []error { return m }

// SysLogListRespValidationError is the validation error returned by
// SysLogListResp.Validate if the designated constraints aren't met.
type SysLogListRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLogListRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLogListRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLogListRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLogListRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLogListRespValidationError) ErrorName() string { return "SysLogListRespValidationError" }

// Error satisfies the builtin error interface
func (e SysLogListRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLogListResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLogListRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLogListRespValidationError{}

// Validate checks the field values on SysLogInfoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysLogInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLogInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysLogInfoReqMultiError, or
// nil if none found.
func (m *SysLogInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLogInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysLogInfoReqMultiError(errors)
	}

	return nil
}

// SysLogInfoReqMultiError is an error wrapping multiple validation errors
// returned by SysLogInfoReq.ValidateAll() if the designated constraints
// aren't met.
type SysLogInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLogInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLogInfoReqMultiError) AllErrors() []error { return m }

// SysLogInfoReqValidationError is the validation error returned by
// SysLogInfoReq.Validate if the designated constraints aren't met.
type SysLogInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLogInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLogInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLogInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLogInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLogInfoReqValidationError) ErrorName() string { return "SysLogInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e SysLogInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLogInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLogInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLogInfoReqValidationError{}

// Validate checks the field values on SysLogInfoResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysLogInfoResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLogInfoResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysLogInfoRespMultiError,
// or nil if none found.
func (m *SysLogInfoResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLogInfoResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysLogInfoRespValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysLogInfoRespMultiError(errors)
	}

	return nil
}

// SysLogInfoRespMultiError is an error wrapping multiple validation errors
// returned by SysLogInfoResp.ValidateAll() if the designated constraints
// aren't met.
type SysLogInfoRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLogInfoRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLogInfoRespMultiError) AllErrors() []error { return m }

// SysLogInfoRespValidationError is the validation error returned by
// SysLogInfoResp.Validate if the designated constraints aren't met.
type SysLogInfoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLogInfoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLogInfoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLogInfoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLogInfoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLogInfoRespValidationError) ErrorName() string { return "SysLogInfoRespValidationError" }

// Error satisfies the builtin error interface
func (e SysLogInfoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLogInfoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLogInfoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLogInfoRespValidationError{}
