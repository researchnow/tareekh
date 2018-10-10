package tareekh_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/researchnow/tareekh"
)

//TestFromDateString
func TestFromDateString(t *testing.T) {
	dt, _ := tareekh.FromDateString("2016-11-22")
	if tareekh.ToShortDate(dt) != "2016-11-22" {
		t.FailNow()
	}

	tests := []struct {
		Date  string
		Valid bool
	}{{
		Date:  "2016-11-22x",
		Valid: false,
	}, {
		Date:  "2016-11",
		Valid: false,
	}, {
		Date:  "",
		Valid: false,
	}, {
		Date:  "2016-11-22",
		Valid: true,
	}}

	for _, tt := range tests {
		_, err := tareekh.FromDateString(tt.Date)
		if tt.Valid && err != nil {
			t.FailNow()
		}

		if !tt.Valid && err == nil {
			t.FailNow()
		}
	}

}

// TestIsDateInFuture
func TestIsDateInFuture(t *testing.T) {
	if tareekh.IsDateInFuture(tareekh.Now().AddDate(0, 0, -1)) {
		t.FailNow()
	}

	if !tareekh.IsDateInFuture(tareekh.Now().AddDate(0, 0, 1)) {
		t.FailNow()
	}
}

// TestFromDayOfMonth
func TestFromDayOfMonth(t *testing.T) {
	day := 2
	tnow := tareekh.Now()
	dm := tareekh.FromDayOfMonth(day)
	if dm.Month() != tnow.Month() || dm.Year() != tnow.Year() || dm.Day() != day {
		t.FailNow()
	}
}

// TestBeginningOfMonth
func TestBeginningOfMonth(t *testing.T) {
	if tareekh.BeginningOfMonth().Day() != 1 {
		t.FailNow()
	}
}

// TestEndOfMonth
func TestEndOfMonth(t *testing.T) {
	if tareekh.EndOfMonth().Day() < tareekh.LeastPossibleDays {
		t.FailNow()
	}
}

// TestDaysAgo
func TestDaysAgo(t *testing.T) {
	today := tareekh.Today()
	days := tareekh.DaysAgo(6)
	diff := today.Day() - days.Day()
	if diff != 6 {
		t.FailNow()
	}
}

// TestBeginningOfDay
func TestBeginningOfDay(t *testing.T) {
	begin := tareekh.BeginningOfDay(tareekh.Now())
	ft := fmt.Sprintf("%d:%d:%d", begin.Hour(), begin.Minute(), begin.Second())
	if ft != "0:0:0" {
		t.FailNow()
	}
}

// TestDifferenceInDays

func TestDifferenceInDays(t *testing.T) {
	format := "2006-01-02T15:04:05.000"
	then, err := time.Parse(format, "2016-12-16T00:14:46.000")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	now, err := time.Parse(format, "2016-12-18T00:14:46.000")
	if err != nil {
		t.FailNow()
	}

	days := tareekh.DifferenceInDays(then, now)
	fmt.Println(days)
	if days != 2 {
		t.FailNow()
	}
}
