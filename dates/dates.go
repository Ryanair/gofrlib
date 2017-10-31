package dates

import (
	"time"
	"log"
)

// Weekday represent the day using ordinary int type.
type Weekday int

// Set of constants that represents days of week in ISO format. Monday represents int type with value 1 and so on
// until Sunday with int 7.
const (
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
	Sunday    Weekday = 7
)

// CreateOfferDates creates offer dates between two dates: startDate and endDate. Created date is then validated
// against passed weekdays, e.g. if weekdays contain a Friday then date 2017-10-13 (Friday) is added to slice.
func CreateOfferDates(startDate string, endDate string, weekdays []Weekday) ([]time.Time, error) {
	var err error
	dateFormat := "2006-01-02"
	start, err := time.Parse(dateFormat, startDate)
	if err != nil {
		log.Println("Cannot parse startDate, ", err)
		return nil, err
	}
	end, err := time.Parse(dateFormat, endDate)
	if err != nil {
		log.Println("Cannot parse endDate, ", err)
		return nil, err
	}
	days := int(end.Sub(start).Hours() / 24)
	times := make([]time.Time, 0)
	tmpDate := start
	for i := 0; i <= days; i++ {
		wd := tmpDate.Weekday()
		if containsWeekday(wd, weekdays) {
			times = append(times, tmpDate)
		}
		tmpDate = tmpDate.Add(24 * time.Hour)
	}
	return times, nil
}

// ToIsoWeekday parse time.Weekday value into dates.Weekday
func ToIsoWeekday(weekday time.Weekday) Weekday {
	if weekday == time.Sunday {
		return Sunday
	}
	return Weekday(int(weekday) % 7)
}

func containsWeekday(weekday time.Weekday, weekdays []Weekday) bool {
	isoWeekday := ToIsoWeekday(weekday)
	for _, weekday := range weekdays {
		if weekday == isoWeekday {
			return true
		}
	}
	return false
}
