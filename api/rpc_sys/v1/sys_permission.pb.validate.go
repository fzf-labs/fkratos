// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_sys/v1/sys_permission.proto

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

// Validate checks the field values on SysPermissionInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionInfoMultiError, or nil if none found.
func (m *SysPermissionInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Pid

	// no validation rules for Type

	// no validation rules for Title

	// no validation rules for UpperName

	// no validation rules for Path

	// no validation rules for Icon

	// no validation rules for MenuType

	// no validation rules for Url

	// no validation rules for Component

	// no validation rules for Keepalive

	// no validation rules for Extend

	// no validation rules for Remark

	// no validation rules for Sort

	// no validation rules for Status

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SysPermissionInfoValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysPermissionInfoValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysPermissionInfoValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SysPermissionInfoMultiError(errors)
	}

	return nil
}

// SysPermissionInfoMultiError is an error wrapping multiple validation errors
// returned by SysPermissionInfo.ValidateAll() if the designated constraints
// aren't met.
type SysPermissionInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionInfoMultiError) AllErrors() []error { return m }

// SysPermissionInfoValidationError is the validation error returned by
// SysPermissionInfo.Validate if the designated constraints aren't met.
type SysPermissionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionInfoValidationError) ErrorName() string {
	return "SysPermissionInfoValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionInfoValidationError{}

// Validate checks the field values on SysPermissionListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionListReqMultiError, or nil if none found.
func (m *SysPermissionListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysPermissionListReqMultiError(errors)
	}

	return nil
}

// SysPermissionListReqMultiError is an error wrapping multiple validation
// errors returned by SysPermissionListReq.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionListReqMultiError) AllErrors() []error { return m }

// SysPermissionListReqValidationError is the validation error returned by
// SysPermissionListReq.Validate if the designated constraints aren't met.
type SysPermissionListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionListReqValidationError) ErrorName() string {
	return "SysPermissionListReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionListReqValidationError{}

// Validate checks the field values on SysPermissionListResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionListResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionListResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionListRespMultiError, or nil if none found.
func (m *SysPermissionListResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionListResp) validate(all bool) error {
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
					errors = append(errors, SysPermissionListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysPermissionListRespValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysPermissionListRespValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SysPermissionListRespMultiError(errors)
	}

	return nil
}

// SysPermissionListRespMultiError is an error wrapping multiple validation
// errors returned by SysPermissionListResp.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionListRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionListRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionListRespMultiError) AllErrors() []error { return m }

// SysPermissionListRespValidationError is the validation error returned by
// SysPermissionListResp.Validate if the designated constraints aren't met.
type SysPermissionListRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionListRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionListRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionListRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionListRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionListRespValidationError) ErrorName() string {
	return "SysPermissionListRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionListRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionListResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionListRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionListRespValidationError{}

// Validate checks the field values on SysPermissionInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionInfoReqMultiError, or nil if none found.
func (m *SysPermissionInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return SysPermissionInfoReqMultiError(errors)
	}

	return nil
}

// SysPermissionInfoReqMultiError is an error wrapping multiple validation
// errors returned by SysPermissionInfoReq.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionInfoReqMultiError) AllErrors() []error { return m }

// SysPermissionInfoReqValidationError is the validation error returned by
// SysPermissionInfoReq.Validate if the designated constraints aren't met.
type SysPermissionInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionInfoReqValidationError) ErrorName() string {
	return "SysPermissionInfoReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionInfoReqValidationError{}

// Validate checks the field values on SysPermissionInfoResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionInfoResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionInfoResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionInfoRespMultiError, or nil if none found.
func (m *SysPermissionInfoResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionInfoResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysPermissionInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysPermissionInfoRespValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysPermissionInfoRespValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysPermissionInfoRespMultiError(errors)
	}

	return nil
}

// SysPermissionInfoRespMultiError is an error wrapping multiple validation
// errors returned by SysPermissionInfoResp.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionInfoRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionInfoRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionInfoRespMultiError) AllErrors() []error { return m }

// SysPermissionInfoRespValidationError is the validation error returned by
// SysPermissionInfoResp.Validate if the designated constraints aren't met.
type SysPermissionInfoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionInfoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionInfoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionInfoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionInfoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionInfoRespValidationError) ErrorName() string {
	return "SysPermissionInfoRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionInfoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionInfoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionInfoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionInfoRespValidationError{}

// Validate checks the field values on SysPermissionStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionStoreReqMultiError, or nil if none found.
func (m *SysPermissionStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Pid

	// no validation rules for Type

	// no validation rules for Title

	// no validation rules for UpperName

	// no validation rules for Path

	// no validation rules for Icon

	// no validation rules for MenuType

	// no validation rules for Url

	// no validation rules for Component

	// no validation rules for Keepalive

	// no validation rules for Extend

	// no validation rules for Remark

	// no validation rules for Sort

	// no validation rules for Status

	if len(errors) > 0 {
		return SysPermissionStoreReqMultiError(errors)
	}

	return nil
}

// SysPermissionStoreReqMultiError is an error wrapping multiple validation
// errors returned by SysPermissionStoreReq.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionStoreReqMultiError) AllErrors() []error { return m }

// SysPermissionStoreReqValidationError is the validation error returned by
// SysPermissionStoreReq.Validate if the designated constraints aren't met.
type SysPermissionStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionStoreReqValidationError) ErrorName() string {
	return "SysPermissionStoreReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionStoreReqValidationError{}

// Validate checks the field values on SysPermissionStoreResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionStoreResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionStoreResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionStoreRespMultiError, or nil if none found.
func (m *SysPermissionStoreResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionStoreResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysPermissionStoreRespMultiError(errors)
	}

	return nil
}

