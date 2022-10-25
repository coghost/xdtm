package xdtm_test

import (
	"testing"
	"time"

	"github.com/coghost/xdtm"
	"github.com/stretchr/testify/suite"
)

type NaturalSuite struct {
	suite.Suite
}

func TestNatural(t *testing.T) {
	suite.Run(t, new(NaturalSuite))
}

func (s *NaturalSuite) SetupSuite() {
}

func (s *NaturalSuite) TearDownSuite() {
}

func (s *NaturalSuite) Test_00_NParseBaseNow() {
	testsIn := []struct {
		raw   string
		wantS string
		wantE error
	}{
		{"string with no datetime", "2009-02-13T23:31:30+00:00", nil},
		// {"", "2009-02-13T23:31:30+00:00", nil},
		// {"1 minute", "2009-02-13T23:31:30+00:00", nil},
	}
	base := time.Unix(1234567890, 0).UTC()
	for _, tt := range testsIn {
		v, e := xdtm.NParseBaseNow(xdtm.WithBaseTime(base))
		s.Equal(tt.wantE, e)
		s.Equal(tt.wantS, v)
	}
}

func (s *NaturalSuite) Test_02_NParse() {
	var past = []struct {
		raw   string
		wantS string
		wantE error
	}{
		// {"2 days ago", "2009-02-11T00:00:00+00:00", nil},
		// {"3 hours ago", "2009-02-13T20:31:30+00:00", nil},
		// {"yesterday", "2009-02-12T00:00:00+00:00", nil},

		// now
		{"now", "2009-02-13T23:31:30+00:00", nil},
		{"right now", "2009-02-13T23:31:30+00:00", nil},
		{"  right  now  ", "2009-02-13T23:31:30+00:00", nil},

		// error
		{"string with no datetime", "", xdtm.ErrNonDateStr},
		{"", "", xdtm.ErrEmptyStr},

		// seconds
		// {"10 second", "2009-02-13T23:30:30+00:00", nil},

		// minutes
		{"1 minute", "2009-02-13T23:30:30+00:00", nil},
		{"next minute", "2009-02-13T23:32:30+00:00", nil},
		{"last minute", "2009-02-13T23:30:30+00:00", nil},
		{"one minute", "2009-02-13T23:30:30+00:00", nil},
		{"1 minute ago", "2009-02-13T23:30:30+00:00", nil},
		{"5 minutes ago", "2009-02-13T23:26:30+00:00", nil},
		{"five minutes ago", "2009-02-13T23:26:30+00:00", nil},
		{"   5    minutes  ago   ", "2009-02-13T23:26:30+00:00", nil},
		{"2 minutes from now", "2009-02-13T23:33:30+00:00", nil},
		{"two minutes from now", "2009-02-13T23:33:30+00:00", nil},
		{"Message me in 2 minutes", "2009-02-13T23:33:30+00:00", nil},
		{"Message me in 3 minutes from now", "2009-02-13T23:34:30+00:00", nil},

		// hours
		{"1 hour", "2009-02-13T22:31:30+00:00", nil},
		{"last hour", "2009-02-13T22:31:30+00:00", nil},
		{"next hour", "2009-02-14T00:31:30+00:00", nil},
		{"1 hour ago", "2009-02-13T22:31:30+00:00", nil},
		{"6 hours ago", "2009-02-13T17:31:30+00:00", nil},
		{"1 hour from now", "2009-02-14T00:31:30+00:00", nil},
		{"Remind me in 1 hour", "2009-02-14T00:31:30+00:00", nil},
		{"Remind me in 1 hour from now", "2009-02-14T00:31:30+00:00", nil},
		{"Remind me in 1 hour and 3 minutes from now", "2009-02-14T00:34:30+00:00", nil},
		{"Remind me in an hour", "2009-02-14T00:31:30+00:00", nil},
		{"Remind me in an hour from now", "2009-02-14T00:31:30+00:00", nil},

		// days
		{"1 day", "2009-02-12T00:00:00+00:00", nil},
		{"next day", "2009-02-14T00:00:00+00:00", nil},
		{"1 day ago", "2009-02-12T00:00:00+00:00", nil},
		{"3 days ago", "2009-02-10T00:00:00+00:00", nil},
		{"3 days ago at 11:25am", "2009-02-10T11:25:00+00:00", nil},
		{"1 day from now", "2009-02-14T23:31:30+00:00", nil},
		{"Remind me one day from now", "2009-02-14T23:31:30+00:00", nil},
		{"Remind me in a day", "2009-02-14T23:31:30+00:00", nil},
		{"Remind me in one day", "2009-02-14T23:31:30+00:00", nil},
		{"Remind me in one day from now", "2009-02-14T23:31:30+00:00", nil},

		// weeks
		{"1 week", "2009-02-06T00:00:00+00:00", nil},
		{"1 week ago", "2009-02-06T00:00:00+00:00", nil},
		{"2 weeks ago", "2009-01-30T00:00:00+00:00", nil},
		{"2 weeks ago at 8am", "2009-01-30T08:00:00+00:00", nil},
		{"next week", "2009-02-20T00:00:00+00:00", nil},
		{"Message me in a week", "2009-02-20T23:31:30+00:00", nil},
		{"Message me in one week", "2009-02-20T23:31:30+00:00", nil},
		{"Message me in one week from now", "2009-02-20T23:31:30+00:00", nil},
		{"Message me in two weeks from now", "2009-02-27T23:31:30+00:00", nil},
		{"Message me two weeks from now", "2009-02-27T23:31:30+00:00", nil},
		{"Message me in two weeks", "2009-02-27T23:31:30+00:00", nil},

		// months
		{"1 month ago", "2009-01-13T23:31:30+00:00", nil},
		{"last month", "2009-01-13T23:31:30+00:00", nil},
		{"next month", "2009-03-13T23:31:30+00:00", nil},
		{"1 month ago at 9:30am", "2009-01-13T09:30:00+00:00", nil},
		{"2 months ago", "2008-12-13T23:31:30+00:00", nil},
		{"12 months ago", "2008-02-13T23:31:30+00:00", nil},
		{"1 month from now", "2009-03-13T23:31:30+00:00", nil},
		{"next 2 months", "2009-04-13T23:31:30+00:00", nil},
		{"2 months from now", "2009-04-13T23:31:30+00:00", nil},
		{"12 months from now at 6am", "2010-02-13T06:00:00+00:00", nil},
		{"Remind me in 12 months from now at 6am", "2010-02-13T06:00:00+00:00", nil},
		{"Remind me in a month", "2009-03-13T23:31:30+00:00", nil},
		{"Remind me in 2 months", "2009-04-13T23:31:30+00:00", nil},
		{"Remind me in a month from now", "2009-03-13T23:31:30+00:00", nil},
		{"Remind me in 2 months from now", "2009-04-13T23:31:30+00:00", nil},

		// years
		{"last year", "2008-02-13T23:31:30+00:00", nil},
		{"next year", "2010-02-13T23:31:30+00:00", nil},
		{"one year ago", "2008-02-13T23:31:30+00:00", nil},
		{"one year from now", "2010-02-13T23:31:30+00:00", nil},
		{"two years ago", "2007-02-13T23:31:30+00:00", nil},
		{"2 years ago", "2007-02-13T23:31:30+00:00", nil},
		{"Remind me in one year from now", "2010-02-13T23:31:30+00:00", nil},
		{"Remind me in a year", "2010-02-13T23:31:30+00:00", nil},
		{"Remind me in a year from now", "2010-02-13T23:31:30+00:00", nil},

		// today
		{"today", "2009-02-13T00:00:00+00:00", nil},
		{"today at 10am", "2009-02-13T10:00:00+00:00", nil},

		// yesterday
		{"yesterday", "2009-02-12T00:00:00+00:00", nil},
		{"yesterday 10am", "2009-02-12T10:00:00+00:00", nil},
		{"yesterday at 10am", "2009-02-12T10:00:00+00:00", nil},
		{"yesterday at 10:15am", "2009-02-12T10:15:00+00:00", nil},

		// tomorrow
		{"tomorrow", "2009-02-14T00:00:00+00:00", nil},
		{"tomorrow 10am", "2009-02-14T10:00:00+00:00", nil},
		{"tomorrow at 10am", "2009-02-14T10:00:00+00:00", nil},
		{"tomorrow at 10:15am", "2009-02-14T10:15:00+00:00", nil},

		// past weekdays
		{"sunday", "2009-02-08T00:00:00+00:00", nil},
		{"monday", "2009-02-09T00:00:00+00:00", nil},
		{"tuesday", "2009-02-10T00:00:00+00:00", nil},
		{"wednesday", "2009-02-11T00:00:00+00:00", nil},
		{"thursday", "2009-02-12T00:00:00+00:00", nil},
		{"friday", "2009-02-06T00:00:00+00:00", nil},
		{"saturday", "2009-02-07T00:00:00+00:00", nil},

		{"last sunday", "2009-02-08T00:00:00+00:00", nil},
		{"past sunday", "2009-02-08T00:00:00+00:00", nil},
		{"last monday", "2009-02-09T00:00:00+00:00", nil},
		{"last tuesday", "2009-02-10T00:00:00+00:00", nil},
		{"last wednesday", "2009-02-11T00:00:00+00:00", nil},
		{"last thursday", "2009-02-12T00:00:00+00:00", nil},
		{"last friday", "2009-02-06T00:00:00+00:00", nil},
		{"last saturday", "2009-02-07T00:00:00+00:00", nil},

		// ordinal dates
		{"november 15th", "2008-11-15T23:31:30+00:00", nil},
		{"december 1st", "2008-12-01T23:31:30+00:00", nil},
		{"december 2nd", "2008-12-02T23:31:30+00:00", nil},
		{"december 3rd", "2008-12-03T23:31:30+00:00", nil},
		{"december 4th", "2008-12-04T23:31:30+00:00", nil},
		{"december 15th", "2008-12-15T23:31:30+00:00", nil},
		{"december 23rd", "2008-12-23T23:31:30+00:00", nil},
		{"december 23rd 5pm", "2008-12-23T17:00:00+00:00", nil},
		{"december 23rd at 5pm", "2008-12-23T17:00:00+00:00", nil},
		{"december 23rd at 5:25pm", "2008-12-23T17:25:00+00:00", nil},

		// 12-hour clock
		{"10am", "2009-02-13T10:00:00+00:00", nil},
		{"10 am", "2009-02-13T10:00:00+00:00", nil},
		{"5pm", "2009-02-13T17:00:00+00:00", nil},
		{"10:25am", "2009-02-13T10:25:00+00:00", nil},
		{"1:05pm", "2009-02-13T13:05:00+00:00", nil},
		{"10:25:10am", "2009-02-13T10:25:10+00:00", nil},
		{"1:05:10pm", "2009-02-13T13:05:10+00:00", nil},

		// 24-hour clock
		{"10", "2009-02-13T10:00:00+00:00", nil},
		{"10:25", "2009-02-13T10:25:00+00:00", nil},
		{"10:25:30", "2009-02-13T10:25:30+00:00", nil},
		{"17", "2009-02-13T17:00:00+00:00", nil},
		{"17:25:30", "2009-02-13T17:25:30+00:00", nil},

		// case sensitivity
		{"December 23rd AT 5:25 PM", "2008-12-23T17:25:00+00:00", nil},
		{"next December 23rd AT 5:25 PM", "2009-12-23T17:25:00+00:00", nil},

		// QA
		{"Restart the server in 2 days from now", "2009-02-15T23:31:30+00:00", nil},
		{"Remind me on the 5th of next month", "2009-03-05T23:31:30+00:00", nil},
		{"Remind me on the 5th of next month at 7am", "2009-03-05T07:00:00+00:00", nil},
		{"Remind me at 7am on the 5th of next month", "2009-03-05T07:00:00+00:00", nil},
		{"Remind me in one month from now", "2009-03-13T23:31:30+00:00", nil},
		{"Remind me in one month from now at 7am", "2009-03-13T07:00:00+00:00", nil},

		//
		// combined
		{"1 hour 10 minutes ago", "2009-02-13T22:21:30+00:00", nil},
		//
		{"next September", "2009-09-13T23:31:30+00:00", nil},

		// errors
		{`10:am`, "", nil},
	}

	base := time.Unix(1234567890, 0).UTC()
	for _, tt := range past {
		v, e := xdtm.NParse(tt.raw, xdtm.WithBaseTime(base))
		s.Equal(tt.wantE, e, colorize(tt.raw))
		s.Equal(tt.wantS, v, colorize(tt.raw))
	}
}

func colorize(str string) string {
	return str
}

func (s *NaturalSuite) Test_03_DateWParse() {
	tests := []struct {
		raw   string
		wantS string
		wantE error
	}{
		// seconds
		{"1 second ago", "2009-02-13T23:31:29+00:00", nil},
		{"10 second ago", "2009-02-13T23:31:20+00:00", nil},
		// {"1 hour 10 minutes ago", "2009-02-13T23:31:20+00:00", nil},
		{"next September", "2009-09-13T23:31:30+00:00", nil},
	}
	base := time.Unix(1234567890, 0).UTC()
	for _, tt := range tests {
		r, err := xdtm.WParse(tt.raw, xdtm.WithBaseTime(base))
		s.Equal(tt.wantE, err, colorize(tt.raw))
		s.Equal(tt.wantS, r, colorize(tt.raw))
	}
}
