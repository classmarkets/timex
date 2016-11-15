// Package timex provides functions for adding days, weeks or months to
// a `time.Time` while handling daylight saving time (DST) correctly.
package timex

import "time"

const (
	day  = 24 * time.Hour
	week = 7 * day
)

// AddDays returns the result of adding a given amount of days to t and
// accounts for any DST changes at location l.
func AddDays(t time.Time, days int, l *time.Location) time.Time {
	return addDuration(t, l, time.Duration(days)*day)
}

// AddWeeks returns the result of adding a given amount of weeks to t and
// accounts for any DST changes at location l.
func AddWeeks(t time.Time, weeks int, l *time.Location) time.Time {
	return addDuration(t, l, time.Duration(weeks)*week)
}

func addDuration(t0 time.Time, l *time.Location, d time.Duration) time.Time {
	t0 = t0.In(l)
	t1 := t0.Add(d)

	// account for changes due to DST
	_, oldOffset := t0.Zone()
	_, newOffset := t1.Zone()
	diff := time.Duration(oldOffset-newOffset) * time.Second

	return t1.Add(diff)
}

// AddMonths adds the given amount of months to t. Additionally if t represents
// the last day of its month then the returned time t' represents the last day
// of the next month after t. Otherwise t.Day() will equal t'.Day().
func AddMonths(t time.Time, months int, l *time.Location) time.Time {
	t = t.In(l)
	t1 := t.AddDate(0, months, 0)
	d0, d1 := t.Day(), t1.Day()
	if d1 != d0 {
		// due to normalization in AddDate we might need to correct
		// t1. See documentation of AddDate.
		if months > 0 && d1 == 1 {
			// we overshot by 1 day
			t1 = t1.Add(-day)
		}
		if months < 0 {
			if t1.Month() == t.Month() {
				// ended up in the same month due to normalization
				// worst case: 2015-03-31 - 1month = 2015-02-31 = 2015-03-03
				diff := time.Duration(t1.Day()) * day
				t1 = t1.Add(-diff)
			}
		}
	}

	if IsLastDayOfMonth(t) {
		_, oldOffset := t1.Zone()
		tmp, m1 := t1, t1.Month()
		for {
			tmp = tmp.Add(day)
			if tmp.Month() != m1 {
				break
			}
			t1 = tmp
		}

		if _, newOffset := t1.Zone(); newOffset != oldOffset {
			diff := time.Duration(oldOffset-newOffset) * time.Second
			t1 = t1.Add(diff)
		}
	}

	return t1
}

// IsLastDayOfMonth returns true if t is any time within the last day of the
// month of t.
func IsLastDayOfMonth(t time.Time) bool {
	if t.Day() < 28 {
		return false
	}

	// by definition t is the last day of the month if adding another day would
	// result in a time within the next month.
	return t.Add(day).Month() != t.Month()
}
