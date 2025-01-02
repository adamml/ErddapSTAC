package stac

// TODO: Add Summaries, Assets and Item Assets

type Collection struct {
	Type             string     `json:"type"`
	StacVersion      string     `json:"stac_version"`
	StacExtensions   []string   `json:"stac_extensions,omitempty"`
	Id               string     `json:"id"`
	Title            string     `json:"title,omitempty"`
	Description      string     `json:"description"`
	Keywords         []string   `json:"keywords,omitempty"`
	License          string     `json:"license"`
	Providers        []Provider `json:"providers,omitempty"`
	CollectionExtent Extent     `json:"extent"`
	Links            []Link     `json:"links"`
}

func NewCollection() Collection {
	return Collection{
		Type:           STAC_TYPES_COLLECTION,
		StacVersion:    STAC_VERSION,
		StacExtensions: make([]string, 0),
		Providers:      make([]Provider, 0),
		Links:          make([]Link, 0),
		Keywords:       make([]string, 0),
	}
}
