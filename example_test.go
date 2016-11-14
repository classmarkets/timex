package timex_test

import (
	"fmt"
	"time"

	"github.com/classmarkets/timex"
)

func ExampleAddDays() {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	t0 := time.Date(2016, 3, 26, 10, 0, 0, 0, berlin)

	// DST in germany in 2016 started at 2016-03-27 02:00
	// Therefore just adding 24 hours would yield the wrong time
	t1 := t0.Add(3 * 24 * time.Hour)
	t2 := timex.AddDays(t0, 3, berlin)

	fmt.Println("wrong:  ", t1)
	fmt.Println("correct:", t2)

	// output:
	// wrong:   2016-03-29 11:00:00 +0200 CEST
	// correct: 2016-03-29 10:00:00 +0200 CEST
}
