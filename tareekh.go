package tareekh

import (
	"fmt"
	"time"
)

var TimeZone = ""

const (
	LeastPossibleDays = 28

	DefaultDateFormat = "2006-01-02"
)

// Now
func Now() time.Time {
	if TimeZone != "" {
		loc, _ := time.LoadLocation(TimeZone)
		return time.Now().In(loc)
	}
	return time.Now()
}

// Today returns the current time
func Today() time.Time {
	return Now()
}

// Yestrday returns time object set to yesterdays (-1 Day)
func Yesterday() time.Time {
	return Today().AddDate(0, 0, -1)
}

// BeginningOfMonth returns time object set to first day of current month
func BeginningOfMonth() time.Time {
	return FromDayOfMonth(1)
}

// EndOfMonth return time object set to last day of current month
func EndOfMonth() time.Time {
	return BeginningOfMonth().AddDate(0, 1, -1)
}

// FromDayOfMonth return time object set to specified day of the current month
func FromDayOfMonth(day int) time.Time {
	tnow := Now()
	return time.Date(tnow.Year(), tnow.Month(), day, 0, 0, 0, 0, tnow.Location())
}

// ToShortDate converts time object to string with format YYYY-MM-DD
func ToShortDate(t time.Time) string {
	return t.Format(DefaultDateFormat)
}

// FromDateString parses date string of format YYYY-MM-DD and returns time object
func FromDateString(dt string) (time.Time, error) {
	localizedFormat := fmt.Sprintf("%s %s", DefaultDateFormat, "MST")
	zone, _ := Now().Zone()
	t, err := time.Parse(localizedFormat, fmt.Sprintf("%s %s", dt, zone))
	if err != nil {
		return t, err
	}
	return t, nil
}

// IsDateInFuture returns true if the date is in future
func IsDateInFuture(t time.Time) bool {
	return Now().Sub(t).Hours() < 0
}
