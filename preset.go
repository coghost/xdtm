package xdtm

import "errors"

var (
	ErrEmptyStr   = errors.New("empty date str found")
	ErrNonDateStr = errors.New("not a date str")
	ErrSameAsBase = errors.New("same as base time, use with caution")
)

const (
	_PRE_None  = 0
	_PRE_Milli = 3
	_PRE_Micro = 6
	_PRE_Nano  = 9
)
