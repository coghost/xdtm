package xdtm

/**
(
	DFmt  = "YYYY-MM-DD HH:mm:ss"
	DFmt1 = "YYYYMMDDHHmmss"
	DFmt2 = "HH:mm:ss"
	DFDay = "YYYYMMDD"

	ZoneFmt = "YYYY-MM-DDTHH:mm:ssZ"
	DFYMDhms = "2006-01-02 15:04:05"
)

StrNow == xutil.StrNow
**/

import (
	"math"
	"time"

	"github.com/dromara/carbon/v2"
)

const (
	RFC3339MicroLayout = carbon.RFC3339MicroLayout
)

// TimestampAsF64
func TimestampAsF64(offsetSeconds int, opts ...DtmOptFunc) float64 {
	opt := DtmOpts{precision: precisionNone}
	bindDtmOpts(&opt, opts...)

	carb := carbon.Now()
	carb = carb.AddSeconds(offsetSeconds)

	timestamp := int64(0)

	switch opt.precision {
	case precisionNone:
		timestamp = carb.Timestamp()
	case precisionMilli:
		timestamp = carb.TimestampMilli()
	case precisionMicro:
		timestamp = carb.TimestampMicro()
	case precisionNano:
		timestamp = carb.TimestampNano()
	default:
		timestamp = 0
	}

	return float64(timestamp) / math.Pow10(opt.precision)
}

func TimestampAsI64(offsetSeconds int, opts ...DtmOptFunc) int64 {
	return int64(TimestampAsF64(offsetSeconds, opts...))
}

// PythonTimeTime returns same format with python's time.time() `1234567890.123456`
//
//	this is a wrapper of carbon.Timestamp-Milli/Micro/Nano
func PythonTimeTime(offsetArgs ...int) float64 {
	offset := 0
	if len(offsetArgs) > 0 {
		offset = offsetArgs[0]
	}

	return TimestampAsF64(offset, WithPrecision(precisionMicro))
}

// Now alias of carbon.Now()
func Now(timezone ...string) (c carbon.Carbon) {
	return carbon.Now(timezone...)
}

// UTCNow alias of carbon.Now(carbon.UTC)
func UTCNow() (c carbon.Carbon) {
	return carbon.Now(carbon.UTC)
}

// UTCToIso8601 outputs a string in "2006-01-02T15:04:05-07:00" layout.
func UTCToIso8601(args ...time.Time) string {
	now := UTCNow()

	if len(args) > 0 {
		utc, _ := time.LoadLocation(carbon.UTC)
		c := carbon.CreateFromStdTime(args[0])
		now = c.SetLocation(utc)
	}

	return now.ToIso8601String()
}

// StrNow alias of carbon.Now().ToDateTimeString()
//   - outputs a string in "2006-01-02 15:04:05" layout.
func StrNow() string {
	return carbon.Now().ToDateTimeString()
}

// StrToday
//   - outputs a string in "2006-01-02" layout.
func StrToday() string {
	return carbon.Now().ToDateString()
}

// StrYesterday
//   - outputs a string in "2006-01-02" layout.
func StrYesterday() string {
	return carbon.Now().SubDay().ToDateString()
}

// ToShortDateTimeString outputs a string in "20060102150405" layout.
func StrNowShortDatetime() string {
	return carbon.Now().ToShortDateTimeString()
}

func CarbonNow() carbon.Carbon {
	return carbon.Now()
}

// Str2Unix returns timestamp of given string
func Str2Unix(str string, opts ...DtmOptFunc) int64 {
	opt := DtmOpts{layout: carbon.DateTimeLayout}
	bindDtmOpts(&opt, opts...)

	return carbon.ParseByLayout(str, opt.layout).Timestamp()
}

// Unix2Str returns a string in "2006-01-02 15:04:05" layout
func Unix2Str(timestamp int64) string {
	return carbon.CreateFromTimestamp(timestamp).ToDateTimeString()
}
