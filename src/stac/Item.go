package stac

type Coordinate []interface{}

type Geometry struct {
	Type        string        `json:"type,omitempty"`
	Coordinates []interface{} `json:"coordinates,omitempty"`
}

type Item struct {
	Type           string           `json:"type"`
	StacVersion    string           `json:"stac_version"`
	StacExtensions []string         `json:"stac_extensions,omitempty"`
	Id             string           `json:"id"`
	Geometry       Geometry         `json:"geometry,omitempty"`
	Bbox           []float32        `json:"bbox,omitempty"`
	Properties     Properties       `json:"properties"`
	Links          []Link           `json:"links"`
	Assets         map[string]Asset `json:"assets"`
	Collection     string           `json:"collection,omitempty"`
}

func NewItem() Item {
	return Item{
		Type:           STAC_TYPES_ITEM,
		StacVersion:    STAC_VERSION,
		StacExtensions: make([]string, 0),
		Assets:         map[string]Asset{},
		Geometry:       Geometry{},
	}
}
