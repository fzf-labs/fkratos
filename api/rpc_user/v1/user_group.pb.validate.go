// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_user/v1/user_group.proto

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

// Validate checks the field values on UserGroupInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserGroupInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserGroupInfoMultiError, or
// nil if none found.
func (m *UserGroupInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return UserGroupInfoMultiError(errors)
	}

	return nil
}

// UserGroupInfoMultiError is an error wrapping multiple validation errors
// returned by UserGroupInfo.ValidateAll() if the designated constraints
// aren't met.
type UserGroupInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupInfoMultiError) AllErrors() []error { return m }

// UserGroupInfoValidationError is the validation error returned by
// UserGroupInfo.Validate if the designated constraints aren't met.
type UserGroupInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupInfoValidationError) ErrorName() string { return "UserGroupInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserGroupInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupInfoValidationError{}

// Validate checks the field values on UserGroupListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserGroupListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupListReqMultiError, or nil if none found.
func (m *UserGroupListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserGroupListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserGroupListReqValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserGroupListReqValidationError{
				field:  "Paginator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserGroupListReqMultiError(errors)
	}

	return nil
}

// UserGroupListReqMultiError is an error wrapping multiple validation errors
// returned by UserGroupListReq.ValidateAll() if the designated constraints
// aren't met.
type UserGroupListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupListReqMultiError) AllErrors() []error { return m }

// UserGroupListReqValidationError is the validation error returned by
// UserGroupListReq.Validate if the designated constraints aren't met.
type UserGroupListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupListReqValidationError) ErrorName() string { return "UserGroupListReqValidationError" }

// Error satisfies the builtin error interface
func (e UserGroupListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupListReqValidationError{}

// Validate checks the field values on UserGroupListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserGroupListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupListReplyMultiError, or nil if none found.
func (m *UserGroupListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserGroupListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserGroupListReplyValidationError{
					field:  "Paginator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserGroupListReplyValidationError{
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
					errors = append(errors, UserGroupListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserGroupListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserGroupListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserGroupListReplyMultiError(errors)
	}

	return nil
}

// UserGroupListReplyMultiError is an error wrapping multiple validation errors
// returned by UserGroupListReply.ValidateAll() if the designated constraints
// aren't met.
type UserGroupListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupListReplyMultiError) AllErrors() []error { return m }

// UserGroupListReplyValidationError is the validation error returned by
// UserGroupListReply.Validate if the designated constraints aren't met.
type UserGroupListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupListReplyValidationError) ErrorName() string {
	return "UserGroupListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserGroupListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupListReplyValidationError{}

// Validate checks the field values on UserGroupStoreReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserGroupStoreReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupStoreReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupStoreReqMultiError, or nil if none found.
func (m *UserGroupStoreReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupStoreReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	if len(errors) > 0 {
		return UserGroupStoreReqMultiError(errors)
	}

	return nil
}

// UserGroupStoreReqMultiError is an error wrapping multiple validation errors
// returned by UserGroupStoreReq.ValidateAll() if the designated constraints
// aren't met.
type UserGroupStoreReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupStoreReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupStoreReqMultiError) AllErrors() []error { return m }

// UserGroupStoreReqValidationError is the validation error returned by
// UserGroupStoreReq.Validate if the designated constraints aren't met.
type UserGroupStoreReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupStoreReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupStoreReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupStoreReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupStoreReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupStoreReqValidationError) ErrorName() string {
	return "UserGroupStoreReqValidationError"
}

// Error satisfies the builtin error interface
func (e UserGroupStoreReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupStoreReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupStoreReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupStoreReqValidationError{}

// Validate checks the field values on UserGroupStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserGroupStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupStoreReplyMultiError, or nil if none found.
func (m *UserGroupStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UserGroupStoreReplyMultiError(errors)
	}

	return nil
}

// UserGroupStoreReplyMultiError is an error wrapping multiple validation
// errors returned by UserGroupStoreReply.ValidateAll() if the designated
// constraints aren't met.
type UserGroupStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupStoreReplyMultiError) AllErrors() []error { return m }

// UserGroupStoreReplyValidationError is the validation error returned by
// UserGroupStoreReply.Validate if the designated constraints aren't met.
type UserGroupStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupStoreReplyValidationError) ErrorName() string {
	return "UserGroupStoreReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserGroupStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupStoreReplyValidationError{}

// Validate checks the field values on UserGroupDelReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserGroupDelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupDelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupDelReqMultiError, or nil if none found.
func (m *UserGroupDelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupDelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserGroupDelReqMultiError(errors)
	}

	return nil
}

// UserGroupDelReqMultiError is an error wrapping multiple validation errors
// returned by UserGroupDelReq.ValidateAll() if the designated constraints
// aren't met.
type UserGroupDelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupDelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupDelReqMultiError) AllErrors() []error { return m }

// UserGroupDelReqValidationError is the validation error returned by
// UserGroupDelReq.Validate if the designated constraints aren't met.
type UserGroupDelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupDelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupDelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupDelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupDelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupDelReqValidationError) ErrorName() string { return "UserGroupDelReqValidationError" }

// Error satisfies the builtin error interface
func (e UserGroupDelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupDelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupDelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupDelReqValidationError{}

// Validate checks the field values on UserGroupDelReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserGroupDelReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGroupDelReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserGroupDelReplyMultiError, or nil if none found.
func (m *UserGroupDelReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGroupDelReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserGroupDelReplyMultiError(errors)
	}

	return nil
}

// UserGroupDelReplyMultiError is an error wrapping multiple validation errors
// returned by UserGroupDelReply.ValidateAll() if the designated constraints
// aren't met.
type UserGroupDelReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGroupDelReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGroupDelReplyMultiError) AllErrors() []error { return m }

// UserGroupDelReplyValidationError is the validation error returned by
// UserGroupDelReply.Validate if the designated constraints aren't met.
type UserGroupDelReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGroupDelReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGroupDelReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGroupDelReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGroupDelReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGroupDelReplyValidationError) ErrorName() string {
	return "UserGroupDelReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UserGroupDelReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGroupDelReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGroupDelReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGroupDelReplyValidationError{}
