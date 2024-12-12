package stac

import (
	"time"
)

type TemporalExtent struct {
	Interval [1][2]time.Time `json:"interval"`
}
