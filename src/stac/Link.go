package stac

// TODO: Add headers and body to the definition

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Method string `json:"method,omitempty"`
}
