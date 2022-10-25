package xdtm

import (
	"github.com/golang-module/carbon/v2"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/tj/go-naturaldate"
)

var (
	wp = &when.Parser{}
)

func init() {
	wp.Add(en.All...)
	wp.Add(common.All...)
}

func NParseBaseNow(opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().Carbon2Time().UTC()}
	bindDtmOpts(&opt, opts...)

	t, err := naturaldate.Parse("now", opt.baseTime)
	if err != nil {
		return "", err
	}
	return UTCToIso8601(t), err
}

// NParse wrapper of naturedate.Parse
// WARN: string like `no datetime str`, will return current datetime and no error
func NParse(raw string, opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().Carbon2Time().UTC()}
	bindDtmOpts(&opt, opts...)

	if raw == "" {
		return "", ErrEmptyStr
	}

	t, err := naturaldate.Parse(raw, opt.baseTime)
	if err != nil {
		return "", err
	}

	str := UTCToIso8601(t)

	base, e := NParseBaseNow(opts...)
	if e != nil {
		return "", e
	}
	// when str parsed is same with base str
	// 1. raw is "now"
	// 2. raw is non date string
	// so we try to parse it with WParse
	if str == base {
		wv, e := WParse(raw, opts...)
		if e == ErrNonDateStr {
			return wv, ErrNonDateStr
		}
		return str, nil
	}

	return str, nil
}

func WParse(raw string, opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().Carbon2Time().UTC()}
	bindDtmOpts(&opt, opts...)
	if raw == "" {
		return "", ErrEmptyStr
	}

	s, err := wp.Parse(raw, opt.baseTime)
	if err != nil {
		return "", err
	}

	if s == nil {
		return "", ErrNonDateStr
	}

	return UTCToIso8601(s.Time), nil
}
