package stac

// Implements the STAC Catalog specification
type Catalog struct {
	Type           string   `json:"type"`
	StacVersion    string   `json:"stac_version"`
	StacExtensions []string `json:"stac_extensions,omitempty"`
	Id             string   `json:"id"`
	Title          string   `json:"title,omitempty"`
	Description    string   `json:"description"`
	Links          []Link   `json:"links"`
}
