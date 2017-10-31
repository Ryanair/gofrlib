package dates

import (
	"testing"
	"time"
)

func TestCalculateDays(t *testing.T) {
	startDate := "2017-10-02"
	endDate := "2017-10-08"
	weekdays := []Weekday{Monday, Wednesday, Friday}
	dates, _ := CreateOfferDates(startDate, endDate, weekdays)

	if len(dates) != 3 {
		t.Errorf("Count of all dates cannot be different than 3! Got: %d, expected: %d", len(dates), 3)
	}
}

func TestToIsoWeekday(t *testing.T) {
	tables := []struct {
		x time.Weekday
		y Weekday
	}{
		{time.Sunday, Sunday},
		{time.Monday, Monday},
		{time.Tuesday, Tuesday},
		{time.Wednesday, Wednesday},
		{time.Thursday, Thursday},
		{time.Friday, Friday},
		{time.Saturday, Saturday},
	}
	for _, table := range tables {
		weekday := ToIsoWeekday(table.x)
		some := table.y
		if weekday != some {
			t.Errorf("ToIsoWeekday of (%v) was incorrect, got: %v, want: %v.", table.x, int(table.x), table.y)
		}
	}
}
