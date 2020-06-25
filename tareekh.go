// Package tareekh is a small wrapper which provides functionality making it
// easy to work with time.Time by implementing commonly used methods
package tareekh

import (
	"fmt"
	"math"
	"time"
)

// TimeZone uses whichever is the default unless specified
var TimeZone = ""

const (
	// LeastPossibleDays in a month
	LeastPossibleDays = 28

	// DefaultDateFormat is the golang date format used by default
	DefaultDateFormat = "2006-01-02"
)

// Now returns time.Now() and is timezone aware
func Now() time.Time {
	if TimeZone != "" {
		loc, _ := time.LoadLocation(TimeZone)
		return time.Now().In(loc)
	}
	return time.Now()
}

// Today returns the current time.
func Today() time.Time {
	return Now()
}

// Yesterday returns time object set to yesterdays (-1 Day)
func Yesterday() time.Time {
	return DaysAgo(1)
}

// BeginningOfMonth returns time object set to first day of current month
func BeginningOfMonth() time.Time {
	return FromDayOfMonth(1)
}

// EndOfMonth return time object set to last day of current month
func EndOfMonth() time.Time {
	return BeginningOfMonth().AddDate(0, 1, -1)
}

// DaysAgo returns time for specified number of days back
func DaysAgo(days int) time.Time {
	return Today().AddDate(0, 0, -1*days)
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
	return Now().Before(t)
}

// IsDateInPast returns true if the date is in past
func IsDateInPast(t time.Time) bool {
	return Now().After(t)
}

// BeginningOfDay resets the time to 00:00:00 while keeping date intact
func BeginningOfDay(t time.Time) time.Time {
	hour := time.Duration(t.Hour()) * time.Hour
	min := time.Duration(t.Minute()) * time.Minute
	sec := time.Duration(t.Second()) * time.Second
	return t.Add(-1 * (hour + min + sec))
}

// DifferenceInDays returns the number of days between two given dates.
func DifferenceInDays(then, now time.Time) int {
	diff := now.Sub(then)
	return int(diff.Hours() / 24)
}

// DifferenceInHours returns the number of hours between two given dates.
func DifferenceInHours(then, now time.Time) int {
	diff := now.Sub(then)
	return int(diff.Hours())
}

// DifferenceInSeconds returns the number of seconds between two given dates.
func DifferenceInSeconds(then, now time.Time) int {
	diff := now.Sub(then)
	return int(math.Round(diff.Seconds()))
}

// Format returns the time in string format.
func Format(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

// ExtractBeginningOfMonth ... extract the beginning of the month for the input time. Example: '01/19/2020 10:00:00' is given in the input, output would be '01/01/2020 00:00:00'
func ExtractBeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

// ExtractEndOfMonth ... extract the ending of the month for the input time. Example: '01/31/2020 23:59:59' is given in the input, output would be '01/01/2020 00:00:00'
func ExtractEndOfMonth(t time.Time) time.Time {
	firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	return lastday
}
