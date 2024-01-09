package xdtm

import "time"

type DtmOpts struct {
	precision int

	baseTime      time.Time
	baseTimestamp int64
	baseTimeStr   string

	layout string
	zone   string

	fallback string

	replacement map[string]string

	bySearch bool
}

type DtmOptFunc func(o *DtmOpts)

func bindDtmOpts(opt *DtmOpts, opts ...DtmOptFunc) {
	for _, f := range opts {
		f(opt)
	}
}

func WithPrecision(i int) DtmOptFunc {
	return func(o *DtmOpts) {
		o.precision = i
	}
}

func WithBaseTime(t time.Time) DtmOptFunc {
	return func(o *DtmOpts) {
		o.baseTime = t
	}
}

func WithBaseTimestamp(i int64) DtmOptFunc {
	return func(o *DtmOpts) {
		o.baseTimestamp = i
	}
}

func WithBaseTimeStr(t string) DtmOptFunc {
	return func(o *DtmOpts) {
		o.baseTimeStr = t
	}
}

func WithLayout(s string) DtmOptFunc {
	return func(o *DtmOpts) {
		o.layout = s
	}
}

func WithZone(s string) DtmOptFunc {
	return func(o *DtmOpts) {
		o.zone = s
	}
}

func WithReplacements(m map[string]string) DtmOptFunc {
	return func(o *DtmOpts) {
		o.replacement = m
	}
}

func WithFallback(s string) DtmOptFunc {
	return func(o *DtmOpts) {
		o.fallback = s
	}
}

func WithBySearch(b bool) DtmOptFunc {
	return func(o *DtmOpts) {
		o.bySearch = b
	}
}
