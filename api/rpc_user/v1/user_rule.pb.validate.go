// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_user/v1/user_rule.proto

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

// Validate checks the field values on UserRuleInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserRuleInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserRuleInfoMultiError, or
// nil if none found.
func (m *UserRuleInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRuleInfoMultiError(errors)
	}

	return nil
}

// UserRuleInfoMultiError is an error wrapping multiple validation errors
// returned by UserRuleInfo.ValidateAll() if the designated constraints aren't met.
type UserRuleInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleInfoMultiError) AllErrors() []error { return m }

// UserRuleInfoValidationError is the validation error returned by
// UserRuleInfo.Validate if the designated constraints aren't met.
type UserRuleInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleInfoValidationError) ErrorName() string { return "UserRuleInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleInfoValidationError{}

// Validate checks the field values on UserRuleListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleListReqMultiError, or nil if none found.
func (m *UserRuleListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserRuleListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserRuleListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserRuleListReqValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserRuleListReqMultiError(errors)
	}

	return nil
}

// UserRuleListReqMultiError is an error wrapping multiple validation errors
// returned by UserRuleListReq.ValidateAll() if the designated constraints
// aren't met.
type UserRuleListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleListReqMultiError) AllErrors() []error { return m }

// UserRuleListReqValidationError is the validation error returned by
// UserRuleListReq.Validate if the designated constraints aren't met.
type UserRuleListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleListReqValidationError) ErrorName() string { return "UserRuleListReqValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleListReqValidationError{}

// Validate checks the field values on UserRuleListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleListReplyMultiError, or nil if none found.
func (m *UserRuleListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserRuleListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserRuleListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserRuleListReplyValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserRuleListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserRuleListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserRuleListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserRuleListReplyMultiError(errors)
	}

	return nil
}

// UserRuleListReplyMultiError is an error wrapping multiple validation errors
// returned by UserRuleListReply.ValidateAll() if the designated constraints
// aren't met.
type UserRuleListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleListReplyMultiError) AllErrors() []error { return m }

// UserRuleListReplyValidationError is the validation error returned by
// UserRuleListReply.Validate if the designated constraints aren't met.
type UserRuleListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleListReplyValidationError) ErrorName() string {
	return "UserRuleListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserRuleListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleListReplyValidationError{}

// Validate checks the field values on UserRuleInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleInfoReqMultiError, or nil if none found.
func (m *UserRuleInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UID

	if len(errors) > 0 {
		return UserRuleInfoReqMultiError(errors)
	}

	return nil
}

// UserRuleInfoReqMultiError is an error wrapping multiple validation errors
// returned by UserRuleInfoReq.ValidateAll() if the designated constraints
// aren't met.
type UserRuleInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleInfoReqMultiError) AllErrors() []error { return m }

// UserRuleInfoReqValidationError is the validation error returned by
// UserRuleInfoReq.Validate if the designated constraints aren't met.
type UserRuleInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleInfoReqValidationError) ErrorName() string { return "UserRuleInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleInfoReqValidationError{}

// Validate checks the field values on UserRuleInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleInfoReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleInfoReplyMultiError, or nil if none found.
func (m *UserRuleInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserRuleInfoReplyValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserRuleInfoReplyValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserRuleInfoReplyValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserRuleInfoReplyMultiError(errors)
	}

	return nil
}

// UserRuleInfoReplyMultiError is an error wrapping multiple validation errors
// returned by UserRuleInfoReply.ValidateAll() if the designated constraints
// aren't met.
type UserRuleInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleInfoReplyMultiError) AllErrors() []error { return m }

// UserRuleInfoReplyValidationError is the validation error returned by
// UserRuleInfoReply.Validate if the designated constraints aren't met.
type UserRuleInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleInfoReplyValidationError) ErrorName() string {
	return "UserRuleInfoReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserRuleInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleInfoReplyValidationError{}

// Validate checks the field values on UserRuleStoreReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleStoreReqMultiError, or nil if none found.
func (m *UserRuleStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRuleStoreReqMultiError(errors)
	}

	return nil
}

// UserRuleStoreReqMultiError is an error wrapping multiple validation errors
// returned by UserRuleStoreReq.ValidateAll() if the designated constraints
// aren't met.
type UserRuleStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleStoreReqMultiError) AllErrors() []error { return m }

// UserRuleStoreReqValidationError is the validation error returned by
// UserRuleStoreReq.Validate if the designated constraints aren't met.
type UserRuleStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleStoreReqValidationError) ErrorName() string { return "UserRuleStoreReqValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleStoreReqValidationError{}

// Validate checks the field values on UserRuleStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserRuleStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleStoreReplyMultiError, or nil if none found.
func (m *UserRuleStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	if len(errors) > 0 {
		return UserRuleStoreReplyMultiError(errors)
	}

	return nil
}

// UserRuleStoreReplyMultiError is an error wrapping multiple validation errors
// returned by UserRuleStoreReply.ValidateAll() if the designated constraints
// aren't met.
type UserRuleStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleStoreReplyMultiError) AllErrors() []error { return m }

// UserRuleStoreReplyValidationError is the validation error returned by
// UserRuleStoreReply.Validate if the designated constraints aren't met.
type UserRuleStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleStoreReplyValidationError) ErrorName() string {
	return "UserRuleStoreReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserRuleStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleStoreReplyValidationError{}

// Validate checks the field values on UserRuleDelReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserRuleDelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleDelReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserRuleDelReqMultiError,
// or nil if none found.
func (m *UserRuleDelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleDelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRuleDelReqMultiError(errors)
	}

	return nil
}

// UserRuleDelReqMultiError is an error wrapping multiple validation errors
// returned by UserRuleDelReq.ValidateAll() if the designated constraints
// aren't met.
type UserRuleDelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleDelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleDelReqMultiError) AllErrors() []error { return m }

// UserRuleDelReqValidationError is the validation error returned by
// UserRuleDelReq.Validate if the designated constraints aren't met.
type UserRuleDelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleDelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleDelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleDelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleDelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleDelReqValidationError) ErrorName() string { return "UserRuleDelReqValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleDelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleDelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleDelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleDelReqValidationError{}

// Validate checks the field values on UserRuleDelReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRuleDelReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRuleDelReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRuleDelReplyMultiError, or nil if none found.
func (m *UserRuleDelReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRuleDelReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRuleDelReplyMultiError(errors)
	}

	return nil
}

// UserRuleDelReplyMultiError is an error wrapping multiple validation errors
// returned by UserRuleDelReply.ValidateAll() if the designated constraints
// aren't met.
type UserRuleDelReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRuleDelReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRuleDelReplyMultiError) AllErrors() []error { return m }

// UserRuleDelReplyValidationError is the validation error returned by
// UserRuleDelReply.Validate if the designated constraints aren't met.
type UserRuleDelReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRuleDelReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRuleDelReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRuleDelReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRuleDelReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRuleDelReplyValidationError) ErrorName() string { return "UserRuleDelReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserRuleDelReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRuleDelReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRuleDelReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRuleDelReplyValidationError{}
