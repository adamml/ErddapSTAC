package stac

type Statistics struct {
	Minimum      float64 `json:"minimum,omitempty"`
	Maximum      float64 `json:"maximum,omitempty"`
	Mean         float64 `json:"mean,omitempty"`
	StdDev       float64 `json:"stddev,omitempty"`
	Count        int64   `json:"stddev,omitempty"`
	ValidPercent float64 `json:"valid_percent,omitempty"`
}
