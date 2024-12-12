package erddap

// EDDJsonTable is used to unmarshal an Erddap JSON table response
type EDDJsonTable struct {
	// Table contains an EDDJsonTableTable struct
	Table EDDJsonTableTable
}
