package xdtm

import (
	"errors"
	"fmt"

	"github.com/dromara/carbon/v2"
	"github.com/ijt/go-anytime"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/tj/go-naturaldate"
)

func NParseBaseNow(opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().StdTime().UTC()}
	bindDtmOpts(&opt, opts...)

	t, err := anytime.Parse("now", opt.baseTime, anytime.DefaultToPast)
	if err != nil {
		return "", err
	}

	return UTCToIso8601(t), err
}

// NParse wrapper of naturedate.Parse
// WARN: string like `no datetime str`, will return current datetime and no error
func NParse(raw string, opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().StdTime().UTC()}
	bindDtmOpts(&opt, opts...)

	if raw == "" {
		return "", ErrEmptyStr
	}

	if opt.baseTimestamp != 0 {
		opt.baseTime = carbon.CreateFromTimestamp(opt.baseTimestamp, carbon.Greenwich).StdTime().UTC()
	}

	if opt.baseTimeStr != "" {
		opt.baseTime = carbon.Parse(opt.baseTimeStr, carbon.Greenwich).StdTime().UTC()
	}

	if t, err := anytime.Parse(raw, opt.baseTime, anytime.DefaultToPast); err == nil {
		return UTCToIso8601(t), nil
	}

	timeStr, err := naturaldate.Parse(raw, opt.baseTime, naturaldate.WithDirection(naturaldate.Past))
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrParseFailed, err)
	}

	base, err := NParseBaseNow(opts...)
	if err != nil {
		return "", err
	}

	// when str parsed is same with base str
	// 1. raw is "now"
	// 2. raw is non date string
	// so we try to parse it with WParse
	str := UTCToIso8601(timeStr)
	if str == base {
		wv, e := WParse(raw, opts...)
		if errors.Is(e, ErrNonDateStr) {
			return wv, ErrNonDateStr
		}

		return str, nil
	}

	return str, nil
}

func NParseToCarbon(raw string, opts ...DtmOptFunc) (carbon.Carbon, error) {
	str, err := NParse(raw, opts...)
	if err != nil {
		return carbon.Carbon{}, err
	}

	return ToCarbon(str)
}

func MustNParseToCarbon(raw string, opts ...DtmOptFunc) carbon.Carbon {
	c, err := NParseToCarbon(raw, opts...)
	if err != nil {
		panic(err)
	}

	return c
}

// WParse is partial of NParse
func WParse(raw string, opts ...DtmOptFunc) (string, error) {
	opt := DtmOpts{baseTime: carbon.Now().StdTime().UTC()}
	bindDtmOpts(&opt, opts...)

	if raw == "" {
		return "", ErrEmptyStr
	}

	if opt.baseTimestamp != 0 {
		opt.baseTime = carbon.CreateFromTimestamp(opt.baseTimestamp, carbon.Greenwich).StdTime().UTC()
	}

	if opt.baseTimeStr != "" {
		opt.baseTime = carbon.Parse(opt.baseTimeStr, carbon.Greenwich).StdTime().UTC()
	}

	whenParser := &when.Parser{}

	whenParser.Add(en.All...)
	whenParser.Add(common.All...)

	res, err := whenParser.Parse(raw, opt.baseTime)
	if err != nil {
		return "", err
	}

	if res == nil {
		return "", ErrNonDateStr
	}

	return UTCToIso8601(res.Time), nil
}
