package tareekh_test

import (
	"testing"

	"github.com/peanut-labs/tareekh"
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

//TestIsDateInFuture
func TestIsDateInFuture(t *testing.T) {
	if tareekh.IsDateInFuture(tareekh.Now().AddDate(0, 0, -1)) {
		t.FailNow()
	}

	if !tareekh.IsDateInFuture(tareekh.Now().AddDate(0, 0, 1)) {
		t.FailNow()
	}
}
