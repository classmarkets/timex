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
	if d1 != d0 && d1 == 1 {
		// we overshot by 1 day due to normalization (see documentation of AddDate)
		t1 = t1.Add(-day)
	}

	if IsLastDayOfMonth(t) {
		// we want to stick to the last day of the month
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

	return t.Add(day).Month() != t.Month()
}
