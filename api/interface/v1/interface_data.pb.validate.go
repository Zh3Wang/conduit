// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interface/v1/interface_data.proto

package interfacePb

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

// Validate checks the field values on SingleArticle with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SingleArticle) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SingleArticle with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SingleArticleMultiError, or
// nil if none found.
func (m *SingleArticle) ValidateAll() error {
	return m.validate(true)
}

func (m *SingleArticle) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Slug

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Body

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for Favorited

	// no validation rules for FavoritesCount

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SingleArticleValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SingleArticleValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SingleArticleValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SingleArticleMultiError(errors)
	}

	return nil
}

// SingleArticleMultiError is an error wrapping multiple validation errors
// returned by SingleArticle.ValidateAll() if the designated constraints
// aren't met.
type SingleArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SingleArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SingleArticleMultiError) AllErrors() []error { return m }

// SingleArticleValidationError is the validation error returned by
// SingleArticle.Validate if the designated constraints aren't met.
type SingleArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SingleArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SingleArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SingleArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SingleArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SingleArticleValidationError) ErrorName() string { return "SingleArticleValidationError" }

// Error satisfies the builtin error interface
func (e SingleArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSingleArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SingleArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SingleArticleValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Token

	// no validation rules for Username

	// no validation rules for Bio

	// no validation rules for Image

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Profile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProfileMultiError, or nil if none found.
func (m *Profile) ValidateAll() error {
	return m.validate(true)
}

func (m *Profile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Bio

	// no validation rules for Image

	// no validation rules for Following

	if len(errors) > 0 {
		return ProfileMultiError(errors)
	}

	return nil
}

// ProfileMultiError is an error wrapping multiple validation errors returned
// by Profile.ValidateAll() if the designated constraints aren't met.
type ProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProfileMultiError) AllErrors() []error { return m }

// ProfileValidationError is the validation error returned by Profile.Validate
// if the designated constraints aren't met.
type ProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileValidationError) ErrorName() string { return "ProfileValidationError" }

// Error satisfies the builtin error interface
func (e ProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileValidationError{}