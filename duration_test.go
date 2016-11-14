package timex_test

import (
	"testing"
	"time"

	"github.com/classmarkets/timex"
)

var berlin *time.Location

func init() {
	var err error
	berlin, err = time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}
}

func TestAddDays(t *testing.T) {
	cases := map[string]struct {
		t, want time.Time
		days    int
	}{
		"no DST change": {
			t:    time.Date(2016, 11, 14, 8, 0, 0, 0, berlin),
			want: time.Date(2016, 11, 15, 8, 0, 0, 0, berlin),
			days: 1,
		},
		"CET → CEST": {
			t:    time.Date(2016, 03, 26, 8, 0, 0, 0, berlin),
			want: time.Date(2016, 03, 27, 8, 0, 0, 0, berlin),
			days: 1,
		},
		"CEST → CET": {
			t:    time.Date(2016, 10, 29, 8, 0, 0, 0, berlin),
			want: time.Date(2016, 10, 30, 8, 0, 0, 0, berlin),
			days: 1,
		},
		"CET → CEST (multiple days)": {
			t:    time.Date(2016, 03, 25, 8, 0, 0, 0, berlin),
			want: time.Date(2016, 03, 28, 8, 0, 0, 0, berlin),
			days: 3,
		},
		"CEST → CET (multiple days)": {
			t:    time.Date(2016, 10, 28, 8, 0, 0, 0, berlin),
			want: time.Date(2016, 10, 31, 8, 0, 0, 0, berlin),
			days: 3,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			have := timex.AddDays(c.t, c.days, berlin)

			if have != c.want {
				t.Errorf("AddDays(%v, %v, %v):", c.t, c.days, berlin)
				t.Errorf("  want=%s", c.want)
				t.Errorf("  have=%s", have)
			} else {
				t.Logf("AddDays(%v, %v, %v): %s", c.t, c.days, berlin, have)
			}
		})
	}
}

func TestAddWeeks(t *testing.T) {
	cases := map[string]struct {
		t, want time.Time
		weeks   int
	}{
		"no DST change": {
			t:     time.Date(2016, 11, 10, 8, 0, 0, 0, berlin),
			want:  time.Date(2016, 11, 17, 8, 0, 0, 0, berlin),
			weeks: 1,
		},
		"CET → CEST": {
			t:     time.Date(2016, 03, 20, 8, 0, 0, 0, berlin),
			want:  time.Date(2016, 03, 27, 8, 0, 0, 0, berlin),
			weeks: 1,
		},
		"CEST → CET": {
			t:     time.Date(2016, 10, 23, 8, 0, 0, 0, berlin),
			want:  time.Date(2016, 10, 30, 8, 0, 0, 0, berlin),
			weeks: 1,
		},
		"CET → CEST multiple weeks": {
			t:     time.Date(2016, 03, 07, 8, 0, 0, 0, berlin),
			want:  time.Date(2016, 03, 28, 8, 0, 0, 0, berlin),
			weeks: 3,
		},
		"CEST → CET multiple weeks": {
			t:     time.Date(2016, 10, 10, 8, 0, 0, 0, berlin),
			want:  time.Date(2016, 10, 31, 8, 0, 0, 0, berlin),
			weeks: 3,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			have := timex.AddWeeks(c.t, c.weeks, berlin)

			if have != c.want {
				t.Errorf("AddWeeks(%v, %v, %v):", c.t, c.weeks, berlin)
				t.Errorf("  want=%s", c.want)
				t.Errorf("  have=%s", have)
			} else {
				t.Logf("AddWeeks(%v, %v, %v): %s", c.t, c.weeks, berlin, have)
			}
		})
	}
}

func TestAddMonths(t *testing.T) {
	cases := map[string]struct {
		t, want time.Time
		months  int
	}{
		"no DST change": {
			t:      time.Date(2016, 11, 10, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 12, 10, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"CET → CEST": {
			t:      time.Date(2016, 03, 01, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 04, 01, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"CEST → CET": {
			t:      time.Date(2016, 10, 01, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 11, 01, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"CET → CEST multiple months": {
			t:      time.Date(2016, 02, 29, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 05, 31, 8, 0, 0, 0, berlin),
			months: 3,
		},
		"CEST → CET multiple months": {
			t:      time.Date(2016, 7, 31, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 10, 31, 8, 0, 0, 0, berlin),
			months: 3,
		},
		"UTC": {
			t:      time.Date(2016, 7, 31, 8, 0, 0, 0, time.UTC),
			want:   time.Date(2016, 10, 31, 10, 0, 0, 0, berlin),
			months: 3,
		},
		"stick to end of the month": {
			t:      time.Date(2016, 6, 30, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 7, 31, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"stick to end of the month 2": {
			t:      time.Date(2016, 7, 31, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 6, 30, 8, 0, 0, 0, berlin),
			months: -1,
		},
		"stick to end of the month 3": {
			t:      time.Date(2016, 10, 31, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 11, 30, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"stick to end of the month 4": {
			t:      time.Date(2016, 11, 30, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 10, 31, 8, 0, 0, 0, berlin),
			months: -1,
		},
		"stick to beginning of the month": {
			t:      time.Date(2016, 10, 1, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 11, 1, 8, 0, 0, 0, berlin),
			months: 1,
		},
		"stick to beginning of the month 2": {
			t:      time.Date(2016, 11, 1, 8, 0, 0, 0, berlin),
			want:   time.Date(2016, 10, 1, 8, 0, 0, 0, berlin),
			months: -1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			have := timex.AddMonths(c.t, c.months, berlin).Format("2006-01-02 15:04")
			want := c.want.Format("2006-01-02 15:04")

			if have != want {
				t.Errorf("AddMonths(%v, %v, %v):", c.t, c.months, berlin)
				t.Errorf("  want=%s", want)
				t.Errorf("  have=%s", have)
			} else {
				t.Logf("AddMonths(%v, %v, %v): %s", c.t, c.months, berlin, have)
			}
		})
	}
}

func TestIsLastDayOfMonth(t *testing.T) {
	cases := []struct {
		t    time.Time
		want bool
	}{
		{t: time.Date(2016, 1, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 2, 29, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 3, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 4, 30, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 5, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 6, 30, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 7, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 8, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 9, 30, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 10, 31, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 11, 30, 8, 0, 0, 0, berlin), want: true},
		{t: time.Date(2016, 12, 31, 8, 0, 0, 0, berlin), want: true},

		{t: time.Date(2016, 1, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 2, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 3, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 4, 31, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 5, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 6, 31, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 7, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 8, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 9, 31, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 10, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 11, 31, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 12, 30, 8, 0, 0, 0, berlin), want: false},
		{t: time.Date(2016, 12, 32, 8, 0, 0, 0, berlin), want: false},
	}

	for _, c := range cases {
		if have := timex.IsLastDayOfMonth(c.t); have != c.want {
			t.Errorf("IsLastDayOfMonth(%v): %v but want %v", c.t, have, c.want)
		}
	}
}
