package xdtm_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/coghost/xdtm"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
)

type CarbonSuite struct {
	suite.Suite
}

func TestCarbon(t *testing.T) {
	suite.Run(t, new(CarbonSuite))
}

func (s *CarbonSuite) SetupSuite() {
}

func (s *CarbonSuite) TearDownSuite() {
}
func (s *CarbonSuite) Test_00_TimeAsF64() {
	// t and t1 should almost equal
	c := xdtm.TimestampAsF64(-3600)
	c1 := xdtm.TimestampAsF64(0) - 3600
	fmt.Printf("c = %#v\n", cast.ToString(c))
	fmt.Printf("c1 = %#v\n", cast.ToString(c1))
}

func (s *CarbonSuite) Test_01_TimeAsPy() {
	c1, c2 := 0.0, 0.0

	go func(v *float64) {
		*v = xdtm.PythonTimeTime(0) - 3600
	}(&c1)

	go func(v *float64) {
		*v = xdtm.PythonTimeTime(-3600)
	}(&c2)

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("c1 = %#v\n", cast.ToString(c1))
	fmt.Printf("c2 = %#v\n", cast.ToString(c2))

	df1 := math.Abs(c1 - c2)
	s.LessOrEqual(df1, 1e-5)
}

func (s *CarbonSuite) Test_02_tostring() {
	v := xdtm.UTCToIso8601()
	s.Contains(v, "T")
}

func (s *CarbonSuite) Test_03_strnow() {
	w1, g1, g2 := "", "", ""
	go func(v *string) {
		*v = ""
	}(&g1)

	go func(g1 *string) {
		// DFmt: carbon.Now().ToDateTimeString()
		// DFmt1: carbon.Now().ToShortDateTimeString()
		// DFmt2: carbon.Now().ToTimeString()
		// DFDay: carbon.Now().ToShortDateString()
		*g1 = carbon.Now().ToShortDateString()
	}(&w1)

	go func(v *string) {
		*v = xdtm.StrNow()
	}(&g2)

	time.Sleep(100 * time.Millisecond)
	s.Equal(w1, g1)

	s.NotEmpty(g1)
	fmt.Println(g1, g2)
}

func (s *CarbonSuite) Test_04_utcstrnow() {
	w1, g1, g2 := "", "", ""

	go func(v *string) {
		// DFmt: carbon.Now().ToDateTimeString()
		// DFmt1: carbon.Now().ToShortDateTimeString()
		// DFmt2: carbon.Now().ToTimeString()
		// DFDay: carbon.Now().ToShortDateString()
		*v = xdtm.UTCToIso8601()
	}(&g1)

	go func(v *string) {
		// *v = xdtm.UTCNow().ToIso8601String()
		*v = ""
	}(&g2)

	time.Sleep(100 * time.Millisecond)
	s.Equal(w1, g1)
	if g2 != "" {
		s.Equal(g1, g2)
	}

	s.NotEmpty(g1)
	fmt.Println(w1, g1)
}

func (s *CarbonSuite) Test_05_unix() {
	str := "2022-10-25 12:27:15"
	ts := int64(1666672035)
	s2 := xdtm.StrNow()

	g1 := xdtm.Str2Unix(s2) - xdtm.Str2Unix(str)
	s.Greater(g1, int64(0))

	g2 := xdtm.Str2Unix(str)
	s.Equal(ts, g2)

	g3 := carbon.CreateFromTimestamp(ts).ToDateTimeString()
	s.Equal(str, g3)
}
