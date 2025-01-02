package xdtm

import (
	"fmt"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/markusmobius/go-dateparser"
)

func Get(raw string, opts ...DtmOptFunc) (c carbon.Carbon, err error) {
	c, err = ToCarbon(raw, opts...)
	return
}

func ToCarbon(raw string, opts ...DtmOptFunc) (c carbon.Carbon, err error) {
	opt := DtmOpts{
		baseTime: time.Now().UTC(),
		dpsConfig: &dateparser.Configuration{
			PreferredDateSource: dateparser.Past,
		},
	}

	bindDtmOpts(&opt, opts...)

	for k, v := range opt.replacement {
		raw = strings.ReplaceAll(raw, k, v)
	}

	cfg := opt.dpsConfig
	cfg.CurrentTime = opt.baseTime

	if opt.bySearch {
		_, res, err := dateparser.Search(cfg, raw)
		if err != nil {
			return c, err
		}

		if len(res) == 0 {
			return c, fmt.Errorf("%w: %s", ErrNonDateStr, raw)
		}

		return carbon.CreateFromStdTime(res[0].Date.Time), err
	}

	if opt.layout == "" {
		dt, err := dateparser.Parse(cfg, raw)
		return carbon.CreateFromStdTime(dt.Time), err
	}

	dt, err := dateparser.Parse(cfg, raw, opt.layout)
	if err != nil {
		return carbon.Carbon{}, fmt.Errorf("cannot parse date: %w", err)
	}

	return carbon.CreateFromStdTime(dt.Time), err
}

func GetDateStr(raw string, opts ...DtmOptFunc) string {
	opt := DtmOpts{}
	bindDtmOpts(&opt, opts...)

	c, err := ToCarbon(raw, opts...)
	if err != nil {
		return IfThenElse(opt.fallback != "", opt.fallback, "")
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
func IfThenElse[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}

	return b
}
