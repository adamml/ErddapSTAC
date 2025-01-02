package geojson

const TYPE_POINT = "Point"
const TYPE_LINESTRING = "LineString"
const TYPE_POLYGON = "Polygon"

type Coordinate interface{}

type Geometry struct {
	Type        string     `json:"type,omitempty"`
	Coordinates Coordinate `json:"coordinates,omitempty"`
}
