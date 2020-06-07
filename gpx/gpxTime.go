package gpx

import (
	"fmt"
	"time"
)

type GpxTime struct {
	time.Time
}

// MarshalText skips the second fractions which are not allowed in GPX format
func (gt *GpxTime) MarshalText() ([]byte, error) {
	if gt.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprint(gt.Time.Format(time.RFC3339))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (gt *GpxTime) IsSet() bool {
	return gt.UnixNano() != nilTime
}
