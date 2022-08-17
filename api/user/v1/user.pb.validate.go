// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user/v1/user.proto

package userPb

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on UserReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserReplyMultiError, or nil
// if none found.
func (m *UserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserReplyValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserReplyMultiError(errors)
	}

	return nil
}

// UserReplyMultiError is an error wrapping multiple validation errors returned
// by UserReply.ValidateAll() if the designated constraints aren't met.
type UserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserReplyMultiError) AllErrors() []error { return m }

// UserReplyValidationError is the validation error returned by
// UserReply.Validate if the designated constraints aren't met.
type UserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserReplyValidationError) ErrorName() string { return "UserReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserReplyValidationError{}

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on GetProfileByUserNameRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileByUserNameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileByUserNameRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileByUserNameRequestMultiError, or nil if none found.
func (m *GetProfileByUserNameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileByUserNameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	if len(errors) > 0 {
		return GetProfileByUserNameRequestMultiError(errors)
	}

	return nil
}

// GetProfileByUserNameRequestMultiError is an error wrapping multiple
// validation errors returned by GetProfileByUserNameRequest.ValidateAll() if
// the designated constraints aren't met.
type GetProfileByUserNameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileByUserNameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileByUserNameRequestMultiError) AllErrors() []error { return m }

// GetProfileByUserNameRequestValidationError is the validation error returned
// by GetProfileByUserNameRequest.Validate if the designated constraints
// aren't met.
type GetProfileByUserNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileByUserNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileByUserNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileByUserNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileByUserNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileByUserNameRequestValidationError) ErrorName() string {
	return "GetProfileByUserNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileByUserNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileByUserNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileByUserNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileByUserNameRequestValidationError{}

// Validate checks the field values on GetProfileByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileByIdRequestMultiError, or nil if none found.
func (m *GetProfileByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProfileByIdRequestMultiError(errors)
	}

	return nil
}

// GetProfileByIdRequestMultiError is an error wrapping multiple validation
// errors returned by GetProfileByIdRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProfileByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileByIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileByIdRequestMultiError) AllErrors() []error { return m }

// GetProfileByIdRequestValidationError is the validation error returned by
// GetProfileByIdRequest.Validate if the designated constraints aren't met.
type GetProfileByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileByIdRequestValidationError) ErrorName() string {
	return "GetProfileByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileByIdRequestValidationError{}

// Validate checks the field values on GetProfileReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProfileReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileReplyMultiError, or nil if none found.
func (m *GetProfileReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProfileReplyValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProfileReplyValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProfileReplyValidationError{
				field:  "Profile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProfileReplyMultiError(errors)
	}

	return nil
}

// GetProfileReplyMultiError is an error wrapping multiple validation errors
// returned by GetProfileReply.ValidateAll() if the designated constraints
// aren't met.
type GetProfileReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileReplyMultiError) AllErrors() []error { return m }

// GetProfileReplyValidationError is the validation error returned by
// GetProfileReply.Validate if the designated constraints aren't met.
type GetProfileReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileReplyValidationError) ErrorName() string { return "GetProfileReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetProfileReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileReplyValidationError{}