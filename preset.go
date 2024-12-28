package xdtm

import "errors"

var (
	ErrEmptyStr    = errors.New("empty date str found")
	ErrNonDateStr  = errors.New("not a date str")
	ErrSameAsBase  = errors.New("same as base time, use with caution")
	ErrParseFailed = errors.New("cannot parse datetime")
)

const (
	precisionNone  = 0
	precisionMilli = 3
	precisionMicro = 6
	precisionNano  = 9
)
