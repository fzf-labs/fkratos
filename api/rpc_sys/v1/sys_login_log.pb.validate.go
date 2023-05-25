// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_sys/v1/sys_login_log.proto

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

// Validate checks the field values on SysLoginLogInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogInfoMultiError, or nil if none found.
func (m *SysLoginLogInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TenantId

	// no validation rules for AdminID

	// no validation rules for Username

	// no validation rules for Ip

	// no validation rules for Useragent

	// no validation rules for LoginTime

	if len(errors) > 0 {
		return SysLoginLogInfoMultiError(errors)
	}

	return nil
}

// SysLoginLogInfoMultiError is an error wrapping multiple validation errors
// returned by SysLoginLogInfo.ValidateAll() if the designated constraints
// aren't met.
type SysLoginLogInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogInfoMultiError) AllErrors() []error { return m }

// SysLoginLogInfoValidationError is the validation error returned by
// SysLoginLogInfo.Validate if the designated constraints aren't met.
type SysLoginLogInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogInfoValidationError) ErrorName() string { return "SysLoginLogInfoValidationError" }

// Error satisfies the builtin error interface
func (e SysLoginLogInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogInfoValidationError{}

// Validate checks the field values on SysLoginLogListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogListReqMultiError, or nil if none found.
func (m *SysLoginLogListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return SysLoginLogListReqMultiError(errors)
	}

	return nil
}

// SysLoginLogListReqMultiError is an error wrapping multiple validation errors
// returned by SysLoginLogListReq.ValidateAll() if the designated constraints
// aren't met.
type SysLoginLogListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogListReqMultiError) AllErrors() []error { return m }

// SysLoginLogListReqValidationError is the validation error returned by
// SysLoginLogListReq.Validate if the designated constraints aren't met.
type SysLoginLogListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogListReqValidationError) ErrorName() string {
	return "SysLoginLogListReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogListReqValidationError{}

// Validate checks the field values on SysLoginLogListResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogListResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogListResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogListRespMultiError, or nil if none found.
func (m *SysLoginLogListResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogListResp) validate(all bool) error {
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
					errors = append(errors, SysLoginLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysLoginLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysLoginLogListRespValidationError{
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
				errors = append(errors, SysLoginLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysLoginLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysLoginLogListRespValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysLoginLogListRespMultiError(errors)
	}

	return nil
}

// SysLoginLogListRespMultiError is an error wrapping multiple validation
// errors returned by SysLoginLogListResp.ValidateAll() if the designated
// constraints aren't met.
type SysLoginLogListRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogListRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogListRespMultiError) AllErrors() []error { return m }

// SysLoginLogListRespValidationError is the validation error returned by
// SysLoginLogListResp.Validate if the designated constraints aren't met.
type SysLoginLogListRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogListRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogListRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogListRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogListRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogListRespValidationError) ErrorName() string {
	return "SysLoginLogListRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogListRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogListResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogListRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogListRespValidationError{}

// Validate checks the field values on SysLoginLogInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogInfoReqMultiError, or nil if none found.
func (m *SysLoginLogInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysLoginLogInfoReqMultiError(errors)
	}

	return nil
}

// SysLoginLogInfoReqMultiError is an error wrapping multiple validation errors
// returned by SysLoginLogInfoReq.ValidateAll() if the designated constraints
// aren't met.
type SysLoginLogInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogInfoReqMultiError) AllErrors() []error { return m }

// SysLoginLogInfoReqValidationError is the validation error returned by
// SysLoginLogInfoReq.Validate if the designated constraints aren't met.
type SysLoginLogInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogInfoReqValidationError) ErrorName() string {
	return "SysLoginLogInfoReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogInfoReqValidationError{}

// Validate checks the field values on SysLoginLogInfoResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogInfoResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogInfoResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogInfoRespMultiError, or nil if none found.
func (m *SysLoginLogInfoResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogInfoResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysLoginLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysLoginLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysLoginLogInfoRespValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysLoginLogInfoRespMultiError(errors)
	}

	return nil
}

// SysLoginLogInfoRespMultiError is an error wrapping multiple validation
// errors returned by SysLoginLogInfoResp.ValidateAll() if the designated
// constraints aren't met.
type SysLoginLogInfoRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogInfoRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogInfoRespMultiError) AllErrors() []error { return m }

// SysLoginLogInfoRespValidationError is the validation error returned by
// SysLoginLogInfoResp.Validate if the designated constraints aren't met.
type SysLoginLogInfoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogInfoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogInfoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogInfoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogInfoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogInfoRespValidationError) ErrorName() string {
	return "SysLoginLogInfoRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogInfoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogInfoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogInfoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogInfoRespValidationError{}

// Validate checks the field values on SysLoginLogStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogStoreReqMultiError, or nil if none found.
func (m *SysLoginLogStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantId

	// no validation rules for AdminID

	// no validation rules for Ip

	// no validation rules for Useragent

	// no validation rules for LoginTime

	if len(errors) > 0 {
		return SysLoginLogStoreReqMultiError(errors)
	}

	return nil
}

// SysLoginLogStoreReqMultiError is an error wrapping multiple validation
// errors returned by SysLoginLogStoreReq.ValidateAll() if the designated
// constraints aren't met.
type SysLoginLogStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogStoreReqMultiError) AllErrors() []error { return m }

// SysLoginLogStoreReqValidationError is the validation error returned by
// SysLoginLogStoreReq.Validate if the designated constraints aren't met.
type SysLoginLogStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogStoreReqValidationError) ErrorName() string {
	return "SysLoginLogStoreReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogStoreReqValidationError{}

// Validate checks the field values on SysLoginLogStoreResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysLoginLogStoreResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysLoginLogStoreResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysLoginLogStoreRespMultiError, or nil if none found.
func (m *SysLoginLogStoreResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysLoginLogStoreResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysLoginLogStoreRespMultiError(errors)
	}

	return nil
}

// SysLoginLogStoreRespMultiError is an error wrapping multiple validation
// errors returned by SysLoginLogStoreResp.ValidateAll() if the designated
// constraints aren't met.
type SysLoginLogStoreRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysLoginLogStoreRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysLoginLogStoreRespMultiError) AllErrors() []error { return m }

// SysLoginLogStoreRespValidationError is the validation error returned by
// SysLoginLogStoreResp.Validate if the designated constraints aren't met.
type SysLoginLogStoreRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysLoginLogStoreRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysLoginLogStoreRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysLoginLogStoreRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysLoginLogStoreRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysLoginLogStoreRespValidationError) ErrorName() string {
	return "SysLoginLogStoreRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysLoginLogStoreRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysLoginLogStoreResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysLoginLogStoreRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysLoginLogStoreRespValidationError{}
