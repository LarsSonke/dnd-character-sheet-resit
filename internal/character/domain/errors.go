package domain

import "errors"

// Domain errors - these represent business rule violations
// Outer layers decide how to present these to users
var (
	// ErrNoSpellSlot indicates a spell slot is not available for casting
	ErrNoSpellSlot = errors.New("spell slot not available")
)
