package helper

import (
	"time"
)

func Zero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func DayEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func Monday(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return t.AddDate(0, 0, offset)
}

func Sunday(t time.Time) time.Time {
	offset := int(7 - t.Weekday())
	if offset == 7 {
		offset = 0
	}

	return t.AddDate(0, 0, offset)
}