// SysPermissionStoreRespMultiError is an error wrapping multiple validation
// errors returned by SysPermissionStoreResp.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionStoreRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionStoreRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionStoreRespMultiError) AllErrors() []error { return m }

// SysPermissionStoreRespValidationError is the validation error returned by
// SysPermissionStoreResp.Validate if the designated constraints aren't met.
type SysPermissionStoreRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionStoreRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionStoreRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionStoreRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionStoreRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionStoreRespValidationError) ErrorName() string {
	return "SysPermissionStoreRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionStoreRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionStoreResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionStoreRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionStoreRespValidationError{}

// Validate checks the field values on SysPermissionDelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionDelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionDelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionDelReqMultiError, or nil if none found.
func (m *SysPermissionDelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionDelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysPermissionDelReqMultiError(errors)
	}

	return nil
}

// SysPermissionDelReqMultiError is an error wrapping multiple validation
// errors returned by SysPermissionDelReq.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionDelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionDelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionDelReqMultiError) AllErrors() []error { return m }

// SysPermissionDelReqValidationError is the validation error returned by
// SysPermissionDelReq.Validate if the designated constraints aren't met.
type SysPermissionDelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionDelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionDelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionDelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionDelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionDelReqValidationError) ErrorName() string {
	return "SysPermissionDelReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionDelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionDelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionDelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionDelReqValidationError{}

// Validate checks the field values on SysPermissionDelResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionDelResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionDelResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionDelRespMultiError, or nil if none found.
func (m *SysPermissionDelResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionDelResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysPermissionDelRespMultiError(errors)
	}

	return nil
}

// SysPermissionDelRespMultiError is an error wrapping multiple validation
// errors returned by SysPermissionDelResp.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionDelRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionDelRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionDelRespMultiError) AllErrors() []error { return m }

// SysPermissionDelRespValidationError is the validation error returned by
// SysPermissionDelResp.Validate if the designated constraints aren't met.
type SysPermissionDelRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionDelRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionDelRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionDelRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionDelRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionDelRespValidationError) ErrorName() string {
	return "SysPermissionDelRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionDelRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionDelResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionDelRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionDelRespValidationError{}

// Validate checks the field values on SysPermissionStatusReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionStatusReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionStatusReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionStatusReqMultiError, or nil if none found.
func (m *SysPermissionStatusReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionStatusReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	if len(errors) > 0 {
		return SysPermissionStatusReqMultiError(errors)
	}

	return nil
}

// SysPermissionStatusReqMultiError is an error wrapping multiple validation
// errors returned by SysPermissionStatusReq.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionStatusReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionStatusReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionStatusReqMultiError) AllErrors() []error { return m }

// SysPermissionStatusReqValidationError is the validation error returned by
// SysPermissionStatusReq.Validate if the designated constraints aren't met.
type SysPermissionStatusReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionStatusReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionStatusReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionStatusReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionStatusReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionStatusReqValidationError) ErrorName() string {
	return "SysPermissionStatusReqValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionStatusReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionStatusReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionStatusReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionStatusReqValidationError{}

// Validate checks the field values on SysPermissionStatusResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysPermissionStatusResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPermissionStatusResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysPermissionStatusRespMultiError, or nil if none found.
func (m *SysPermissionStatusResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPermissionStatusResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysPermissionStatusRespMultiError(errors)
	}

	return nil
}

// SysPermissionStatusRespMultiError is an error wrapping multiple validation
// errors returned by SysPermissionStatusResp.ValidateAll() if the designated
// constraints aren't met.
type SysPermissionStatusRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPermissionStatusRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPermissionStatusRespMultiError) AllErrors() []error { return m }

// SysPermissionStatusRespValidationError is the validation error returned by
// SysPermissionStatusResp.Validate if the designated constraints aren't met.
type SysPermissionStatusRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPermissionStatusRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPermissionStatusRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPermissionStatusRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPermissionStatusRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPermissionStatusRespValidationError) ErrorName() string {
	return "SysPermissionStatusRespValidationError"
}

// Error satisfies the builtin error interface
func (e SysPermissionStatusRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPermissionStatusResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPermissionStatusRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPermissionStatusRespValidationError{}
