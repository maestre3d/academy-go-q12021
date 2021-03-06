package domain

import (
	"strconv"
	"strings"
)

const (
	// error types groups
	domain         = "domain"
	infrastructure = "infrastructure"
	// specific error types
	notFound      = "not found"
	alreadyExists = "already exists"
	outOfRange    = "out of range"
	invalidFormat = "invalid format"
	required      = "required"
)

// Error contains specific mechanisms useful for further error mapping and other
// specific use cases
type Error struct {
	group       string
	kind        string
	entity      string
	description string
}

// Error returns the current error description
func (e Error) Error() string {
	return e.description
}

func (e Error) Entity() string {
	return e.entity
}

// IsDomain checks if the current error belongs to Domain error group
func (e Error) IsDomain() bool {
	return e.group == domain
}

// IsDomain checks if the current error belongs to Infrastructure error group
func (e Error) IsInfrastructure() bool {
	return e.group == infrastructure
}

// IsNotFound checks if the current error belongs to Not Found error types
func (e Error) IsNotFound() bool {
	return e.kind == notFound
}

// IsAlreadyExists checks if the current error belongs to Already Exists error types
func (e Error) IsAlreadyExists() bool {
	return e.kind == alreadyExists
}

// IsOutOfRange checks if the current error belongs to Out of Range error types
func (e Error) IsOutOfRange() bool {
	return e.kind == outOfRange
}

// IsInvalidFormat checks if the current error belongs to Invalid Format error types
func (e Error) IsInvalidFormat() bool {
	return e.kind == invalidFormat
}

// IsRequired checks if the current error belongs to Required error types
func (e Error) IsRequired() bool {
	return e.kind == required
}

// NewDomain creates an Error for Domain generic use cases
func NewDomain(e, d string) Error {
	return Error{
		group:       domain,
		entity:      e,
		description: d,
	}
}

// NewInfrastructure creates an Error for Infrastructure generic use cases
func NewInfrastructure(d string) Error {
	return Error{
		group:       infrastructure,
		description: d,
	}
}

// NewNotFound creates an Error for Not Found use cases
func NewNotFound(e string) Error {
	entityDesc := ""
	if e != "" {
		entityDesc = e + " "
	}
	return Error{
		group:       domain,
		kind:        notFound,
		entity:      e,
		description: entityDesc + "not found",
	}
}

// NewAlreadyExists creates an Error for Already Exists use cases
func NewAlreadyExists(e string) Error {
	entityDesc := ""
	if e != "" {
		entityDesc = e + " "
	}
	return Error{
		group:       domain,
		kind:        alreadyExists,
		entity:      e,
		description: entityDesc + "already exists",
	}
}

// NewOutOfRange creates an Error for Out of Range use cases
func NewOutOfRange(e string, a, b int) Error {
	entityDesc := ""
	if e != "" {
		entityDesc = e + " is "
	}
	return Error{
		group:       domain,
		kind:        outOfRange,
		entity:      e,
		description: entityDesc + "out of range [" + strconv.Itoa(a) + "," + strconv.Itoa(b) + ")",
	}
}

// NewInvalidFormat creates an Error for Invalid Format use cases
func NewInvalidFormat(e string, exp ...string) Error {
	entityDesc := ""
	if e != "" {
		entityDesc = e + " contains an "
	}
	return Error{
		group:       domain,
		kind:        invalidFormat,
		entity:      e,
		description: entityDesc + "invalid format, expected [" + strings.Join(exp, ",") + "]",
	}
}

// NewRequired creates an Error for Required use cases
func NewRequired(e string) Error {
	entityDesc := ""
	if e != "" {
		entityDesc = e + " is "
	}
	return Error{
		group:       domain,
		kind:        required,
		entity:      e,
		description: entityDesc + "required",
	}
}
