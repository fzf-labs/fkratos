// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/usertoken.proto

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

// Validate checks the field values on CreateUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUsertokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUsertokenRequestMultiError, or nil if none found.
func (m *CreateUsertokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUsertokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateUsertokenRequestMultiError(errors)
	}

	return nil
}

// CreateUsertokenRequestMultiError is an error wrapping multiple validation
// errors returned by CreateUsertokenRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateUsertokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUsertokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUsertokenRequestMultiError) AllErrors() []error { return m }

// CreateUsertokenRequestValidationError is the validation error returned by
// CreateUsertokenRequest.Validate if the designated constraints aren't met.
type CreateUsertokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUsertokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUsertokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUsertokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUsertokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUsertokenRequestValidationError) ErrorName() string {
	return "CreateUsertokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUsertokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUsertokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUsertokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUsertokenRequestValidationError{}

// Validate checks the field values on CreateUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUsertokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUsertokenReplyMultiError, or nil if none found.
func (m *CreateUsertokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUsertokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateUsertokenReplyMultiError(errors)
	}

	return nil
}

// CreateUsertokenReplyMultiError is an error wrapping multiple validation
// errors returned by CreateUsertokenReply.ValidateAll() if the designated
// constraints aren't met.
type CreateUsertokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUsertokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUsertokenReplyMultiError) AllErrors() []error { return m }

// CreateUsertokenReplyValidationError is the validation error returned by
// CreateUsertokenReply.Validate if the designated constraints aren't met.
type CreateUsertokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUsertokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUsertokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUsertokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUsertokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUsertokenReplyValidationError) ErrorName() string {
	return "CreateUsertokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUsertokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUsertokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUsertokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUsertokenReplyValidationError{}

// Validate checks the field values on UpdateUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUsertokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUsertokenRequestMultiError, or nil if none found.
func (m *UpdateUsertokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUsertokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateUsertokenRequestMultiError(errors)
	}

	return nil
}

// UpdateUsertokenRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateUsertokenRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateUsertokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUsertokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUsertokenRequestMultiError) AllErrors() []error { return m }

// UpdateUsertokenRequestValidationError is the validation error returned by
// UpdateUsertokenRequest.Validate if the designated constraints aren't met.
type UpdateUsertokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUsertokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUsertokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUsertokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUsertokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUsertokenRequestValidationError) ErrorName() string {
	return "UpdateUsertokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUsertokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUsertokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUsertokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUsertokenRequestValidationError{}

// Validate checks the field values on UpdateUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUsertokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUsertokenReplyMultiError, or nil if none found.
func (m *UpdateUsertokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUsertokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateUsertokenReplyMultiError(errors)
	}

	return nil
}

// UpdateUsertokenReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateUsertokenReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateUsertokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUsertokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUsertokenReplyMultiError) AllErrors() []error { return m }

// UpdateUsertokenReplyValidationError is the validation error returned by
// UpdateUsertokenReply.Validate if the designated constraints aren't met.
type UpdateUsertokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUsertokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUsertokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUsertokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUsertokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUsertokenReplyValidationError) ErrorName() string {
	return "UpdateUsertokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUsertokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUsertokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUsertokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUsertokenReplyValidationError{}

// Validate checks the field values on DeleteUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUsertokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUsertokenRequestMultiError, or nil if none found.
func (m *DeleteUsertokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUsertokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUsertokenRequestMultiError(errors)
	}

	return nil
}

// DeleteUsertokenRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteUsertokenRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteUsertokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUsertokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUsertokenRequestMultiError) AllErrors() []error { return m }

// DeleteUsertokenRequestValidationError is the validation error returned by
// DeleteUsertokenRequest.Validate if the designated constraints aren't met.
type DeleteUsertokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUsertokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUsertokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUsertokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUsertokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUsertokenRequestValidationError) ErrorName() string {
	return "DeleteUsertokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUsertokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUsertokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUsertokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUsertokenRequestValidationError{}

// Validate checks the field values on DeleteUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUsertokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUsertokenReplyMultiError, or nil if none found.
func (m *DeleteUsertokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUsertokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUsertokenReplyMultiError(errors)
	}

	return nil
}

// DeleteUsertokenReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteUsertokenReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteUsertokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUsertokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUsertokenReplyMultiError) AllErrors() []error { return m }

// DeleteUsertokenReplyValidationError is the validation error returned by
// DeleteUsertokenReply.Validate if the designated constraints aren't met.
type DeleteUsertokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUsertokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUsertokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUsertokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUsertokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUsertokenReplyValidationError) ErrorName() string {
	return "DeleteUsertokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUsertokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUsertokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUsertokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUsertokenReplyValidationError{}

// Validate checks the field values on GetUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUsertokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUsertokenRequestMultiError, or nil if none found.
func (m *GetUsertokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUsertokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUsertokenRequestMultiError(errors)
	}

	return nil
}

// GetUsertokenRequestMultiError is an error wrapping multiple validation
// errors returned by GetUsertokenRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUsertokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUsertokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUsertokenRequestMultiError) AllErrors() []error { return m }

// GetUsertokenRequestValidationError is the validation error returned by
// GetUsertokenRequest.Validate if the designated constraints aren't met.
type GetUsertokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsertokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsertokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsertokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsertokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsertokenRequestValidationError) ErrorName() string {
	return "GetUsertokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsertokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsertokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsertokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsertokenRequestValidationError{}

// Validate checks the field values on GetUsertokenReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUsertokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUsertokenReplyMultiError, or nil if none found.
func (m *GetUsertokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUsertokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUsertokenReplyMultiError(errors)
	}

	return nil
}

// GetUsertokenReplyMultiError is an error wrapping multiple validation errors
// returned by GetUsertokenReply.ValidateAll() if the designated constraints
// aren't met.
type GetUsertokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUsertokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUsertokenReplyMultiError) AllErrors() []error { return m }

// GetUsertokenReplyValidationError is the validation error returned by
// GetUsertokenReply.Validate if the designated constraints aren't met.
type GetUsertokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsertokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsertokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsertokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsertokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsertokenReplyValidationError) ErrorName() string {
	return "GetUsertokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsertokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsertokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsertokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsertokenReplyValidationError{}

// Validate checks the field values on ListUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUsertokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsertokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsertokenRequestMultiError, or nil if none found.
func (m *ListUsertokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsertokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUsertokenRequestMultiError(errors)
	}

	return nil
}

// ListUsertokenRequestMultiError is an error wrapping multiple validation
// errors returned by ListUsertokenRequest.ValidateAll() if the designated
// constraints aren't met.
type ListUsertokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsertokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsertokenRequestMultiError) AllErrors() []error { return m }

// ListUsertokenRequestValidationError is the validation error returned by
// ListUsertokenRequest.Validate if the designated constraints aren't met.
type ListUsertokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsertokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsertokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsertokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsertokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsertokenRequestValidationError) ErrorName() string {
	return "ListUsertokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsertokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsertokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsertokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsertokenRequestValidationError{}

// Validate checks the field values on ListUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUsertokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsertokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsertokenReplyMultiError, or nil if none found.
func (m *ListUsertokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsertokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUsertokenReplyMultiError(errors)
	}

	return nil
}

// ListUsertokenReplyMultiError is an error wrapping multiple validation errors
// returned by ListUsertokenReply.ValidateAll() if the designated constraints
// aren't met.
type ListUsertokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsertokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsertokenReplyMultiError) AllErrors() []error { return m }

// ListUsertokenReplyValidationError is the validation error returned by
// ListUsertokenReply.Validate if the designated constraints aren't met.
type ListUsertokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsertokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsertokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsertokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsertokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsertokenReplyValidationError) ErrorName() string {
	return "ListUsertokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsertokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsertokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsertokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsertokenReplyValidationError{}