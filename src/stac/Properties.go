package stac

// TODO: Add Band, Instrument, Data Values, Statistics, Data Type

type Properties struct {
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
	Roles        []string      `json:"roles,omitempty"`
	Datetime     string        `json:"datetime"`
	Created      string        `json:"created,omitempty"`
	Updated      string        `json:"updated,omitempty"`
	StartTime    string        `json:"start_datetime,omitempty"`
	EndTime      string        `json:"end_datetime,omitempty"`
	License      string        `json:"license,omitempty"`
	Providers    []Provider    `json:"providers,omitempty"`
	TableColumns []TableColumn `json:"table:columns,omitempty"`
}
