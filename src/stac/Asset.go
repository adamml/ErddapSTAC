package stac

type Asset struct {
	Href        string   `json:"href"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}
