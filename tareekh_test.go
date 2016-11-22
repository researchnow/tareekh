package tareekh_test

import (
	"testing"

	"github.com/peanut-labs/tareekh"
)

//TestFromDateString
func TestFromDateString(t *testing.T) {
	dt, err := tareekh.FromDateString("2016-11-22")
	if err != nil {
		t.FailNow()
	}
	if tareekh.ToShortDate(dt) != "2016-11-22" {
		t.FailNow()
	}

	_, err = tareekh.FromDateString("2016-11-22x")
	if err == nil {
		t.FailNow()
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
