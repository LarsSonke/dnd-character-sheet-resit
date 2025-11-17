package domain

import "errors"

// Domain errors - these represent business rule violations
// Outer layers can decide how to present these to users
var (
	// ErrNoSpellSlot indicates no spell slot is available for the requested level
	ErrNoSpellSlot = errors.New("no spell slot available")
)
