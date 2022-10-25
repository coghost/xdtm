package xdtm

import "time"

type DtmOpts struct {
	precision int

	baseTime      time.Time
	baseTimestamp int64

	layout string
	zone   string
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
