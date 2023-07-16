package xdtm

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/markusmobius/go-dateparser"
)

func Get(raw string, opts ...DtmOptFunc) (c carbon.Carbon, err error) {
	c, err = ToCarbon(raw, opts...)
	return
}

func ToCarbon(raw string, opts ...DtmOptFunc) (c carbon.Carbon, err error) {
	opt := DtmOpts{baseTime: time.Now().UTC()}
	bindDtmOpts(&opt, opts...)

	for k, v := range opt.replacement {
		raw = strings.ReplaceAll(raw, k, v)
	}

	cfg := &dateparser.Configuration{
		CurrentTime: opt.baseTime,
	}

	if opt.bySearch {
		_, res, err := dateparser.Search(cfg, raw)
		if err != nil {
			return c, err
		}
		if len(res) == 0 {
			return c, errors.New("no date str found")
		}
		dts := res[0]
		return carbon.Time2Carbon(dts.Date.Time), err
	}

	if opt.layout == "" {
		dt, err := dateparser.Parse(cfg, raw)
		return carbon.Time2Carbon(dt.Time), err
	}

	dt, err := dateparser.Parse(cfg, raw, opt.layout)
	return carbon.Time2Carbon(dt.Time), err
}

func GetDateStr(raw string, opts ...DtmOptFunc) string {
	opt := DtmOpts{baseTime: time.Now().UTC()}
	bindDtmOpts(&opt, opts...)

	c, err := ToCarbon(raw, opts...)
	if err != nil {
		return IfThenElse(opt.fallback != "", opt.fallback, "").(string)
	}
	return c.ToDateString()
}

func GetDateTimeStr(raw string, opts ...DtmOptFunc) string {
	c, err := ToCarbon(raw, opts...)
	if err != nil {
		return ""
	}
	return c.ToDateTimeString()
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
