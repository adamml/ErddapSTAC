package stac

type SpatialExtent struct {
	Bbox [1][4]float32 `json:"bbox"`
}
