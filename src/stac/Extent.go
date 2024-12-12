package stac

type Extent struct {
	Spatial  SpatialExtent  `json:"spatial"`
	Temporal TemporalExtent `json:"temporal"`
}
