// Package timex provides functions for adding days, weeks or months to
// a `time.Time` while handling daylight saving time (DST) correctly.
package timex

import "time"

const day = 24 * time.Hour

// AddDays returns the result of adding a given amount of days to t and
// accounts for any DST changes at location l.
func AddDays(t time.Time, days int, l *time.Location) time.Time {
	return t.In(l).AddDate(0, 0, days)
}

// AddWeeks returns the result of adding a given amount of weeks to t and
// accounts for any DST changes at location l.
func AddWeeks(t time.Time, weeks int, l *time.Location) time.Time {
	return t.In(l).AddDate(0, 0, weeks*7)
}

// AddMonths adds the given amount of months to t. Additionally if t represents
// the last day of its month then the returned time t' represents the last day
// of the next month after t. Otherwise t.Day() will equal t'.Day().
func AddMonths(t time.Time, months int, l *time.Location) time.Time {
	t1 := t.In(l).AddDate(0, months, 0)
	if !IsLastDayOfMonth(t) {
		return t1
	}

	d0 := t.Day()
	d1 := t1.Day()

	if d1 != d0 && d1 == 1 {
		// we overshot by 1 day due to normalization (see documentation of AddDate)
		t1 = t1.Add(-day)
	}

	// we want to stick to the last day of the month
	tmp, m1 := t1, t1.Month()
	for {
		tmp = tmp.Add(day)
		if tmp.Month() != m1 {
			break
		}
		t1 = tmp
	}

	return t1
}

// IsLastDayOfMonth returns true if t is any time within the last day of the
// month of t.
func IsLastDayOfMonth(t time.Time) bool {
	if t.Day() < 29 {
		return false
	}

	return t.Add(day).Month() != t.Month()
}
