package date

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month time.Month
	day   int
}

func (d Date) ToTime() time.Time {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, time.UTC)
}

func FromTime(t time.Time) Date {
	return Date{t.Year(), t.Month(), t.Day()}
}

func (d Date) ToString() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}
