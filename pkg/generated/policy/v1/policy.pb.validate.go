// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: policy/v1/policy.proto

package policyv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	sharedv1 "github.com/charithe/menshen/pkg/generated/shared/v1"
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
	_ = ptypes.DynamicAny{}

	_ = sharedv1.Effect(0)

	_ = sharedv1.Effect(0)

	_ = sharedv1.Effect(0)
)

// define the regex for a UUID once up-front
var _policy_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetApiVersion() != "paams.dev/v1" {
		return PolicyValidationError{
			field:  "ApiVersion",
			reason: "value must equal paams.dev/v1",
		}
	}

	// no validation rules for Disabled

	// no validation rules for Description

	switch m.PolicyType.(type) {

	case *Policy_ResourcePolicy:

		if v, ok := interface{}(m.GetResourcePolicy()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolicyValidationError{
					field:  "ResourcePolicy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Policy_PrincipalPolicy:

		if v, ok := interface{}(m.GetPrincipalPolicy()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolicyValidationError{
					field:  "PrincipalPolicy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Policy_DerivedRoles:

		if v, ok := interface{}(m.GetDerivedRoles()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PolicyValidationError{
					field:  "DerivedRoles",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return PolicyValidationError{
			field:  "PolicyType",
			reason: "value is required",
		}

	}

	return nil
}

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyValidationError) ErrorName() string { return "PolicyValidationError" }

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyValidationError{}

// Validate checks the field values on ResourcePolicy with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResourcePolicy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetResource()) < 1 {
		return ResourcePolicyValidationError{
			field:  "Resource",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_ResourcePolicy_Resource_Pattern.MatchString(m.GetResource()) {
		return ResourcePolicyValidationError{
			field:  "Resource",
			reason: "value does not match regex pattern \"^[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*(\\\\:[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*)*$\"",
		}
	}

	if !_ResourcePolicy_Version_Pattern.MatchString(m.GetVersion()) {
		return ResourcePolicyValidationError{
			field:  "Version",
			reason: "value does not match regex pattern \"^[[:word:]]+$\"",
		}
	}

	_ResourcePolicy_ImportDerivedRoles_Unique := make(map[string]struct{}, len(m.GetImportDerivedRoles()))

	for idx, item := range m.GetImportDerivedRoles() {
		_, _ = idx, item

		if _, exists := _ResourcePolicy_ImportDerivedRoles_Unique[item]; exists {
			return ResourcePolicyValidationError{
				field:  fmt.Sprintf("ImportDerivedRoles[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_ResourcePolicy_ImportDerivedRoles_Unique[item] = struct{}{}
		}

		if !_ResourcePolicy_ImportDerivedRoles_Pattern.MatchString(item) {
			return ResourcePolicyValidationError{
				field:  fmt.Sprintf("ImportDerivedRoles[%v]", idx),
				reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
			}
		}

	}

	if len(m.GetRules()) < 1 {
		return ResourcePolicyValidationError{
			field:  "Rules",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourcePolicyValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResourcePolicyValidationError is the validation error returned by
// ResourcePolicy.Validate if the designated constraints aren't met.
type ResourcePolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourcePolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourcePolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourcePolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourcePolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourcePolicyValidationError) ErrorName() string { return "ResourcePolicyValidationError" }

// Error satisfies the builtin error interface
func (e ResourcePolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourcePolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourcePolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourcePolicyValidationError{}

var _ResourcePolicy_Resource_Pattern = regexp.MustCompile("^[[:alpha:]][[:word:]\\@\\.\\-]*(\\:[[:alpha:]][[:word:]\\@\\.\\-]*)*$")

var _ResourcePolicy_Version_Pattern = regexp.MustCompile("^[[:word:]]+$")

var _ResourcePolicy_ImportDerivedRoles_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

// Validate checks the field values on ResourceRule with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResourceRule) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAction()) < 1 {
		return ResourceRuleValidationError{
			field:  "Action",
			reason: "value length must be at least 1 runes",
		}
	}

	_ResourceRule_DerivedRoles_Unique := make(map[string]struct{}, len(m.GetDerivedRoles()))

	for idx, item := range m.GetDerivedRoles() {
		_, _ = idx, item

		if _, exists := _ResourceRule_DerivedRoles_Unique[item]; exists {
			return ResourceRuleValidationError{
				field:  fmt.Sprintf("DerivedRoles[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_ResourceRule_DerivedRoles_Unique[item] = struct{}{}
		}

		if !_ResourceRule_DerivedRoles_Pattern.MatchString(item) {
			return ResourceRuleValidationError{
				field:  fmt.Sprintf("DerivedRoles[%v]", idx),
				reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
			}
		}

	}

	_ResourceRule_Roles_Unique := make(map[string]struct{}, len(m.GetRoles()))

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if _, exists := _ResourceRule_Roles_Unique[item]; exists {
			return ResourceRuleValidationError{
				field:  fmt.Sprintf("Roles[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_ResourceRule_Roles_Unique[item] = struct{}{}
		}

		if !_ResourceRule_Roles_Pattern.MatchString(item) {
			return ResourceRuleValidationError{
				field:  fmt.Sprintf("Roles[%v]", idx),
				reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
			}
		}

	}

	if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceRuleValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _ResourceRule_Effect_InLookup[m.GetEffect()]; !ok {
		return ResourceRuleValidationError{
			field:  "Effect",
			reason: "value must be in list [1 2]",
		}
	}

	return nil
}

// ResourceRuleValidationError is the validation error returned by
// ResourceRule.Validate if the designated constraints aren't met.
type ResourceRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceRuleValidationError) ErrorName() string { return "ResourceRuleValidationError" }

// Error satisfies the builtin error interface
func (e ResourceRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceRuleValidationError{}

var _ResourceRule_DerivedRoles_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

var _ResourceRule_Roles_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

var _ResourceRule_Effect_InLookup = map[sharedv1.Effect]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on PrincipalPolicy with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PrincipalPolicy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPrincipal()) < 1 {
		return PrincipalPolicyValidationError{
			field:  "Principal",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_PrincipalPolicy_Principal_Pattern.MatchString(m.GetPrincipal()) {
		return PrincipalPolicyValidationError{
			field:  "Principal",
			reason: "value does not match regex pattern \"^[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*(\\\\:[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*)*$\"",
		}
	}

	if !_PrincipalPolicy_Version_Pattern.MatchString(m.GetVersion()) {
		return PrincipalPolicyValidationError{
			field:  "Version",
			reason: "value does not match regex pattern \"^[[:word:]]+$\"",
		}
	}

	if len(m.GetRules()) < 1 {
		return PrincipalPolicyValidationError{
			field:  "Rules",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PrincipalPolicyValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PrincipalPolicyValidationError is the validation error returned by
// PrincipalPolicy.Validate if the designated constraints aren't met.
type PrincipalPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrincipalPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrincipalPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrincipalPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrincipalPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrincipalPolicyValidationError) ErrorName() string { return "PrincipalPolicyValidationError" }

// Error satisfies the builtin error interface
func (e PrincipalPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipalPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrincipalPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrincipalPolicyValidationError{}

var _PrincipalPolicy_Principal_Pattern = regexp.MustCompile("^[[:alpha:]][[:word:]\\@\\.\\-]*(\\:[[:alpha:]][[:word:]\\@\\.\\-]*)*$")

var _PrincipalPolicy_Version_Pattern = regexp.MustCompile("^[[:word:]]+$")

// Validate checks the field values on PrincipalRule with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PrincipalRule) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetResource()) < 1 {
		return PrincipalRuleValidationError{
			field:  "Resource",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_PrincipalRule_Resource_Pattern.MatchString(m.GetResource()) {
		return PrincipalRuleValidationError{
			field:  "Resource",
			reason: "value does not match regex pattern \"^[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*(\\\\:[[:alpha:]][[:word:]\\\\@\\\\.\\\\-]*)*$\"",
		}
	}

	if len(m.GetActions()) < 1 {
		return PrincipalRuleValidationError{
			field:  "Actions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PrincipalRuleValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PrincipalRuleValidationError is the validation error returned by
// PrincipalRule.Validate if the designated constraints aren't met.
type PrincipalRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrincipalRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrincipalRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrincipalRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrincipalRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrincipalRuleValidationError) ErrorName() string { return "PrincipalRuleValidationError" }

// Error satisfies the builtin error interface
func (e PrincipalRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipalRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrincipalRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrincipalRuleValidationError{}

var _PrincipalRule_Resource_Pattern = regexp.MustCompile("^[[:alpha:]][[:word:]\\@\\.\\-]*(\\:[[:alpha:]][[:word:]\\@\\.\\-]*)*$")

// Validate checks the field values on DerivedRoles with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DerivedRoles) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return DerivedRolesValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_DerivedRoles_Name_Pattern.MatchString(m.GetName()) {
		return DerivedRolesValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
		}
	}

	if len(m.GetDefinitions()) < 1 {
		return DerivedRolesValidationError{
			field:  "Definitions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDefinitions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DerivedRolesValidationError{
					field:  fmt.Sprintf("Definitions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DerivedRolesValidationError is the validation error returned by
// DerivedRoles.Validate if the designated constraints aren't met.
type DerivedRolesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DerivedRolesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DerivedRolesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DerivedRolesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DerivedRolesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DerivedRolesValidationError) ErrorName() string { return "DerivedRolesValidationError" }

// Error satisfies the builtin error interface
func (e DerivedRolesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDerivedRoles.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DerivedRolesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DerivedRolesValidationError{}

var _DerivedRoles_Name_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

// Validate checks the field values on RoleDef with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RoleDef) Validate() error {
	if m == nil {
		return nil
	}

	if !_RoleDef_Name_Pattern.MatchString(m.GetName()) {
		return RoleDefValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
		}
	}

	if len(m.GetParentRoles()) < 1 {
		return RoleDefValidationError{
			field:  "ParentRoles",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_RoleDef_ParentRoles_Unique := make(map[string]struct{}, len(m.GetParentRoles()))

	for idx, item := range m.GetParentRoles() {
		_, _ = idx, item

		if _, exists := _RoleDef_ParentRoles_Unique[item]; exists {
			return RoleDefValidationError{
				field:  fmt.Sprintf("ParentRoles[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_RoleDef_ParentRoles_Unique[item] = struct{}{}
		}

		if !_RoleDef_ParentRoles_Pattern.MatchString(item) {
			return RoleDefValidationError{
				field:  fmt.Sprintf("ParentRoles[%v]", idx),
				reason: "value does not match regex pattern \"^[[:word:]\\\\-\\\\.]+$\"",
			}
		}

	}

	if v, ok := interface{}(m.GetComputation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleDefValidationError{
				field:  "Computation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RoleDefValidationError is the validation error returned by RoleDef.Validate
// if the designated constraints aren't met.
type RoleDefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleDefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleDefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleDefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleDefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleDefValidationError) ErrorName() string { return "RoleDefValidationError" }

// Error satisfies the builtin error interface
func (e RoleDefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoleDef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleDefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleDefValidationError{}

var _RoleDef_Name_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

var _RoleDef_ParentRoles_Pattern = regexp.MustCompile("^[[:word:]\\-\\.]+$")

// Validate checks the field values on Computation with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Computation) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Computation.(type) {

	case *Computation_Match:

		if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ComputationValidationError{
					field:  "Match",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Computation_Script:
		// no validation rules for Script

	default:
		return ComputationValidationError{
			field:  "Computation",
			reason: "value is required",
		}

	}

	return nil
}

// ComputationValidationError is the validation error returned by
// Computation.Validate if the designated constraints aren't met.
type ComputationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComputationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComputationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComputationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComputationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComputationValidationError) ErrorName() string { return "ComputationValidationError" }

// Error satisfies the builtin error interface
func (e ComputationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComputation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComputationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComputationValidationError{}

// Validate checks the field values on Match with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Match) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetExpr()) < 1 {
		return MatchValidationError{
			field:  "Expr",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetExpr() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return MatchValidationError{
				field:  fmt.Sprintf("Expr[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// MatchValidationError is the validation error returned by Match.Validate if
// the designated constraints aren't met.
type MatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchValidationError) ErrorName() string { return "MatchValidationError" }

// Error satisfies the builtin error interface
func (e MatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchValidationError{}

// Validate checks the field values on TestSuite with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TestSuite) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return TestSuiteValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Description

	// no validation rules for Skip

	// no validation rules for SkipReason

	if len(m.GetTests()) < 1 {
		return TestSuiteValidationError{
			field:  "Tests",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTests() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TestSuiteValidationError{
					field:  fmt.Sprintf("Tests[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TestSuiteValidationError is the validation error returned by
// TestSuite.Validate if the designated constraints aren't met.
type TestSuiteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestSuiteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestSuiteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestSuiteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestSuiteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestSuiteValidationError) ErrorName() string { return "TestSuiteValidationError" }

// Error satisfies the builtin error interface
func (e TestSuiteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestSuite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestSuiteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestSuiteValidationError{}

// Validate checks the field values on Test with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Test) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return TestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Description

	// no validation rules for Skip

	// no validation rules for SkipReason

	if m.GetRequest() == nil {
		return TestValidationError{
			field:  "Request",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := sharedv1.Effect_name[int32(m.GetExpectedEffect())]; !ok {
		return TestValidationError{
			field:  "ExpectedEffect",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// TestValidationError is the validation error returned by Test.Validate if the
// designated constraints aren't met.
type TestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestValidationError) ErrorName() string { return "TestValidationError" }

// Error satisfies the builtin error interface
func (e TestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestValidationError{}

// Validate checks the field values on PrincipalRule_Action with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PrincipalRule_Action) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAction()) < 1 {
		return PrincipalRule_ActionValidationError{
			field:  "Action",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PrincipalRule_ActionValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _PrincipalRule_Action_Effect_InLookup[m.GetEffect()]; !ok {
		return PrincipalRule_ActionValidationError{
			field:  "Effect",
			reason: "value must be in list [1 2]",
		}
	}

	return nil
}

// PrincipalRule_ActionValidationError is the validation error returned by
// PrincipalRule_Action.Validate if the designated constraints aren't met.
type PrincipalRule_ActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrincipalRule_ActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrincipalRule_ActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrincipalRule_ActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrincipalRule_ActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrincipalRule_ActionValidationError) ErrorName() string {
	return "PrincipalRule_ActionValidationError"
}

// Error satisfies the builtin error interface
func (e PrincipalRule_ActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipalRule_Action.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrincipalRule_ActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrincipalRule_ActionValidationError{}

var _PrincipalRule_Action_Effect_InLookup = map[sharedv1.Effect]struct{}{
	1: {},
	2: {},
}
