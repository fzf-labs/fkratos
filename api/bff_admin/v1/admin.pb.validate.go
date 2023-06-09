// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bff_admin/v1/admin.proto

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

// Validate checks the field values on AdminInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AdminInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AdminInfoMultiError, or nil
// if none found.
func (m *AdminInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AdminInfoMultiError(errors)
	}

	return nil
}

// AdminInfoMultiError is an error wrapping multiple validation errors returned
// by AdminInfo.ValidateAll() if the designated constraints aren't met.
type AdminInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminInfoMultiError) AllErrors() []error { return m }

// AdminInfoValidationError is the validation error returned by
// AdminInfo.Validate if the designated constraints aren't met.
type AdminInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminInfoValidationError) ErrorName() string { return "AdminInfoValidationError" }

// Error satisfies the builtin error interface
func (e AdminInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminInfoValidationError{}

// Validate checks the field values on CreateAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateAdminReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAdminReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateAdminReqMultiError,
// or nil if none found.
func (m *CreateAdminReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAdminReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateAdminReqMultiError(errors)
	}

	return nil
}

// CreateAdminReqMultiError is an error wrapping multiple validation errors
// returned by CreateAdminReq.ValidateAll() if the designated constraints
// aren't met.
type CreateAdminReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAdminReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAdminReqMultiError) AllErrors() []error { return m }

// CreateAdminReqValidationError is the validation error returned by
// CreateAdminReq.Validate if the designated constraints aren't met.
type CreateAdminReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAdminReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAdminReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAdminReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAdminReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAdminReqValidationError) ErrorName() string { return "CreateAdminReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateAdminReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAdminReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAdminReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAdminReqValidationError{}

// Validate checks the field values on CreateAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateAdminReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAdminReplyMultiError, or nil if none found.
func (m *CreateAdminReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAdminReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateAdminReplyMultiError(errors)
	}

	return nil
}

// CreateAdminReplyMultiError is an error wrapping multiple validation errors
// returned by CreateAdminReply.ValidateAll() if the designated constraints
// aren't met.
type CreateAdminReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAdminReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAdminReplyMultiError) AllErrors() []error { return m }

// CreateAdminReplyValidationError is the validation error returned by
// CreateAdminReply.Validate if the designated constraints aren't met.
type CreateAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAdminReplyValidationError) ErrorName() string { return "CreateAdminReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAdminReplyValidationError{}

// Validate checks the field values on UpdateAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateAdminReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAdminReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateAdminReqMultiError,
// or nil if none found.
func (m *UpdateAdminReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAdminReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAdminInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateAdminReqValidationError{
					field:  "AdminInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateAdminReqValidationError{
					field:  "AdminInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAdminInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateAdminReqValidationError{
				field:  "AdminInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateAdminReqMultiError(errors)
	}

	return nil
}

// UpdateAdminReqMultiError is an error wrapping multiple validation errors
// returned by UpdateAdminReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateAdminReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAdminReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAdminReqMultiError) AllErrors() []error { return m }

// UpdateAdminReqValidationError is the validation error returned by
// UpdateAdminReq.Validate if the designated constraints aren't met.
type UpdateAdminReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAdminReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAdminReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAdminReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAdminReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAdminReqValidationError) ErrorName() string { return "UpdateAdminReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateAdminReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAdminReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAdminReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAdminReqValidationError{}

// Validate checks the field values on UpdateAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateAdminReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAdminReplyMultiError, or nil if none found.
func (m *UpdateAdminReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAdminReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateAdminReplyMultiError(errors)
	}

	return nil
}

// UpdateAdminReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateAdminReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateAdminReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAdminReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAdminReplyMultiError) AllErrors() []error { return m }

// UpdateAdminReplyValidationError is the validation error returned by
// UpdateAdminReply.Validate if the designated constraints aren't met.
type UpdateAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAdminReplyValidationError) ErrorName() string { return "UpdateAdminReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAdminReplyValidationError{}

// Validate checks the field values on DeleteAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteAdminReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAdminReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteAdminReqMultiError,
// or nil if none found.
func (m *DeleteAdminReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAdminReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteAdminReqMultiError(errors)
	}

	return nil
}

