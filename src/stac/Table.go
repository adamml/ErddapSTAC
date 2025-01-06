package stac

type TableColumn struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Type        string     `json:"type,omitempty"`
	NoData      any        `json:"nodata,omitempty"`
	DataType    string     `json:"data_type,omitempty"`
	Unit        string     `json:"unit,omitempty"`
	Statistics  Statistics `json:"statistics,omitempty"`
}
