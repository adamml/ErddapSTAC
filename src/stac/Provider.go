package stac

type Provider struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Url         string   `json:"url,omitempty"`
}