// DeleteAdminReqMultiError is an error wrapping multiple validation errors
// returned by DeleteAdminReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteAdminReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAdminReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAdminReqMultiError) AllErrors() []error { return m }

// DeleteAdminReqValidationError is the validation error returned by
// DeleteAdminReq.Validate if the designated constraints aren't met.
type DeleteAdminReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAdminReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAdminReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAdminReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAdminReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAdminReqValidationError) ErrorName() string { return "DeleteAdminReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteAdminReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAdminReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAdminReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAdminReqValidationError{}

// Validate checks the field values on DeleteAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteAdminReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAdminReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAdminReplyMultiError, or nil if none found.
func (m *DeleteAdminReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAdminReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteAdminReplyMultiError(errors)
	}

	return nil
}

// DeleteAdminReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteAdminReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteAdminReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAdminReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAdminReplyMultiError) AllErrors() []error { return m }

// DeleteAdminReplyValidationError is the validation error returned by
// DeleteAdminReply.Validate if the designated constraints aren't met.
type DeleteAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAdminReplyValidationError) ErrorName() string { return "DeleteAdminReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAdminReplyValidationError{}

// Validate checks the field values on GetAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetAdminReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetAdminReqMultiError, or
// nil if none found.
func (m *GetAdminReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAdminReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAdminReqMultiError(errors)
	}

	return nil
}

// GetAdminReqMultiError is an error wrapping multiple validation errors
// returned by GetAdminReq.ValidateAll() if the designated constraints aren't met.
type GetAdminReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAdminReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAdminReqMultiError) AllErrors() []error { return m }

// GetAdminReqValidationError is the validation error returned by
// GetAdminReq.Validate if the designated constraints aren't met.
type GetAdminReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAdminReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAdminReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAdminReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAdminReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAdminReqValidationError) ErrorName() string { return "GetAdminReqValidationError" }

// Error satisfies the builtin error interface
func (e GetAdminReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAdminReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAdminReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAdminReqValidationError{}

// Validate checks the field values on GetAdminReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetAdminReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetAdminReplyMultiError, or
// nil if none found.
func (m *GetAdminReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAdminReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAdminReplyMultiError(errors)
	}

	return nil
}

// GetAdminReplyMultiError is an error wrapping multiple validation errors
// returned by GetAdminReply.ValidateAll() if the designated constraints
// aren't met.
type GetAdminReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAdminReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAdminReplyMultiError) AllErrors() []error { return m }

// GetAdminReplyValidationError is the validation error returned by
// GetAdminReply.Validate if the designated constraints aren't met.
type GetAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAdminReplyValidationError) ErrorName() string { return "GetAdminReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAdminReplyValidationError{}

// Validate checks the field values on ListAdminReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListAdminReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAdminReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListAdminReqMultiError, or
// nil if none found.
func (m *ListAdminReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAdminReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListAdminReqMultiError(errors)
	}

	return nil
}

// ListAdminReqMultiError is an error wrapping multiple validation errors
// returned by ListAdminReq.ValidateAll() if the designated constraints aren't met.
type ListAdminReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAdminReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAdminReqMultiError) AllErrors() []error { return m }

// ListAdminReqValidationError is the validation error returned by
// ListAdminReq.Validate if the designated constraints aren't met.
type ListAdminReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAdminReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAdminReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAdminReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAdminReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAdminReqValidationError) ErrorName() string { return "ListAdminReqValidationError" }

// Error satisfies the builtin error interface
func (e ListAdminReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAdminReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAdminReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAdminReqValidationError{}

// Validate checks the field values on ListAdminReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListAdminReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAdminReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListAdminReplyMultiError,
// or nil if none found.
func (m *ListAdminReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAdminReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListAdminReplyMultiError(errors)
	}

	return nil
}

// ListAdminReplyMultiError is an error wrapping multiple validation errors
// returned by ListAdminReply.ValidateAll() if the designated constraints
// aren't met.
type ListAdminReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAdminReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAdminReplyMultiError) AllErrors() []error { return m }

// ListAdminReplyValidationError is the validation error returned by
// ListAdminReply.Validate if the designated constraints aren't met.
type ListAdminReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAdminReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAdminReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAdminReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAdminReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAdminReplyValidationError) ErrorName() string { return "ListAdminReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListAdminReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAdminReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAdminReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAdminReplyValidationError{}
