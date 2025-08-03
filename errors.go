package gameboy

import "errors"

var (
	ErrNoMoreInstructions = errors.New("no more instructions")
	ErrStackUnderflow     = errors.New("stack underflow")
	ErrStackOverflow      = errors.New("stack overflow")
)
