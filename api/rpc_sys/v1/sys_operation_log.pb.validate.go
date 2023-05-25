// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_sys/v1/sys_operation_log.proto

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

// Validate checks the field values on SysOperationLogInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogInfoMultiError, or nil if none found.
func (m *SysOperationLogInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogInfo) validate(all bool) error {
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

	// no validation rules for Req

	// no validation rules for Resp

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return SysOperationLogInfoMultiError(errors)
	}

	return nil
}

// SysOperationLogInfoMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogInfo.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogInfoMultiError) AllErrors() []error { return m }

// SysOperationLogInfoValidationError is the validation error returned by
// SysOperationLogInfo.Validate if the designated constraints aren't met.
type SysOperationLogInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogInfoValidationError) ErrorName() string {
	return "SysOperationLogInfoValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogInfoValidationError{}

// Validate checks the field values on SysOperationLogListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogListReqMultiError, or nil if none found.
func (m *SysOperationLogListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return SysOperationLogListReqMultiError(errors)
	}

	return nil
}

// SysOperationLogListReqMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogListReq.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogListReqMultiError) AllErrors() []error { return m }

// SysOperationLogListReqValidationError is the validation error returned by
// SysOperationLogListReq.Validate if the designated constraints aren't met.
type SysOperationLogListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogListReqValidationError) ErrorName() string {
	return "SysOperationLogListReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogListReqValidationError{}

// Validate checks the field values on SysOperationLogListResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogListResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogListResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogListRespMultiError, or nil if none found.
func (m *SysOperationLogListResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogListResp) validate(all bool) error {
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
					errors = append(errors, SysOperationLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysOperationLogListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysOperationLogListRespValidationError{
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
				errors = append(errors, SysOperationLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysOperationLogListRespValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysOperationLogListRespValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysOperationLogListRespMultiError(errors)
	}

	return nil
}

// SysOperationLogListRespMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogListResp.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogListRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogListRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogListRespMultiError) AllErrors() []error { return m }

// SysOperationLogListRespValidationError is the validation error returned by
// SysOperationLogListResp.Validate if the designated constraints aren't met.
type SysOperationLogListRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogListRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogListRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogListRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogListRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogListRespValidationError) ErrorName() string {
	return "SysOperationLogListRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogListRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogListResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogListRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogListRespValidationError{}

// Validate checks the field values on SysOperationLogInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogInfoReqMultiError, or nil if none found.
func (m *SysOperationLogInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysOperationLogInfoReqMultiError(errors)
	}

	return nil
}

// SysOperationLogInfoReqMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogInfoReq.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogInfoReqMultiError) AllErrors() []error { return m }

// SysOperationLogInfoReqValidationError is the validation error returned by
// SysOperationLogInfoReq.Validate if the designated constraints aren't met.
type SysOperationLogInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogInfoReqValidationError) ErrorName() string {
	return "SysOperationLogInfoReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogInfoReqValidationError{}

// Validate checks the field values on SysOperationLogInfoResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogInfoResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogInfoResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogInfoRespMultiError, or nil if none found.
func (m *SysOperationLogInfoResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogInfoResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysOperationLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysOperationLogInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysOperationLogInfoRespValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysOperationLogInfoRespMultiError(errors)
	}

	return nil
}

// SysOperationLogInfoRespMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogInfoResp.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogInfoRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogInfoRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogInfoRespMultiError) AllErrors() []error { return m }

// SysOperationLogInfoRespValidationError is the validation error returned by
// SysOperationLogInfoResp.Validate if the designated constraints aren't met.
type SysOperationLogInfoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogInfoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogInfoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogInfoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogInfoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogInfoRespValidationError) ErrorName() string {
	return "SysOperationLogInfoRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogInfoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogInfoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogInfoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogInfoRespValidationError{}

// Validate checks the field values on SysOperationLogStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogStoreReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogStoreReqMultiError, or nil if none found.
func (m *SysOperationLogStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AdminID

	// no validation rules for Username

	// no validation rules for Ip

	// no validation rules for Uri

	// no validation rules for UriDesc

	// no validation rules for Useragent

	// no validation rules for Req

	// no validation rules for Resp

	if len(errors) > 0 {
		return SysOperationLogStoreReqMultiError(errors)
	}

	return nil
}

// SysOperationLogStoreReqMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogStoreReq.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogStoreReqMultiError) AllErrors() []error { return m }

// SysOperationLogStoreReqValidationError is the validation error returned by
// SysOperationLogStoreReq.Validate if the designated constraints aren't met.
type SysOperationLogStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogStoreReqValidationError) ErrorName() string {
	return "SysOperationLogStoreReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogStoreReqValidationError{}

// Validate checks the field values on SysOperationLogStoreResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysOperationLogStoreResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysOperationLogStoreResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysOperationLogStoreRespMultiError, or nil if none found.
func (m *SysOperationLogStoreResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysOperationLogStoreResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysOperationLogStoreRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysOperationLogStoreRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysOperationLogStoreRespValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysOperationLogStoreRespMultiError(errors)
	}

	return nil
}

// SysOperationLogStoreRespMultiError is an error wrapping multiple validation
// errors returned by SysOperationLogStoreResp.ValidateAll() if the designated
// constraints aren't met.
type SysOperationLogStoreRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysOperationLogStoreRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysOperationLogStoreRespMultiError) AllErrors() []error { return m }

// SysOperationLogStoreRespValidationError is the validation error returned by
// SysOperationLogStoreResp.Validate if the designated constraints aren't met.
type SysOperationLogStoreRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysOperationLogStoreRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysOperationLogStoreRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysOperationLogStoreRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysOperationLogStoreRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysOperationLogStoreRespValidationError) ErrorName() string {
	return "SysOperationLogStoreRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysOperationLogStoreRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysOperationLogStoreResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysOperationLogStoreRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysOperationLogStoreRespValidationError{}
