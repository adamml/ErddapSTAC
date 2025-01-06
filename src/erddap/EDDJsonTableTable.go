package erddap

// EDDJsonTableTable is used for unmarshaling the table key of raw Erddap table JSON responses.
type EDDJsonTableTable struct {
	// Coulmnnames lists the names of the columns in the Erddap table response.
	Columnnames []string
	// Columntypes lists the data types of the columns in the Erddap table response.
	Columntypes []string
	// Column units lists the units of measure of each column in the Erddap table response.
	Columnunits []string
	// Rows is a multi-dimensional array of the table rows in the Erddap table response.
	// Each slice in Row contains items corresponding to the spec in Columnnames and Columntypes.
	Rows [][]string
}
